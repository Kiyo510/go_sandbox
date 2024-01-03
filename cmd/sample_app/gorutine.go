package main

import (
	"fmt"
	"time"
)

func main() {
	channel1 := make(chan string)

	go func() {
		time.Sleep(5 * time.Second)
		close(channel1)
	}()

	workerCounter := 0
loop:
	for {
		select {
		case <-channel1:
			fmt.Println("channel1を受信")
			break loop
		default:
			workerCounter++
			time.Sleep(1 * time.Second)
		}
	}

	fmt.Println("default:", workerCounter)
}
