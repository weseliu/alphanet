package protocal

import (
	"github.com/weseliu/alphanet/codec"
	_ "github.com/weseliu/alphanet/codec/json"
)


type UserAuth struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Channel  string `json:"channel"`
	DeviceId string `json:"deviceId"`
	Platform string `json:"platform"`
}

type AuthResult struct {
	Token     string `json:"token"`
}

func init(){
	codec.RegisterMessageMeta("json", (*UserAuth)(nil))
}
