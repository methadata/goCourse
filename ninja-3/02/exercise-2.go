package main

import "fmt"

func main() {
	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for j := 1; j < 4; j++ {
			fmt.Printf("\t%U\t%c\n", i, i)
		}
	}
}
