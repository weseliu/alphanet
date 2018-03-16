package codec

import (
	"errors"
	"go/types"
)

type Codec interface {
	Name() string
	Encode(msg interface{}) ([]byte, error)
	Decode(data []byte)  (interface{}, error)
}

type codecManager struct {
	codecMap map[string] Codec
	messageMeta map[string] map[string] types.Type
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

func (Self *codecManager)GetCodec(codecName string) Codec{
	return Self.codecMap[codecName]
}

func (Self *codecManager)Encode(codecName string, msg interface{}) ([]byte, error){
	if Self.codecMap[codecName] != nil {
		return Self.codecMap[codecName].Encode(msg)
	}
	return nil, errors.New("encode codec not exist")
}

func (Self *codecManager)Decode(codecName string, data []byte) (interface{}, error){
	if Self.codecMap[codecName] != nil {
		return Self.codecMap[codecName].Decode(data)
	}
	return nil, errors.New("decode codec not exist")
}