package array

import "testing"

func Test_MaxProduct(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "example 1 - positive and negative",
			nums: []int{2, 3, -2, 4},
			want: 6,
		},
		{
			name: "example 2 - with zero",
			nums: []int{-2, 0, -1},
			want: 0,
		},
		{
			name: "single element",
			nums: []int{5},
			want: 5,
		},
		{
			name: "single negative element",
			nums: []int{-5},
			want: -5,
		},
		{
			name: "all positive",
			nums: []int{1, 2, 3, 4},
			want: 24,
		},
		{
			name: "all negative - even count",
			nums: []int{-1, -2, -3, -4},
			want: 24,
		},
		{
			name: "all negative - odd count",
			nums: []int{-1, -2, -3},
			want: 6,
		},
		{
			name: "two negatives make positive",
			nums: []int{-2, 3, -4},
			want: 24,
		},
		{
			name: "zero resets product",
			nums: []int{2, 0, 3, 4},
			want: 12,
		},
		{
			name: "negative at start",
			nums: []int{-2, 3, 4},
			want: 12,
		},
		{
			name: "negative at end",
			nums: []int{2, 3, -4},
			want: 6,
		},
		{
			name: "multiple zeros",
			nums: []int{0, 0, 0, 0},
			want: 0,
		},
		{
			name: "alternating positive negative",
			nums: []int{2, -3, 2, -3},
			want: 36,
		},
		{
			name: "large positive product",
			nums: []int{1, 2, 3, 4, 5},
			want: 120,
		},
		{
			name: "negative in middle",
			nums: []int{2, 3, -2, 4},
			want: 6,
		},
		{
			name: "two negatives separated",
			nums: []int{-2, 1, -3, 4},
			want: 24,
		},
		{
			name: "complex case 1",
			nums: []int{2, -5, -2, -4, 3},
			want: 24,
		},
		{
			name: "complex case 2",
			nums: []int{-1, -2, -9, -6},
			want: 108,
		},
		{
			name: "with one",
			nums: []int{-2, 1, -1},
			want: 2,
		},
		{
			name: "edge case - all ones",
			nums: []int{1, 1, 1, 1},
			want: 1,
		},
		{
			name: "edge case - one and negative",
			nums: []int{1, -1, 1, -1},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("Input nums: %v", tt.nums)

			got := maxProduct(tt.nums)

			t.Logf("Output result: %d", got)

			if got != tt.want {
				t.Fatalf("maxProduct(%v) = %d; want %d", tt.nums, got, tt.want)
			}
		})
	}
}
