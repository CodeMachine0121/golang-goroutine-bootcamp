package main

import (
	"fmt"
	"time"
)

func main() {

	tasks := make(chan int, 3)

	// 啟動 producer goroutine
	// 當我們把雙向的 tasks channel 傳遞給 producer 時，
	// 它被隱式轉換成了 chan<- int 型別
	go producer(tasks)

	// 啟動 consumer goroutine
	// 同樣地，這裡它被轉換成了 <-chan int 型別
	go consumer(tasks)

	// 給 goroutines 一些時間執行
	// 注意：在真實應用中我們應該使用 WaitGroup
	time.Sleep(3 * time.Second)
	fmt.Println("Main: Done.")
}

// consumer 函式現在明確表示，它只會從 'tasks' channel 接收資料 (<-chan)
func consumer(tasks <-chan int) {
	// 使用 for...range 優雅地遍歷 channel
	for task := range tasks {
		fmt.Printf("Consumer: Received task %d\n", task)
	}
	fmt.Println("Consumer: Loop finished.")
}

// producer 函式現在明確表示，它只會向 'tasks' channel 發送資料 (chan<-)
func producer(tasks chan<- int) {
	defer close(tasks)
	for i := 1; i <= 5; i++ {
		fmt.Printf("Producer: Sending task %d\n", i)
		tasks <- i
		time.Sleep(200 * time.Millisecond)
	}
	fmt.Println("Producer: All tasks sent.")
}
