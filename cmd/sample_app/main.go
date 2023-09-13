package main

import (
	"encoding/json"
	"fmt"
)

type Bottle struct {
	Name  string `json:"name"`
	Price int    `json:"price,omitempty"`
	KCal  *int   `json:"kcal,omitempty"`
}

func main() {
	b := Bottle{
		Name:  "Tarou",
		Price: 0,
		KCal:  Int(0),
	}
	out, _ := json.Marshal(b)
	fmt.Println(string(out))
}

func Int(v int) *int {
	return &v
}
