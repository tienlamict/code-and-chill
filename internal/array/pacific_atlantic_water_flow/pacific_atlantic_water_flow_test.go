package array

import (
	"reflect"
	"sort"
	"testing"
)

// sortResult sắp xếp kết quả để so sánh
func sortResult(result [][]int) {
	sort.Slice(result, func(i, j int) bool {
		if result[i][0] != result[j][0] {
			return result[i][0] < result[j][0]
		}
		return result[i][1] < result[j][1]
	})
}

// equalResults kiểm tra xem hai kết quả có bằng nhau không (không quan tâm thứ tự)
func equalResults(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	// Tạo bản sao để sắp xếp
	aCopy := make([][]int, len(a))
	bCopy := make([][]int, len(b))
	for i := range a {
		aCopy[i] = make([]int, len(a[i]))
		copy(aCopy[i], a[i])
	}
	for i := range b {
		bCopy[i] = make([]int, len(b[i]))
		copy(bCopy[i], b[i])
	}

	sortResult(aCopy)
	sortResult(bCopy)

	return reflect.DeepEqual(aCopy, bCopy)
}

func Test_PacificAtlantic(t *testing.T) {

	tests := []struct {
		name    string
		heights [][]int
		want    [][]int
	}{
		{
			name: "example 1",
			heights: [][]int{
				{1, 2, 2, 3, 5},
				{3, 2, 3, 4, 4},
				{2, 4, 5, 3, 1},
				{6, 7, 1, 4, 5},
				{5, 1, 1, 2, 4},
			},
			want: [][]int{{0, 4}, {1, 3}, {1, 4}, {2, 2}, {3, 0}, {3, 1}, {4, 0}},
		},
		{
			name:    "example 2",
			heights: [][]int{{1}},
			want:    [][]int{{0, 0}},
		},
		{
			name: "single row",
			heights: [][]int{
				{1, 2, 3, 4, 5},
			},
			want: [][]int{{0, 0}, {0, 1}, {0, 2}, {0, 3}, {0, 4}},
		},
		{
			name: "single column",
			heights: [][]int{
				{1},
				{2},
				{3},
				{4},
				{5},
			},
			want: [][]int{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}},
		},
		{
			name: "all same height",
			heights: [][]int{
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
			},
			want: [][]int{{0, 0}, {0, 1}, {0, 2}, {1, 0}, {1, 1}, {1, 2}, {2, 0}, {2, 1}, {2, 2}},
		},
		{
			name: "increasing from top-left",
			heights: [][]int{
				{1, 2, 3},
				{2, 3, 4},
				{3, 4, 5},
			},
			want: [][]int{{0, 2}, {1, 2}, {2, 0}, {2, 1}, {2, 2}},
		},
		{
			name: "decreasing from top-left",
			heights: [][]int{
				{5, 4, 3},
				{4, 3, 2},
				{3, 2, 1},
			},
			want: [][]int{{0, 0}, {0, 1}, {0, 2}, {1, 0}, {2, 0}},
		},
		{
			name: "mountain in middle",
			heights: [][]int{
				{1, 1, 1},
				{1, 5, 1},
				{1, 1, 1},
			},
			want: [][]int{{0, 0}, {0, 1}, {0, 2}, {1, 0}, {1, 1}, {1, 2}, {2, 0}, {2, 1}, {2, 2}},
		},
		{
			name: "valley in middle",
			heights: [][]int{
				{5, 5, 5},
				{5, 1, 5},
				{5, 5, 5},
			},
			want: [][]int{{0, 0}, {0, 1}, {0, 2}, {1, 0}, {1, 2}, {2, 0}, {2, 1}, {2, 2}},
		},
		{
			name: "two rows",
			heights: [][]int{
				{1, 2, 3},
				{3, 2, 1},
			},
			want: [][]int{{0, 1}, {0, 2}, {1, 0}, {1, 1}},
		},
		{
			name: "two columns",
			heights: [][]int{
				{1, 3},
				{2, 2},
				{3, 1},
			},
			want: [][]int{{0, 1}, {1, 0}, {1, 1}, {2, 0}},
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			t.Logf("Input heights: %v", tt.heights)

			got := pacificAtlantic(tt.heights)

			t.Logf("Output: %v", got)
			t.Logf("Expected: %v", tt.want)

			if !equalResults(got, tt.want) {

				t.Fatalf(
					"pacificAtlantic(%v) = %v; want %v",
					tt.heights, got, tt.want,
				)

			}

		})

	}

}
