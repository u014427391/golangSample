package main

import "golangSample/mysql"

func main() {
	mysql.InitDB()
	mysql.QueryAll()
	mysql.DB.Close()
}
