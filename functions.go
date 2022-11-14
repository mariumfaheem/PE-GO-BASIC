package main

import (
	"fmt"
	"strings"
)

func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initial []string
	for _, v := range names {
		initial = append(initial, v[:1])
	}
	if len(initial) > 1 {
		return initial[0], initial[1]
	}

	return initial[0], "_"

}
func mutiple_return_values() {
	rec1, rec2 := getInitials("tifa lockhart")
	fmt.Println(rec1, rec2)

}
