// package main

// import (
// 	"fmt"
// 	"runtime"
// 	"time"
// )

// func main() {
// 	// 用go创建一个行参为空，返回值为空的一个函数
// 	go func() {
// 		defer fmt.Println("A.defer")
// 		func() {
// 			defer fmt.Println("B.defer")
// 			runtime.Goexit()
// 			fmt.Println("B")
// 		}()

// 		fmt.Println("A")
// 	}()

// 	// go创建有参函数
// 	go func(a int, b int) bool {
// 		fmt.Println("a =", a, "b =", b)
// 		return true
// 	}(10, 20)

// 	// 死循环
// 	for {
// 		time.Sleep(1 * time.Second)
// 	}
// }
