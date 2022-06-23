package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := "1 2 3 4 5"
	var t = []int32{}
	var result string

	ss := strings.Fields(input)
	for _, v := range ss {
		a, _ := strconv.Atoi(v)
		t = append(t, int32(a))
	}

	if len(t) == 1 {
		result = fmt.Sprint(t[0])
	} else {
		max, min := t[0], t[0]
		for _, val := range t {
			if val > max {
				max = val
			}
			if val < min {
				min = val
			}
		}
		result = fmt.Sprint(max, min)
	}
	fmt.Println(result)
}
