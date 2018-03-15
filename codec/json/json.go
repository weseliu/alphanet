package json

import (
	"encoding/json"
	"github.com/weseliu/alphanet/codec"
	"reflect"
)

type jsonCodec struct {
	name  string
	data  []byte
}

func (Self *jsonCodec) Name() string {
	return "json"
}

func (Self *jsonCodec) Encode(msgObj interface{}) (data []byte, err error) {
	v := reflect.TypeOf(msgObj)
	Self.name = v.Name()
	Self.data, err = json.Marshal(msgObj)
	if err != nil {
		return nil, err
	}
	return json.Marshal(Self)
}

func (Self *jsonCodec) Decode(data []byte, msgObj interface{}) (err error) {
	err = json.Unmarshal(data, Self)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, msgObj)
}

func init() {
	codec.CodecManager().RegisterCodec(&jsonCodec{})
}
