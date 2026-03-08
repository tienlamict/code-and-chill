package dynamic_programming

import "testing"

func Test_coinChange(t *testing.T) {
	tests := []struct {
		name   string
		coins  []int
		amount int
		want   int
	}{
		{
			name:   "example 1 - [1,2,5] amount 11",
			coins:  []int{1, 2, 5},
			amount: 11,
			want:   3,
		},
		{
			name:   "example 2 - impossible",
			coins:  []int{2},
			amount: 3,
			want:   -1,
		},
		{
			name:   "example 3 - amount 0",
			coins:  []int{1},
			amount: 0,
			want:   0,
		},
		{
			name:   "single coin exact",
			coins:  []int{5},
			amount: 5,
			want:   1,
		},
		{
			name:   "single coin multiple",
			coins:  []int{2},
			amount: 8,
			want:   4,
		},
		{
			name:   "only coin 1",
			coins:  []int{1},
			amount: 10,
			want:   10,
		},
		{
			name:   "coins larger than amount",
			coins:  []int{5, 10},
			amount: 3,
			want:   -1,
		},
		{
			name:   "one coin equals amount",
			coins:  []int{1, 2, 3},
			amount: 2,
			want:   1,
		},
		{
			name:   "duplicate denominations",
			coins:  []int{1, 1, 2, 2, 5},
			amount: 11,
			want:   3,
		},
		{
			name:   "greedy fails - need non-greedy",
			coins:  []int{1, 3, 4},
			amount: 6,
			want:   2, // 3+3, not 4+1+1
		},
		{
			name:   "large amount",
			coins:  []int{1, 2, 5},
			amount: 100,
			want:   20,
		},
		{
			name:   "single coin 1 amount 1",
			coins:  []int{1},
			amount: 1,
			want:   1,
		},
		{
			name:   "empty amount with coins",
			coins:  []int{1, 2, 5},
			amount: 0,
			want:   0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := coinChange(tt.coins, tt.amount)
			if got != tt.want {
				t.Errorf("coinChange(%v, %d) = %d, want %d", tt.coins, tt.amount, got, tt.want)
			}
		})
	}
}
