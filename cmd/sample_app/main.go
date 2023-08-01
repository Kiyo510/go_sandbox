package main

import "fmt"

type Author struct {
	name string
	age  int
}
type Book struct {
	name   string
	author Author
}

func main() {
	b := &Book{
		name: "Hurry Potter",
		author: Author{
			name: "John",
			age:  28,
		},
	}

	fmt.Println(b.author)
	fmt.Println((*b).author)

}
