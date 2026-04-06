package house_robber_II

import "testing"

func Test_rob(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{2, 3, 2},
			want: 3,
		},
		{
			name: "Example 2",
			nums: []int{1, 2, 3, 1},
			want: 4,
		},
		{
			name: "Example 3",
			nums: []int{1, 2, 3},
			want: 3,
		},
		{
			name: "Single house",
			nums: []int{5},
			want: 5,
		},
		{
			name: "Two houses",
			nums: []int{1, 2},
			want: 2,
		},
		{
			name: "All same values",
			nums: []int{2, 2, 2, 2},
			want: 4,
		},
		{
			name: "Zero money in houses",
			nums: []int{0, 0, 0},
			want: 0,
		},
		{
			name: "Large values",
			nums: []int{100, 1, 1, 100},
			want: 101, // 100 + 1 (last and first cannot be together, so 100 + 1)
		},
		{
			name: "Irregular values",
			nums: []int{6, 3, 10, 8, 2, 10, 3, 5, 10, 5, 3},
			want: 36,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob(tt.nums); got != tt.want {
				t.Errorf("rob() = %v, want %v", got, tt.want)
			}
		})
	}
}
