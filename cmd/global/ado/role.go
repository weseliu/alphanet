package ado

import (
	"github.com/weseliu/alphanet/db"
	"log"
)

type RoleModel struct {
	Id 	int64 `table:"role"`
	Name string
	Age int
}

func Role() *RoleModel {
	return nil
}

func (Self *RoleModel) GetRole(roleId int64) *RoleModel {
	data := db.Instance().Query(&RoleModel{}, "select * from role where Id = ?", roleId)
	if data != nil {
		return data.(*RoleModel)
	}
	return nil
}

func (Self *RoleModel) AddRole(user *RoleModel) bool {
	_, err := db.Instance().Insert(user)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
