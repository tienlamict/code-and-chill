package array

import "testing"

func TestLengthOfLIS(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1: [10,9,2,5,3,7,101,18]",
			nums:     []int{10, 9, 2, 5, 3, 7, 101, 18},
			expected: 4,
		},
		{
			name:     "Example 2: [0,1,0,3,2,3]",
			nums:     []int{0, 1, 0, 3, 2, 3},
			expected: 4,
		},
		{
			name:     "Example 3: Tất cả phần tử giống nhau",
			nums:     []int{7, 7, 7, 7, 7, 7, 7},
			expected: 1,
		},
		{
			name:     "Mảng một phần tử",
			nums:     []int{1},
			expected: 1,
		},
		{
			name:     "Mảng giảm dần",
			nums:     []int{5, 4, 3, 2, 1},
			expected: 1,
		},
		{
			name:     "Mảng tăng dần",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 5,
		},
		{
			name:     "Có số âm",
			nums:     []int{-10, -5, 0, 3, 5},
			expected: 5,
		},
		{
			name:     "Số âm và dương",
			nums:     []int{-2, -1, 0, 1, 2},
			expected: 5,
		},
		{
			name:     "Mảng có phần tử trùng lặp",
			nums:     []int{1, 3, 5, 4, 7, 7, 8},
			expected: 5,
		},
		{
			name:     "Mảng lớn với pattern phức tạp",
			nums:     []int{1, 3, 6, 7, 9, 4, 10, 5, 6},
			expected: 6,
		},
		{
			name:     "Mảng với số lớn",
			nums:     []int{10000, 10001, 9999, 10002, 10003},
			expected: 4,
		},
		{
			name:     "Mảng với số âm lớn",
			nums:     []int{-10000, -9999, -10001, -9998},
			expected: 3,
		},
		{
			name:     "Mảng có nhiều phần tử giống nhau xen kẽ",
			nums:     []int{1, 2, 2, 3, 3, 3, 4},
			expected: 4,
		},
		{
			name:     "Mảng với phần tử đầu lớn nhất",
			nums:     []int{10, 1, 2, 3, 4, 5},
			expected: 5,
		},
		{
			name:     "Mảng với phần tử cuối nhỏ nhất",
			nums:     []int{1, 2, 3, 4, 5, 0},
			expected: 5,
		},
		{
			name:     "Mảng hai phần tử tăng",
			nums:     []int{1, 2},
			expected: 2,
		},
		{
			name:     "Mảng hai phần tử giảm",
			nums:     []int{2, 1},
			expected: 1,
		},
		{
			name:     "Mảng hai phần tử bằng nhau",
			nums:     []int{2, 2},
			expected: 1,
		},
		{
			name:     "Mảng với pattern zigzag",
			nums:     []int{1, 3, 2, 4, 3, 5, 4, 6},
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lengthOfLIS(tt.nums)
			if result != tt.expected {
				t.Errorf("lengthOfLIS(%v) = %d, expected %d", tt.nums, result, tt.expected)
			}
		})
	}
}

func TestLengthOfLISDP(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1: [10,9,2,5,3,7,101,18]",
			nums:     []int{10, 9, 2, 5, 3, 7, 101, 18},
			expected: 4,
		},
		{
			name:     "Example 2: [0,1,0,3,2,3]",
			nums:     []int{0, 1, 0, 3, 2, 3},
			expected: 4,
		},
		{
			name:     "Example 3: Tất cả phần tử giống nhau",
			nums:     []int{7, 7, 7, 7, 7, 7, 7},
			expected: 1,
		},
		{
			name:     "Mảng một phần tử",
			nums:     []int{1},
			expected: 1,
		},
		{
			name:     "Mảng giảm dần",
			nums:     []int{5, 4, 3, 2, 1},
			expected: 1,
		},
		{
			name:     "Mảng tăng dần",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lengthOfLISDP(tt.nums)
			if result != tt.expected {
				t.Errorf("lengthOfLISDP(%v) = %d, expected %d", tt.nums, result, tt.expected)
			}
		})
	}
}

// Benchmark để so sánh hiệu suất giữa hai cách giải
func BenchmarkLengthOfLIS(b *testing.B) {
	nums := make([]int, 2500)
	for i := range nums {
		nums[i] = i % 1000
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lengthOfLIS(nums)
	}
}

func BenchmarkLengthOfLISDP(b *testing.B) {
	nums := make([]int, 2500)
	for i := range nums {
		nums[i] = i % 1000
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lengthOfLISDP(nums)
	}
}
