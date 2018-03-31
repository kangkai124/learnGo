package main

import "fmt"

func printSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}

func main() {
	var numbers []int
	printSlice(numbers)

	numbers = append(numbers, 0)
	numbers = append(numbers, 1)
	numbers = append(numbers, 2,3,4)
	printSlice(numbers)

	numbers1 := make([]int, len(numbers), (cap(numbers))*2)
	copy(numbers1,numbers)
	printSlice(numbers1)
}

