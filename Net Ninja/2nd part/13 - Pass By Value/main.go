package main

import "fmt"

func updateName(x string) string {
	x = "virgil"
	return x
}

func main() {
	// group A types -> strings, ints, bools, floats, arrays, structs
	name := "dante"

	name = updateName(name)

	fmt.Println(name)

	// group B types -> slices, maps, functions
	menu := map[string]float64{
		"pie":       5.59,
		"ice cream": 4.88,
	}
}
