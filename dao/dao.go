package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB *sql.DB

func InitDB() {
	//访问本地数据库无密码//
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/message_borad")
	if err != nil {
		log.Fatalf("connect mysql error : %v", err)
	}
	DB = db
	//检验连接//
	fmt.Println(db.Ping())
}
