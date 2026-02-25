package array

import "testing"

func Test_CanJump(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{
			name: "example 1 - có thể nhảy đến cuối",
			nums: []int{2, 3, 1, 1, 4},
			want: true,
		},
		{
			name: "example 2 - không thể nhảy đến cuối",
			nums: []int{3, 2, 1, 0, 4},
			want: false,
		},
		{
			name: "mảng một phần tử - đã ở cuối",
			nums: []int{0},
			want: true,
		},
		{
			name: "mảng hai phần tử - có thể nhảy",
			nums: []int{1, 0},
			want: true,
		},
		{
			name: "mảng hai phần tử - không thể nhảy",
			nums: []int{0, 1},
			want: false,
		},
		{
			name: "tất cả phần tử đều 0 - không thể nhảy",
			nums: []int{0, 0, 0, 0},
			want: false,
		},
		{
			name: "có thể nhảy ngay từ đầu",
			nums: []int{5, 0, 0, 0, 0},
			want: true,
		},
		{
			name: "nhảy từng bước một",
			nums: []int{1, 1, 1, 1, 1},
			want: true,
		},
		{
			name: "nhảy qua nhiều bước",
			nums: []int{2, 0, 0},
			want: true,
		},
		{
			name: "không thể nhảy qua số 0",
			nums: []int{1, 0, 2},
			want: false,
		},
		{
			name: "có thể nhảy qua số 0",
			nums: []int{2, 0, 1, 0},
			want: true,
		},
		{
			name: "mảng lớn - có thể nhảy",
			nums: []int{3, 2, 1, 0, 4, 2, 1, 0, 1},
			want: false,
		},
		{
			name: "mảng lớn - không thể nhảy",
			nums: []int{2, 1, 0, 0, 0, 0, 0},
			want: false,
		},
		{
			name: "nhảy lớn ở giữa",
			nums: []int{1, 2, 3, 4, 5},
			want: true,
		},
		{
			name: "số 0 ở vị trí cuối",
			nums: []int{2, 3, 1, 1, 0},
			want: true,
		},
		{
			name: "số 0 ở vị trí đầu",
			nums: []int{0, 2, 3},
			want: false,
		},
		{
			name: "mảng tăng dần",
			nums: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			want: true,
		},
		{
			name: "mảng giảm dần",
			nums: []int{5, 4, 3, 2, 1, 0},
			want: true,
		},
		{
			name: "mảng giảm dần không thể nhảy",
			nums: []int{5, 4, 3, 2, 1, 0, 0},
			want: false,
		},
		{
			name: "nhảy lớn ở cuối",
			nums: []int{1, 1, 1, 1, 1, 10},
			want: true,
		},
		{
			name: "mảng với nhiều số 0 liên tiếp",
			nums: []int{2, 0, 0, 0, 1},
			want: false,
		},
		{
			name: "không thể nhảy qua nhiều số 0",
			nums: []int{3, 0, 0, 0, 1},
			want: false,
		},
		{
			name: "có thể nhảy qua nhiều số 0",
			nums: []int{4, 0, 0, 0, 1},
			want: true,
		},
		{
			name: "mảng một phần tử lớn",
			nums: []int{100},
			want: true,
		},
		{
			name: "mảng với giá trị lớn",
			nums: []int{1, 100, 0, 0, 0},
			want: true,
		},
		{
			name: "mảng với giá trị lớn không thể nhảy",
			nums: []int{1, 100, 0, 0, 0, 0},
			want: true, // Có thể nhảy từ vị trí 1 (100) đến cuối
		},
		{
			name: "edge case - mảng rỗng (theo constraints không xảy ra nhưng test an toàn)",
			nums: []int{},
			want: true, // Mảng rỗng, coi như đã ở cuối
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("Input nums: %v", tt.nums)

			got := canJump(tt.nums)

			t.Logf("Output: %v", got)
			t.Logf("Expected: %v", tt.want)

			if got != tt.want {
				t.Fatalf(
					"canJump(%v) = %v; want %v",
					tt.nums, got, tt.want,
				)
			}
		})
	}
}
