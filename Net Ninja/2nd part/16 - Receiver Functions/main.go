package main

import (
	"fmt"
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
	fs := "Bill Breakdown:\n"
	total := 0.0

	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", k, v)
		total += v
	}

	fs += fmt.Sprintf("Total: ...$%.2f", total)
	return fs
}

// update tip
func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

// add item to bill
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

func main() {
	myBill := newBill("Mario's bill")

	myBill.addItem("пирог", 5.99)
	myBill.addItem("торт", 3.99)

	fmt.Println(myBill.format())
}
