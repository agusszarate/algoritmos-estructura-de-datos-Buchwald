package main

import "testing"

func TestMinimoExcluido(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		expected int
	}{
		{
			name:     "Falta el 0",
			arr:      []int{1, 2, 3, 4},
			expected: 0,
		},
		{
			name:     "Falta el último + 1",
			arr:      []int{0, 1, 2, 3},
			expected: 4,
		},
		{
			name:     "Falta en el medio",
			arr:      []int{0, 1, 3, 4, 5},
			expected: 2,
		},
		{
			name:     "Falta el primero de varios",
			arr:      []int{2, 3, 4, 5, 6},
			expected: 0,
		},
		{
			name:     "Falta el segundo",
			arr:      []int{0, 2, 3, 4, 5},
			expected: 1,
		},
		{
			name:     "Arreglo con un solo elemento (0)",
			arr:      []int{0},
			expected: 1,
		},
		{
			name:     "Arreglo con un solo elemento (mayor a 0)",
			arr:      []int{5},
			expected: 0,
		},
		{
			name:     "Faltan varios al principio",
			arr:      []int{3, 4, 5, 6, 7},
			expected: 0,
		},
		{
			name:     "Falta al final de secuencia larga",
			arr:      []int{0, 1, 2, 3, 4, 5, 6, 8, 9},
			expected: 7,
		},
		{
			name:     "Arreglo vacío",
			arr:      []int{},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MinimoExcluido(tt.arr)
			if result != tt.expected {
				t.Errorf("MinimoExcluido(%v) = %d, expected %d", tt.arr, result, tt.expected)
			}
		})
	}
}

// Test adicional para verificar casos edge
func TestMinimoExcluidoEdgeCases(t *testing.T) {
	// Test con números más grandes
	bigArray := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		bigArray[i] = i
	}
	result := MinimoExcluido(bigArray)
	expected := 1000
	if result != expected {
		t.Errorf("MinimoExcluido(big array 0-999) = %d, expected %d", result, expected)
	}

	// Test con salto grande al inicio
	jumpArray := []int{100, 101, 102, 103}
	result = MinimoExcluido(jumpArray)
	expected = 0
	if result != expected {
		t.Errorf("MinimoExcluido(%v) = %d, expected %d", jumpArray, result, expected)
	}
}
