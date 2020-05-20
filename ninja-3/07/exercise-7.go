package main

import "fmt"

func main() {
	a := false
	if a == true {
		fmt.Println(true)
	} else if a != true {
		fmt.Printf("elsif %v\n", a)
	} else {
		fmt.Println("never here")
	}
}
