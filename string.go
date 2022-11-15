package main

import "fmt"

func stringfunc() {

	//String
	nameFour := "marium"
	age := 22
	name := "marium"

	fmt.Print("you scored %f points! \n", 225.55)
	fmt.Print("you scored %0.1f points! \n", 225.55)
	fmt.Println("hello", nameFour, nameFour)
	fmt.Println("my age is %v and my name is %v", age, name)

	var str = fmt.Sprintf("my age is %v and my name is %v", age, name)
	fmt.Println(str)
}
