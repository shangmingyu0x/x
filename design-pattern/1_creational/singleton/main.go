package main

import (
	"fmt"
	"sync"
)

// 单例是一种创建型设计模式， 让你能够保证一个类只有一个实例，并提供一个访问该实例的全局节点。

// 单例拥有与全局变量相同的优缺点。 尽管它们非常有用， 但却会破坏代码的模块化特性。
// 在某些其他上下文中，你不能使用依赖于单例的类。
// 你也将必须使用单例类。绝大多数情况下， 该限制会在创建单元测试时出现。

// 方式一：通常而言，单例实例会在结构体首次初始化时创建。
// 		 为了实现这一操作， 我们在结构体中定义一个 get­Instance获取实例方法。
// 		 该方法将负责创建和返回单例实例。创建后，每次调用 get­Instance时都会返回相同的单例实例。
// 一些值得注意的地方：
//		1、最开始时会有 nil检查， 确保 single­Instance单例实例在最开始时为空
//		这是为了防止在每次调用 get­Instance方法时都去执行消耗巨大的锁定操作。
// 		如果检查不通过， 则就意味着 single­Instance字段已被填充。
// 		2、single­Instance结构体将在锁定期间创建。
// 		3、在获取到锁后还会有另一个 nil检查。
//		  这是为了确保即便是有多个协程绕过了第一次检查，也只能有一个可以创建单例实例。
//		  否则， 所有协程都会创建自己的单例结构体实例。
// 方法二：
//		1、init函数
//		我们可以在 init函数中创建单例实例。这仅适用于实例的早期初始化工作已经确定时。 ​
//		init函数仅会在包中的每个文件里调用一次， 所以我们可以确定其只会创建一个实例。
//		2、sync.Once
//		sync.Once仅会执行一次操作。 可查看下面的代码：
func main() {

	for i := 0; i < 30; i++ {
		go getInstance()
		go getInstanceOnce(i)
	}

	// Scanln is similar to Scan, but stops scanning at a newline and
	// after the final item there must be a newline or EOF.
	fmt.Scanln()
}

// 方式一
var lock = &sync.Mutex{}

type single struct {
}

var singleInstance *single

func getInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating single instance now.")
			singleInstance = &single{}
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance
}

// 方法二
var once sync.Once

type singleOnce struct {
}

var singleInstanceOnce *singleOnce

func getInstanceOnce(i int) *singleOnce {
	fmt.Println(i)
	if singleInstanceOnce == nil {
		once.Do(
			func() {
				fmt.Println("Creating once single instance now.")
				singleInstanceOnce = &singleOnce{}
			})
	} else {
		fmt.Println("Once single instance already created.")
	}

	return singleInstanceOnce
}
