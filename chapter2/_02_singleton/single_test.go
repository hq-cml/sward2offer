package _02_singleton

import (
	"testing"
)

func BenchmarkGetIns1(b *testing.B) {
	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		GetIns1()
	}
}

func BenchmarkGetIns2(b *testing.B) {
	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		GetIns2()
	}
}

func BenchmarkGetIns3(b *testing.B) {
	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		GetIns3()
	}
}