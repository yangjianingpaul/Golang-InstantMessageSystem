package main

// import (
// 	"GolangDemo/lib1"
// 	"GolangDemo/lib2"
// 	// _ "GolangDemo/lib1"	Import not used
// 	// mylib2 "GolangDemo/lib2"	aliasing
// 	// . "GolangDemo/lib1"	直接使用包函数
// )

// // 2）To declare global variables, methods 1, 2, and 3 are OK
// var gA int = 100
// var gB = 200
// // It can only be declared in a function body
// gC:=300

// **********************************************************

// // 3)enumeration definition
// const (
// 	//The keyword iota, which defaults to 0 for the first line and increments to 1 for each subsequent line, can only be used in const()
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

// // 4)function
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

func main() { //The '{' of the function must be on the same line as the function name, otherwise it will compile incorrectly
	// //1)The expression in golang, with or without ";" Either is fine, but not recommended
	// fmt.Println("hello Go!")
	// //import Pour the time package
	// time.Sleep(1 * time.Second)

	// **************************************************

	// //2)declare variables:
	// // Method 1: Declare that the default value of the variable is 0
	// var a int
	// fmt.Println("a =", a)

	// // Method 2: Variable initialization
	// var b int = 100
	// fmt.Println("b =", b)

	// // Method 3: Omit the variable type, not recommended
	// var c = 100
	// fmt.Println("c =", c)
	// fmt.Printf("type of c = %T\n", c) //打印变量类型

	// // Method 4: Omit the var keyword, automatic matching (common method) does not support global
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

	// // Declare multiple variables
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

	// // 3)Constant (read-only attribute)
	// const length int = 10

	// ********************************************************

	// // 4)function
	// c := foo1("abc", 555)
	// fmt.Println("c=", c)

	// ret1, ret2 := foo2("abc", 999)
	// fmt.Println("ret1=", ret1, "ret2=", ret2)

	// ret1, ret2 = foo3("foo3", 333)
	// fmt.Println("ret1=", ret1, "ret2=", ret2)

	// ********************************************************

	// 5)init function
	lib1.Lib1Test()
	lib2.Lib2Test()

	// mylib2.Lib2Test()	//aliasing
	// Lib2Test()
}
