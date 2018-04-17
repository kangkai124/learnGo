package main

import "fmt"

// slice 是对数组的一个view

// updateSlice(s []int) {}    []int  是slice
func updateSlice(s []int) {
	s[0] = 666
}

func main() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("arr[2, 6] =", arr[:])
	fmt.Println("arr[2, 6] =", arr[2:6])
	fmt.Println("arr[2, 6] =", arr[:6])
	fmt.Println("arr[2, 6] =", arr[2:])

	s1 := arr[2:]
	s2 := arr[:]
	fmt.Println("s1=", s1)
	fmt.Println("s2=", s2)

	fmt.Println("After UpdateSlice(s1)")
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)

	fmt.Println("After UpdateSlice(s2)")
	updateSlice(s2)
	fmt.Println(s2)
	fmt.Println(arr)

	fmt.Println("Extending Slice")
	arr[0] = 1; arr[2] = 3
	fmt.Println("arr =", arr)
	m1 := arr[2:6]
	m2 := m1[3:5]
	fmt.Println(m1, m2)  // [3, 4, 5, 6] [6, 7]
	// slice 可以向后扩展，不可以向前扩展
	fmt.Printf("m1=%v, len(m1)=%d, cap(m1)=%d\n",
		m1, len(m1), cap(m1))
	fmt.Printf("m2=%v, len(m2)=%d, cap(m2)=%d\n",
		m2, len(m2), cap(m2))
	/*
	 * m1的cap=6, 即容量为6， 即[3, 4, 5, 6, 7, 8]
	 * m2的cap=3, 即容量为6， 即[6, 7, 8]
	 */

	m3 := append(m2, 10)
	m4 := append(m3, 11)
	m5 := append(m4, 12)
	// m4, m5 不再是arr的一个view
	fmt.Println("m3, m4, m5=", m3, m4, m5)
	fmt.Println("arr=", arr)

	// slice添加元素时如果超过cap，系统会重新分配更大的底层数组
}
