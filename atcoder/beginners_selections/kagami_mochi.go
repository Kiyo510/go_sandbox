package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	m := make(map[int]bool)
	for i := 0; i < N; i++ {
		var d int
		fmt.Scan(&d)
		m[d] = true
	}

	fmt.Println(len(m))
}
