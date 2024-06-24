// package main

// import "fmt"

// // func printArray(myArray [4]int) {
// // 	for index, value := range myArray {
// // 		fmt.Println("index=", index, "value=", value)
// // 	}
// // }

// // 3.Slice passing, reference passing
// func printArray(myArray []int) {
// 	for _, v := range myArray {
// 		fmt.Println("value=", v)
// 	}

// 	myArray[0] = 100
// }

// func main() {
// 	// 1.Fixed length array
// 	var myArray1 [10]int
// 	for i := 0; i < len(myArray1); i++ {
// 		fmt.Println(myArray1[i])
// 	}
// 	fmt.Printf("myArray1 type=%T\n", myArray1)

// 	//  2.range traversal
// 	myArray2 := [10]int{1, 2, 3, 4}
// 	for index, value := range myArray2 {
// 		fmt.Println("index=", index, ",value=", value)
// 	}
// 	fmt.Printf("myArray2 type=%T\n", myArray2)

// 	myArray3 := [4]int{11, 22, 33, 44}
// 	fmt.Printf("myArray3 type=%T\n", myArray3)

// 	myArray4 := []int{111, 222, 333, 444}
// 	printArray(myArray4)

// 	// 3.Slice transfer
// 	slice1 := []int{1, 2, 3}
// 	fmt.Printf("myArray type is %T\n", slice1)
// 	printArray(slice1)
// 	fmt.Println("==========")
// 	for _, value := range slice1 {
// 		fmt.Println("value =", value)
// 	}

// 	// 4.slice initialization
// 	// 1)
// 	var slice1 []int
// 	slice1 = make([]int, 3)

// 	// 2)
// 	var slice1 []int = make([]int, 3)

// 	// 3)
// 	slice1 := make([]int, 3)

// 	fmt.Printf("len=%d,slice=%v\n", len(slice1), slice1)

// 	// 5.Determines whether slice is empty
// 	if slice1 == nil {
// 		fmt.Println("slice1 is an empty slice")
// 	} else {
// 		fmt.Println("slice1 has space")
// 	}

// 	// 6.Slice append, defines slice capacity
// 	var number = make([]int, 3, 5)
// 	fmt.Printf("len=%d,cap=%d,slice=%v\n", len(number), cap(number), number)

// 	number = append(number, 2)
// 	number = append(number, 3)
// 	number = append(number, 1)
// 	fmt.Printf("len=%d,cap=%d,slice=%v\n", len(number), cap(number), number)

// 	// 7.The amputation of slices
// 	slice := []int{1, 2, 3}
// 	s1 := slice[0:2]
// 	fmt.Println(s1)

// 	s1[0] = 100
// 	fmt.Println(slice)
// 	fmt.Println(s1)

// 	// Copy the slices of the underlying array together
// 	s2 := make([]int, 3)

// 	// Copy the values in s to s2 in turn
// 	copy(s2, slice)
// 	s2[0] = 200
// 	fmt.Println(s2)
// 	fmt.Println(slice)
// }
