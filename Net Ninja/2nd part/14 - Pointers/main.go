package main

import (
	"fmt"
)

func updateName(x *string) {
	*x = "virgil"
}

func main() {
	name := "dante"

	// updateName(name)

	// fmt.Println("memory address of name is:", &name)

	m := &name
	// fmt.Println("memory address of name is:", m)
	// fmt.Println("value at memory address:", *m)

	fmt.Println(name)
	updateName(m)
	fmt.Println(name)
	// fmt.Println(name)
}
