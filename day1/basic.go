package main

import "fmt"

var aa = 3	// 函数外，包内部的变量，只能用var
var ss = "kkk"

var (
	kk = "kk"
	rr = "rr"
)

func variableZeroValue() {
	var a int	// 变量名 变量类型
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b = 3, 4	// var a, b int = 3, 4  当赋值时， type可以省略
	var s = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "hah"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	a, b, c, s := 3, 4, false, "hah"	// :定义，可省略var，但是冒号只能在函数里使用
	fmt.Println(a, b, c, s)
}

func main() {
	fmt.Println("hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(kk, rr)
}
