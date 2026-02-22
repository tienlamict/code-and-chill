package array

import "testing"

func Test_EraseOverlapIntervals(t *testing.T) {
	tests := []struct {
		name       string
		intervals  [][]int
		want       int
	}{
		{
			name:       "example 1",
			intervals:  [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}},
			want:       1,
		},
		{
			name:       "example 2 - all duplicates",
			intervals:  [][]int{{1, 2}, {1, 2}, {1, 2}},
			want:       2,
		},
		{
			name:       "example 3 - no overlap",
			intervals:  [][]int{{1, 2}, {2, 3}},
			want:       0,
		},
		{
			name:       "single interval",
			intervals:  [][]int{{1, 2}},
			want:       0,
		},
		{
			name:       "two overlapping intervals",
			intervals:  [][]int{{1, 3}, {2, 4}},
			want:       1,
		},
		{
			name:       "two non-overlapping intervals",
			intervals:  [][]int{{1, 2}, {3, 4}},
			want:       0,
		},
		{
			name:       "intervals touching at point",
			intervals:  [][]int{{1, 2}, {2, 3}, {3, 4}},
			want:       0,
		},
		{
			name:       "nested intervals",
			intervals:  [][]int{{1, 5}, {2, 3}, {3, 4}},
			want:       1,
		},
		{
			name:       "multiple overlaps",
			intervals:  [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}},
			want:       2,
		},
		{
			name:       "complex case 1",
			intervals:  [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}, {2, 4}},
			want:       2,
		},
		{
			name:       "complex case 2",
			intervals:  [][]int{{1, 100}, {11, 22}, {1, 11}, {2, 12}},
			want:       2,
		},
		{
			name:       "all intervals overlap",
			intervals:  [][]int{{1, 3}, {2, 4}, {3, 5}, {4, 6}},
			want:       2,
		},
		{
			name:       "negative numbers",
			intervals:  [][]int{{-5, -2}, {-3, 0}, {-4, -1}},
			want:       2,
		},
		{
			name:       "mixed positive and negative",
			intervals:  [][]int{{-2, 1}, {0, 3}, {2, 5}},
			want:       1,
		},
		{
			name:       "large intervals",
			intervals:  [][]int{{1, 100}, {50, 150}, {100, 200}},
			want:       1,
		},
		{
			name:       "same start different end",
			intervals:  [][]int{{1, 2}, {1, 3}, {1, 4}},
			want:       2,
		},
		{
			name:       "same end different start",
			intervals:  [][]int{{1, 5}, {2, 5}, {3, 5}},
			want:       2,
		},
		{
			name:       "empty array",
			intervals:  [][]int{},
			want:       0,
		},
		{
			name:       "one interval covers all",
			intervals:  [][]int{{1, 10}, {2, 3}, {4, 5}, {6, 7}},
			want:       1,
		},
		{
			name:       "alternating overlaps",
			intervals:  [][]int{{1, 3}, {2, 4}, {3, 5}, {4, 6}, {5, 7}},
			want:       2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("Input intervals: %v", tt.intervals)

			got := eraseOverlapIntervals(tt.intervals)

			t.Logf("Output: %d", got)

			if got != tt.want {
				t.Fatalf(
					"eraseOverlapIntervals(%v) = %d; want %d",
					tt.intervals, got, tt.want,
				)
			}
		})
	}
}
