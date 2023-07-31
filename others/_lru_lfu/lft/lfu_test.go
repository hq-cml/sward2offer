package lft

import (
	"log"
	"testing"
)

func TestLFUCache_A(t *testing.T) {
	cache := NewLFU(3)
	cache.Put(1,10)
	cache.Put(2,20)
	cache.Put(3,30)

	log.Printf("key = 1;value = %d",cache.Get(1))
	log.Printf("key = 2;value = %d",cache.Get(2))
	log.Printf("key = 3;value = %d",cache.Get(3))
	log.Printf("key = 3;value = %d",cache.Get(3))
	//2020/08/07 10:39:47 key = 1;value = 10
	//2020/08/07 10:39:47 key = 2;value = 20

	// 容量已满 所以会删除 key = 1
	cache.Put(4,40)
	log.Printf("key = 2;value = %d",cache.Get(4))
	// 容量已满 所以会删除 key = 2
	cache.Put(5,50)

	log.Printf("key = 1;value = %d",cache.Get(1))

}

// 2 1 2 1 2 3 4 -> 删除3
func TestLFUCache_B(t *testing.T) {
	cache := NewLFU(3)
	cache.Put(1,1)
	cache.Put(2,2)
	cache.Put(3,3)

	log.Printf("key = 2;value = %d",cache.Get(2))
	log.Printf("key = 1;value = %d",cache.Get(1))
	log.Printf("key = 2;value = %d",cache.Get(2))
	log.Printf("key = 1;value = %d",cache.Get(1))
	log.Printf("key = 2;value = %d",cache.Get(2))
	log.Printf("key = 3;value = %d",cache.Get(3))

	// debug 查看到：删除了3
	cache.Put(4, 4)
}