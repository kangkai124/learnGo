package main

import "fmt"

// Go 的数组是值类型

func printArray(arr [5]int) {	// 会拷贝数组
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func printArrayByPointer(arr *[5]int) {	// 会拷贝数组
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	// 定义数组的几种方法
	var arr1 [5]int
	arr2 := [3]int{1, 2, 3}
	arr3 := [...]int{2, 4, 6, 8, 10}
	var grid [4][5]int

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	for i, v := range arr3 {	// key value
		fmt.Println(i, v)
	}
	for _, v := range arr3 {
		fmt.Println(v)
	}

	//printArray(arr2)  [3]int 和 [5]int 是不同类型

	fmt.Println("printArray(arr1)")
	printArray(arr1)

	fmt.Println("printArray(arr3)")
	printArray(arr3)

	fmt.Println("arr1 and arr3")
	fmt.Println(arr1, arr3)		// 由于值类型，arr1和arr3没有改变

	var (
		arr4 = [5]int{1, 2, 3, 4, 5}
		arr5 = [5]int{2, 3, 4, 5, 6}
	)
	fmt.Println("printArrayByPointer(arr1)")
	printArrayByPointer(&arr4)

	fmt.Println("printArrayByPointer(arr3)")
	printArrayByPointer(&arr5)

	fmt.Println("arr1 and arr3")
	fmt.Println(arr4, arr5)		// 由于值类型，arr1和arr3没有改变
}
