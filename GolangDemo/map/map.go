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
	// // 1.第一种声明方式，key是string，value也是string
	// var myMap1 map[string]string
	// if myMap1 == nil {
	// 	fmt.Println("empty map")
	// }

	// myMap1 = make(map[string]string, 10)
	// myMap1["one"] = "java"
	// myMap1["two"] = "c++"
	// myMap1["three"] = "python"

	// fmt.Println(myMap1)

	// 2.第二种声明方式，用make给map分配数据空间
	// var myMap2 map[int]string
	// myMap2 := make(map[int]string, 10)
	// myMap2[1] = "java"
	// myMap2[2] = "c++"
	// myMap2[3] = "python"

	// fmt.Println(myMap2)

	// // 3.第三种声明方式
	// myMap3 := map[string]string{
	// 	"one":   "php",
	// 	"two":   "c++",
	// 	"three": "python",
	// }
	// fmt.Println(myMap3)

	// 4.增删改查
	// var cityMap map[string]string
	// 新增
	cityMap := make(map[string]string)
	cityMap["China"] = "Beijing"
	cityMap["Japan"] = "Tokyo"
	cityMap["USA"] = "NewYork"
	// 删除
	delete(cityMap, "China")
	cityMap["USA"] = "DC"
	// 查询
	printMap(cityMap)
	// 修改
	ChangeValue(cityMap)
}
