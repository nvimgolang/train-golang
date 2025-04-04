package main

import "fmt"

func main() {
	list := []int{1, 2, 3, 4}
	fmt.Println("before", list)
	handle(list[:1])
	fmt.Println("after", list)
}

func handle(list []int) {
	// list[1] = 123

	newList := make([]int, len(list))
	copy(newList, list)

	newList = append(newList, 5)
	fmt.Println("append", newList)
}
