package main

import "fmt"

func main() {
	// 判断7是否为偶数（7%2计算余数，若余数为0则是偶数）
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd") // 7%2=1，执行else分支输出奇数
	}

	// 判断8是否能被4整除（8%4余数为0表示能整除）
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4") // 条件成立，输出可整除信息
	}

	// 短变量声明num并初始化为9，随后进行多条件判断
	if num := 9; num < 0 { // 第一条件：判断是否为负数
		fmt.Println(num, "is negative")
	} else if num < 10 { // 第二条件：判断是否为1位数（小于10）
		fmt.Println(num, "has 1 digit") // 9<10，执行此分支
	} else {
		fmt.Println(num, "has multiple digits") // 其他情况（≥10）
	}
}
