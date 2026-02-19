package array

import (
	"reflect"
	"testing"
)

// copyMatrix tạo bản sao sâu của ma trận 2D
func copyMatrix(m [][]int) [][]int {
	if m == nil {
		return nil
	}
	cp := make([][]int, len(m))
	for i := range m {
		cp[i] = make([]int, len(m[i]))
		copy(cp[i], m[i])
	}
	return cp
}

// matricesEqual so sánh hai ma trận 2D có bằng nhau không
func matricesEqual(a, b [][]int) bool {
	return reflect.DeepEqual(a, b)
}

func TestRotate(t *testing.T) {
	tests := []struct {
		name     string
		matrix   [][]int
		expected [][]int
	}{
		{
			name: "Example 1: 3x3 matrix",
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: [][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
		{
			name: "Example 2: 4x4 matrix",
			matrix: [][]int{
				{5, 1, 9, 11},
				{2, 4, 8, 10},
				{13, 3, 6, 7},
				{15, 14, 12, 16},
			},
			expected: [][]int{
				{15, 13, 2, 5},
				{14, 3, 4, 1},
				{12, 6, 8, 9},
				{16, 7, 10, 11},
			},
		},
		{
			name: "Ma trận 1x1",
			matrix: [][]int{
				{1},
			},
			expected: [][]int{
				{1},
			},
		},
		{
			name: "Ma trận 2x2",
			matrix: [][]int{
				{1, 2},
				{3, 4},
			},
			expected: [][]int{
				{3, 1},
				{4, 2},
			},
		},
		{
			name: "Ma trận có số âm",
			matrix: [][]int{
				{-1, -2},
				{-3, -4},
			},
			expected: [][]int{
				{-3, -1},
				{-4, -2},
			},
		},
		{
			name: "Ma trận có số 0",
			matrix: [][]int{
				{0, 1},
				{2, 0},
			},
			expected: [][]int{
				{2, 0},
				{0, 1},
			},
		},
		{
			name: "Ma trận 5x5",
			matrix: [][]int{
				{1, 2, 3, 4, 5},
				{6, 7, 8, 9, 10},
				{11, 12, 13, 14, 15},
				{16, 17, 18, 19, 20},
				{21, 22, 23, 24, 25},
			},
			expected: [][]int{
				{21, 16, 11, 6, 1},
				{22, 17, 12, 7, 2},
				{23, 18, 13, 8, 3},
				{24, 19, 14, 9, 4},
				{25, 20, 15, 10, 5},
			},
		},
		{
			name: "Ma trận với giá trị lớn",
			matrix: [][]int{
				{1000, -1000},
				{0, 500},
			},
			expected: [][]int{
				{0, 1000},
				{500, -1000},
			},
		},
		{
			name: "Ma trận tất cả phần tử giống nhau",
			matrix: [][]int{
				{7, 7, 7},
				{7, 7, 7},
				{7, 7, 7},
			},
			expected: [][]int{
				{7, 7, 7},
				{7, 7, 7},
				{7, 7, 7},
			},
		},
		{
			name: "Ma trận đơn vị 3x3",
			matrix: [][]int{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
			expected: [][]int{
				{0, 0, 1},
				{0, 1, 0},
				{1, 0, 0},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Tạo bản sao vì rotate sửa đổi in-place
			matrix := copyMatrix(tt.matrix)
			rotate(matrix)

			if !matricesEqual(matrix, tt.expected) {
				t.Errorf("rotate() = %v, expected %v", matrix, tt.expected)
			}
		})
	}
}

// TestRotateInPlace kiểm tra rotate thực sự sửa đổi ma trận gốc
func TestRotateInPlace(t *testing.T) {
	matrix := [][]int{
		{1, 2},
		{3, 4},
	}
	expected := [][]int{
		{3, 1},
		{4, 2},
	}

	rotate(matrix)

	if !matricesEqual(matrix, expected) {
		t.Errorf("rotate() should modify matrix in-place. Got %v, expected %v", matrix, expected)
	}
}

// TestRotateFourTimes trả về ma trận ban đầu khi xoay 4 lần
func TestRotateFourTimes(t *testing.T) {
	original := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	matrix := copyMatrix(original)

	for i := 0; i < 4; i++ {
		rotate(matrix)
	}

	if !matricesEqual(matrix, original) {
		t.Errorf("rotate 4 times should return to original. Got %v, expected %v", matrix, original)
	}
}
