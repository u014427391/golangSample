package main

import "fmt"

func main() {
	res := addUpper(10)
	fmt.Printf("返回值res:%v", res)
}

func add(n1 int , n2 int) int {
	return n1 + n2
}

func addUpper(n int) int {
	res := 0
	for i := 0 ; i <= n ; i++ {
		res += i
	}
	return res
}

