package main

import (
	"fmt"
	"go/format"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return b
}

// format the bill
func (b bill) format() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0

	// list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%v ...$%v \n", k+":", v)
		total += v
	}

	// totol
	fs += fmt.Sprintf("%v ...$0.2f", "total: ", total)
}
