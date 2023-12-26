package main

import (
	"fmt"
	"strings"
)

func main() {
	var S string
	fmt.Scan(&S)

	var result strings.Builder
	for i, v := range S {
		if i > 0 {
			result.WriteRune(' ')
		}
		result.WriteRune(v)
	}

	fmt.Println(result.String())
}
