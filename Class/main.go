package main

import "fmt"

type Hero struct {
	Name  string
	Ad    int
	Level int
}

// 创建Hero的方法
func (this Hero) GetName() {
	fmt.Println("英雄的名字是", this.Name)
}

// 新命名的方法
func (this *Hero) SetName(newName string) {
	this.Name = newName
}

func (this Hero) Eat(Food string) {
	fmt.Println("英雄", this.Name, "正在吃", Food)
}

func test1() {
	hero := Hero{
		Name:  "zhang3",
		Ad:    100,
		Level: 1,
	}

	hero.GetName()
	hero.SetName("li4")
	hero.GetName()
	hero.Eat("包子")
}

// 对于下面的类创建一个方法
type SuperHero struct {
	Hero //这是对英雄类的继承
	//值得注意的是，Golang中的子类无法重写父类的属性

	Ability string
	Ad      int
}

func (this SuperHero) Eat(Food string) {
	fmt.Println("超级英雄", this.Name, "正在吃", Food)

}

func (this SuperHero) AbilityPrint(ability string) {
	fmt.Println("超级英雄", this.Name, "有", ability, "的能力")
}

func test2() {
	SuperHero := SuperHero{
		Hero: Hero{
			Name: "zhang3",
			Ad:   100,
		},
		Ability: "super strength",
		Ad:      1, //我们这个实验可以看到 两个重命名的Ad 程序会默认使用外面（“重写”）的这个
	}

	SuperHero.Eat("包子")
	SuperHero.SetName("li4")
	SuperHero.AbilityPrint(SuperHero.Ability)
	println(SuperHero.Ad)
}

func main() {
	//test1()

	test2()
}
