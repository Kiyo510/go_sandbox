package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var dataReady bool
	cond := sync.NewCond(&sync.Mutex{}) //(1)

	// データが準備されるのを待つゴルーチン
	waitForData := func(i int) {
		defer wg.Done()
		cond.L.Lock()
		for !dataReady {
			fmt.Printf("ゴルーチン%d: データを待っています\n", i)
			cond.Wait() //(2)
		}
		fmt.Printf("ゴルーチン%d: データが準備されました\n", i)
		cond.L.Unlock()
	}

	// 5つのゴルーチンを起動し、データが準備されるのを待つ
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go waitForData(i)
	}

	// データを準備する
	go func() {
		fmt.Println("データの準備")
		time.Sleep(3 * time.Second) // データ準備に2秒かかるとします
		cond.L.Lock()
		dataReady = true
		cond.Broadcast() // (3)すべての待機中のゴルーチンに通知
		cond.L.Unlock()
	}()

	wg.Wait()
}
