package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	channel := make(chan string)
// 	defer close(channel)

// 	go OnlyIn(channel)
// 	go OnlyOut(channel)

// 	time.Sleep(5 * time.Second)
// }

// // channel yang digunakan hanya untuk masuk
// func OnlyIn(channel chan<- string) {
// 	time.Sleep(2 * time.Second)
// 	channel <- "Aliefsufi Uzan Kafil Ardi"
// }

// // channel yang digunakan untuk terima data
// func OnlyOut(channel <-chan string) {
// 	data := <-channel
// 	fmt.Println(data)
// }
