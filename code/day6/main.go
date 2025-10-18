package main

import (
	"fmt"
	"time"
)

func producer(tasks chan<- int) {
	defer close(tasks)

	// 5　只有發送方才知道
	for i := 1; i <= 5; i++ {

		fmt.Printf("Producer: Sending task %d\n", i)
		tasks <- i
		time.Sleep(200 * time.Millisecond)
	}
	fmt.Println("Producer: All tasks sent. Channel closed.")
}

func main() {

	tasks := make(chan int, 3)
	go producer(tasks)

	// 使用 for...range 優雅地遍歷 channel
	// 這個迴圈會一直執行，直到 'tasks' channel 被關閉
	for task := range tasks {
		fmt.Printf("Consumer: Received task %d\n", task)
		time.Sleep(500 * time.Millisecond)
	}

	fmt.Println("Consumer: Loop finished. All tasks processed.")

}
