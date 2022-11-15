package main

import "fmt"

type bill1 struct {
	name string
	item map[string]float64
	tip  float64
}

func newBill1(name string) bill1 {
	b := bill1{
		name: name,
		item: map[string]float64{
			"pie":  5.99,
			"cake": 3.99,
		},
		tip: 0,
	}
	return b
}

func (b bill1) fomat1() string {
	fs := "bill Breakdown\n"
	var total float64 = 0

	for k, v := range b.item {
		fs += fmt.Sprintf("%v.....%v", k+":", +v)
		total += v
	}

	fs += fmt.Sprintf("%v ....$%0.2f", "total : ", total)
	return fs

}
func main() {

	b := newBill1("marium")
	fmt.Println(b.fomat1())
}
