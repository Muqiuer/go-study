package main

import "fmt"

type AnimalIF interface {
	Sleep()
	GetColor() string // 获取颜色

	GetType() string // 获取种类
}

type Cat struct {
	color string
}

func (this *Cat) Sleep() {
	fmt.Println("Cat is sleeping")
}

func (this *Cat) GetColor() string {
	return this.color
}

func (this *Cat) GetType() string {
	return "Cat"
}

type Dog struct {
	color string
}

func (this *Dog) Sleep() {
	fmt.Println("Dog is sleeping")
}

func (this *Dog) GetColor() string {
	return this.color
}

func (this *Dog) GetType() string {
	return "Dog"
}

func showAnimal(animal AnimalIF) {
	animal.Sleep()
	fmt.Println("Color=", animal.GetColor())
	fmt.Println("Type=", animal.GetType())
}

func main() {
	// var animal AnimalIF

	// animal = &Cat{
	// 	color: "Yellow",
	// }
	// animal.Sleep()
	// fmt.Println("Color=", animal.GetColor())
	// fmt.Println("Type=", animal.GetType())

	// animal = &Dog{
	// 	color: "Black",
	// }
	// animal.Sleep()
	// fmt.Println("Color=", animal.GetColor())
	// fmt.Println("Type=", animal.GetType())
	cat := &Cat{
		color: "Yellow",
	}
	dog := &Dog{
		color: "Black",
	}
	showAnimal(cat)
	showAnimal(dog)

}
