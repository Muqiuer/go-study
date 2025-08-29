package main

import (
	"fmt"
)

type Human struct {
	Name string
	sex  string
}

func (this *Human) Eat() {
	fmt.Println("Human.Eat()...")
}

func (this *Human) Walk() {
	fmt.Println("Human.Walk()...")
}

type SuperMan struct {
	Human // SuperMan类继承了Human类
	Level int
}

// 重定义父类的方法Eat()
func (this *SuperMan) Eat() {
	fmt.Println("SuperMan.Eat()...")
}

// 子类新方法
func (this *SuperMan) Fly() {
	fmt.Println("SuperMan.Fly()...")
}

func (this *SuperMan) Show() {
	fmt.Println("Name=", this.Name)
	fmt.Println("Level=", this.Level)
	fmt.Println("sex=", this.sex)

}

func main() {
	h := Human{
		Name: "张三",
		sex:  "男",
	}
	h.Eat()
	h.Walk()

	// s := SuperMan{
	// 	Human: Human{
	// 		Name: "张三",
	// 		sex:  "男",
	// 	},
	// 	Level: 1,
	// }

	var s SuperMan
	s.Name = "李四" // 直接访问父类的属性
	s.sex = "男"
	s.Level = 1

	s.Eat()  // 子类重定义了父类的方法
	s.Walk() // 子类可以直接访问父类的方法
	s.Fly()  // 子类自己的方法
	s.Show()
}
