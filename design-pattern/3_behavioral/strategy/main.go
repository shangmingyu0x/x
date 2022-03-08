package main

import "fmt"

/*
策略是一种行为设计模式， 它将一组行为转换为对象， 并使其在原始上下文对象内部能够相互替换。

原始对象被称为上下文， 它包含指向策略对象的引用并将执行行为的任务分派给策略对象。
为了改变上下文完成其工作的方式，其他对象可以使用另一个对象来替换当前链接的策略对象。
*/
func main() {
	lfu := &lfu{}
	cache := initCache(lfu)

	cache.add("a", "1")
	cache.add("b", "2")

	cache.add("c", "3")

	lru := &lru{}
	cache.setEvictionAlgo(lru)

	cache.add("d", "4")

	fifo := &fifo{}
	cache.setEvictionAlgo(fifo)

	cache.add("e", "5")

}

// 策略接口
type evictionAlgo interface {
	evict(c *cache)
}

// 具体策略
type fifo struct {
}

func (l *fifo) evict(c *cache) {
	fmt.Println("Evicting by fifo strtegy")
}

// 具体策略
type lru struct {
}

func (l *lru) evict(c *cache) {
	fmt.Println("Evicting by lru strtegy")
}

// 具体策略
type lfu struct {
}

func (l *lfu) evict(c *cache) {
	fmt.Println("Evicting by lfu strtegy")
}

//  背景
type cache struct {
	storage      map[string]string
	evictionAlgo evictionAlgo
	capacity     int
	maxCapacity  int
}

func initCache(e evictionAlgo) *cache {
	storage := make(map[string]string)
	return &cache{
		storage:      storage,
		evictionAlgo: e,
		capacity:     0,
		maxCapacity:  2,
	}
}

func (c *cache) setEvictionAlgo(e evictionAlgo) {
	c.evictionAlgo = e
}

func (c *cache) add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.evict()
	}
	c.capacity++
	c.storage[key] = value
}

func (c *cache) get(key string) {
	delete(c.storage, key)
}

func (c *cache) evict() {
	c.evictionAlgo.evict(c)
	c.capacity--
}
