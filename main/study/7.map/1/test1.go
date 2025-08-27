package main

import (
	"fmt"
)

func main() {
	// 声明map1是一种map类型，key是int类型，value是string类型

	//在使用map前，需要先用make给map分配数据空间
	var myMap1 = make(map[string]string, 10)

	myMap1["one"] = "java"
	myMap1["two"] = "python"
	myMap1["three"] = "golang"

	fmt.Println(myMap1)

	//第二种声明方式
	myMap2 := make(map[int]string)
	myMap2[1] = "java"
	myMap2[2] = "python"
	myMap2[3] = "golang"
	fmt.Println(myMap2)

	//第三种声明方式
	var myMap3 = map[string]string{
		"one":   "java",
		"two":   "python",
		"three": "golang",
	}
	fmt.Println(myMap3)

}
