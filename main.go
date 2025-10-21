package main

import (
	"fmt"
)

func main() {

	//go run compile + run
	//go build only comple and run the binary later
	// var ptr *int
	// fmt.Println("value of pointer is ", ptr)

	for i := range 6 {
		fmt.Println(i)
	}
}
