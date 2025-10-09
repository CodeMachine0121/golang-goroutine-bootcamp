package main

import (
	"fmt"
	"time"
)

func calcuate(resultChan chan int) {
	fmt.Println("Worker: Starting calculation...")
	time.Sleep(2 * time.Second)
	result := 42
	fmt.Println("Worker: Calculation finished. Sending result.")

	// 將計算結果發送到 channel
	resultChan <- result
}

func main() {

	// 建立一個 channel　用來接收結果
	resultChannel := make(chan int)

	// 啟動 worker goroutine, 並把 channel　傳給他
	go calcuate(resultChannel)

	fmt.Println("Main: Waiting for result...")

	// 從 channel 接收結果　這一行會被阻塞
	finalResult := <-resultChannel
	fmt.Printf("Main: Received result: %d\n", finalResult)
}
