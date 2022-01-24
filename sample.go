package main

import (
	"fmt"
	"golangSample/mysql"
	"math"
	"unsafe"
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

	var r rune = 1
	fmt.Printf("r类型%T \n" , r)
	var b byte = 1
	fmt.Printf("b类型%T \n" , b)

	var n int16
	fmt.Printf("n变量占用的字节数:%d \n" , unsafe.Sizeof(n))

	//var intV1 int16
	//intV2 := 10
	//intV1 = intV2
	//fmt.Println(intV1)

	var fNum1 float32 = -123.0000990
	var fNum2 float64 = -123.0000990
	fmt.Println("fNum1=" , fNum1 , "fNum2=" , fNum2)

	// 最小误差值
	p := 0.000001
	// 判断两个浮点数误差是否在误差值之间
	if math.Dim(float64(fNum1), fNum2) < p {
		fmt.Println("fNum1 和 fNum2 相等")
	}

	// 字符 byte
	var c byte = 'a'
	fmt.Println("c=" , c)
	fmt.Printf("c=%c \n" , c)

	//var ct byte = '，'
	//fmt.Println("ct=" , ct)

	var bol bool = true
	fmt.Println("bol=" , bol)
	fmt.Println("size of bool=" , unsafe.Sizeof(bol))


	var str string = "golang"
	//cannot assign to str[0]
	// str[0] = 's'
	fmt.Printf("%s length is %d \n" , str , len(str))
	// 反引号
	res := `programming language :
		Java
		Golang
		Python
		PHP
		...
		`
	res1 := "programming language :\n" +
		"Java\n" +
		"Golang\n" +
		"Python\n" +
		"PHP\n" +
		"...\n"
	fmt.Printf("%s" , res)
	fmt.Printf("%s" , res1)
	// 字符串拼接
	str = str + ",learning(学习)"
	fmt.Printf("%s\n" , str)
	// [0:6) 左闭右开区间
	str1 := str[:6]
	// [7:len(str)]
	str2 := str[7:]
	// 注意，这里也是[0:6)
	str3 := str[0:6]
	fmt.Println("str1=" , str1 , "str2=" , str2 , "str3=", str3)
	// 遍历字符串
	for i := 0 ; i <len(str) ; i++ {
		ch := str[i]
		fmt.Printf("%d %v %c\n" , i , ch,  ch)
	}
	fmt.Println("range遍历")
	for i , ch := range str {
		fmt.Printf("%d %v %c\n" , i , ch,  ch)
	}



}

func GetUserInfo() (userName string , password string) {
	return "admin", "pwd"
}
