package hash_table

import "testing"

func Test_MissingNumber(t *testing.T) {

	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "example 1",
			nums: []int{3, 0, 1},
			want: 2,
		},
		{
			name: "example 2",
			nums: []int{0, 1},
			want: 2,
		},
		{
			name: "example 3",
			nums: []int{9, 6, 4, 2, 3, 5, 7, 0, 1},
			want: 8,
		},
		{
			name: "missing 0",
			nums: []int{1, 2},
			want: 0,
		},
		{
			name: "missing last number",
			nums: []int{0, 1, 2, 3},
			want: 4,
		},
		{
			name: "single element missing 0",
			nums: []int{1},
			want: 0,
		},
		{
			name: "single element missing 1",
			nums: []int{0},
			want: 1,
		},
		{
			name: "missing middle number",
			nums: []int{0, 1, 3, 4},
			want: 2,
		},
		{
			name: "large array",
			nums: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99},
			want: 100,
		},
		{
			name: "missing first number in large array",
			nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want: 0,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			t.Logf("Input nums: %v", tt.nums)

			got := missingNumber(tt.nums)

			t.Logf("Output result: %d", got)

			if got != tt.want {

				t.Fatalf(
					"missingNumber(%v) = %d; want %d",
					tt.nums, got, tt.want,
				)

			}

		})

	}

}
