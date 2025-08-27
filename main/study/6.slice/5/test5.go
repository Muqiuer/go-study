package main

import "fmt"

func main() {
	s := []int{1, 2, 3} //len=3, cap=3
	s1 := s[0:2]        //[1,2] len=2, cap=3

	fmt.Println(s1)

	s1[0] = 100
	fmt.Println(s)
	fmt.Println(s1) //s跟s1指向的是同一个数组

	//copy 可以将底层数组的slice一起进行拷贝
	s2 := make([]int, 3)
	//将s中的值依次拷贝到s2中
	copy(s2, s)
	fmt.Println(s2)

	s2[1] = 200
	fmt.Println(s)
	fmt.Println(s2)
}
