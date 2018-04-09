package interfaces

import (
	"time"
	"github.com/weseliu/alphanet/util"
	"encoding/json"
	"encoding/base64"
)

var encryptKey = []byte("Edas@3ad")

type IdentityToken struct {
	Account string
	Channel string
	Time    time.Time
}

func (Self *IdentityToken)Encrypt() (string, error){
	data, err := json.Marshal(Self)
	if err == nil{
		bytes, err := util.DesEncrypt(data, encryptKey)
		return base64.StdEncoding.EncodeToString(bytes), err
	}
	return "", err
}

func (Self *IdentityToken)Decrypt(base64String string) error {
	bytes, err := base64.StdEncoding.DecodeString(base64String)
	if err == nil {
		data, err := util.DesDecrypt(bytes, encryptKey)
		if err == nil {
			err = json.Unmarshal(data, Self)
			return err
		}
	}
	return err
}