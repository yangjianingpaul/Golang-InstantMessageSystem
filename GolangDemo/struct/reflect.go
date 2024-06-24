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
// 	// 获取input的type
// 	inputType := reflect.TypeOf(input)
// 	fmt.Println("inputType is:", inputType)
// 	// 获取input的value
// 	inputValue := reflect.ValueOf(input)
// 	fmt.Println("inputValue is:", inputValue)

// 	// 通过type获取里面的字段
// 	// 1.获取interface的reflect.type()方法，通过type得到NumField，进行遍历
// 	// 2.得到每一个field，数据类型
// 	// 3.通过field有一个Interface()方法得到对应的value
// 	// for i := 0; i < inputType.NumField(); i++ {
// 	// 	field := inputType.Field(i)
// 	// 	value := inputValue.Field(i).Interface()

// 	// 	fmt.Printf("%s:%v=%v\n", field.Name, field.Type, value)
// 	// }

// 	// 通过type获取里面的方法，调用
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
