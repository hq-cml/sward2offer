package _go

import (
	"fmt"
	"sync"
)

// 两个协程交替打印奇偶数
// 双channel版本，实现PV操作
func OddEven1() {
	var wg sync.WaitGroup
	wg.Add(2)
	notifyOdd := make(chan struct{}, 1) // chan需要有1个容量，否则会有问题
	notifyEven := make(chan struct{}, 1)
	var max = 10

	// 打印奇数
	go func() {
		defer wg.Done()
		var i = 1
		for {
			<-notifyOdd
			fmt.Println("Odd: ", i)
			notifyEven <- struct{}{}
			i += 2
			if i > max {
				close(notifyEven) // close通常由写的一方来实现
				break
			}
		}
	}()

	// 偶数
	go func() {
		defer wg.Done()
		var i = 2
		for {
			<-notifyEven
			fmt.Println("Even: ", i)
			notifyOdd <- struct{}{}
			i += 2
			if i > max {
				close(notifyOdd)
				break
			}
		}
	}()

	notifyOdd <- struct{}{}
	wg.Wait()
}
