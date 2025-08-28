package main

import "fmt"

//声明一种行的数据类型myInt，是int的别名
type myInt int

type Book struct {
	title  string
	author string
}

func changeBook1(book Book) {
	//传递一个book的副本
	book.author = "abc"

}

func changeBook2(book *Book) {
	//传递一个book的指针
	book.author = "abc"
}

func main() {
	// var a myInt = 10
	// fmt.Println("a=", a)
	// fmt.Printf("type of a = %T\n", a)

	var book1 Book
	book1.title = "Golang"
	book1.author = "zhangsan"

	fmt.Printf("%v\n", book1)

	changeBook1(book1)
	fmt.Printf("%v\n", book1)
	changeBook2(&book1)
	fmt.Printf("%v\n", book1)

}
