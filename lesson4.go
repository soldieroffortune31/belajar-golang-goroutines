package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	channel := make(chan string)
// 	defer close(channel)

// 	go GiveMeResponse(channel)
// 	data := <-channel
// 	fmt.Println(data)

// }

// func GiveMeResponse(channel chan string) {
// 	time.Sleep(2 * time.Second)
// 	channel <- "Aliefsufi Uzan Kafil Ardi"
// }
