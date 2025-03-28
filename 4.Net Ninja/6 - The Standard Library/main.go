package main

import (
	"fmt"
	"sort"
	// "strings"
)

func main() {
	// greeting := "hello there homies"

	// fmt.Println(strings.Contains(greeting, "hello"))
	// fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	// fmt.Println(strings.ToUpper(greeting))

	// fmt.Println(strings.Index(greeting, "ll"))

	// fmt.Println(strings.Split(greeting, " "))

	// original one
	// fmt.Println("original string value =", greeting)

	ages := []int{45, 50, 55, 60, 65, 70, 75, 80}

	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 75)
	fmt.Println(index)

	names := []string{"sam", "sane", "dante", "pupa"}

	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "pupa"))
}
