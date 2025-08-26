package main

import (
	"fmt"
)

func foo1(a string, b int) int {
	fmt.Println("----foo1----")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	c := 100
	return c
}

// 返回多个返回值，匿名的
func foo2(a string, b int) (int, int) {
	fmt.Println("----foo2----")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	return 1, 2
}

// 返回多个返回值，有形参名的
func foo3(a string, b int) (r1 int, r2 int) {
	fmt.Println("----foo3----")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	//r1 r2是属于foo3的形参，初始化默认值为0
	fmt.Println("r1 = ", r1)
	fmt.Println("r2 = ", r2)

	r1 = 100
	r2 = 200
	return
}
func foo4(a string, b int) (r1, r2 int) {
	fmt.Println("----foo4----")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	r1 = 100
	r2 = 200
	return
}

func main() {
	c := foo1("abc", 111)

	fmt.Println("c = ", c)

	ret1, ret2 := foo2("aaa", 222)
	fmt.Println("ret1 = ", ret1)
	fmt.Println("ret2 = ", ret2)

	ret1, ret2 = foo3("bbb", 333)
	fmt.Println("ret1 = ", ret1)
	fmt.Println("ret2 = ", ret2)

	ret1, ret2 = foo4("ccc", 444)
	fmt.Println("ret1 = ", ret1)
	fmt.Println("ret2 = ", ret2)
}
