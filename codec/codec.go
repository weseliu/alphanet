package codec

import "errors"

type Codec interface {
	Name() string
	Encode(msg interface{}) ([]byte, error)
	Decode(data []byte, msg interface{}) error
}

type codecManager struct {
	codecMap map[string] Codec
}

var instance *codecManager
func CodecManager() *codecManager{
	if instance == nil{
		instance = &codecManager{}
		instance.codecMap = make(map[string] Codec)
	}
	return instance
}

func (Self *codecManager)RegisterCodec(codec Codec){
	if codec == nil {
		return
	}
	Self.codecMap[codec.Name()] = codec
}

func (Self *codecManager)Encode(codecName string, msg interface{}) ([]byte, error){
	if Self.codecMap[codecName] != nil {
		return Self.codecMap[codecName].Encode(msg)
	}
	return nil, errors.New("encode codec not exist")
}

func (Self *codecManager)Decode(codecName string, data []byte, msg interface{}) error{
	if Self.codecMap[codecName] != nil {
		return Self.codecMap[codecName].Decode(data, msg)
	}
	return errors.New("decode codec not exist")
}