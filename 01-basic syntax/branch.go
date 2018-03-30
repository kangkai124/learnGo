package main

import (
	"fmt"
	"io/ioutil"
)

func check() {
	// if条件里可以赋值
	// if条件里赋值的变量的作用域就在这个if语句里
	if contents, err := ioutil.ReadFile("abc.txt"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}

func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf(
			"Wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}

func main() {
	fmt.Println(
		grade(59),
		grade(69),
		grade(88),
		grade(92),
	)
	check()
}
