// package main

// import "fmt"

// // // 1.channel定义
// // func main() {
// // 	c := make(chan int)
// // 	go func() {
// // 		defer fmt.Println("goroutine end")

// // 		fmt.Println("goroutine running...")
// // 		c <- 666 //将666发送给c
// // 	}()

// // 	num := <-c //从c中接受数据，并赋值给num

// // 	fmt.Println("num=", num)
// // 	fmt.Println("main goroutine ending...")
// // }

// // // 2.channel已满再写数据就会阻塞
// // func main() {
// // 	c := make(chan int, 3) //带有缓冲的channel
// // 	fmt.Println("len(c)=", len(c), "cap(c)", cap(c))

// // 	go func() {
// // 		defer fmt.Println("subthread end")

// // 		for i := 0; i < 4; i++ {
// // 			c <- i
// // 			fmt.Println("subthread running:发送的元素=", i, "len(c)=", len(c), "cap(c)=", cap(c))
// // 		}
// // 	}()

// // 	time.Sleep(2 * time.Second)

// // 	for i := 0; i < 4; i++ {
// // 		num := <-c
// // 		fmt.Println("num=", num)
// // 	}

// // 	fmt.Println("main end")
// // }

// // // 3.channel关闭
// // func main() {
// // 	c := make(chan int)

// // 	go func() {
// // 		for i := 0; i < 5; i++ {
// // 			c <- i
// // 		}
// // 		close(c) //关闭一个channel
// // 	}()

// // 	// for {
// // 	// 	// ok如果为true表示channel没有关闭，如果为false表示channel已经关闭
// // 	// 	if data, ok := <-c; ok {
// // 	// 		fmt.Println(data)
// // 	// 	} else {
// // 	// 		break
// // 	// 	}
// // 	// }

// // 	// for循环简化
// // 	for data := range c {
// // 		fmt.Println(data)
// // 	}

// // 	fmt.Println("Main Finished...")
// // }

// // 4.select/case处理多路channel
// func fibonacii(c, quit chan int) {
// 	x, y := 1, 1

// 	for {
// 		select {
// 		case c <- x:
// 			//如果c可写，则该case就会进来
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
