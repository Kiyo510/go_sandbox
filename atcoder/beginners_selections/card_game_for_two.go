package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int
	fmt.Scan(&N)

	arr := make([]int, N)
	for i := range arr {
		var A int
		fmt.Scan(&A)
		arr[i] = A
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})

	var alice, bob int
	for i, a := range arr {
		if i%2 == 0 {
			alice += a
		} else {
			bob += a
		}
	}

	fmt.Println(alice - bob)
}
