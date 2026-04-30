package hash_table

import (
	"sort"
	"testing"
)

func Test_TopKFrequent(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{
			name: "example 1",
			nums: []int{1, 1, 1, 2, 2, 3},
			k:    2,
			want: []int{1, 2},
		},
		{
			name: "example 2 - single element",
			nums: []int{1},
			k:    1,
			want: []int{1},
		},
		{
			name: "example 3 - tie frequencies",
			nums: []int{1, 2, 1, 2, 1, 2, 3, 1, 3, 2},
			k:    2,
			want: []int{1, 2},
		},
		{
			name: "k equals unique count",
			nums: []int{1, 2, 3},
			k:    3,
			want: []int{1, 2, 3},
		},
		{
			name: "all same elements",
			nums: []int{5, 5, 5, 5},
			k:    1,
			want: []int{5},
		},
		{
			name: "negative numbers",
			nums: []int{-1, -1, -2, -2, -2, -3},
			k:    2,
			want: []int{-2, -1},
		},
		{
			name: "mix of negative and positive",
			nums: []int{-1, 1, -1, 1, -1},
			k:    1,
			want: []int{-1},
		},
		{
			name: "large frequency difference",
			nums: []int{1, 1, 1, 1, 1, 2, 2, 3},
			k:    1,
			want: []int{1},
		},
		{
			name: "k=2 from three unique",
			nums: []int{3, 3, 3, 2, 2, 1},
			k:    2,
			want: []int{3, 2},
		},
		{
			name: "boundary values",
			nums: []int{-10000, 10000, -10000, 10000, -10000},
			k:    1,
			want: []int{-10000},
		},
		{
			name: "two elements equal frequency k=1",
			nums: []int{1, 2},
			k:    1,
			want: []int{1},
		},
		{
			name: "longer array",
			nums: []int{4, 4, 4, 6, 6, 7, 7, 7, 7, 9},
			k:    2,
			want: []int{7, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("Input nums: %v, k: %d", tt.nums, tt.k)

			got := topKFrequent(tt.nums, tt.k)
			t.Logf("Output: %v", got)

			if !equalIgnoreOrder(got, tt.want) {
				t.Fatalf("topKFrequent(%v, %d) = %v; want %v", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}

// equalIgnoreOrder so sánh hai slice số nguyên không phân biệt thứ tự
func equalIgnoreOrder(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	aCopy := make([]int, len(a))
	bCopy := make([]int, len(b))
	copy(aCopy, a)
	copy(bCopy, b)
	sort.Ints(aCopy)
	sort.Ints(bCopy)
	for i := range aCopy {
		if aCopy[i] != bCopy[i] {
			return false
		}
	}
	return true
}
