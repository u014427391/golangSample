package main

import (
	"fmt"
	"golangSample/mysql"
)


var v1 int            // 整型
var v2 string         // 字符串
var v3 bool           // 布尔型
var v4 [10]int        // 数组，类型为整型
var v5 struct {       // 结构体
	v float32
}
var v6 *int           // 指针，指向整型
var v7 map[string]int   // map key为字符串，value为整型
var v8 func(a int) int // 函数



type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

const (
	_ = 1 << (10 * iota)
	KiB // 1024
	MiB // 1048576
	GiB // 1073741824
	TiB // 1099511627776             (exceeds 1 << 32)
	PiB // 1125899906842624
	EiB // 1152921504606846976
	ZiB // 1180591620717411303424    (exceeds 1 << 64)
	YiB // 1208925819614629174706176
)


const (
	c1 = iota // 0
	c2        // 1
	c3		  // 2
)

const (
	c4 = 1 << iota // 1
	c5			   // 2
	c6             // 4
)

const (
	c7 = iota * 2 // 0
	c8            // 2
	c9            // 4
)

func main() {
	mysql.InitDB()
	mysql.QueryAll()
	mysql.DB.Close()


	// 匿名变量
	userName,_ := GetUserInfo()
	fmt.Println(userName)

	fmt.Printf("v1: %v\n", v1)
	fmt.Printf("v2: %v\n", v2)
	fmt.Printf("v3: %v\n", v3)
	fmt.Printf("v4: %v\n", v4)
	fmt.Printf("v5: %v\n", v5)
	fmt.Printf("v6: %v\n", v6)
	fmt.Printf("v7: %v\n", v7)
	fmt.Printf("v8: %v\n", v8)

	fmt.Printf("c1:%v\n" , c1)
	fmt.Printf("c2:%v\n" , c2)
	fmt.Printf("c3:%v\n" , c3)
	fmt.Printf("c4:%v\n" , c4)
	fmt.Printf("c5:%v\n" , c5)
	fmt.Printf("c6:%v\n" , c6)
	fmt.Printf("c7:%v\n" , c7)
	fmt.Printf("c8:%v\n" , c8)
	fmt.Printf("c9:%v\n" , c9)

	var i int = 10
	fmt.Printf("i数据类型:%T \n" , i)

}

func GetUserInfo() (userName string , password string) {
	return "admin", "pwd"
}
