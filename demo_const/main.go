package main

import "fmt"

// 常量
const PI = 3.12415
const Name = "cao"

// 性别 枚举
const (
	Unknow = 0
	Female = 1
	Male   = 2
)

// iota
const (
	a = iota
	b
	c
)

// iota 运算后的值,每行都会增加1000
const (
	YearOne = iota + 1000
	YearTwo
	YearThree
	YearFour
)

/**
3.12415
cao

unknow = 0;Female = 1,Male=2
a = 0;b = 1,c=2
YearOne = 1000;YearTwo = 1001,YearThree=1002;YearFour=1003
*/
func main() {
	fmt.Println(PI)
	fmt.Println(Name)
	fmt.Printf("unknow = %d;Female = %d,Male=%d\n", Unknow, Female, Male)
	fmt.Printf("a = %d;b = %d,c=%d\n", a, b, c)
	fmt.Printf("YearOne = %d;YearTwo = %d,YearThree=%d;YearFour=%d\n", YearOne, YearTwo, YearThree, YearFour)
}
