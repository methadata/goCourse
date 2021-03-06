package main

import "fmt"

func main() {

	songs := map[string][]string{
		"bond_james":      []string{"Shaken", "not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being evil", "Ice cream", "Sunsets"},
	}
	fmt.Println(songs)

	songs["yo"] = []string{"La bamba", "Consejo de sabios", "twist & shout"}

	delete(songs, "no_dr")

	for k, v := range songs {
		fmt.Println(k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}
