package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

const (
	userName   = "root"
	password   = ""
	ip         = "127.0.0.1"
	port       = "3306"
	dbName     = "shop"
	driverName = "mysql"
)

var DB *sql.DB

type User struct {
	id int64
	name  string
	age int
	email string
	contactNumber string
	password string
	sex int
}

func InitDB() {
	dataSourceName := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	DB, _ = sql.Open(driverName, dataSourceName)
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	// 验证连接
	if err := DB.Ping(); err != nil {
		fmt.Println("open database fail")
		return
	}
	fmt.Println("connect success")


}

func QueryAll() {
	var user User
	rows, err := DB.Query("select * from user")
	if err != nil {
		fmt.Println("query failed.")
	}
	for rows.Next() {
		rows.Scan(&user.id , &user.name ,&user.age ,&user.email , &user.contactNumber , &user.password , &user.sex)
		fmt.Println(user)
	}
	rows.Close()
}

