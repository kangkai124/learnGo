package main

import (
	"fmt"
	"math/cmplx"
	"math"
)

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

func euler() {
	fmt.Println(cmplx.Exp(1i * math.Pi) + 1)
	fmt.Println(cmplx.Pow(math.E, 1i * math.Pi) + 1)
	fmt.Printf("%.3f\n", cmplx.Exp(1i * math.Pi) + 1)
}

/*
 * 变量定义
 * 变量类型写在变量名之后
 * 编译器可推测变量类型
 * 没有char, 只有rune
 * 原生支持复数类型
 */

func trangle() {
	var a, b int = 3, 4
	var c int
	// c = math.Sqrt(a * a + b * b)
	c = int(math.Sqrt(float64(a * a + b * b)))
	fmt.Println(c)
}

func consts() {
	const filename = "abc.txt"
	/*
	 * const a, b = 3, 4 没有定义类型，会根据调用的时候判断类型，如调用math.Sqrt函数a, b类型判断为float64
	 * const a, b int = 3, 4 如果定义了类型，下面调用math.Sqrt则需要再转为float64类型
	 */
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a * a + b * b))
	fmt.Println(filename, c)
}

func enums() {
	const (
		cpp = iota
		java
		_
		python
		golang
		javascript
	)
	fmt.Println(cpp, java, python, golang, javascript)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		pb
	)
	fmt.Println(b, kb, mb, gb, pb)
}

func main() {
	fmt.Println("hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(kk, rr)

	euler()
	trangle()

	consts()

	enums()
}
