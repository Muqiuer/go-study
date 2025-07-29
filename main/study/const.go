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

func main() {
	fmt.Println(HUNAN)
	fmt.Println(HUBEI)
	fmt.Println(BEIJING)
	fmt.Println(SHANGHAI)
}
