package main

import "testing"

func TestIsPrime(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{"Prime number 2", 2, true},
		{"Prime number 3", 3, true},
		{"Prime number 17", 17, true},
		{"Prime number 97", 97, true},
		{"Non-prime 1", 1, false},
		{"Non-prime 4", 4, false},
		{"Non-prime 15", 15, false},
		{"Non-prime 100", 100, false},
		{"Large prime 104729", 104729, true},     // 10,000th prime number
		{"Large non-prime 104730", 104730, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPrime(tt.input); got != tt.expected {
				t.Errorf("IsPrime(%d) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}

func BenchmarkIsPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPrime(104729) // Test with a large prime number
	}
} 