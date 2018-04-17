package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

type iAdder func(int) (int, iAdder)

func advancedAdder(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, advancedAdder(base + v)
	}
}

func main() {
	a := adder()
	for i := 0; i < 10; i++ {
		fmt.Printf("0+1+...+%d = %d\n", i, a(i))
	}

	fmt.Println("advanced adder...")
	v := advancedAdder(0)
	for i := 0; i < 10; i++ {
		var u int
		u, v = v(i)
		fmt.Printf("0+1+...+%d = %d\n", i, u)
	}
}
