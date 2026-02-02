package linked_list

import "testing"

// createList tạo một linked list từ mảng các giá trị
func createList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	head := &ListNode{Val: vals[0]}
	curr := head
	for i := 1; i < len(vals); i++ {
		curr.Next = &ListNode{Val: vals[i]}
		curr = curr.Next
	}

	return head
}

// listToSlice chuyển linked list thành mảng để so sánh
func listToSlice(head *ListNode) []int {
	result := []int{}
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

// equalSlices kiểm tra xem hai slice có bằng nhau không
func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func Test_MergeKLists(t *testing.T) {

	tests := []struct {
		name  string
		lists [][]int
		want  []int
	}{
		{
			name:  "example 1",
			lists: [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}},
			want:  []int{1, 1, 2, 3, 4, 4, 5, 6},
		},
		{
			name:  "example 2",
			lists: [][]int{},
			want:  []int{},
		},
		{
			name:  "example 3",
			lists: [][]int{{}},
			want:  []int{},
		},
		{
			name:  "single list",
			lists: [][]int{{1, 2, 3}},
			want:  []int{1, 2, 3},
		},
		{
			name:  "two lists",
			lists: [][]int{{1, 3, 5}, {2, 4, 6}},
			want:  []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:  "three lists with empty",
			lists: [][]int{{1, 2}, {}, {3, 4}},
			want:  []int{1, 2, 3, 4},
		},
		{
			name:  "all empty lists",
			lists: [][]int{{}, {}, {}},
			want:  []int{},
		},
		{
			name:  "single element lists",
			lists: [][]int{{1}, {2}, {3}},
			want:  []int{1, 2, 3},
		},
		{
			name:  "four lists",
			lists: [][]int{{1, 5}, {2, 6}, {3, 7}, {4, 8}},
			want:  []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			name:  "lists with duplicates",
			lists: [][]int{{1, 1, 2}, {1, 2, 2}, {2, 3, 3}},
			want:  []int{1, 1, 1, 2, 2, 2, 2, 3, 3},
		},
		{
			name:  "one list empty",
			lists: [][]int{{1, 2, 3}, {}},
			want:  []int{1, 2, 3},
		},
		{
			name:  "negative numbers",
			lists: [][]int{{-1, 0, 1}, {-2, 2}, {-3, 3}},
			want:  []int{-3, -2, -1, 0, 1, 2, 3},
		},
		{
			name:  "five lists",
			lists: [][]int{{1}, {2}, {3}, {4}, {5}},
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			name:  "different lengths",
			lists: [][]int{{1, 5, 9}, {2, 3}, {4, 6, 7, 8}},
			want:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			// Tạo mảng các linked lists
			lists := make([]*ListNode, len(tt.lists))
			for i, vals := range tt.lists {
				lists[i] = createList(vals)
			}

			t.Logf("Input lists: %v", tt.lists)

			got := mergeKLists(lists)

			gotSlice := listToSlice(got)

			t.Logf("Output: %v", gotSlice)
			t.Logf("Expected: %v", tt.want)

			if !equalSlices(gotSlice, tt.want) {

				t.Fatalf(
					"mergeKLists(%v) = %v; want %v",
					tt.lists, gotSlice, tt.want,
				)

			}

		})

	}

}
