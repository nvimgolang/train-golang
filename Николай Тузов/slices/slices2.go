package main

import (
	// "container/list"
	"fmt"
)

func main() {
	// list := make([]int, 5, 10)
	// fmt.Println(len(list), cap(list))

	list := []int{1, 2, 3, 4, 5, 9} // we need to double it
	fmt.Println(double(list))
}

func double(nums []int) []int {
	// var res []int

	// res := make([]int, 0, len(nums))

	res := make([]int, len(nums))

	for i, num := range nums {
		res[i] = num * 2
		// res = append(res, num*2)
	}
	return res
}
