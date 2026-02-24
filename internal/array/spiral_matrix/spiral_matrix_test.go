package array

import (
	"reflect"
	"testing"
)

func Test_SpiralOrder(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   []int
	}{
		{
			name:   "example 1 - 3x3 matrix",
			matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			want:   []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			name:   "example 2 - 3x4 matrix",
			matrix: [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
			want:   []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
		{
			name:   "single element",
			matrix: [][]int{{1}},
			want:   []int{1},
		},
		{
			name:   "single row",
			matrix: [][]int{{1, 2, 3, 4}},
			want:   []int{1, 2, 3, 4},
		},
		{
			name:   "single column",
			matrix: [][]int{{1}, {2}, {3}, {4}},
			want:   []int{1, 2, 3, 4},
		},
		{
			name:   "2x2 matrix",
			matrix: [][]int{{1, 2}, {3, 4}},
			want:   []int{1, 2, 4, 3},
		},
		{
			name:   "2x3 matrix",
			matrix: [][]int{{1, 2, 3}, {4, 5, 6}},
			want:   []int{1, 2, 3, 6, 5, 4},
		},
		{
			name:   "3x2 matrix",
			matrix: [][]int{{1, 2}, {3, 4}, {5, 6}},
			want:   []int{1, 2, 4, 6, 5, 3},
		},
		{
			name:   "4x4 matrix",
			matrix: [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}},
			want:   []int{1, 2, 3, 4, 8, 12, 16, 15, 14, 13, 9, 5, 6, 7, 11, 10},
		},
		{
			name:   "1x5 matrix",
			matrix: [][]int{{1, 2, 3, 4, 5}},
			want:   []int{1, 2, 3, 4, 5},
		},
		{
			name:   "5x1 matrix",
			matrix: [][]int{{1}, {2}, {3}, {4}, {5}},
			want:   []int{1, 2, 3, 4, 5},
		},
		{
			name:   "matrix with negative numbers",
			matrix: [][]int{{-1, -2, -3}, {-4, -5, -6}, {-7, -8, -9}},
			want:   []int{-1, -2, -3, -6, -9, -8, -7, -4, -5},
		},
		{
			name:   "matrix with mixed positive and negative",
			matrix: [][]int{{1, -2, 3}, {-4, 5, -6}, {7, -8, 9}},
			want:   []int{1, -2, 3, -6, 9, -8, 7, -4, 5},
		},
		{
			name:   "matrix with zeros",
			matrix: [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
			want:   []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name:   "matrix with large numbers",
			matrix: [][]int{{100, 200}, {300, 400}},
			want:   []int{100, 200, 400, 300},
		},
		{
			name:   "matrix with duplicate elements",
			matrix: [][]int{{1, 1, 1}, {2, 2, 2}, {3, 3, 3}},
			want:   []int{1, 1, 1, 2, 3, 3, 3, 2, 2},
		},
		{
			name:   "rectangular matrix 2x5",
			matrix: [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}},
			want:   []int{1, 2, 3, 4, 5, 10, 9, 8, 7, 6},
		},
		{
			name:   "rectangular matrix 5x2",
			matrix: [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9, 10}},
			want:   []int{1, 2, 4, 6, 8, 10, 9, 7, 5, 3},
		},
		{
			name:   "matrix with boundary values",
			matrix: [][]int{{-100, -100}, {-100, -100}},
			want:   []int{-100, -100, -100, -100},
		},
		{
			name:   "matrix 3x5",
			matrix: [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}},
			want:   []int{1, 2, 3, 4, 5, 10, 15, 14, 13, 12, 11, 6, 7, 8, 9},
		},
		{
			name:   "matrix 5x3",
			matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}, {13, 14, 15}},
			want:   []int{1, 2, 3, 6, 9, 12, 15, 14, 13, 10, 7, 4, 5, 8, 11},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("Input matrix: %v", tt.matrix)
			t.Logf("Matrix size: %dx%d", len(tt.matrix), len(tt.matrix[0]))

			got := spiralOrder(tt.matrix)

			t.Logf("Output result: %v", got)
			t.Logf("Expected result: %v", tt.want)

			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("spiralOrder() = %v; want %v", got, tt.want)
			}
		})
	}
}

// Test edge cases
func Test_SpiralOrder_EdgeCases(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   []int
	}{
		{
			name:   "empty matrix - no rows",
			matrix: [][]int{},
			want:   []int{},
		},
		{
			name:   "empty matrix - empty row",
			matrix: [][]int{{}},
			want:   []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := spiralOrder(tt.matrix)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("spiralOrder() = %v; want %v", got, tt.want)
			}
		})
	}
}

// Benchmark test
func Benchmark_SpiralOrder(b *testing.B) {
	matrix := [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}, {21, 22, 23, 24, 25}}
	for i := 0; i < b.N; i++ {
		spiralOrder(matrix)
	}
}
