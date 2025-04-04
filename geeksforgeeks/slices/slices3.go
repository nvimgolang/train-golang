package main

import "fmt"

func main() {
	// creating an array
	arr := [7]string{"This", "is", "the", "tutorial", "of", "Go", "language"}

	// display array
	fmt.Println("Array:", arr)

	// creating a slice
	myslice := arr[1:6]

	// display a slice
	fmt.Println("Slice:", myslice)

	// display a lenght of the slice
	fmt.Println("Lenght of the slice:", len(myslice))

	// display the capacity of the slice
	fmt.Printf("\nCapacity of the slice: %d", cap(myslice))
}
