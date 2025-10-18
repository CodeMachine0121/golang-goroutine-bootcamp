package main

import (
	"fmt"
	"time"
)

func producer(ch chan<- string, name string, delay time.Duration) {
	for i := 1; ; i++ {
		ch <- fmt.Sprintf("From %s: Message %d", name, i)
		time.Sleep(delay)
	}
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go producer(ch1, "Producer 1", 500*time.Millisecond)
	go producer(ch2, "Producer 2", 1*time.Second)

	// 使用 for + select 來不斷接收來自任一 channel 的消息
	for range 10 { // 為了讓範例能結束，我們只接收10次
		select {
		case msg1 := <-ch1:
			fmt.Println("Received:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Received:", msg2)
		}
	}

	fmt.Println("Main: Finished receiving messages.")
}
