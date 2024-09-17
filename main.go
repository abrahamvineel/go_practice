package main

import (
	"fmt"
)

func main() {
	// var ptr *int
	// fmt.Println("value of pointer is ", ptr)

	myNumber := 23
	var ptr = &myNumber

	// fmt.Println("value of actual pointer is", ptr)
	// fmt.Println("value of actual pointer is", *ptr)

	*ptr = *ptr * 2
	// fmt.Println("new value of actual pointer is", myNumber)

	//slices

	var fruitList = []string{"Apple"}
	fmt.Printf("type of fruit list %T\n", fruitList)

	fruitList = append(fruitList, "banana")

	fmt.Println(fruitList)

	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)
}
