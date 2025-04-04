package main

import "fmt"

func main() {
	// creating s slice
	oRignAl_slice := []int{90, 60, 40, 50, 34, 49, 30}

	// creating slices from the given slice
	my_slice_1 := oRignAl_slice[1:5]
	my_slice_2 := oRignAl_slice[0:]
	my_slice_3 := oRignAl_slice[:6]
	my_slice_4 := oRignAl_slice[:]
	my_slice_5 := oRignAl_slice[2:4]

	// display the result
	fmt.Println("Original Slice:", oRignAl_slice)
	fmt.Println("My slice 1:", my_slice_1)
	fmt.Println("My slice 2:", my_slice_2)
	fmt.Println("My slice 3", my_slice_3)
	fmt.Println("My slice 4", my_slice_4)
	fmt.Println("My slice 5", my_slice_5)
}
