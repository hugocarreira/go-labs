package main

import (
	"fmt"
	"time"
)

func test(i int) {
	switch x := i; x {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	default:
		fmt.Println(x)
	}
}

func test2() {
	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}
}

func main() {
	fmt.Print("The number is: ")
	test(2)

	fmt.Print("Hour Message: ")
	test2()
}
