package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	maxNum := 100 // 定义随机数的最大值（生成范围：0~99）
	// 初始化随机数种子（使用当前时间的纳秒级时间戳确保每次运行随机数不同）
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum) // 生成0~maxNum-1的随机数作为目标数字
	//fmt.Println("The secret number is", secretNumber) // 调试用：打印目标数字（可注释）

	fmt.Println("Please input your guess") // 提示用户输入猜测值
	reader := bufio.NewReader(os.Stdin)    // 创建标准输入读取器（用于接收用户输入）

	// 无限循环：持续接收输入直到猜中目标数字
	for {
		// 读取用户输入（以换行符'\n'为结束符）
		input, err := reader.ReadString('\n')
		if err != nil { // 处理读取错误
			fmt.Println("An error occurred while reading input:", err)
			continue // 跳过当前循环，继续等待新输入
		}
		input = strings.TrimSuffix(input, "\n") // 去除输入末尾的换行符（避免影响后续处理）

		// 将输入字符串转换为整数（验证输入格式）
		guess, err := strconv.Atoi(input)
		if err != nil { // 转换失败（非数字输入）
			fmt.Println("Invalid input. Please enter an integer value")
			continue // 跳过当前循环，提示重新输入
		}

		fmt.Println("Your guess is", guess) // 打印用户输入的猜测值

		// 比较猜测值与目标数字
		if guess > secretNumber {
			fmt.Println("Your guess is bigger than the secret number. Please try again") // 猜测过大提示
		} else if guess < secretNumber {
			fmt.Println("Your guess is smaller than the secret number. Please try again") // 猜测过小提示
		} else {
			fmt.Println("Correct, you Legend!") // 猜中目标数字
			break                               // 退出循环，结束程序
		}
	}
}
