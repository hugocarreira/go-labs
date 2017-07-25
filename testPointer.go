package main

import "fmt"

func main() {
	i := 0
	p := &i

	fmt.Println(*p)
}
