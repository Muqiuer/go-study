package main

import (
	"fmt"
)

// 如果类名首字母大写，则表示对外部包可见，否则只能内部访问
type Hero struct {
	Name  string
	Ad    int
	Level int
}

func (this *Hero) Show() {
	fmt.Println("Name=", this.Name)
	fmt.Println("Ad=", this.Ad)
	fmt.Println("Level=", this.Level)
}

func (this *Hero) GetName() string {
	fmt.Println("Name=", this.Name)
	return this.Name
}

func (this *Hero) SetName(newName string) {
	this.Name = newName
}

func main() {
	hero := Hero{
		Name:  "张三",
		Ad:    100,
		Level: 1,
	}
	hero.Show()

	hero.SetName("李四")
	hero.Show()

}
