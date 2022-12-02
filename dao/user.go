package dao

import (
	"main/model"
)

// SearchUserByUserName 查询数据库//
func SearchUserByUserName(name string) (u model.User, err error) {
	row := DB.QueryRow("select id,name,password from user where name=?", name)
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&u.ID, &u.UserName, &u.Password)
	return
}

// InsertUser 插入数据库注册数据//
func InsertUser(u model.User) (err error) {
	_, err = DB.Exec("insert into user(name,password) values (?,?)", u.UserName, u.Password)
	if err != nil {
		return
	}
	return err
}

// AlterPasseword 修改密码//
func AlterPasseword(name string) (u *model.User, err error) {
	_, err = DB.Exec("update user set password=? where name=?", name, u.UserName)
	if err != nil {
		return
	}
	return
}

func SearchID(name string) (u *model.Message, err error) {
	row := DB.QueryRow("select id from user where name=?;", name)
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&u.OwnerUID)
	return
}

func SearchMessageByID(ID int64) (u *model.Message, err error) {
	row := DB.QueryRow("select detail from user where send_uid=?", ID)
	if err = row.Err(); row.Err() != nil {
		return nil, nil
	}
	err = row.Scan(&u.Detail)
	return nil, nil
}
