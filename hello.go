package main

import "fmt"

func main() {
	fmt.Println("hello world")

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	if num := 9; num < 0 {
		fmt.Println("negative number")
	} else if num < 10 {
		fmt.Println(num, " has 1 digit")
	}
}
