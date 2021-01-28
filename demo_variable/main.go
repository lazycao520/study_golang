// go 变量
package main

import (
	"fmt"
	"os"
	"runtime"
)

// 全局变量
var globalAge int
var b = 200

// 值类型 引用类型
var name = "cao"
var globalA = "G"

var c string

func n() {
	fmt.Println(globalA)
}
func m() {
	globalA = "O"
	fmt.Println(globalA)
}

func f1() {
	c := "O"
	print(c)
	f2()
}
func f2() {
	print(c)
}

func main() {
	// 局部变量
	var a int
	globalAge = 20
	fmt.Printf("globalAge = %d", globalAge)
	fmt.Println(a)

	var goos string = runtime.GOOS
	fmt.Printf("The operating system is: %s\n", goos)
	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)

	var last_name = name
	// 变量作用域
	fmt.Println(&name, &last_name)
	n()
	m()
	n()

	//
	c = "G"
	print(c)
	f1()
}
