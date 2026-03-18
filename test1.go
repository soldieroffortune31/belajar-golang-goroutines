package main

import "fmt"

func main() {

	channel := make(chan int)
	// defer close(channel)

	go SetDataChannel(channel)

	for i := 1; i <= 5; i++ {
		fmt.Println(<-channel)
	}
}

func SetDataChannel(channel chan int) {
	for i := 1; i <= 5; i++ {
		channel <- i
	}
	close(channel)
}
