package main

import "fmt"

// 装饰是一种结构设计模式， 允许你通过将对象放入特殊封装对象中来为原对象增加新的行为。
func main() {

	pizza := &veggeMania{}

	//Add cheese topping
	pizzaWithCheese := &cheeseTopping{
		pizza: pizza,
	}

	//Add tomato topping
	pizzaWithCheeseAndTomato := &tomatoTopping{
		pizza: pizzaWithCheese,
	}

	fmt.Printf("Price of veggeMania with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.getPrice())
}

// 零件接口
type pizza interface {
	getPrice() int
}

// 具体零件
type veggeMania struct {
}

func (p *veggeMania) getPrice() int {
	return 15
}

// 具体装饰
type tomatoTopping struct {
	pizza pizza
}

func (c *tomatoTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 7
}

// 具体装饰
type cheeseTopping struct {
	pizza pizza
}

func (c *cheeseTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 10
}
