package array

import (
	"reflect"
	"testing"
)

func Test_Merge(t *testing.T) {
	tests := []struct {
		name       string
		intervals  [][]int
		want       [][]int
	}{
		{
			name:       "example 1",
			intervals:  [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			want:       [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			name:       "example 2 - intervals chạm nhau",
			intervals:  [][]int{{1, 4}, {4, 5}},
			want:       [][]int{{1, 5}},
		},
		{
			name:       "example 3 - intervals chưa sắp xếp",
			intervals:  [][]int{{4, 7}, {1, 4}},
			want:       [][]int{{1, 7}},
		},
		{
			name:       "một interval",
			intervals:  [][]int{{1, 3}},
			want:       [][]int{{1, 3}},
		},
		{
			name:       "hai intervals không chồng chéo",
			intervals:  [][]int{{1, 2}, {3, 4}},
			want:       [][]int{{1, 2}, {3, 4}},
		},
		{
			name:       "hai intervals chồng chéo hoàn toàn",
			intervals:  [][]int{{1, 4}, {2, 3}},
			want:       [][]int{{1, 4}},
		},
		{
			name:       "hai intervals chồng chéo một phần",
			intervals:  [][]int{{1, 3}, {2, 6}},
			want:       [][]int{{1, 6}},
		},
		{
			name:       "nhiều intervals gộp thành một",
			intervals:  [][]int{{1, 4}, {2, 5}, {3, 6}, {4, 7}},
			want:       [][]int{{1, 7}},
		},
		{
			name:       "intervals lồng nhau",
			intervals:  [][]int{{1, 10}, {2, 3}, {4, 5}, {6, 7}},
			want:       [][]int{{1, 10}},
		},
		{
			name:       "mảng rỗng",
			intervals:  [][]int{},
			want:       [][]int{},
		},
		{
			name:       "hai intervals giống nhau",
			intervals:  [][]int{{1, 3}, {1, 3}},
			want:       [][]int{{1, 3}},
		},
		{
			name:       "intervals chạm tại điểm",
			intervals:  [][]int{{1, 2}, {2, 3}, {3, 4}},
			want:       [][]int{{1, 4}},
		},
		{
			name:       "thứ tự ngược",
			intervals:  [][]int{{15, 18}, {8, 10}, {1, 3}, {2, 6}},
			want:       [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			name:       "interval nhỏ nằm trong interval lớn",
			intervals:  [][]int{{1, 5}, {2, 4}},
			want:       [][]int{{1, 5}},
		},
		{
			name:       "cùng start khác end",
			intervals:  [][]int{{1, 2}, {1, 4}, {1, 3}},
			want:       [][]int{{1, 4}},
		},
		{
			name:       "cùng end khác start",
			intervals:  [][]int{{1, 5}, {2, 5}, {3, 5}},
			want:       [][]int{{1, 5}},
		},
		{
			name:       "intervals với giá trị 0",
			intervals:  [][]int{{0, 1}, {1, 2}, {2, 3}},
			want:       [][]int{{0, 3}},
		},
		{
			name:       "interval đơn tại 0",
			intervals:  [][]int{{0, 0}},
			want:       [][]int{{0, 0}},
		},
		{
			name:       "giá trị lớn",
			intervals:  [][]int{{1, 10000}, {5000, 10000}, {9999, 10000}},
			want:       [][]int{{1, 10000}},
		},
		{
			name:       "nhiều nhóm không chồng chéo",
			intervals:  [][]int{{1, 2}, {5, 6}, {10, 11}, {15, 16}},
			want:       [][]int{{1, 2}, {5, 6}, {10, 11}, {15, 16}},
		},
		{
			name:       "gộp từng cặp",
			intervals:  [][]int{{1, 2}, {2, 3}, {5, 6}, {6, 7}},
			want:       [][]int{{1, 3}, {5, 7}},
		},
		{
			name:       "interval đầu tiên chứa tất cả",
			intervals:  [][]int{{1, 100}, {2, 3}, {4, 5}, {6, 7}},
			want:       [][]int{{1, 100}},
		},
		{
			name:       "interval cuối mở rộng",
			intervals:  [][]int{{1, 3}, {2, 6}, {5, 10}},
			want:       [][]int{{1, 10}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("Input intervals: %v", tt.intervals)

			got := merge(tt.intervals)

			t.Logf("Output: %v", got)
			t.Logf("Expected: %v", tt.want)

			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf(
					"merge(%v) = %v; want %v",
					tt.intervals, got, tt.want,
				)
			}
		})
	}
}
