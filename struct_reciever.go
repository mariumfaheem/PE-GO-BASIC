package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//blueprint for the type of data
type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

//we can create functions or method that are associated with bill

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{"pie": 5.99, "cake": 3.99},
		tip:   0,
	}
	return b

}

//format the bill
func (b bill) format() string {
	fs := "Bill Breakdown: \n"
	var total float64 = 0

	for k, v := range b.items {
		fs += fmt.Sprintf("%v ...$%v", k+":", v)
		total += v
	}
	//total
	fs += fmt.Sprintf("%v ... $%0.2f", "total:", total)
	return fs

}

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Printf(prompt)
	input, error := r.ReadString('\n')
	return input, error
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	name, _ := getInput("create a new bill name: ", reader)
	//it will trimp all spaces
	name = strings.TrimSpace(name)
	b := newBill(name)
	fmt.Println("created the bill - ", b.name)
	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose Options (a -add items, s - save the bill , t - add tips): ", reader)

	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)
		fmt.Println(name, price)
	case "t":
		tip, _ := getInput("Enter tip amount ($): ", reader)
		fmt.Println(tip)
	case "s":
		fmt.Println("you chose t")
	default:
		fmt.Println("Not a valid option")
		promptOptions(b)

	}

}
func struct_and_reciever() {
	bill := newBill("mario bill")
	fmt.Println(bill.format())

	mybill := createBill()
	fmt.Println(mybill)
	promptOptions(mybill)

}
