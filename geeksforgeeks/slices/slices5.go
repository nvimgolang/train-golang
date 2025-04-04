package main

import "fmt"

func main() {
	// creating an arr
	arr := [4]string{"Geeks", "for", "Geeks", "GFG"}

	// creating slices from the given arr
	my_slice_1 := arr[1:2]
	my_slice_2 := arr[0:]
	my_slice_3 := arr[:2]
	my_slice_4 := arr[:]

	// display the result
	fmt.Println("My array:", arr)
	fmt.Println("My slice 1:", my_slice_1)
	fmt.Println("My slice 2:", my_slice_2)
	fmt.Println("My slice 3:", my_slice_3)
	fmt.Println("My slice 4:", my_slice_4)
}
