package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

func fundBook(isbn string) error {
	return ErrNotFound

}

func main() {
	err := fundBook("hoge")
	fmt.Println(err)
}
