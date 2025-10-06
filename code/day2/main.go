package main

import "fmt"

func sayHello() {
	fmt.Println("Hello from sayHello function")
}

func main() {
	go sayHello()
	fmt.Println("Hello from main function")
}
