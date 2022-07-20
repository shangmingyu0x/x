package main

import (
	"fmt"
	"math"
	"runtime"
)

// 模拟执行业务的 goroutine
func doBusiness(ch chan bool, i int) {
	fmt.Println("go func:", i, "goroutine count:", runtime.NumGoroutine())
	<-ch
}

// https://mp.weixin.qq.com/s/yOGNaNiolWsjRwZx27xNwg
func main() {
	maxCnt := math.MaxInt64
	// max_cnt := 10
	fmt.Println(maxCnt)

	ch := make(chan bool, 3)

	for i := 0; i < maxCnt; i++ {
		ch <- true
		go doBusiness(ch, i)
	}
}
