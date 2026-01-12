package hash_table

import "testing"

func Test_LongestConsecutive(t *testing.T) {

	tests := []struct {
		name string

		nums []int

		want int
	}{

		{

			name: "example 1",

			nums: []int{100, 4, 200, 1, 3, 2},

			want: 4,
		},

		{

			name: "example 2",

			nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},

			want: 9,
		},

		{

			name: "empty array",

			nums: []int{},

			want: 0,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			t.Logf("Input nums: %v", tt.nums)

			got := longestConsecutive(tt.nums)

			t.Logf("Output result: %d", got)

			if got != tt.want {

				t.Fatalf(

					"longestConsecutive(%v) = %d; want %d",

					tt.nums, got, tt.want,
				)

			}

		})

	}

}
