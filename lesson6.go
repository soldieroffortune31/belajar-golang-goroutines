package main

// import (
// 	"fmt"
// 	"time"
// )

/*
*
1. ketika channel tidak ada yang mengisi datanya maka error
2. ketika channel tidak ada yang menerima datanya juga error
3. tapi ketika channel di beri buffer, maka gak wajib ada pemerima data
*/
// func main() {

// 	channel := make(chan string, 3)
// 	defer close(channel)

// 	go func() {
// 		channel <- "Alief"
// 		channel <- "Uzan"
// 		channel <- "Ardi"
// 	}()

// 	go func() {
// 		fmt.Println(<-channel)
// 		fmt.Println(<-channel)
// 		fmt.Println(<-channel)
// 	}()

// 	// fmt.Println(cap(channel)) // melihat panjang buffer
// 	// fmt.Println(len(channel)) // melihat panjang data di buffer
// 	time.Sleep(2 * time.Second)
// 	fmt.Println("Selesai")
// }
