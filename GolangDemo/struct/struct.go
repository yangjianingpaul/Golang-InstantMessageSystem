// package main

// import "fmt"

// // // 1.Define structure
// // type Book struct {
// // 	title string
// // 	auth  string
// // }

// // func changeBook(book *Book) {
// // 	book.auth = "666666"
// // }

// // func changeBook2(book *Book) {
// // 	book.auth = "777"
// // }

// // // 2.Encapsulation: class name, attribute name, and method name, with the first letter uppercase, indicates that external (other packages) can access it. Otherwise, this package can access it
// // type Hero struct {
// // 	Name  string
// // 	Ad    int
// // 	level int
// // }

// // func (this *Hero) Show() {
// // 	fmt.Println("Name =", this.Name)
// // 	fmt.Println("Ad =", this.Ad)
// // 	fmt.Println("Level =", this.level)
// // }

// // func (this *Hero) GetName() string {
// // 	return this.Name
// // }

// // func (this *Hero) SetName(newName string) {
// // 	this.Name = newName
// // }

// // // 3.inherit
// // type Human struct {
// // 	name string
// // 	sex  string
// // }

// // func (this *Human) Eat() {
// // 	fmt.Println("Human.Eat()...")
// // }

// // func (this *Human) Walk() {
// // 	fmt.Println("Human.Walk()...")
// // }

// // type SuperMan struct {
// // 	Human //Inherits methods from the Human class
// // 	level int
// // }

// // // Redefine the Eat() method
// // func (this *SuperMan) Eat() {
// // 	fmt.Println("SuperMan.Eat()")
// // }

// // func (this *SuperMan) Fly() {
// // 	fmt.Println("SuperMan.Fly()...")
// // }

// // func (this *SuperMan) Print() {
// // 	fmt.Println("name =", this.name)
// // 	fmt.Println("sex =", this.sex)
// // 	fmt.Println("level =", this.level)
// // }

// // 4.Polymorphism, in essence, is a pointer
// type AnimalIF interface {
// 	Sleep()
// 	GetColor() string
// 	GetType() string
// }

// type Cat struct {
// 	color string
// }

// func (this *Cat) Sleep() {
// 	fmt.Println("Cat is Sleep")
// }

// func (this *Cat) GetColor() string {
// 	return this.color
// }

// func (this *Cat) GetType() string {
// 	return "Cat"
// }

// type Dog struct {
// 	color string
// }

// func (this *Dog) Sleep() {
// 	fmt.Println("Dog is Sleep")
// }

// func (this *Dog) GetColor() string {
// 	return this.color
// }

// func (this *Dog) GetType() string {
// 	return "Dog"
// }

// func showAnimal(animal AnimalIF) {
// 	animal.Sleep()
// 	fmt.Println("color=", animal.GetColor())
// 	fmt.Println("kind=", animal.GetType())
// }

// func main() {
// 	// 1.Define structure
// 	// var book1 Book
// 	// book1.title = "Golang"
// 	// book1.auth = "zhang3"
// 	// changeBook(&book1)
// 	// fmt.Printf("%v\n", book1)

// 	// // 3.inherit
// 	// h := Human{"zhang3", "female"}
// 	// h.Eat()
// 	// h.Walk()

// 	// // s := SuperMan{Human{"li4", "male"}, 1}
// 	// var s SuperMan
// 	// s.name = "li"
// 	// s.sex = "female"
// 	// s.level = 88

// 	// s.Eat()
// 	// s.Walk()
// 	// s.Fly()
// 	// s.Print()

// 	// 4.polymorphic
// 	var animal AnimalIF
// 	animal = &Cat{"Green"}
// 	animal.Sleep()
// 	animal = &Dog{"Yellow"}
// 	animal.Sleep()

// 	cat := Cat{"Green"}
// 	dog := Dog{"Yellow"}
// 	showAnimal(&cat)
// 	showAnimal(&dog)
// }
