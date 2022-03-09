package main

import "fmt"

// 工厂方法是一种创建型设计模式，解决了在【不指定具体类】的情况下【创建产品对象的问题】。

// 工厂方法定义了一个方法，且必须使用该方法代替通过直接调用构造函数来创建对象 （ new操作符） 的方式。
// 子类可重写该方法来更改将被创建的对象所属类。
// 由于 Go 中缺少类和继承等 OOP 特性， 所以无法使用 Go 来实现经典的工厂方法模式。
// 不过， 我们仍然能实现模式的基础版本， 即简单工厂。
func main() {
	ak47, _ := getGun("ak47")
	musket, _ := getGun("musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g iGun) {
	fmt.Printf("Gun: %s", g.getName())
	fmt.Println()
	fmt.Printf("Power: %d", g.getPower())
	fmt.Println()
}

// 产品接口
type iGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

// 具体产品
type gun struct {
	name  string
	power int
}

func (g *gun) setName(name string) {
	g.name = name
}

func (g *gun) getName() string {
	return g.name
}

func (g *gun) setPower(power int) {
	g.power = power
}

func (g *gun) getPower() int {
	return g.power
}

// 具体产品
type ak47 struct {
	gun
}

func newAk47() iGun {
	return &ak47{
		gun: gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}

// 具体产品
type musket struct {
	gun
}

func newMusket() iGun {
	return &musket{
		gun: gun{
			name:  "Musket gun",
			power: 1,
		},
	}
}

// 工厂（方法）
func getGun(gunType string) (iGun, error) {
	if gunType == "ak47" {
		return newAk47(), nil
	}
	if gunType == "musket" {
		return newMusket(), nil
	}
	return nil, fmt.Errorf("Wrong gun type passed")
}
