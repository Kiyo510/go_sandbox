package main

import "fmt"

//go:generate stringer -type Fruit

type Fruit int

const (
	Apple Fruit = iota
	Orange
	Banana
)

func main() {
	var myFruit Fruit
	myFruit = Banana
	fmt.Println(myFruit)
}
