package dynamic_programming

import "testing"

func Test_uniquePaths(t *testing.T) {
	tests := []struct {
		name string
		m    int
		n    int
		want int
	}{
		{
			name: "example 1 - LeetCode",
			m:    3,
			n:    7,
			want: 28,
		},
		{
			name: "example 2 - LeetCode",
			m:    3,
			n:    2,
			want: 3,
		},
		{
			name: "single row",
			m:    1,
			n:    10,
			want: 1,
		},
		{
			name: "single column",
			m:    10,
			n:    1,
			want: 1,
		},
		{
			name: "1x1 grid",
			m:    1,
			n:    1,
			want: 1,
		},
		{
			name: "2x2 grid",
			m:    2,
			n:    2,
			want: 2,
		},
		{
			name: "2x3 grid",
			m:    2,
			n:    3,
			want: 3,
		},
		{
			name: "3x3 grid",
			m:    3,
			n:    3,
			want: 6,
		},
		{
			name: "large grid 7x3",
			m:    7,
			n:    3,
			want: 28,
		},
		{
			name: "large grid 10x10",
			m:    10,
			n:    10,
			want: 48620,
		},
		{
			name: "rectangular grid 5x4",
			m:    5,
			n:    4,
			want: 35,
		},
		{
			name: "rectangular grid 4x5",
			m:    4,
			n:    5,
			want: 35,
		},
		{
			name: "small grid 2x4",
			m:    2,
			n:    4,
			want: 4,
		},
		{
			name: "medium grid 5x5",
			m:    5,
			n:    5,
			want: 70,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("Input: m = %d, n = %d", tt.m, tt.n)

			got := uniquePaths(tt.m, tt.n)

			t.Logf("Output: %d", got)
			t.Logf("Expected: %d", tt.want)

			if got != tt.want {
				t.Errorf("uniquePaths(%d, %d) = %d; want %d", tt.m, tt.n, got, tt.want)
			}
		})
	}
}
