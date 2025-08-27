package main

import (
	"fmt"
)

func printMap(cityMap map[string]string) {
	//这里的cityMap是引用传递
	for key, value := range cityMap {
		fmt.Println("key=", key, "value=", value)
	}
}

func changeValue(cityMap map[string]string) {
	cityMap["no4"] = "武汉"
}

func main() {
	cityMap := make(map[string]string)
	//添加
	cityMap["no1"] = "北京"
	cityMap["no2"] = "上海"
	cityMap["no3"] = "深圳"
	//遍历
	printMap(cityMap)
	//删除
	delete(cityMap, "no2")
	//修改
	cityMap["no1"] = "长沙"
	changeValue(cityMap)

	fmt.Println("--------------------")

	//遍历
	printMap(cityMap)

}
