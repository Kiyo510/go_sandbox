package main

import (
	"fmt"
)

func fanIn(ch1, ch2 <-chan string) <-chan string {
	newCh := make(chan string)
	go func() {
		for {
			newCh <- <-ch1
		}
	}()
	go func() {
		for {
			newCh <- <-ch2
		}
	}()

	return newCh
}

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
	ch := fanIn(generator("hello"), generator("bye"))
	for i := 0; ; i++ {
		fmt.Println(<-ch)
	}
}
