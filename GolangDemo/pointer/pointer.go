package main

import "fmt"

func swap(a *int, b *int) {
	c := *a
	*a = *b
	*b = c
}

func main() {
	var a int = 10
	var b int = 20

	swap(&a, &b)

	fmt.Println("a =", a, "b =", b)

	var p *int
	p = &a
	fmt.Println(&a)
	fmt.Println(p)

	var pp **int
	pp = &p
	fmt.Println(&p)
	fmt.Println(pp)
}
