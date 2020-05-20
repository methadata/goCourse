package main

import "fmt"

func main() {

	x1 := []string{"James", "Bond", "Shaken, not stirred"}
	x2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	fmt.Println(x1)
	fmt.Println(x2)

	x3 := [][]string{x1, x2}
	fmt.Println(x3)

	for i, v := range x3 {
		fmt.Println("record: ", i)
		for j, v2 := range v {
			fmt.Printf("\t index position: %v \t value: \t %v \n", j, v2)
		}
	}
}
