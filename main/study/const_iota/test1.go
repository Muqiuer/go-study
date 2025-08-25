package main

import (
	"fmt"
)

const (
	HUNAN = iota //iota=0,并在后面的常量累加
	HUBEI
	BEIJING
	SHANGHAI
)

const (
	a, b = iota + 1, iota + 2 //iota=0,a=1,b=2
	c, d                      //iota=1,c=iota+1=2,d=iota+2=3
	e, f                      //iota=2,e=iota+1=3,f=iota+2=4

	g, h = iota * 2, iota * 3 //iota=3,g=iota*2=6,h=iota*3=9
	i, j                      //iota=4,i=iota*2=8,j=iota*3=12
)

func main() {
	fmt.Println("HUNAN = ", HUNAN)
	fmt.Println("HUBEI = ", HUBEI)
	fmt.Println("BEIJING = ", BEIJING)
	fmt.Println("SHANGHAI = ", SHANGHAI)

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Println("c = ", c)
	fmt.Println("d = ", d)
	fmt.Println("e = ", e)
	fmt.Println("f = ", f)
	fmt.Println("g = ", g)
	fmt.Println("h = ", h)
	fmt.Println("i = ", i)
	fmt.Println("j = ", j)
}
