package ado

import (
	"github.com/weseliu/alphanet/db"
	"log"
)

type RegisterModel struct {
	Account 	string `table:"register"`
	Channel     string
	Password    string
}

func Register() *RegisterModel {
	return nil
}

func (Self *RegisterModel) GetUser(channel string, account string) *RegisterModel {
	data := db.Instance().Query(&RegisterModel{}, "select * from register where Account = ? and Channel = ?", channel, account)
	if data != nil {
		return data.(*RegisterModel)
	}
	return nil
}

func (Self *RegisterModel) AddUser(user *RegisterModel) bool {
	_, err := db.Instance().Insert(user)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
