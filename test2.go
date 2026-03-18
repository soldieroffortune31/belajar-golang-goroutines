package main

// import "fmt"

// func main() {

// 	channel := make(chan string)

// 	datas := []string{"A", "B", "C", "D", "E", "F"}

// 	go func() {
// 		for i := 0; i < 4; i++ {
// 			channel <- datas[i]
// 		}
// 		close(channel)
// 	}()

// 	for data := range channel {
// 		fmt.Println(data)
// 	}

// }
