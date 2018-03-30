package main

import (
	"fmt"
	"reflect"
	"runtime"
	"math"
)

/*
 * 函数可以返回多个值
 * Go 函数式一等公民
 * 函数可作为参数
 * 没有默认参数，可选参数
 */

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf(
			"unsupported operation: %s", op)
	}
}

func div(a, b int) (q, r int) {
	return a / b, a % b
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling funciton %s with args" + "(%d, %d)\n", opName, a, b)
	return op(a, b)
}

func pow (a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

// 可变参数列表
func sumArgs(values ...int) int {
	sum := 0
	for i :=range values {
		sum += values[i]
	}
	return sum
}

// 指针
func swap(a, b *int) {
	*a, *b = *b, *a
}

func swap2(a, b int) (int, int) {
	return b, a
}

func main() {
	if result, err := eval(14, 3, "x"); err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println(result)
	}
	q, r := div(13, 3)
	one, _ := div(15, 7)
	fmt.Println(q, r, one)

	fmt.Println(apply(pow, 3, 4))
	fmt.Println(apply(
		func(a, b int) int {
			return int(math.Pow(
				float64(a), float64(b)))
		}, 3, 4))

	fmt.Println(sumArgs(1, 2, 3, 4, 5))

	a, b := 3, 4
	swap(&a, &b)
	c, d := 4, 5
	c, d = swap2(c, d)
	fmt.Println(a, b, c, d)
}
