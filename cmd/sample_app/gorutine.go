package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	var lock sync.Mutex

	increment := func() {
		lock.Lock() // (1)
		defer lock.Unlock() // (2)
		count++
		fmt.Printf("increment: %d\n", count)
	}

	decrement := func() {
		lock.Lock() // (1)
		defer lock.Unlock() // (2)
		count--
		fmt.Printf("decrement %d\n", count)
	}

	// Increment
	var arithmetic sync.WaitGroup
	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			increment()
		}()
	}

	// Decrement
	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			decrement()
		}()
	}

	arithmetic.Wait()

	fmt.Println("arithmetic complete")
}
