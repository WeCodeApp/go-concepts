package advanced

import (
	"fmt"
	"testing"
)

func Add(a, b int) int {
	return a + b
}

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("Add(2, 3) = %d; want %d", result, expected)
	}
}

// ============ TESTING
func TestAddTableDriven(t *testing.T) {
	tests := []struct{ a, b, expected int }{
		{2, 3, 5},
		{0, 1, 1},
		{-1, 1, 0},
	}

	for _, test := range tests {
		result := Add(test.a, test.b)
		if result != test.expected {
			t.Errorf("Add(%d, %d) = %d; want %d", test.a, test.b, result, test.expected)
		}
	}
}

func TestAddSubtests(t *testing.T) {
	tests := []struct{ a, b, expected int }{
		{2, 3, 5},
		{0, 0, 0},
		{-1, 1, 0},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Add(%d, %d)", test.a, test.b), func(t *testing.T) {
			result := Add(test.a, test.b)
			if result != test.expected {
				t.Errorf("result = %d; want %d", result, test.expected)
			}
		})
	}
}

// ========= BENCHMARKING
func BenchmarkAddSmallInput(b *testing.B) {
	// for i := 0; i < b.N ; i++
	for range b.N {
		Add(2, 3)
	}
}

func BenchmarkAddMediumInput(b *testing.B) {
	// for i := 0; i < b.N ; i++
	for range b.N {
		Add(200, 300)
	}
}

func BenchmarkAddLargeInput(b *testing.B) {
	// for i := 0; i < b.N ; i++
	for range b.N {
		Add(2000, 3000)
	}
}
