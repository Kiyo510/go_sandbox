package main

import (
	"fmt"
	"math/rand"
)

func fanIn(ch1, ch2 <-chan string) <-chan string {
	newCh := make(chan string)
	go func() {
		for {
			select {
			case s := <-ch1:
				newCh <- s
			case s := <-ch2:
				newCh <- s
			default:
				//update_progressbar()
			}
		}
	}()

	return newCh
}

func main() {
	quit := make(chan string)
	ch := generator("Hi!", quit)
	for i := rand.Intn(50); i >= 0; i-- {
		fmt.Println(<-ch, i)
	}

	quit <- "Bye"
	fmt.Printf("%s", <-quit)
}

func generator(msg string, quit chan string) <-chan string {
	ch := make(chan string)

	go func() {
		for i := 0; ; i++ {
			select {
			case ch <- fmt.Sprintf("%s", msg):
			case <-quit:
				quit <- "See you!"
				return
			}
		}
	}()
	return ch
}
