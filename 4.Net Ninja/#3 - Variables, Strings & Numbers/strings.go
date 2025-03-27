package main

import "fmt"

// var someName := "hello"

func main() {
	var nameOne string = "mario"
	nameTwo := "luigi"
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "peach"
	nameThree = "bowser"

	fmt.Println(nameOne, nameTwo, nameThree)

	nameFour := "yoshi"
	fmt.Println(nameFour)
}
