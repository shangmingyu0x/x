package main

import "fmt"

/*
观察者模式: 建议为发布者类添加订阅机制， 让每个对象都能订阅或取消订阅发布者事件流。
实际上, 该机制包括 1） 一个用于存储订阅者对象引用的列表成员变量； 2） 几个用于添加或删除该列表中订阅者的公有方法。

实际应用中可能会有十几个不同的订阅者类跟踪着同一个发布者类的事件， 不希望【发布者】与【所有这些类】相耦合的。
此外如果他人会使用发布者类， 那么你甚至可能会对其中的一些类一无所知。

因此，所有订阅者都必须实现同样的接口， 发布者仅通过该接口与订阅者交互。
接口中必须声明通知方法及其参数， 这样发布者在发出通知时还能传递一些上下文数据。

如果你的应用中有多个不同类型的发布者， 且希望订阅者可兼容所有发布者， 那么你甚至可以进一步让所有订阅者遵循同样的接口。
该接口仅需描述几个订阅方法即可。 这样订阅者就能在不与具体发布者类耦合的情况下通过接口观察发布者的状态。


观察者模式结构
	1、发布者（Publisher）会向其他对象发送值得关注的事件(主题)。
		事件会在发布者自身状态改变或执行特定行为后发生。
		发布者中包含一个允许新订阅者加入和当前订阅者离开列表的订阅构架。
	2、当新事件发生时，发送者会遍历订阅列表并调用每个订阅者对象的通知方法。该方法是在订阅者接口中声明的。
	3、订阅者（Subscriber）接口声明了通知接口。
		在绝大多数情况下， 该接口仅包含一个 update更新方法。
		该方法可以拥有多个参数， 使发布者能在更新时传递事件的详细信息。
	4、具体订阅者（Concrete Subscribers）可以执行一些操作来回应发布者的通知。
		所有具体订阅者类都实现了同样的接口， 因此发布者不需要与具体类相耦合。
	5、订阅者通常需要一些上下文信息来正确地处理更新。
		因此，发布者通常会将一些上下文数据作为通知方法的参数进行传递。
		发布者也可将自身作为参数进行传递，使订阅者直接获取所需的数据。
	6、客户端（Client）会分别创建发布者和订阅者对象，然后为订阅者注册发布者更新。
*/
func main() {

	shirtItem := newItem("Nike Shirt")

	observerFirst := &customer{id: "abc@gmail.com"}
	observerSecond := &customer{id: "xyz@gmail.com"}

	shirtItem.register(observerFirst)
	shirtItem.register(observerSecond)

	shirtItem.updateAvailability()
}

// 观察者关注的主体（或发布订阅的主题）
type subject interface {
	register(Observer observer)
	deregister(Observer observer)
	notifyAll()
}

// 具体的主体
type item struct {
	observerList []observer
	name         string
	inStock      bool
}

func newItem(name string) *item {
	return &item{
		name: name,
	}
}
func (i *item) updateAvailability() {
	fmt.Printf("Item %s is now in stock\n", i.name)
	i.inStock = true
	i.notifyAll()
}

func (i *item) register(o observer) {
	i.observerList = append(i.observerList, o)
}

func (i *item) deregister(o observer) {
	i.observerList = removeFromslice(i.observerList, o)
}

func (i *item) notifyAll() {
	for _, observer := range i.observerList {
		observer.update(i.name)
	}
}

func removeFromslice(observerList []observer, observerToRemove observer) []observer {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.getID() == observer.getID() {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}

// 观察者(订阅者)
type observer interface {
	update(string)
	getID() string
}

// 具体观察者（订阅这个）
type customer struct {
	id string
}

func (c *customer) update(itemName string) {
	fmt.Printf("Sending email to customer %s for item %s\n", c.id, itemName)
}

func (c *customer) getID() string {
	return c.id
}
