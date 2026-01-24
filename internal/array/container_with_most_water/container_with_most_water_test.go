package array

import "testing"

func Test_MaxArea(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{
			name:   "example 1",
			height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			want:   49,
		},
		{
			name:   "example 2",
			height: []int{1, 1},
			want:   1,
		},
		{
			name:   "all zeros",
			height: []int{0, 0, 0, 0},
			want:   0,
		},
		{
			name:   "simple symmetric",
			height: []int{2, 0, 2},
			want:   4,
		},
		{
			name:   "strictly increasing",
			height: []int{1, 2, 3, 4, 5},
			want:   6, // min(2,5)*3 = 6 or min(3,5)*2 = 6
		},
		{
			name:   "strictly decreasing",
			height: []int{5, 4, 3, 2, 1},
			want:   6, // min(5,3)*2 = 6 or min(5,2)*3 = 6
		},
		{
			name:   "tall in the middle doesn't help if edges short",
			height: []int{1, 100, 1},
			want:   2,
		},
		{
			name:   "many equal heights",
			height: []int{4, 4, 4, 4},
			want:   12,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("Input height: %v", tt.height)

			got := maxArea(tt.height)

			t.Logf("Output result: %d", got)

			if got != tt.want {
				t.Fatalf("maxArea(%v) = %d; want %d", tt.height, got, tt.want)
			}
		})
	}
}


