package json

import (
	"encoding/json"
	"github.com/weseliu/alphanet/codec"
	"reflect"
)

type jsonCodec struct {
	MsgName string `json:"name"`
	MsgData string `json:"data"`
}

func (Self *jsonCodec) Name() string {
	return "json"
}

func (Self *jsonCodec) Encode(msgObj interface{}) (data []byte, err error) {
	v := reflect.TypeOf(msgObj).Elem()
	Self.MsgName = v.Name()
	data, err = json.Marshal(msgObj)
	Self.MsgData = string(data)
	if err != nil {
		return nil, err
	}
	return json.Marshal(Self)
}

func (Self *jsonCodec) Decode(data []byte) (msgObj interface{}, err error) {
	err = json.Unmarshal(data, Self)
	if err != nil {
		return nil, err
	}

	msgObj = codec.BuildMessage(Self.MsgName)
	err = json.Unmarshal([]byte(Self.MsgData), msgObj)
	return msgObj, err
}

func init() {
	codec.CodecManager().RegisterCodec(&jsonCodec{})
}
