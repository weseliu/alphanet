package json

import (
	"encoding/json"
	"github.com/weseliu/alphanet/codec"
)

type jsonCodec struct {
}

func (Self *jsonCodec) Name() string {
	return "json"
}

func (Self *jsonCodec) Encode(msgObj interface{}) ([]byte, error) {
	return json.Marshal(msgObj)
}

func (Self *jsonCodec) Decode(data []byte, msgObj interface{}) error {
	return json.Unmarshal(data, msgObj)
}

func init() {
	codec.CodecManager().RegisterCodec(&jsonCodec{})
}
