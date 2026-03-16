package hash_table

import (
	"reflect"
	"testing"
)

// compare2DIntSlices so sánh hai ma trận int ([][]int).
func compare2DIntSlices(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !reflect.DeepEqual(a[i], b[i]) {
			return false
		}
	}
	return true
}

func Test_setZeroes(t *testing.T) {
	tests := []struct {
		name   string
		input  [][]int
		expect [][]int
	}{
		{
			name:   "example 1",
			input:  [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}},
			expect: [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}},
		},
		{
			name:   "example 2",
			input:  [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}},
			expect: [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}},
		},
		{
			name:   "single element non-zero",
			input:  [][]int{{5}},
			expect: [][]int{{5}},
		},
		{
			name:   "single element zero",
			input:  [][]int{{0}},
			expect: [][]int{{0}},
		},
		{
			name:   "no zeros",
			input:  [][]int{{1, 2}, {3, 4}},
			expect: [][]int{{1, 2}, {3, 4}},
		},
		{
			name:   "all zeros",
			input:  [][]int{{0, 0}, {0, 0}},
			expect: [][]int{{0, 0}, {0, 0}},
		},
		{
			name:   "zero in first row only",
			input:  [][]int{{1, 0, 3}, {4, 5, 6}, {7, 8, 9}},
			expect: [][]int{{0, 0, 0}, {4, 0, 6}, {7, 0, 9}},
		},
		{
			name:   "zero in first column only",
			input:  [][]int{{1, 2, 3}, {0, 5, 6}, {7, 8, 9}},
			expect: [][]int{{0, 2, 3}, {0, 0, 0}, {0, 8, 9}},
		},
		{
			name:   "zeros in first row and first column",
			input:  [][]int{{0, 2, 3}, {4, 5, 0}, {7, 8, 9}},
			expect: [][]int{{0, 0, 0}, {0, 5, 0}, {0, 8, 0}},
		},
		{
			name:   "negative numbers and zeros",
			input:  [][]int{{-1, -2, 0}, {-3, -4, -5}, {0, -6, -7}},
			expect: [][]int{{0, 0, 0}, {0, -4, 0}, {0, 0, 0}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Sao chép input để log dễ hơn nếu cần, ở đây thao tác trực tiếp là đủ
			matrix := tt.input

			setZeroes(matrix)

			if !compare2DIntSlices(matrix, tt.expect) {
				t.Fatalf("setZeroes() = %v, expect %v", matrix, tt.expect)
			}
		})
	}
}

