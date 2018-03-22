package protocal

import (
	"github.com/weseliu/alphanet/codec"
	_ "github.com/weseliu/alphanet/codec/json"
)


type UserAuth struct {
	Id		 int64 `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Channel  string `json:"channel"`
	DeviceId string `json:"deviceId"`
	Platform string `json:"platform"`
}

type AuthResult struct {
	Ret       int		`json:"ret"`
	Msg    string `json:"msg"`
	Token     string `json:"token"`
}

func init(){
	codec.RegisterMessageMeta("json", (*UserAuth)(nil))
}
