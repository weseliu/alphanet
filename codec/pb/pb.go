package pbcodec

import (
	"github.com/golang/protobuf/proto"
	"github.com/weseliu/alphanet/codec"
)

type pbCodec struct {
}

func (Self *pbCodec) Name() string {
	return "pb"
}

func (Self *pbCodec) Encode(msgObj interface{}) ([]byte, error) {
	msg := msgObj.(proto.Message)
	data, err := proto.Marshal(msg)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (Self *pbCodec) Decode(data []byte) (msgObj interface{}, err error) {
	err = proto.Unmarshal(data, msgObj.(proto.Message))
	if err != nil {
		return nil, err
	}
	return msgObj, nil
}

func init() {
	codec.CodecManager().RegisterCodec(&pbCodec{})
}
