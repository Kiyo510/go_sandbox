package main

import (
	"bytes"
	"fmt"
	"sync"
)

func main() {
	printData := func(wg *sync.WaitGroup, data []byte) {
		defer wg.Done()
		var buffer bytes.Buffer
		for _, d := range data {
			fmt.Fprintf(&buffer, "%c", d)
		}
		fmt.Println(buffer.String())
	}

	var wg sync.WaitGroup
	wg.Add(2)
	data := []byte("golang")
	printData(&wg, data[:3])
	printData(&wg, data[3:])
	wg.Wait()
}
