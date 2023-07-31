package lru

import (
	"log"
	"testing"
)

// 2 1 2 1 2 3 4 -> 删除1
func TestNewLRU_A(t *testing.T) {
	cache := NewLRU(3)
	cache.Put(1,1)
	cache.Put(2,2)
	cache.Put(3,3)

	log.Printf("key = 2;value = %d",cache.Get(2))
	log.Printf("key = 1;value = %d",cache.Get(1))
	log.Printf("key = 2;value = %d",cache.Get(2))
	log.Printf("key = 1;value = %d",cache.Get(1))
	log.Printf("key = 2;value = %d",cache.Get(2))
	log.Printf("key = 3;value = %d",cache.Get(3))

	// debug 查看到：删除了1
	cache.Put(4, 4)
}