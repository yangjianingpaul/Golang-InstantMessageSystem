package main

import (
	"GolangDemo/lib1"
	"GolangDemo/lib2"
	// _ "GolangDemo/lib1"	导入不使用
	// mylib2 "GolangDemo/lib2"	取别名
	// . "GolangDemo/lib1"	直接使用包函数
)

// // 2）声明全局变量，方法一、二、三是可以的
// var gA int = 100
// var gB = 200
// 只能够在函数体内声明
// gC:=300

// **********************************************************

// // 3)枚举定义
// const (
// 	//关键字iota，第一行默认值0，之后每行累加1,只能在const()中使用
// 	// BEIJING  = 0
// 	// SHANGHAI = 1
// 	// SHENZHEN = 2

// 	BEIJING = iota
// 	SHANGHAI
// 	SHENZHEN
// )

// const (
// 	a, b = iota + 1, iota + 2
// 	c, d
// 	e, f

// 	g, h = iota * 2, iota * 3
// 	i, k
// )

// ************************************************************

// // 4)函数
// func foo1(a string, b int) int {
// 	fmt.Println("a=", a)
// 	fmt.Println("b=", b)
// 	c := 100
// 	return c
// }

// func foo2(a string, b int) (int, int) {
// 	fmt.Println("a=", a)
// 	fmt.Println("b=", b)
// 	return 666, 777
// }

// // func foo4  (a string, b int) (r1, r2 int)
// func foo3(a string, b int) (r1 int, r2 int) {
// 	fmt.Println("---foo3---")
// 	fmt.Println("a=", a)
// 	fmt.Println("b=", b)

// 	r1 = 1000
// 	r2 = 2000
// 	return
// }

func main() { //函数的'{'一定是和函数名在同一行的，否则编译错误
	// //1)golang中的表达式，加不加“;”都可以，建议不加
	// fmt.Println("hello Go!")
	// //import倒入time包
	// time.Sleep(1 * time.Second)

	// **************************************************

	// //2)声明变量：
	// // 方法一：声明变量 默认值是0
	// var a int
	// fmt.Println("a =", a)

	// // 方法二：变量初始化
	// var b int = 100
	// fmt.Println("b =", b)

	// // 方法三：省去变量类型，不推荐
	// var c = 100
	// fmt.Println("c =", c)
	// fmt.Printf("type of c = %T\n", c) //打印变量类型

	// // 方法四：省去var关键字，自动匹配（常用方法）不支持全局
	// e := 100
	// fmt.Println("e =", e)
	// fmt.Printf("type of e = %T\n", e)

	// f := "abcd"
	// fmt.Println("f =", f)
	// fmt.Printf("type of f = %T\n", f)

	// g := 3.14
	// fmt.Println("g =", g)
	// fmt.Printf("type of g = %T\n", g)

	// // fmt.Println("gA =", gA, "gB =", gB)

	// // 声明多个变量
	// var xx, yy int = 100, 200
	// fmt.Println("xx =", xx, "yy =", yy)
	// var kk, ll = 100, "abcd"
	// fmt.Println("kk =", kk, "ll =", ll)

	// var (
	// 	vv int  = 100
	// 	jj bool = true
	// )

	// fmt.Println("vv =", vv, "jj =", jj)

	// ***************************************************

	// // 3)常量（只读属性）
	// const length int = 10

	// ********************************************************

	// // 4)函数
	// c := foo1("abc", 555)
	// fmt.Println("c=", c)

	// ret1, ret2 := foo2("abc", 999)
	// fmt.Println("ret1=", ret1, "ret2=", ret2)

	// ret1, ret2 = foo3("foo3", 333)
	// fmt.Println("ret1=", ret1, "ret2=", ret2)

	// ********************************************************

	// 5)init函数
	lib1.Lib1Test()
	lib2.Lib2Test()

	// mylib2.Lib2Test()	//取别名
	// Lib2Test()
}
