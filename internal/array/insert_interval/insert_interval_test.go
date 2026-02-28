package array

import (
	"reflect"
	"testing"
)

func Test_Insert(t *testing.T) {
	tests := []struct {
		name        string
		intervals   [][]int
		newInterval []int
		want        [][]int
	}{
		{
			name:        "example 1",
			intervals:   [][]int{{1, 3}, {6, 9}},
			newInterval: []int{2, 5},
			want:        [][]int{{1, 5}, {6, 9}},
		},
		{
			name:        "example 2",
			intervals:   [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
			newInterval: []int{4, 8},
			want:        [][]int{{1, 2}, {3, 10}, {12, 16}},
		},
		{
			name:        "mảng rỗng",
			intervals:   [][]int{},
			newInterval: []int{1, 3},
			want:        [][]int{{1, 3}},
		},
		{
			name:        "chèn vào đầu",
			intervals:   [][]int{{5, 7}, {10, 12}},
			newInterval: []int{1, 3},
			want:        [][]int{{1, 3}, {5, 7}, {10, 12}},
		},
		{
			name:        "chèn vào cuối",
			intervals:   [][]int{{1, 3}, {5, 7}},
			newInterval: []int{10, 12},
			want:        [][]int{{1, 3}, {5, 7}, {10, 12}},
		},
		{
			name:        "chèn vào giữa không chồng chéo",
			intervals:   [][]int{{1, 3}, {10, 12}},
			newInterval: []int{5, 7},
			want:        [][]int{{1, 3}, {5, 7}, {10, 12}},
		},
		{
			name:        "gộp với một interval",
			intervals:   [][]int{{1, 3}, {6, 9}},
			newInterval: []int{2, 5},
			want:        [][]int{{1, 5}, {6, 9}},
		},
		{
			name:        "gộp với nhiều intervals",
			intervals:   [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
			newInterval: []int{4, 8},
			want:        [][]int{{1, 2}, {3, 10}, {12, 16}},
		},
		{
			name:        "gộp tất cả intervals",
			intervals:   [][]int{{1, 2}, {3, 5}, {6, 7}},
			newInterval: []int{0, 10},
			want:        [][]int{{0, 10}},
		},
		{
			name:        "interval chạm đầu",
			intervals:   [][]int{{1, 3}, {6, 9}},
			newInterval: []int{3, 5},
			want:        [][]int{{1, 5}, {6, 9}},
		},
		{
			name:        "interval chạm cuối",
			intervals:   [][]int{{1, 3}, {6, 9}},
			newInterval: []int{4, 6},
			want:        [][]int{{1, 3}, {4, 9}},
		},
		{
			name:        "interval chạm cả đầu và cuối",
			intervals:   [][]int{{1, 3}, {6, 9}},
			newInterval: []int{3, 6},
			want:        [][]int{{1, 9}},
		},
		{
			name:        "interval nằm hoàn toàn trong một interval",
			intervals:   [][]int{{1, 5}, {10, 15}},
			newInterval: []int{2, 4},
			want:        [][]int{{1, 5}, {10, 15}},
		},
		{
			name:        "interval bao phủ một interval",
			intervals:   [][]int{{1, 3}, {6, 9}},
			newInterval: []int{0, 5},
			want:        [][]int{{0, 5}, {6, 9}},
		},
		{
			name:        "interval bao phủ nhiều intervals",
			intervals:   [][]int{{1, 2}, {3, 4}, {5, 6}},
			newInterval: []int{0, 7},
			want:        [][]int{{0, 7}},
		},
		{
			name:        "một interval duy nhất",
			intervals:   [][]int{{1, 3}},
			newInterval: []int{2, 5},
			want:        [][]int{{1, 5}},
		},
		{
			name:        "một interval không chồng chéo",
			intervals:   [][]int{{1, 3}},
			newInterval: []int{5, 7},
			want:        [][]int{{1, 3}, {5, 7}},
		},
		{
			name:        "interval với giá trị 0",
			intervals:   [][]int{{0, 1}, {3, 4}},
			newInterval: []int{1, 3},
			want:        [][]int{{0, 4}},
		},
		{
			name:        "interval tại 0",
			intervals:   [][]int{{1, 3}},
			newInterval: []int{0, 0},
			want:        [][]int{{0, 0}, {1, 3}},
		},
		{
			name:        "interval lớn",
			intervals:   [][]int{{1, 2}, {100, 200}},
			newInterval: []int{50, 150},
			want:        [][]int{{1, 2}, {50, 200}},
		},
		{
			name:        "gộp với interval đầu tiên",
			intervals:   [][]int{{1, 5}, {10, 15}},
			newInterval: []int{0, 3},
			want:        [][]int{{0, 5}, {10, 15}},
		},
		{
			name:        "gộp với interval cuối cùng",
			intervals:   [][]int{{1, 5}, {10, 15}},
			newInterval: []int{12, 20},
			want:        [][]int{{1, 5}, {10, 20}},
		},
		{
			name:        "interval giống hệt một interval",
			intervals:   [][]int{{1, 3}, {6, 9}},
			newInterval: []int{1, 3},
			want:        [][]int{{1, 3}, {6, 9}},
		},
		{
			name:        "interval chạm tại điểm đầu",
			intervals:   [][]int{{1, 3}, {6, 9}},
			newInterval: []int{3, 4},
			want:        [][]int{{1, 4}, {6, 9}},
		},
		{
			name:        "interval chạm tại điểm cuối",
			intervals:   [][]int{{1, 3}, {6, 9}},
			newInterval: []int{4, 6},
			want:        [][]int{{1, 3}, {4, 9}},
		},
		{
			name:        "nhiều intervals liên tiếp",
			intervals:   [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}},
			newInterval: []int{2, 7},
			want:        [][]int{{1, 8}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("Input intervals: %v, newInterval: %v", tt.intervals, tt.newInterval)

			got := insert(tt.intervals, tt.newInterval)

			t.Logf("Output: %v", got)
			t.Logf("Expected: %v", tt.want)

			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf(
					"insert(%v, %v) = %v; want %v",
					tt.intervals, tt.newInterval, got, tt.want,
				)
			}
		})
	}
}
