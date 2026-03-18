package main

// import "fmt"

// func main() {

// 	channel := make(chan int)
// 	defer close(channel)

// 	go SendData(channel, 1, 3)
// 	go SendData(channel, 4, 6)
// 	go SendData(channel, 7, 9)

// 	for i := 0; i < 9; i++ {
// 		fmt.Println(<-channel)
// 	}

// }

// func SendData(channel chan int, x int, n int) {
// 	for i := x; i <= n; i++ {
// 		channel <- i
// 	}
// }
