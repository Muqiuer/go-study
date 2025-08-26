package main

import (
	"fmt"
)

func printmyArray(myArray []int) {
	//引用传递，实际传输为指针
	//_表示匿名的变量
	for _, value := range myArray {
		fmt.Println("value =", value)
	}

	myArray[0] = 100
}

func main() {
	myArray := []int{1, 2, 3, 4} //动态数组，切片 slice

	fmt.Printf("myArray types = %T\n", myArray)

	printmyArray(myArray)
	fmt.Println("-----------------")
	for _, value := range myArray {
		fmt.Println("value =", value)
	}
}
