package util

import (
	"fmt"
	"testing"
)

func TestIsPrime(t *testing.T) {
	tests := []struct {
		input    int
		expected bool
	}{
		{2, true},
		{4, false},
		{7, true},
		{9, false},
		{11, true},
	}

	for _, test := range tests {
		t.Run("IsPrime", func(t *testing.T) {
			result := IsPrime(test.input)
			if result != test.expected {
				t.Errorf("IsPrime(%d) = %v; expected %v", test.input, result, test.expected)
			}
		})
	}
}

func TestReverseString(t *testing.T) {
	result := ReverseString("hello")
	expected := "olleh"
	if result != expected {
		t.Errorf("ReverseString('hello') = %s; expected %s", result, expected)
	}
}

func BenchmarkReverseString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReverseString("hello world")
	}
}

func TestFactorial(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{-5, -1},
		{0, 1},
		{1, 1},
		{5, 120},
		{7, 5040},
	}

	for _, test := range tests {
		t.Run("Factorial", func(t *testing.T) {
			result := Factorial(test.input)
			if result != test.expected {
				t.Errorf("Factorial(%d) = %d; expected %d", test.input, result, test.expected)
			}
		})
	}
}

func TestSumOfSlice(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{}, 0},
		{[]int{3, 9, 5}, 17},
		{[]int{-3, -3, -3}, -9},
		{[]int{100}, 100},
	}

	for _, test := range tests {
		t.Run("SumOfSlice", func(t *testing.T) {
			result := SumOfSlice(test.input)
			if result != test.expected {
				t.Errorf("SumOfSlice(%v) = %d; expected %d", test.input, result, test.expected)
			}
		})
	}
}

func BenchmarkIsPrime(b *testing.B) {
	numbers := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29} // Test with small prime numbers
	for i := 0; i < b.N; i++ {
		for _, num := range numbers {
			IsPrime(num)
		}
	}
}
func BenchmarkSumOfSlice(b *testing.B) {
	for size := 100; size <= 10000; size *= 10 { // Test slices with 100, 1000, and 10000 elements
		b.Run(fmt.Sprintf("SliceSize_%d", size), func(b *testing.B) {
			slice := make([]int, size)
			for i := range slice {
				slice[i] = i
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				SumOfSlice(slice)
			}
		})
	}
}
