// package main

// import (
// 	"fmt"
// 	"time"
// )

// // sub goroutine
// func newTask() {
// 	i := 0
// 	for {
// 		i++
// 		fmt.Printf("new Croutine:i=%d\n", i)
// 		time.Sleep(1 * time.Second)
// 	}
// }

// // main goroutine
// func main() {
// 	// Create a go procedure to execute the newTask() process
// 	go newTask()

// 	i := 0
// 	for {
// 		i++
// 		fmt.Printf("main groutine:i=%d\n", i)
// 		time.Sleep(1 * time.Second)
// 	}
// }
