package ado

import (
	"github.com/weseliu/alphanet/db"
	"log"
)

type UserModel struct {
	Id       int64 `table:"user"`
	Name     string
	Password string
	Age      int
	Address  string
}

func User() *UserModel {
	return nil
}

func (Self *UserModel) GetUser(id int64) *UserModel {
	data := db.Instance().Query(&UserModel{}, "select * from user where Id = ?", id)
	if data != nil {
		return data.(*UserModel)
	}
	return nil
}

func (Self *UserModel) AddUser(user *UserModel) (id int64) {
	id, err := db.Instance().Insert(user)
	if err != nil {
		log.Println(err)
	}
	return
}
