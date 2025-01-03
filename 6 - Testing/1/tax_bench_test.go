package tax

import "testing"

func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500.00)
	}
}

func BenchmarkCalculateTaxLazy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTaxLazy(500.00)
	}
}
