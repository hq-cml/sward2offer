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
	notifyOdd := make(chan struct{})
	notifyEven := make(chan struct{})
	var max = 100
	go func() {
		defer wg.Done()
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Odd err:", r)
			}
		}()
		var i = 1
		for {
			<-notifyOdd
			fmt.Println("Odd: ", i)
			notifyEven <- struct{}{}
			i += 2
			if i > max {
				close(notifyOdd)
				break
			}
		}
	}()

	go func() {
		defer wg.Done()
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Even err:", r)
			}
		}()
		var i = 2
		for {
			<-notifyEven
			fmt.Println("Even: ", i)
			//time.Sleep(1 * time.Second)
			notifyOdd <- struct{}{}
			i += 2
			if i > max {
				close(notifyEven)
				break
			}
		}
	}()

	notifyOdd <- struct{}{}
	wg.Wait()
}
