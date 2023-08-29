package main

import (
	"encoding/json"
	"fmt"
)

type ip struct {
	Origin string `json:"origin"`
	URL    string `json:"url"`
}

func main() {
	s := `{"origin": "255.111.111.111", "url": "https://httpbin.org/get"}`
	var resp ip
	if err := json.Unmarshal([]byte(s), &resp); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", resp)
}
