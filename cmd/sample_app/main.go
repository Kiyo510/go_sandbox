package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type ip struct {
	Origin string `json:"origin"`
	URL    string `json:"url"`
}

func main() {
	f, err := os.Open("ip.json")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	var resp ip
	if err := json.NewDecoder(f).Decode(&resp); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", resp)
}
