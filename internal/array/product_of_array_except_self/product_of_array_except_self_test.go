package product_of_array_except_self

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "Ví dụ 1",
			nums:     []int{1, 2, 3, 4},
			expected: []int{24, 12, 8, 6},
		},
		{
			name:     "Ví dụ 2",
			nums:     []int{-1, 1, 0, -3, 3},
			expected: []int{0, 0, 9, 0, 0},
		},
		{
			name:     "Mảng có 2 phần tử",
			nums:     []int{10, 20},
			expected: []int{20, 10},
		},
		{
			name:     "Mảng có nhiều số 0",
			nums:     []int{1, 0, 3, 0},
			expected: []int{0, 0, 0, 0},
		},
		{
			name:     "Mảng có số âm",
			nums:     []int{-2, -1, -3},
			expected: []int{3, 6, 2},
		},
		{
			name:     "Mảng toàn số 1",
			nums:     []int{1, 1, 1, 1},
			expected: []int{1, 1, 1, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := ProductExceptSelf(tt.nums)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("ProductExceptSelf(%v) = %v, expected %v", tt.nums, actual, tt.expected)
			}
		})
	}
}
