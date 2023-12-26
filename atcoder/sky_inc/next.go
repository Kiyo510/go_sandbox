package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int
	fmt.Scan(&N)

	seen := make(map[int]bool)
	var result []int
	for i := 0; i < N; i++ {
		var A int
		fmt.Scan(&A)

		if !seen[A] {
			seen[A] = true
			result = append(result, A)
		}
	}

	if len(result) == 1 {
		fmt.Println(result[0])
		return
	}

	sort.Slice(result, func(i, j int) bool { return result[i] > result[j] })
	fmt.Println(result[1])
}
