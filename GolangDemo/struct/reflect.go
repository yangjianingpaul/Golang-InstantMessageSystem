// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// func reflectNum(arg interface{}) {
// 	fmt.Println("type:", reflect.TypeOf(arg))
// 	fmt.Println("value:", reflect.ValueOf(arg))
// }

// type User struct {
// 	Id   int
// 	Name string
// 	Age  int
// }

// func (this User) Call() {
// 	fmt.Println("user is called...")
// 	fmt.Printf("%v\n", this)
// }

// func DoFiledAndMethod(input interface{}) {
// 	// Gets the type of input
// 	inputType := reflect.TypeOf(input)
// 	fmt.Println("inputType is:", inputType)
// 	// Gets the value of input
// 	inputValue := reflect.ValueOf(input)
// 	fmt.Println("inputValue is:", inputValue)

// 	// Get the fields by type
// 	// 1.Obtain the reflect.type() method of interface, obtain NumField by type, and traverse
// 	// 2.Get each field, data type
// 	// 3.The field has an Interface() method to get the corresponding value
// 	// for i := 0; i < inputType.NumField(); i++ {
// 	// 	field := inputType.Field(i)
// 	// 	value := inputValue.Field(i).Interface()

// 	// 	fmt.Printf("%s:%v=%v\n", field.Name, field.Type, value)
// 	// }

// 	// Get the method inside by type, call
// 	for i := 0; i < inputType.NumMethod(); i++ {
// 		m := inputType.Method(i)
// 		fmt.Println("%s:%v\n", m.Name, m.Type)
// 	}
// }

// func main() {
// 	var num float64 = 1.2345
// 	reflectNum(num)

// 	user := User{1, "Aceld", 18}
// 	DoFiledAndMethod(user)
// }
