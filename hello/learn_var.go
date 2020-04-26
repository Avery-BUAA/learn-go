package main

import (
	"fmt"
)

func main() {
	var num1 int
	num1 = 300
	fmt.Printf("shuzhishi: %d\n", num1)
	fmt.Printf("shuzhishi: %d\n", num1)

	var num2 int = 400
	fmt.Printf("shuzhishi: %d\n", num2)

	//类型自动推断
	var name = "王二狗"
	fmt.Printf("类型是%T，数值是%s\n ",name,name)

	//简短声明
	age := 18
	fmt.Print(age)

	//多个变量同时声明
	var m,n = 100,200
	fmt.Print(m,n)
}