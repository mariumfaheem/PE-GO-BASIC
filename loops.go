package main

import "fmt"

func loop() {

	//Loops

	x := 0
	for x < 5 {
		fmt.Println("value of x is ", x)
		x++
	}

	for i := 0; i <= 5; i++ {
		fmt.Println("value of i is ", i)

	}

	names := []string{"mario", "lungi", "yoshi", "peach"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])

	}

	//With range

	for index, value := range names {
		fmt.Printf("The value of %v is %v \n", index, value)
	}

	for _, value := range names {
		fmt.Printf("The value is %v \n", value)
	}
}
