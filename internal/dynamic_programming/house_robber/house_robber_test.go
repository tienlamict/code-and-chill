package house_robber

import "testing"

func Test_rob(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "example 1 - [1,2,3,1]",
			nums: []int{1, 2, 3, 1},
			want: 4,
		},
		{
			name: "example 2 - [2,7,9,3,1]",
			nums: []int{2, 7, 9, 3, 1},
			want: 12,
		},
		{
			name: "một nhà",
			nums: []int{5},
			want: 5,
		},
		{
			name: "hai nhà - lấy nhà nhiều tiền hơn",
			nums: []int{2, 7},
			want: 7,
		},
		{
			name: "hai nhà bằng nhau",
			nums: []int{3, 3},
			want: 3,
		},
		{
			name: "ba nhà - trộm nhà 0 và 2",
			nums: []int{1, 2, 3},
			want: 4,
		},
		{
			name: "trường hợp biên - toàn 0",
			nums: []int{0, 0, 0, 0},
			want: 0,
		},
		{
			name: "phần tử trùng lặp",
			nums: []int{2, 2, 2, 2},
			want: 4,
		},
		{
			name: "số lớn - gần constraint max 400",
			nums: []int{400, 400, 400},
			want: 800,
		},
		{
			name: "đan xen - chọn mọi nhà chẵn",
			nums: []int{1, 0, 1, 0, 1},
			want: 3,
		},
		{
			name: "nhiều nhà - kiểm tra lựa chọn tối ưu",
			nums: []int{2, 1, 1, 2},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rob(tt.nums)
			if got != tt.want {
				t.Errorf("rob(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

func Test_rob_empty(t *testing.T) {
	// Mảng rỗng (ngoài constraint 1<=len nhưng đảm bảo code an toàn)
	got := rob([]int{})
	if got != 0 {
		t.Errorf("rob([]) = %d, want 0", got)
	}
}
