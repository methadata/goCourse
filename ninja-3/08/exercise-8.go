package main

import "fmt"

func main() {
	a := false
	switch {
	case a == true:
		fmt.Println("a is true")
	case a == false:
		fmt.Println("a is false")

	default:
		fmt.Println("nothing to see here")

	}
}
