package main

import (
	"fmt"
)

func digitSum(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

func main() {
	var N, A, B int
	fmt.Scan(&N, &A, &B)

	total := 0
	for i := 1; i <= N; i++ {
		sum := digitSum(i)
		if sum >= A && sum <= B {
			total += i
		}
	}

	fmt.Println(total)
}
