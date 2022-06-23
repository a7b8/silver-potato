package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{4, 1, 4, -4, 6, 3, 8, 8}
	var result = []int{}

	sort.Ints(arr)
	result = append(result, arr[0])

	for i := 0; i < len(arr); i++ {
		if arr[i] != result[len(result)-1] {
			result = append(result, arr[i])
		}
	}
	fmt.Println(result)
}
