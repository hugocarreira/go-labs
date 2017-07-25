package main

import "fmt"

func main () {
	group := []int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println(group)

	var one []int = group[8:10]
	fmt.Println(one)

	fmt.Println(len(group))
}
