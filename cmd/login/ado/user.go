package ado

import "github.com/weseliu/alphanet/db"

type UserModel struct {
	Id       int64
	Name     string
	Password string
	Agg      int
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

func (Self *UserModel) AddUser(user *UserModel) {

}

func init() {
	db.Instance().RegisterModel(&UserModel{})
}
