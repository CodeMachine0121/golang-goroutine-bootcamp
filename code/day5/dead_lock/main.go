package main

import "fmt"

func main() {

	ch := make(chan int)

	fmt.Println("Sending 42 to channel...")

	ch <- 42 // 程式碼會卡在這邊

	fmt.Println("Send operation finished.")

}
