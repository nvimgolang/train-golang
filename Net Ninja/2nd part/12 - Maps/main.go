package main

import (
	"fmt"
)

func main() {
	menu := map[string]float64{
		"soup":                 4.99,
		"pie":                  7.99,
		"strawberry ice cream": 3.99,
		"salad":                6.99,
	}

	fmt.Println(menu)
	fmt.Println(menu["strawberry ice cream"])

	// looping maps
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	// ints as key type
	phonebook := map[int]string{
		7777777777: "virgil",
		6666666666: "dante",
		3333333333: "nero",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[6666666666])

	phonebook[6666666666] = "devil"
	fmt.Println(phonebook)

	phonebook[7777777777] = "death"
	fmt.Println(phonebook)
}
