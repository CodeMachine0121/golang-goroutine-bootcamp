package main

import (
	"fmt"
	"runtime"
	"time"
)

// 這個函式完全不讀取 channel
func doNothingWithChannel(messages <-chan string) {
	// 假裝很忙
	time.Sleep(1 * time.Second)
	fmt.Println("doNothingWithChannel is done, but it read nothing.")
}

func main() {
	fmt.Printf("Initial Goroutines: %d\n", runtime.NumGoroutine())

	// 建立一個容量為 1 的 channel
	messages := make(chan string, 1)

	// 啟動一個 goroutine，它會發送兩條消息
	go func() {
		messages <- "Hello"
		fmt.Println("Goroutine: Sent 'Hello'")

		// 在發送 "World" 時，緩衝區已滿，且 processFirstMessage 已退出
		// 所以這裡會永久阻塞
		messages <- "World"
		fmt.Println("Goroutine: Sent 'World'") // 這行永遠不會執行
	}()

	time.Sleep(100 * time.Millisecond)
	doNothingWithChannel(messages)

	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutines before exit: %d\n", runtime.NumGoroutine())
}
