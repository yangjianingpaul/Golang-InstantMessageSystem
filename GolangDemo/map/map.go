package main

import "fmt"

func printMap(cityMap map[string]string) {
	for key, value := range cityMap {
		fmt.Println("key=", key)
		fmt.Println("value=", value)
	}
}

func ChangeValue(cityMap map[string]string) {
	cityMap["England"] = "London"
}

func main() {
	// // 1.In the first way, key is a string and value is a string
	// var myMap1 map[string]string
	// if myMap1 == nil {
	// 	fmt.Println("empty map")
	// }

	// myMap1 = make(map[string]string, 10)
	// myMap1["one"] = "java"
	// myMap1["two"] = "c++"
	// myMap1["three"] = "python"

	// fmt.Println(myMap1)

	// 2.The second declaration uses make to allocate data space to a map
	// var myMap2 map[int]string
	// myMap2 := make(map[int]string, 10)
	// myMap2[1] = "java"
	// myMap2[2] = "c++"
	// myMap2[3] = "python"

	// fmt.Println(myMap2)

	// // 3.The third way to declare
	// myMap3 := map[string]string{
	// 	"one":   "php",
	// 	"two":   "c++",
	// 	"three": "python",
	// }
	// fmt.Println(myMap3)

	// 4.CRUD
	// var cityMap map[string]string
	// create
	cityMap := make(map[string]string)
	cityMap["China"] = "Beijing"
	cityMap["Japan"] = "Tokyo"
	cityMap["USA"] = "NewYork"
	// delete
	delete(cityMap, "China")
	cityMap["USA"] = "DC"
	// read
	printMap(cityMap)
	// update
	ChangeValue(cityMap)
}
