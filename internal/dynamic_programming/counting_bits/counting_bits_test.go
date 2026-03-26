package counting_bits

import (
	"reflect"
	"testing"
)

func TestCountBits(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected []int
	}{
		{
			name:     "n = 2",
			n:        2,
			expected: []int{0, 1, 1},
		},
		{
			name:     "n = 5",
			n:        5,
			expected: []int{0, 1, 1, 2, 1, 2},
		},
		{
			name:     "n = 0",
			n:        0,
			expected: []int{0},
		},
		{
			name:     "n = 1",
			n:        1,
			expected: []int{0, 1},
		},
		{
			name:     "n = 10",
			n:        10,
			expected: []int{0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countBits(tt.n)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("countBits() = %v, want %v", got, tt.expected)
			}
		})
	}
}
