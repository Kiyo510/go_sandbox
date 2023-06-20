package main

import (
	"fmt"
	"golang.org/x/exp/maps"
)

func main() {
	persons := map[string]int{"伊藤": 10, "佐藤": 22, "鈴木": 2999}
	fmt.Println(maps.Keys(persons))
}
