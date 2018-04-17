package main

import "fmt"

func printSliceInfo(s []int) {
	fmt.Printf("s=%v, len=%d, cap=%d\n", s, len(s), cap(s))
}

func main() {
	fmt.Println("Creating slice")
	var s []int  //Zero value for slice is nil

	for i := 0; i < 100; i++ {
		printSliceInfo(s)
		s = append(s, 2 * i + 1)
	}
	fmt.Println(s)

	s1 := []int{2, 4, 6, 8}
	printSliceInfo(s1)

	// make([]T, length, capacity)
	s2 := make([]int, 16)
	s3 := make([]int, 10, 32)
	printSliceInfo(s2)
	printSliceInfo(s3)

	fmt.Println("copying slice")
	copy(s2, s1)
	printSliceInfo(s2)

	fmt.Println("Deleting elements from slice")
	// s2[:3] + s2[4:]
	s2 = append(s2[:3], s2[4:]...)
	fmt.Println("after delete 8")
	printSliceInfo(s2)

	fmt.Println("Popping from front")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println(front)
	printSliceInfo(s2)	// 删掉前面，cap减少

	fmt.Println("Popping from back")
	tail := s2[len(s2) - 1]
	s2 = s2[:len(s2) - 1]
	fmt.Println(tail)
	printSliceInfo(s2)
}