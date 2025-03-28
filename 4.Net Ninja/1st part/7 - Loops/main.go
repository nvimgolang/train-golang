package main

import (
	"fmt"
	// "go/format"
)

func main() {
	// x := 0
	// for x < 5 {
	// 	fmt.Println("value of x is:", x)
	// 	x++
	// }

	// for i := 0; i < 5; i++ {
	// 	fmt.Println("value of i is :", i)
	// }

	names := []string{"kinas", "mason", "reaper", "dre"}

	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }

	// for index, value := range names {
	// 	fmt.Printf("the value at index %v is %v", index, value)
	// }

	for _, value := range names {
		fmt.Printf("the value is %v \n", value)
		value = "new string"
	}

	fmt.Println(names)
}
