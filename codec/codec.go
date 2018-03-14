package codec

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

func (Self *codecManager)Encode(msg interface{}) ([]byte, error){
	return nil, nil
}

func (Self *codecManager)Decode(data []byte, msg interface{}) error{
	return nil
}