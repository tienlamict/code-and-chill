package dynamic_programming

import "testing"

func Test_climbStairs(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "example 1 - n=2",
			n:    2,
			want: 2,
		},
		{
			name: "example 2 - n=3",
			n:    3,
			want: 3,
		},
		{
			name: "n=1",
			n:    1,
			want: 1,
		},
		{
			name: "n=4",
			n:    4,
			want: 5,
		},
		{
			name: "n=5",
			n:    5,
			want: 8,
		},
		{
			name: "n=6",
			n:    6,
			want: 13,
		},
		{
			name: "constraint max n=45",
			n:    45,
			want: 1836311903,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := climbStairs(tt.n)
			if got != tt.want {
				t.Errorf("climbStairs(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}
