package main

import (
	"fmt"
)

// interface{}是万能的数据类型
func myFunc(arg interface{}) {
	fmt.Println("myFunc() is called...")
	fmt.Println(arg)

	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string type")
	} else {
		fmt.Println("arg is string type, value=", value)
		fmt.Printf("value type is %T\n", value)
	}
}

type Book struct {
	Author string
}

func main() {
	book := Book{
		Author: "张三",
	}
	myFunc(book)
	myFunc(100)
	myFunc("hello")
}
