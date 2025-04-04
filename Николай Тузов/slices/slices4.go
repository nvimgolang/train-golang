package main

import (
	// "container/list"
	"fmt"
)

// результат работы функции append нужно присваивать той же переменной
func main() {
	list := make([]int, 4, 4)

	list := append(list, 1)

	list[0] = 5
	list2[0] = 9

	fmt.Println(list)
	fmt.Println(list2)
}
