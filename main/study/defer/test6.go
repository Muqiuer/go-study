package main

import (
	"fmt"
)

func deferFunc() int {
	fmt.Println("defer func called...")
	return 0
}

func returnFunc() int {
	fmt.Println("return func called...")
	return 0
}

func retrunAndDefer() int {
	defer deferFunc()
	return returnFunc()
}

func main() {
	retrunAndDefer()
}
