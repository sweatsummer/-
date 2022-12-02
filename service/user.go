package service

import (
	"main/dao"
	"main/model"
)

func IsExistUserByName(name string) (u model.User, err error) {
	u, err = dao.SearchUserByUserName(name)
	return
}

func CreateUser(u model.User) error {
	err := dao.InsertUser(u)
	return err
}

func UpdataPasseword(name string) (u *model.User, err error) {
	u, err = dao.AlterPasseword(name)
	return
}

func FindID(name string) (u *model.Message, err error) {
	u, err = dao.SearchID(name)
	return
}

func Lookmessage(ID int64) (u *model.Message, err error) {
	u, err = dao.SearchMessageByID(ID)
	return
}
