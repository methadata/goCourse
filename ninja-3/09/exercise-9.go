package main

import (
	"fmt"
	"strings"
)

func main() {
	favSport := "basketbALL"
	switch strings.ToLower(favSport) {
	case "soccer", "football":
		fmt.Println("Play with a ball on our foot")
	case "basketball":
		fmt.Printf("Let's play %v\n", favSport)
	case "golf":
		fmt.Printf("Let's play %v\n", favSport)
	case "tennis":
		fmt.Printf("Let's play %v\n", favSport)
	default:
		fmt.Println("not know sport detected")

	}
}
