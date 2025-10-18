package main

import (
	"fmt"
	"time"
)

func longRunningTask(resultChan chan<- string) {
	// 模擬一個需要 3 秒才能完成的任務
	time.Sleep(3 * time.Second)
	resultChan <- "Task finished successfully!"
}

func main() {
	result := make(chan string)
	go longRunningTask(result)

	select {
	case res := <-result:
		fmt.Println(res)
	case <-time.After(2 * time.Second): // 設定一個 2 秒的超時
		fmt.Println("Timeout! The task took too long.")
	}
}
