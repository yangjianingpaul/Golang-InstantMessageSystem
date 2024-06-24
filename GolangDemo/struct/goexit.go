// package main

// import (
// 	"fmt"
// 	"runtime"
// 	"time"
// )

// func main() {
// 	// Create a function with go that takes an empty line argument and returns an empty value
// 	go func() {
// 		defer fmt.Println("A.defer")
// 		func() {
// 			defer fmt.Println("B.defer")
// 			runtime.Goexit()
// 			fmt.Println("B")
// 		}()

// 		fmt.Println("A")
// 	}()

// 	// go creates a function with arguments
// 	go func(a int, b int) bool {
// 		fmt.Println("a =", a, "b =", b)
// 		return true
// 	}(10, 20)

// 	// endless loop
// 	for {
// 		time.Sleep(1 * time.Second)
// 	}
// }
