package SCC

import "testing"

func BenchmarkNumReader(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NumReader(922)
	}
}
