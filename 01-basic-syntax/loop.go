package main

import (
	"strconv"
	"fmt"
	"os"
	"bufio"
)

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}

func main() {
	fmt.Println(
		convertToBin(13),
		convertToBin(53354),
		convertToBin(0),
	)
	printFile("./abc.txt")
}


