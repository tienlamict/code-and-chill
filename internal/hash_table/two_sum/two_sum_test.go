package hash_table

import (
	"reflect"
	"testing"
)

func Test_TwoSum(t *testing.T) {

	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			name:   "example 1",
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		{
			name:   "example 2",
			nums:   []int{3, 2, 4},
			target: 6,
			want:   []int{1, 2},
		},
		{
			name:   "example 3",
			nums:   []int{3, 3},
			target: 6,
			want:   []int{0, 1},
		},
		{
			name:   "negative numbers",
			nums:   []int{-1, -2, -3, -4, -5},
			target: -8,
			want:   []int{2, 4},
		},
		{
			name:   "mixed positive and negative",
			nums:   []int{-3, 4, 3, 90},
			target: 0,
			want:   []int{0, 2},
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			t.Logf("Input nums: %v, target: %d", tt.nums, tt.target)

			got := twoSum(tt.nums, tt.target)

			t.Logf("Output result: %v", got)

			if !reflect.DeepEqual(got, tt.want) {

				t.Fatalf(
					"twoSum(%v, %d) = %v; want %v",
					tt.nums, tt.target, got, tt.want,
				)

			}

		})

	}

}
