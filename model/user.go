package model

type User struct {
	ID          int64  `json:"id"`
	UserName    string `json:"user_name"`
	Password    string `json:"password"`
	NewPassword string `json:"new_password"`
}
