package array

import "testing"

func Test_Search(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "example 1 - target in rotated part",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 0,
			want:   4,
		},
		{
			name:   "example 2 - target not found",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 3,
			want:   -1,
		},
		{
			name:   "example 3 - single element not found",
			nums:   []int{1},
			target: 0,
			want:   -1,
		},
		{
			name:   "single element found",
			nums:   []int{1},
			target: 1,
			want:   0,
		},
		{
			name:   "target in sorted left half",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 5,
			want:   1,
		},
		{
			name:   "target in sorted right half",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 1,
			want:   5,
		},
		{
			name:   "target at first position",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 4,
			want:   0,
		},
		{
			name:   "target at last position",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 2,
			want:   6,
		},
		{
			name:   "target at middle",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 7,
			want:   3,
		},
		{
			name:   "not rotated - target found",
			nums:   []int{1, 2, 3, 4, 5},
			target: 3,
			want:   2,
		},
		{
			name:   "not rotated - target not found",
			nums:   []int{1, 2, 3, 4, 5},
			target: 6,
			want:   -1,
		},
		{
			name:   "rotated by 1",
			nums:   []int{5, 1, 2, 3, 4},
			target: 1,
			want:   1,
		},
		{
			name:   "rotated by n-1",
			nums:   []int{2, 3, 4, 5, 1},
			target: 1,
			want:   4,
		},
		{
			name:   "two elements - found first",
			nums:   []int{1, 3},
			target: 1,
			want:   0,
		},
		{
			name:   "two elements - found second",
			nums:   []int{1, 3},
			target: 3,
			want:   1,
		},
		{
			name:   "two elements - not found",
			nums:   []int{1, 3},
			target: 2,
			want:   -1,
		},
		{
			name:   "target smaller than all",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: -1,
			want:   -1,
		},
		{
			name:   "target larger than all",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 10,
			want:   -1,
		},
		{
			name:   "complex rotation 1",
			nums:   []int{7, 8, 9, 0, 1, 2, 3, 4, 5, 6},
			target: 3,
			want:   6,
		},
		{
			name:   "complex rotation 2",
			nums:   []int{7, 8, 9, 0, 1, 2, 3, 4, 5, 6},
			target: 8,
			want:   1,
		},
		{
			name:   "complex rotation 3",
			nums:   []int{7, 8, 9, 0, 1, 2, 3, 4, 5, 6},
			target: 0,
			want:   3,
		},
		{
			name:   "negative numbers",
			nums:   []int{3, 1},
			target: 1,
			want:   1,
		},
		{
			name:   "negative numbers rotated",
			nums:   []int{-1, 0, 1, 2, -2},
			target: -2,
			want:   4,
		},
		{
			name:   "negative numbers rotated 2",
			nums:   []int{-1, 0, 1, 2, -2},
			target: 0,
			want:   1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("Input nums: %v, target: %d", tt.nums, tt.target)

			got := search(tt.nums, tt.target)

			t.Logf("Output index: %d", got)

			if got != tt.want {
				t.Fatalf("search(%v, %d) = %d; want %d", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
