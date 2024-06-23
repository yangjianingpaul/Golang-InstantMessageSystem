// package main

// import "fmt"

// func myFunc(arg interface{}) {
// 	fmt.Println("myFunc is called...")
// 	fmt.Println(arg)

// 	// 断言机制
// 	value, ok := arg.(string)

// 	if !ok {
// 		fmt.Println("arg is not string type")
// 	} else {
// 		fmt.Println("arg is string type, value=", value)
// 	}
// }

// type Book struct {
// 	auth string
// }

// func main() {

// 	var a string
// 	a = "aceld"

// 	var allType interface{}
// 	allType = a

// 	str, _ := allType.(string)
// 	fmt.Println(str)

// 	book := Book{"Golang"}
// 	myFunc(book)
// 	myFunc(100)
// 	myFunc("abc")
// 	myFunc(3.14)
// }
