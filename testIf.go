package main

import "fmt"

func test(x float64) {
	if x < 0 {
		fmt.Println("< 0")
	} else if x > 0 {
		fmt.Println("> 0")
	} else {
		fmt.Println("w/e")
	}
	return
}

func main() {
	test(0)
}