package main

import (
	"fmt"
	"time"
)

func main() {

	ticker := time.NewTicker(500 * time.Millisecond)
	// 建立一個每 500 毫秒觸發一次的 Ticker
	// 建立一個在 3 秒後觸發的 channel，用來停止 ticker
	stopper := time.After(3 * time.Second)

	// 使用 defer 來確保 ticker 在 main 函式結束時被停止
	// 這非常重要，可以防止 goroutine 洩漏
	defer ticker.Stop()

	fmt.Println("Ticker started. Will stop after 3 seconds.")

	for {
		select {
		case t := <-ticker.C:
			// 每次 ticker 觸發，就會執行這裡
			fmt.Println("Tick at", t.Format("15:04:05.000"))
		case <-stopper:
			// 3 秒時間到，停止訊號來了
			fmt.Println("Ticker stopped.")
			return // 結束函式
		}
	}

}

func threeSecondTimer() {

	timer := time.NewTimer(3 * time.Second)
	someOtherChan := make(chan int, 1)

	select {
	case <-someOtherChan:
		// 另一個 channel 先到了
		// 我們不再需要這個 timer 了，最好停掉它
		if !timer.Stop() {
			<-timer.C // 如果 Stop 返回 false，說明 timer 可能已經觸發了，需要手動排空 channel
		}
	case <-timer.C:
		// timer 觸發了
		fmt.Println("Timer fired!")
	}
}
