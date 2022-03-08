package iterator

import "fmt"

/*
迭代器是一种行为设计模式，让你能在不暴露复杂数据结构内部细节的情况下遍历其中所有的元素。
在迭代器的帮助下， 客户端可以用一个迭代器接口以相似的方式遍历不同集合中的元素。

迭代器模式的主要思想是将集合背后的迭代逻辑提取至不同的、名为迭代器的对象中。
此迭代器提供了一种泛型方法，可用于在集合上进行迭代， 而又不受其类型影响。
*/
func main() {

	user1 := &user{
		name: "a",
		age:  30,
	}
	user2 := &user{
		name: "b",
		age:  20,
	}

	userCollection := &userCollection{
		users: []*user{user1, user2},
	}

	iterator := userCollection.createIterator()

	for iterator.hasNext() {
		user := iterator.getNext()
		fmt.Printf("User is %+v\n", user)
	}
}

type user struct {
	name string
	age  int
}

// 集合
type collection interface {
	createIterator() iterator
}

// 具体集合
type userCollection struct {
	users []*user
}

func (u *userCollection) createIterator() iterator {
	return &userIterator{
		users: u.users,
	}
}

// 迭代器
type iterator interface {
	hasNext() bool
	getNext() *user
}

// 具体迭代器
type userIterator struct {
	index int
	users []*user
}

func (u *userIterator) hasNext() bool {
	if u.index < len(u.users) {
		return true
	}
	return false

}
func (u *userIterator) getNext() *user {
	if u.hasNext() {
		user := u.users[u.index]
		u.index++
		return user
	}
	return nil
}
