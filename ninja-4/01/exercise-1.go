package main

import (
	"fmt"
)

func main() {
	x := []int{1983, 1986, 1985, 1992, 2019}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)
}
