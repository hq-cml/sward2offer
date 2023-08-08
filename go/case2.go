package _go

import (
	"fmt"
	"sync"
)

// 交替打印 A B C
var wg sync.WaitGroup

// 3个协程轮流按顺序打印字符串
func ThreeGo() {
	Achan := make(chan int, 1)
	Bchan := make(chan int)
	Cchan := make(chan int)
	wg.Add(3)
	go A(Achan, Bchan)
	go B(Bchan, Cchan)
	go C(Cchan, Achan)
	//<-Achan
	wg.Wait()
}

func A(achan chan int, bchan chan int) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		achan <- 1
		fmt.Println("A", i)
		<-bchan
	}
	return
}
func B(bchan chan int, cchan chan int) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		bchan <- 1
		fmt.Println("B", i)
		<-cchan
	}
	return
}
func C(cchan chan int, achan chan int) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		cchan <- 1
		fmt.Println("C", i)
		<-achan
	}
	return
}
