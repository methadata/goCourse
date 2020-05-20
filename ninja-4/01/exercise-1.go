package main

import (
	"fmt"
)

func main() {
	x := []int{1983, 1986, 1985}
	fmt.Println(x)
	y := []int{1965, 1967}
	x = append(x, y...)
	fmt.Println(x)
}
