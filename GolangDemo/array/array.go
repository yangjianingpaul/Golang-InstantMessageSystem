package main

// func printArray(myArray [4]int) {
// 	for index, value := range myArray {
// 		fmt.Println("index=", index, "value=", value)
// 	}
// }

// // 3.切片传递，引用传递
// func printArray(myArray []int) {
// 	for _, v := range myArray {
// 		fmt.Println("value=", v)
// 	}

// 	myArray[0] = 100
// }

func main() {
	// // 1.固定长度数组
	// var myArray1 [10]int
	// for i := 0; i < len(myArray1); i++ {
	// 	fmt.Println(myArray1[i])
	// }
	// fmt.Printf("myArray1 type=%T\n", myArray1)

	//  2.range遍历
	// myArray2 := [10]int{1, 2, 3, 4}
	// for index, value := range myArray2 {
	// 	fmt.Println("index=", index, ",value=", value)
	// }
	// fmt.Printf("myArray2 type=%T\n", myArray2)

	// myArray3 := [4]int{11, 22, 33, 44}
	// fmt.Printf("myArray3 type=%T\n", myArray3)

	// myArray4 := []int{111, 222, 333, 444}
	// printArray(myArray4)

	// // 3.切片传递
	// slice1 := []int{1, 2, 3}
	// fmt.Printf("myArray type is %T\n", slice1)
	// printArray(slice1)
	// fmt.Println("==========")
	// for _, value := range slice1 {
	// 	fmt.Println("value =", value)
	// }

	// 4.slice初始化
	// 1)
	// var slice1 []int
	// slice1 = make([]int, 3)

	// 2)
	// var slice1 []int = make([]int, 3)

	// 3)
	// slice1 := make([]int, 3)

	// fmt.Printf("len=%d,slice=%v\n", len(slice1), slice1)

	// 5.判断slice是否为空
	// if slice1 == nil {
	// 	fmt.Println("slice1是个空切片")
	// } else {
	// 	fmt.Println("slice1有空间")
	// }

	// 6.切片追加，定义切片容量
	// var number = make([]int, 3, 5)
	// fmt.Printf("len=%d,cap=%d,slice=%v\n", len(number), cap(number), number)

	// number = append(number, 2)
	// number = append(number, 3)
	// number = append(number, 1)
	// fmt.Printf("len=%d,cap=%d,slice=%v\n", len(number), cap(number), number)

	// 7.切片的截取
	// slice := []int{1, 2, 3}
	// s1 := slice[0:2]
	// fmt.Println(s1)

	// s1[0] = 100
	// fmt.Println(slice)
	// fmt.Println(s1)

	// // 将底层数组的slice一起进行拷贝
	// s2 := make([]int, 3)

	// // 将s中的值依次拷贝到s2中
	// copy(s2, slice)
	// s2[0] = 200
	// fmt.Println(s2)
	// fmt.Println(slice)
}
