package main

import (
	"fmt"
	"strings"
)

func canForm(s string) string {
	words := []string{"dream", "dreamer", "erase", "eraser"}

	for len(s) > 0 {
		matched := false
		for _, word := range words {
			if strings.HasSuffix(s, word) {
				s = s[:len(s)-len(word)]
				matched = true
				break
			}
		}
		if !matched {
			return "NO"
		}
	}
	return "YES"
}

func main() {
	var S string
	fmt.Scan(&S)

	result := canForm(S)
	fmt.Println(result)
}
