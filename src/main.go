package main

import (
	"fmt"
)

func generator(msg string) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("%s %d", msg, i)
		}
	}()
	return ch
}

func main() {
	ch := generator("Hello")
	for i := 0; ; i++ {
		fmt.Println(<-ch)
	}
}
