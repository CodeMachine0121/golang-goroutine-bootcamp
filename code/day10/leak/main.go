package main

import (
	"fmt"
	"runtime"
	"time"
)

func leakyWorker() {
	ch := make(chan int)

	// 這個 goroutine 嘗試發送數據，但 main 函式並沒有接收
	// 它將會永遠阻塞在這裡
	go func() {
		fmt.Println("Goroutine: Waiting to send data...")
		ch <- 10
		fmt.Println("Goroutine: Data sent!") // 這行永遠不會執行
	}()

	// 注意：leakyWorker 函式本身會立刻返回
	// 它不會等待內部的 goroutine
}

func main() {
	// 在程式開始時，打印 goroutine 數量
	// 通常是 1 (main goroutine)
	fmt.Printf("Initial Goroutines: %d\n", runtime.NumGoroutine())

	leakyWorker()

	// 給一點時間讓 goroutine 啟動
	time.Sleep(1 * time.Second)

	// 洩漏發生了！
	// 你會看到 goroutine 數量變成了 2
	fmt.Printf("Goroutines after leakyWorker: %d\n", runtime.NumGoroutine())

	// 在這裡可以做其他事情，但那個洩漏的 goroutine 會一直存在
	time.Sleep(5 * time.Second)
	fmt.Printf("Goroutines before exit: %d\n", runtime.NumGoroutine())
}
