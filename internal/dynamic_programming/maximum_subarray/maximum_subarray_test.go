package array

import "testing"

func Test_MaxSubArray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "example 1 - mixed positive and negative",
			nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			want: 6,
		},
		{
			name: "example 2 - single element",
			nums: []int{1},
			want: 1,
		},
		{
			name: "example 3 - all positive",
			nums: []int{5, 4, -1, 7, 8},
			want: 23,
		},
		{
			name: "single negative element",
			nums: []int{-5},
			want: -5,
		},
		{
			name: "all negative",
			nums: []int{-3, -2, -5, -1, -4},
			want: -1,
		},
		{
			name: "all positive numbers",
			nums: []int{1, 2, 3, 4, 5},
			want: 15,
		},
		{
			name: "negative at start",
			nums: []int{-2, 3, 4, -1, 5},
			want: 11,
		},
		{
			name: "negative at end",
			nums: []int{3, 4, -1, 5, -10},
			want: 11,
		},
		{
			name: "alternating positive negative",
			nums: []int{1, -1, 1, -1, 1},
			want: 1,
		},
		{
			name: "large positive in middle",
			nums: []int{-2, -3, 10, -1, -2},
			want: 10,
		},
		{
			name: "two elements positive",
			nums: []int{1, 2},
			want: 3,
		},
		{
			name: "two elements negative",
			nums: []int{-1, -2},
			want: -1,
		},
		{
			name: "two elements mixed",
			nums: []int{-1, 2},
			want: 2,
		},
		{
			name: "zeros in array",
			nums: []int{-2, 0, -1},
			want: 0,
		},
		{
			name: "zeros and positive",
			nums: []int{0, 1, 0},
			want: 1,
		},
		{
			name: "all zeros",
			nums: []int{0, 0, 0},
			want: 0,
		},
		{
			name: "large array with mixed values",
			nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4, 6, -2, 3},
			want: 12,
		},
		{
			name: "subarray at start",
			nums: []int{5, 4, 3, -10, 1, 2},
			want: 12,
		},
		{
			name: "subarray at end",
			nums: []int{-5, -4, 3, 4, 5},
			want: 12,
		},
		{
			name: "entire array is max",
			nums: []int{1, 2, 3, 4, 5, 6},
			want: 21,
		},
		{
			name: "negative then positive sequence",
			nums: []int{-10, -5, 1, 2, 3, 4},
			want: 10,
		},
		{
			name: "positive then negative sequence",
			nums: []int{5, 4, 3, -10, -20},
			want: 12,
		},
		{
			name: "complex case 1",
			nums: []int{8, -19, 5, -4, 20},
			want: 21,
		},
		{
			name: "complex case 2",
			nums: []int{-2, -1},
			want: -1,
		},
		{
			name: "edge case - max at boundaries",
			nums: []int{100, -50, 50, -50, 100},
			want: 150,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("Input nums: %v", tt.nums)

			got := maxSubArray(tt.nums)

			t.Logf("Output result: %d", got)

			if got != tt.want {
				t.Fatalf("maxSubArray(%v) = %d; want %d", tt.nums, got, tt.want)
			}
		})
	}
}

func Test_MaxSubArrayDivideConquer(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "example 1 - mixed positive and negative",
			nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			want: 6,
		},
		{
			name: "example 2 - single element",
			nums: []int{1},
			want: 1,
		},
		{
			name: "example 3 - all positive",
			nums: []int{5, 4, -1, 7, 8},
			want: 23,
		},
		{
			name: "single negative element",
			nums: []int{-5},
			want: -5,
		},
		{
			name: "all negative",
			nums: []int{-3, -2, -5, -1, -4},
			want: -1,
		},
		{
			name: "all positive numbers",
			nums: []int{1, 2, 3, 4, 5},
			want: 15,
		},
		{
			name: "negative at start",
			nums: []int{-2, 3, 4, -1, 5},
			want: 11,
		},
		{
			name: "negative at end",
			nums: []int{3, 4, -1, 5, -10},
			want: 11,
		},
		{
			name: "alternating positive negative",
			nums: []int{1, -1, 1, -1, 1},
			want: 1,
		},
		{
			name: "large positive in middle",
			nums: []int{-2, -3, 10, -1, -2},
			want: 10,
		},
		{
			name: "two elements positive",
			nums: []int{1, 2},
			want: 3,
		},
		{
			name: "two elements negative",
			nums: []int{-1, -2},
			want: -1,
		},
		{
			name: "two elements mixed",
			nums: []int{-1, 2},
			want: 2,
		},
		{
			name: "zeros in array",
			nums: []int{-2, 0, -1},
			want: 0,
		},
		{
			name: "zeros and positive",
			nums: []int{0, 1, 0},
			want: 1,
		},
		{
			name: "all zeros",
			nums: []int{0, 0, 0},
			want: 0,
		},
		{
			name: "large array with mixed values",
			nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4, 6, -2, 3},
			want: 12,
		},
		{
			name: "subarray at start",
			nums: []int{5, 4, 3, -10, 1, 2},
			want: 12,
		},
		{
			name: "subarray at end",
			nums: []int{-5, -4, 3, 4, 5},
			want: 12,
		},
		{
			name: "entire array is max",
			nums: []int{1, 2, 3, 4, 5, 6},
			want: 21,
		},
		{
			name: "complex case 1",
			nums: []int{8, -19, 5, -4, 20},
			want: 21,
		},
		{
			name: "complex case 2",
			nums: []int{-2, -1},
			want: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("Input nums: %v", tt.nums)

			got := maxSubArrayDivideConquer(tt.nums)

			t.Logf("Output result: %d", got)

			if got != tt.want {
				t.Fatalf("maxSubArrayDivideConquer(%v) = %d; want %d", tt.nums, got, tt.want)
			}
		})
	}
}

// Benchmark để so sánh hiệu suất giữa hai phương pháp
func Benchmark_MaxSubArray(b *testing.B) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	for i := 0; i < b.N; i++ {
		maxSubArray(nums)
	}
}

func Benchmark_MaxSubArrayDivideConquer(b *testing.B) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	for i := 0; i < b.N; i++ {
		maxSubArrayDivideConquer(nums)
	}
}
