package main

import "fmt"

func test(x string) (y string) {
	y = x
    return
}

func main() {
	fmt.Println(test("Hugo Carreira"))
}