package main

import "fmt"

func main() {
	year := 1983

	for {
		fmt.Println(year)
		year++
		if year > 2020 {
			break
		}
	}
}
