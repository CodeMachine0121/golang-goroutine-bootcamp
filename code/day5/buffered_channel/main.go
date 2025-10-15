package main

import "fmt"

func main() {

	reportBox := make(chan string, 3)

	reportBox <- "Report 1"
	fmt.Println("Sent Report 1.")
	reportBox <- "Report 2"
	fmt.Println("Sent Report 2.")
	reportBox <- "Report 3"
	fmt.Println("Sent Report 3.")

	// 此時文件匣已滿。如果再嘗試發送，程式就會阻塞
	// reportBox <- "Report 4" // 取消這行的註解會導致 deadlock

	// 經理現在開始讀取報告
	fmt.Println("Manager is reading:", <-reportBox)
	fmt.Println("Manager is reading:", <-reportBox)
	fmt.Println("Manager is reading:", <-reportBox)

}
