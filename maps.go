package main

import "fmt"

func maps_learning() {

	menu := map[string]float64{
		"soup":          4.99,
		"pie":           7.99,
		"salad":         6.99,
		"toffe pudding": 3.66,
	}
	fmt.Println(menu)
	fmt.Print(menu["soup"])

	//looping maps

	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	//int as key type

	phonebook := map[int]string{
		1234567: "mario",
		2345678: "luigi",
		3456778: "peach",
	}
	fmt.Println(phonebook)

	for k1, v1 := range phonebook {
		println(k1, "-", v1)
	}
}
