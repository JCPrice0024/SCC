package SCC

import "testing"

func BenchmarkNumBreaker(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NumBreaker(1234567)
	}
}

func BenchmarkNumBreaker2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NumBreaker2(1234567)
	}
}
func BenchmarkNumBreaker3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NumBreaker3(1234567)
	}
}
