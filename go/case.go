package _go

import (
	"fmt"
	"sync"
)

const max = 100

// 两个协程交替打印奇偶数
// 双channel版本，实现PV操作
func OddEven1() {
	var wg sync.WaitGroup
	wg.Add(2)
	notifyOdd := make(chan struct{}, 1) // chan需要有1个容量，否则会有问题
	notifyEven := make(chan struct{}, 1)

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

// 两个协程交替打印奇偶数
// 单channel版本，实现PV操作
// 这个方法不好！！不清晰
// 因为首先因为变量ch是一个无缓冲的channel, 所以只有读写同时就绪时才不会阻塞。所以两个goroutine会同时进入各自的 if 语句（此时 i 是相同的）
// 但是此时只能有一个 if 是成立的，不管哪个goroutine快，都会由于读channel或写channel导致阻塞，因此程序会交替打印1-100且有顺序。
func OddEven2() {
	var wg sync.WaitGroup
	wg.Add(2)
	signal := make(chan struct{})
	// Odd
	go func() {
		defer wg.Done()
		for i := 1; i < 101; i++ {
			signal <- struct{}{}
			if i % 2 == 1 {
				fmt.Println("Odd: ", i)
			}
		}
	}()
	// Even
	go func() {
		defer wg.Done()
		for i := 1; i < 101; i++ {
			<-signal
			if i % 2 == 0 {
				fmt.Println("Even: ", i)
			}
		}
	}()
	wg.Wait()
}