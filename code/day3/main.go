package main

import (
	"fmt"
	"sync" // 引入 sync 套件
	"time"
)

// 我們讓 worker 函式接收一個指向 WaitGroup 的指標
func worker(id int, wg *sync.WaitGroup) {
	// defer 關鍵字確保在函式結束時，一定會執行 Done()
	// 這樣無論函式是正常結束還是中途發生 panic，都能確實通知 WaitGroup
	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)

	// 模擬一個耗時的任務
	time.Sleep(time.Second)

	fmt.Printf("Worker %d done\n", id)
}

func main() {
	// 宣告一個 WaitGroup
	var waitGroup sync.WaitGroup

	// 我們要啟動 3 個 worker goroutine
	for i := 1; i <= 3; i++ {
		// 在每次啟動 goroutine 前，計數器 +1
		waitGroup.Add(1)

		// 啟動 goroutine，並將 WaitGroup 的記憶體位址傳入
		go worker(i, &waitGroup)
	}

	fmt.Println("Main: Waiting for workers to finish...")
	// Wait() 會阻塞在這裡，直到計數器歸零
	waitGroup.Wait()

	fmt.Println("Main: All workers have finished. Exiting.")
}
