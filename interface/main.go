package main

import "fmt"

// 接口的本质是一个指针
type Animal interface {
	Yell() //叫声
	//GetColor() string
	//GetName() string //获得动物种类
	Eat()
	SizePrint()
}

//对于一个接口而言 没有所谓的继承 而是实现了这个接口所有方法的类 就可以使用这个接口

// 下面 我创建一个案例 来实现三种动物的叫声 进食和体型
type Dog struct {
	Color string
	Size  string
	Sound string
	Food  string
}
type Cat struct {
	Food  string
	Size  string
	Sound string
	Like  string
}
type Pig struct {
	Food    string
	Size    string
	Sound   string
	Ability string
}

// 接下来 我们来实现猫的方法
func (c Cat) Eat() {
	fmt.Println("猫猫爱吃", c.Food)
}
func (c Cat) SizePrint() {
	fmt.Println("猫猫的体型是", c.Size)
}
func (c Cat) Yell() {
	fmt.Println(c.Sound)
}

// 接下来 我们来实现狗的方法
func (d *Dog) Yell() {
	println(d.Sound)
}
func (d *Dog) Eat() {
	println("狗狗吃", d.Food)
}
func (d *Dog) SizePrint() {
	fmt.Println("狗狗的体型是", d.Size)
}

//func (d *Dog) GetColor(color string) {
//	d.Color = color
//}
//
//func (d *Dog) GetName(name string) {
//	d.Name = name
//}

// 接下来 我们来实现猪的方法
func (p *Pig) Yell() {
	println(p.Sound)
}
func (p *Pig) Eat() {
	println("猪猪吃", p.Food)
}
func (p *Pig) SizePrint() {
	fmt.Println("狗狗的体型是", p.Size)
}

func test1() {
	dog1 := Dog{
		Color: "red",
		Size:  "big",
		Sound: "wangwang",
		Food:  "bone",
	} //这里 我们创建的是一个狗的实例 现在我们尝试使用接口类去调用方法

	var animal Animal
	animal = &dog1 //我们可以讲这个借口指针指向dog1这个实例
	animal.Yell()
	animal.Eat()
	animal.SizePrint()

	animal = &Cat{
		Food:  "fish",
		Size:  "small",
		Sound: "miao",
		Like:  "mouse",
	}
	animal.Yell()
	animal.Eat()
	animal.SizePrint()
	//通过上面的两个案例，我们可以看出，接口指向一个已有实例或者指向一个创建的实例

}

// 下面 我们实现一下多态的案例
type Animal2 interface {
	Yell()
	Eat()
	SizePrint()
	GetSize() string
}

func (d Dog) GetSize() string {
	return d.Size
}
func (c Cat) GetSize() string {
	return c.Size
}
func (p Pig) GetSize() string {
	return p.Size
}

func ShowAnimal(animal Animal2) {
	animal.Yell()
	animal.Eat()
	//animal.SizePrint()
	fmt.Println("输入的动物体型为", animal.GetSize())
}

func test2() {
	dog2 := Dog{Size: "small", Food: "bone", Sound: "wangwang"}
	cat2 := Cat{Size: "big", Food: "fish", Sound: "miao", Like: "mouse"}

	//这就实现了一个多态的方法 即同一个方法可以被不同的类调用 （只要这些类属于一个接口）
	ShowAnimal(&dog2)
	ShowAnimal(&cat2)
}

//接下来 我想说明一点 就是一个类可以属于多个接口 我们拿智能手机来举例子
// 在智能手机中 我们可以实现“电话的功能”（接听 拨打 发信息） 也可以实现“相机的功能”（拍照 录像） 还可以实现视频播放器的功能（看电视剧）

type Phone interface {
	//这个接口有智能手机和移动电话两个类
	Listen()
	Call()
	Message()
}

type Camera interface {
	//这个接口有智能手机和相机两个类
	TakePhoto()
	RecordVideo()
}

type VideoPlayer interface {
	//这个接口有智能手机和电视两个类
	PlayVideo()
}

type MobilePhone struct {
	Sound string
	Tel   string
	Text  string
	Name  string
}
type camera struct {
	Take   string
	Record string
	Name   string
}
type TV struct {
	Play string
	Name string
}

type Smartphone struct {
	MobilePhone
	camera
	TV
	Name string
}

// 下面三个实现的是一个手机类接口的方法
func (m MobilePhone) Listen() {
	println(m.Name, "发出", m.Sound, "的声响")
}
func (m MobilePhone) Call() {
	println(m.Name, m.Tel, "可以拨打电话")
}
func (m MobilePhone) Message() {
	println(m.Name, "可以发送信息，内容为", m.Text)
}

// 下面三个实现的是一个相机类接口的方法
func (c camera) TakePhoto() {
	println(c.Name, "可以", c.Take)
}
func (c camera) RecordVideo() {
	println(c.Name, "可以", c.Record)
}

// 下面三个实现的是一个电视类接口的方法
func (t TV) PlayVideo() {
	println(t.Name, "可以", t.Play)
}

// 智能手机类方案
type SmartPhone interface {
	Listen()
	Call()
	Message()
	TakePhoto()
	RecordVideo()
	PlayVideo()
}

func (smartphone Smartphone) Listen() {
	println(smartphone.Name, "发出", smartphone.Sound, "的声响")
}
func (smartphone Smartphone) Call() {
	println(smartphone.Name, smartphone.Tel, "可以拨打电话")
}

// 下面是三个多态函数
func ShowPhone(phone Phone) {
	phone.Listen()
	phone.Call()
	phone.Message()
} //由于传入参数的限制 这个函数可以被所有手机类调用
func ShowCamera(camera Camera) {
	camera.TakePhoto()
	camera.RecordVideo()
}
func ShowTV(TV VideoPlayer) {
	TV.PlayVideo()
}

func test3() {
	smartphone := Smartphone{
		MobilePhone: MobilePhone{Sound: "beep", Tel: "123456", Text: "hello", Name: "智能手机"},
		camera:      camera{Take: "take", Record: "record", Name: "智能手机"},
		TV:          TV{Play: "play", Name: "智能手机"},
		Name:        "智能手机外部",
	} //创建一个智能手机实例

	ShowPhone(&smartphone)
	ShowCamera(&smartphone)
	ShowTV(&smartphone)
}

func main() {
	//test1()
	//test2()
	test3()
}
