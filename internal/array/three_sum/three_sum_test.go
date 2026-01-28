package array

import "testing"

// equalSlices kiểm tra xem hai slice có chứa các phần tử giống nhau không (không quan tâm thứ tự)
func equalSlices(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	// Tạo map để đếm số lần xuất hiện của mỗi triplet trong a
	countA := make(map[[3]int]int)
	for _, triplet := range a {
		if len(triplet) != 3 {
			return false
		}
		key := [3]int{triplet[0], triplet[1], triplet[2]}
		countA[key]++
	}

	// Kiểm tra các triplet trong b
	countB := make(map[[3]int]int)
	for _, triplet := range b {
		if len(triplet) != 3 {
			return false
		}
		key := [3]int{triplet[0], triplet[1], triplet[2]}
		countB[key]++
	}

	// So sánh hai map
	if len(countA) != len(countB) {
		return false
	}

	for key, count := range countA {
		if countB[key] != count {
			return false
		}
	}

	return true
}

func Test_ThreeSum(t *testing.T) {

	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "example 1",
			nums: []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name: "example 2",
			nums: []int{0, 1, 1},
			want: [][]int{},
		},
		{
			name: "example 3",
			nums: []int{0, 0, 0},
			want: [][]int{{0, 0, 0}},
		},
		{
			name: "empty array",
			nums: []int{},
			want: [][]int{},
		},
		{
			name: "less than 3 elements",
			nums: []int{1, 2},
			want: [][]int{},
		},
		{
			name: "no solution",
			nums: []int{1, 2, 3},
			want: [][]int{},
		},
		{
			name: "all zeros",
			nums: []int{0, 0, 0, 0},
			want: [][]int{{0, 0, 0}},
		},
		{
			name: "multiple solutions",
			nums: []int{-2, 0, 1, 1, 2},
			want: [][]int{{-2, 0, 2}, {-2, 1, 1}},
		},
		{
			name: "duplicate triplets",
			nums: []int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4},
			want: [][]int{{-4, 0, 4}, {-4, 1, 3}, {-3, -1, 4}, {-3, 0, 3}, {-3, 1, 2}, {-2, -1, 3}, {-2, 0, 2}, {-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name: "all negative",
			nums: []int{-1, -2, -3},
			want: [][]int{},
		},
		{
			name: "all positive",
			nums: []int{1, 2, 3},
			want: [][]int{},
		},
		{
			name: "mixed with duplicates",
			nums: []int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4},
			want: [][]int{{-4, 0, 4}, {-4, 1, 3}, {-3, -1, 4}, {-3, 0, 3}, {-3, 1, 2}, {-2, -1, 3}, {-2, 0, 2}, {-1, -1, 2}, {-1, 0, 1}},
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			t.Logf("Input nums: %v", tt.nums)

			got := threeSum(tt.nums)

			t.Logf("Output result: %v", got)
			t.Logf("Expected result: %v", tt.want)

			if !equalSlices(got, tt.want) {

				t.Fatalf(
					"threeSum(%v) = %v; want %v",
					tt.nums, got, tt.want,
				)

			}

		})

	}

}
