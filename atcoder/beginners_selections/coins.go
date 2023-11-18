package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	minCount := int(^uint(0) >> 1)

	for i := 0; i < N; i++ {
		var A int
		fmt.Scan(&A)

		count := 0
		for A%2 == 0 {
			count++
			A >>= 1
		}

		if count < minCount {
			minCount = count
		}
	}

	fmt.Println(minCount)
}
