package array

import "testing"

func Test_FindMin(t *testing.T) {

	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "example 1",
			nums: []int{3, 4, 5, 1, 2},
			want: 1,
		},
		{
			name: "example 2",
			nums: []int{4, 5, 6, 7, 0, 1, 2},
			want: 0,
		},
		{
			name: "example 3",
			nums: []int{11, 13, 15, 17},
			want: 11,
		},
		{
			name: "single element",
			nums: []int{1},
			want: 1,
		},
		{
			name: "two elements rotated",
			nums: []int{2, 1},
			want: 1,
		},
		{
			name: "two elements not rotated",
			nums: []int{1, 2},
			want: 1,
		},
		{
			name: "minimum at start",
			nums: []int{1, 2, 3, 4, 5},
			want: 1,
		},
		{
			name: "minimum at end",
			nums: []int{2, 3, 4, 5, 1},
			want: 1,
		},
		{
			name: "minimum in middle",
			nums: []int{4, 5, 1, 2, 3},
			want: 1,
		},
		{
			name: "large rotation",
			nums: []int{5, 6, 7, 8, 9, 1, 2, 3, 4},
			want: 1,
		},
		{
			name: "negative numbers",
			nums: []int{-1, 0, 1, 2, -2},
			want: -2,
		},
		{
			name: "all negative",
			nums: []int{-5, -4, -3, -2, -1},
			want: -5,
		},
		{
			name: "three elements rotated",
			nums: []int{3, 1, 2},
			want: 1,
		},
		{
			name: "three elements not rotated",
			nums: []int{1, 2, 3},
			want: 1,
		},
		{
			name: "four elements rotated",
			nums: []int{4, 1, 2, 3},
			want: 1,
		},
		{
			name: "duplicate pattern but unique",
			nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want: 1,
		},
		{
			name: "rotated once",
			nums: []int{10, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			want: 1,
		},
		{
			name: "rotated almost full",
			nums: []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 1},
			want: 1,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			t.Logf("Input nums: %v", tt.nums)

			got := findMin(tt.nums)

			t.Logf("Output result: %d", got)

			if got != tt.want {

				t.Fatalf(
					"findMin(%v) = %d; want %d",
					tt.nums, got, tt.want,
				)

			}

		})

	}

}
