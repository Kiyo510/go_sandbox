package main

import "fmt"

func main() {
	var n int
	var s []byte
	fmt.Scan(&n, &s)

	m := make(map[byte]int, 30)
	var t int

	t++
	m[s[0]] = t
	for i := 1; i < n; i++ {
		if s[i] == s[i-1] {
			t++
		} else {
			t = 1
		}
		if t > m[s[i]] {
			m[s[i]] = t
		}
	}

	var ans int
	for _, v := range m {
		ans += v
	}

	fmt.Println(ans)
}
