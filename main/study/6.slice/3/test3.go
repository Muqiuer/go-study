package main

import (
	"fmt"
)

func main() {
	//声明slice1是一个切片，并且初始化，默认值是1，2，3，长度len是3
	// slice1 := []int{1, 2, 3}

	// 声明slice1是一个切片，但是给silce分配空间
	var slice1 []int
	// slice1 = make([]int, 3)//开辟3个空间，默认值为0

	//声明slice1是一个切片，同时给它分配了3个空间，初始化值为0，通过:=推导出slice1是一个切片
	// var slcie1 []int = make([]int, 3)

	fmt.Printf("len = %d, slcie = %v\n", len(slice1), slice1)

	//判断一个slice是否为0
	if slice1 == nil {
		fmt.Println("slice1 has no space")
	} else {
		fmt.Println("slice1 has space")
	}
}
