package main

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"fmt"
	"strings"
)

const (
	userName = "root"
	password = ""
	ip = "127.0.0.1"
	port = "3306"
	dbName = "shop"
	driverName = "mysql"
)

func main() {
	initDB()
}

func initDB()  {
	dataSourceName := strings.Join([]string{userName, ":", password, "@tcp(",ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	DB, _ :=  sql.Open(driverName , dataSourceName )
	DB.SetConnMaxLifetime(100)
	DB.SetMaxIdleConns(10)
	if err := DB.Ping(); err != nil {
		fmt.Println("open database fail")
		return
	}
	fmt.Println("connnect success")
}
