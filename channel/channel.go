package channel

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"log"
	"net"
	"time"
	"github.com/pkg/errors"
	"unsafe"
)

type PackageHead struct {
	Size int32
	Seq  int32
	Flag int32
}

type Package struct {
	*PackageHead
	Data []byte
}

func NewPackage(seq int32, msg []byte) *Package {
	return &Package{
		PackageHead: &PackageHead{
			Size: int32(len(msg)),
			Seq:  seq,
			Flag: 0,
		},
		Data: msg,
	}
}

func (Self *Package) Encode() []byte {
	Self.Size = int32(len(Self.Data) + int(unsafe.Sizeof(Self.Seq) + unsafe.Sizeof(Self.Flag)))

	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, Self.Size)
	binary.Write(buffer, binary.LittleEndian, Self.Seq)
	binary.Write(buffer, binary.LittleEndian, Self.Flag)
	binary.Write(buffer, binary.LittleEndian, Self.Data)

	return buffer.Bytes()
}

func (Self *Package) Decode(data []byte) (err error) {
	buffer := new(bytes.Buffer)
	buffer.Write(data)

	err = binary.Read(buffer, binary.LittleEndian, &Self.Size)
	if err != nil {
		return
	}

	Self.Data = make([]byte, Self.Size - int32(unsafe.Sizeof(Self.Seq) + unsafe.Sizeof(Self.Flag)))

	err = binary.Read(buffer, binary.LittleEndian, &Self.Seq)
	if err != nil {
		return
	}
	err = binary.Read(buffer, binary.LittleEndian, &Self.Flag)
	if err != nil {
		return
	}
	return binary.Read(buffer, binary.LittleEndian, Self.Data)
}

func (Self *Package)Split(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF {
		return 0, nil, errors.New("conn closed!")
	}

	var sizeLen = int(unsafe.Sizeof(Self.Size))
	if len(data) > sizeLen {
		var length int32
		binary.Read(bytes.NewReader(data[0:sizeLen]), binary.LittleEndian, &length)
		pkgSize := int(length) + sizeLen
		if pkgSize <= len(data) {
			return pkgSize, data[:pkgSize], nil
		}
	}
	return
}

type Channel struct {
	listen     net.Listener
	conn       net.Conn
	address    string
	timeout    time.Duration
	inputChan  chan *Package
	outputChan chan []byte
	seq		int32
	remoteSeq int32
}

func (Self *Channel) Listener() (err error) {
	Self.listen, err = net.Listen("tcp", Self.address)
	if err == nil {
		go Self.accept()
	}
	return err
}

func (Self *Channel) Connect() (err error) {
	Self.conn, err = net.Dial("tcp", Self.address)
	if err != nil {
		return
	}

	go Self.handlerRead(Self.conn)
	go Self.handlerWrite(Self.conn)
	return
}

func (Self *Channel) Send(msg []byte, callback func(result int)) {
	go Self.SendSync(msg, callback)
}

func (Self *Channel)SendSync(msg []byte, callback func(result int)) {
	Self.seq++
	pkg := NewPackage(Self.seq, msg)

	select {
	case Self.outputChan <- pkg.Encode():
		if callback != nil {
			callback(0)
		}
	case <-time.After(Self.timeout * time.Second):
		if callback != nil {
			callback(-1)
		}
	}
}

func (Self *Channel) ReadLoop(reader func([]byte)) {
	for {
		select {
		case data, ok := <-Self.inputChan:
			if ok {
				reader(data.Data)
			} else {
				break
			}
		case <-time.After(Self.timeout * time.Second):
			log.Println("read idle!")
			continue
		}
	}
}

func (Self *Channel) accept() {
	defer close(Self.inputChan)
	defer close(Self.outputChan)
	defer Self.listen.Close()
	for {
		conn, err := Self.listen.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go Self.handlerRead(conn)
		go Self.handlerWrite(conn)
	}
}

func (Self *Channel) handlerRead(conn net.Conn) {
	pkg := &Package{PackageHead : &PackageHead{}}
	scanner := bufio.NewScanner(conn)
	scanner.Split(pkg.Split)

	for {
		scanner.Scan()
		err := scanner.Err()
		if err != nil{
			conn.Close()
			log.Println(err)
			return
		}

		pkg.Decode(scanner.Bytes())
		if pkg.Seq < Self.remoteSeq || Self.remoteSeq == 0 {
			Self.inputChan <- pkg
		} else {
			log.Println("package seq check fail! remote : ", pkg.Seq, ", local : ", Self.remoteSeq)
		}
		/*select {
		case Self.inputChan <- pkg:
		case <-time.After(Self.timeout * time.Second):
			log.Println("inputChan timepout!")
			continue
		}*/
	}
}

func (Self *Channel) handlerWrite(conn net.Conn) {
	for {
		select {
		case output := <-Self.outputChan:
			_, err := conn.Write([]byte(output))
			if err != nil {
				conn.Close()
				return
			}
		default:
		}
	}
}

func NewChannel(address string, timeout time.Duration, inputChanSize int, outputChanSize int) *Channel {
	channel := &Channel{
		inputChan : make(chan *Package, inputChanSize),
		outputChan : make(chan []byte, outputChanSize),
		timeout : timeout,
		address : address,
	}
	return channel
}