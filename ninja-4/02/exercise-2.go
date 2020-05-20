package main

import (
	"fmt"
)

func main() {
	x := []int{1983, 1986, 1985, 1992, 2019, 2001, 2020, 1934, 1922, 1900}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)
}
