package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	channel := make(chan string)
// 	defer close(channel)

// 	go func() {
// 		time.Sleep(2 * time.Second)
// 		channel <- "Aliefsufi Uzan"
// 		channel <- "Kafil Ardi"
// 		fmt.Println("selesai mengirim data ke channel")
// 	}()

// 	data := <-channel
// 	data2 := <-channel
// 	fmt.Println(data)
// 	fmt.Println(data2)

// 	time.Sleep(5 * time.Second)

// 	/* insert data ke channel*/
// 	// channel <- "Alief"

// 	/* cara 1 ambil data dari channel */
// 	// name := <-channel
// 	// fmt.Println(name)

// 	/* cara 2 ambil data dari channel */
// 	// var data string

// 	// data = <-channel

// 	/* close channel */
// }
