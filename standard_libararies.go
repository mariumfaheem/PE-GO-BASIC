package main

import (
	"fmt"
	"sort"
	"strings"
)

func standard_libarary() {
	//Standard libarary
	greetings := "hello there friends!"
	fmt.Println(strings.Contains(greetings, "hello"))
	//it will not change greeting value
	fmt.Println(strings.ReplaceAll(greetings, "hello", "hi"))
	fmt.Println(strings.Index(greetings, "th"))
	fmt.Println(strings.Split(greetings, " "))

	ages := []int{45, 20, 35, 75, 60, 50, 25}

	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 30)
	fmt.Println(index)

}
