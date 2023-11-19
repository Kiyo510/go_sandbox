package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	fmt.Scan(&N)

	prevX, prevY, prevT := 0, 0, 0

	for i := 0; i < N; i++ {
		var t, x, y int
		fmt.Scan(&t, &x, &y)

		dist := abs(x-prevX) + abs(y-prevY)
		time := t - prevT

		if dist > time || (dist%2) != (time%2) {
			fmt.Println("No")
			return
		}

		prevX, prevY, prevT = x, y, t
	}

	fmt.Println("Yes")
}

// 絶対値を返すヘルパー関数
func abs(x int) int {
	return int(math.Abs(float64(x)))
}
