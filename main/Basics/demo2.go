package main

import (
	"fmt"
	"math"
)

func main() {
	// 声明字符串变量a并初始化为"initial"（类型由右侧值自动推断）
	var a = "initial"

	// 同时声明两个int类型变量b和c，并分别初始化为1和2
	var b, c int = 1, 2

	// 声明布尔变量d并初始化为true（类型由右侧值自动推断）
	var d = true

	// 声明float64类型变量e（未显式初始化时默认值为0）
	var e float64

	// 使用短变量声明（:=）将e转换为float32类型后赋值给f
	f := float32(e)

	// 字符串拼接：将a和"foo"拼接后赋值给g
	g := a + "foo"

	// 打印所有变量的当前值（a:字符串, b/c:int, d:布尔, e:float64, f:float32）
	fmt.Println(a, b, c, d, e, f)
	// 打印拼接后的字符串g
	fmt.Println(g)

	// 声明字符串常量s（显式指定类型）
	const s string = "constant"
	// 声明数值常量h（类型由上下文自动推断）
	const h = 500000000
	// 声明计算型常量i（通过3e20除以h得到）
	const i = 3e20 / h
	// 打印常量值及它们的正弦值（math.Sin需要float64类型参数）
	fmt.Println(s, h, i, math.Sin(h), math.Sin(i))
}
