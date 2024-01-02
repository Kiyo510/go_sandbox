package main

import (
	"fmt"
	"sync"
)

func main() {
	begin := make(chan int)
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-begin
			fmt.Printf("Goroutine %d\n", i)
		}(i)
	}

	fmt.Println("Unblocking goroutin!")
	close(begin)
	wg.Wait()
}
