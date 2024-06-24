// package main

// import "fmt"

// // // 1.channel definition
// // func main() {
// // 	c := make(chan int)
// // 	go func() {
// // 		defer fmt.Println("goroutine end")

// // 		fmt.Println("goroutine running...")
// // 		c <- 666 //Send 666 to c
// // 	}()

// // 	num := <-c //Takes data from c and assigns it to num

// // 	fmt.Println("num=", num)
// // 	fmt.Println("main goroutine ending...")
// // }

// // // 2.When the channel is full, write data will be blocked
// // func main() {
// // 	c := make(chan int, 3) //channel with buffering
// // 	fmt.Println("len(c)=", len(c), "cap(c)", cap(c))

// // 	go func() {
// // 		defer fmt.Println("subthread end")

// // 		for i := 0; i < 4; i++ {
// // 			c <- i
// // 			fmt.Println("subthread running:Sent element=", i, "len(c)=", len(c), "cap(c)=", cap(c))
// // 		}
// // 	}()

// // 	time.Sleep(2 * time.Second)

// // 	for i := 0; i < 4; i++ {
// // 		num := <-c
// // 		fmt.Println("num=", num)
// // 	}

// // 	fmt.Println("main end")
// // }

// // // 3.channel off
// // func main() {
// // 	c := make(chan int)

// // 	go func() {
// // 		for i := 0; i < 5; i++ {
// // 			c <- i
// // 		}
// // 		close(c) //Close a channel
// // 	}()

// // 	// for {
// // 	// 	// ok If true indicates that the channel is not closed, if false indicates that the channel is closed
// // 	// 	if data, ok := <-c; ok {
// // 	// 		fmt.Println(data)
// // 	// 	} else {
// // 	// 		break
// // 	// 	}
// // 	// }

// // 	// for loop simplification
// // 	for data := range c {
// // 		fmt.Println(data)
// // 	}

// // 	fmt.Println("Main Finished...")
// // }

// // 4.select/case Processing multiple channels
// func fibonacii(c, quit chan int) {
// 	x, y := 1, 1

// 	for {
// 		select {
// 		case c <- x:
// 			//If c is writable, the case will come in
// 			x = y
// 			y = x + y

// 		case <-quit:
// 			fmt.Println("quit")
// 			return
// 		}
// 	}
// }

// func main() {
// 	c := make(chan int)
// 	quit := make(chan int)

// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			fmt.Println(<-c)
// 		}

// 		quit <- 0
// 	}()

// 	fibonacii(c, quit)
// }
