package main

import (
	"fmt"
)

// func main() {
// 	for i := 0; i < 10000; i++ {
// 		go DisplayNumber(i)
// 	}

// 	time.Sleep(10 * time.Second)
// }

func DisplayNumber(number int) {
	fmt.Println("Display ", number)
}
