package main

import "fmt"

func array_slices() {
	//Array and slices
	var ages [3]int = [3]int{22, 23, 24}
	names := [4]string{"yoshi", "mario", "peach", "bowser"}

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	var scores = []int{100, 50, 60}
	scores[2] = 25
	scores = append(scores, 86)
	fmt.Println(scores, len(scores))

	rangeOne := names[0:3]
	fmt.Println(rangeOne)

}
