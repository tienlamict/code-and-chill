package array

import (
	"reflect"
	"sort"
	"testing"
)

// sortResult sắp xếp kết quả để so sánh
func sortResult(result [][]int) {
	// Sắp xếp từng combination
	for _, comb := range result {
		sort.Ints(comb)
	}

	// Sắp xếp các combinations
	sort.Slice(result, func(i, j int) bool {
		if len(result[i]) != len(result[j]) {
			return len(result[i]) < len(result[j])
		}
		for k := 0; k < len(result[i]); k++ {
			if result[i][k] != result[j][k] {
				return result[i][k] < result[j][k]
			}
		}
		return false
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

func Test_CombinationSum(t *testing.T) {

	tests := []struct {
		name       string
		candidates []int
		target     int
		want       [][]int
	}{
		{
			name:       "example 1",
			candidates: []int{2, 3, 6, 7},
			target:     7,
			want:       [][]int{{2, 2, 3}, {7}},
		},
		{
			name:       "example 2",
			candidates: []int{2, 3, 5},
			target:     8,
			want:       [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
		{
			name:       "example 3",
			candidates: []int{2},
			target:     1,
			want:       [][]int{},
		},
		{
			name:       "single candidate matches",
			candidates: []int{2},
			target:     2,
			want:       [][]int{{2}},
		},
		{
			name:       "single candidate multiple times",
			candidates: []int{2},
			target:     8,
			want:       [][]int{{2, 2, 2, 2}},
		},
		{
			name:       "no solution",
			candidates: []int{3, 5, 7},
			target:     2,
			want:       [][]int{},
		},
		{
			name:       "all candidates needed",
			candidates: []int{1, 2, 3},
			target:     6,
			want:       [][]int{{1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 2}, {1, 1, 1, 3}, {1, 1, 2, 2}, {1, 2, 3}, {2, 2, 2}, {3, 3}},
		},
		{
			name:       "large target",
			candidates: []int{2, 3, 5},
			target:     10,
			want:       [][]int{{2, 2, 2, 2, 2}, {2, 2, 3, 3}, {2, 3, 5}, {5, 5}},
		},
		{
			name:       "duplicate candidates not allowed",
			candidates: []int{2, 3, 5},
			target:     7,
			want:       [][]int{{2, 2, 3}, {2, 5}},
		},
		{
			name:       "one candidate multiple solutions",
			candidates: []int{1},
			target:     3,
			want:       [][]int{{1, 1, 1}},
		},
		{
			name:       "sorted candidates",
			candidates: []int{2, 4, 6, 8},
			target:     8,
			want:       [][]int{{2, 2, 2, 2}, {2, 2, 4}, {2, 6}, {4, 4}, {8}},
		},
		{
			name:       "unsorted candidates",
			candidates: []int{7, 3, 2},
			target:     18,
			want:       [][]int{{2, 2, 2, 2, 2, 2, 2, 2, 2}, {2, 2, 2, 2, 2, 2, 3, 3}, {2, 2, 2, 2, 3, 7}, {2, 2, 2, 3, 3, 3, 3}, {2, 2, 7, 7}, {2, 3, 3, 3, 7}, {3, 3, 3, 3, 3, 3}},
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			t.Logf("Input candidates: %v, target: %d", tt.candidates, tt.target)

			got := combinationSum(tt.candidates, tt.target)

			t.Logf("Output: %v", got)
			t.Logf("Expected: %v", tt.want)

			if !equalResults(got, tt.want) {

				t.Fatalf(
					"combinationSum(%v, %d) = %v; want %v",
					tt.candidates, tt.target, got, tt.want,
				)

			}

		})

	}

}
