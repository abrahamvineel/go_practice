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

	//no inheritance in golang; no super or parent

	abraham := User{"Abraham", "abc.xyz.com", true, 16}
	fmt.Printf("%+v\n", abraham)
	fmt.Println("Name is %v and email is %v", abraham.Name, abraham.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
