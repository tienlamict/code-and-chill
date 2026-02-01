package linked_list

import (
	"reflect"
	"testing"
)

// createList tạo một linked list từ mảng các giá trị
func createList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	head := &ListNode{Val: vals[0]}
	current := head

	for i := 1; i < len(vals); i++ {
		current.Next = &ListNode{Val: vals[i]}
		current = current.Next
	}

	return head
}

// listToSlice chuyển đổi linked list thành mảng để so sánh
func listToSlice(head *ListNode) []int {
	result := make([]int, 0)
	current := head

	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}

	return result
}

func Test_MergeTwoLists(t *testing.T) {
	tests := []struct {
		name   string
		list1  []int
		list2  []int
		want   []int
	}{
		{
			name:  "example 1 - both lists have values",
			list1: []int{1, 2, 4},
			list2: []int{1, 3, 4},
			want:  []int{1, 1, 2, 3, 4, 4},
		},
		{
			name:  "example 2 - both lists are empty",
			list1: []int{},
			list2: []int{},
			want:  []int{},
		},
		{
			name:  "example 3 - list1 is empty",
			list1: []int{},
			list2: []int{0},
			want:  []int{0},
		},
		{
			name:  "list2 is empty",
			list1: []int{1, 2, 3},
			list2: []int{},
			want:  []int{1, 2, 3},
		},
		{
			name:  "list1 has smaller values",
			list1: []int{1, 2, 3},
			list2: []int{4, 5, 6},
			want:  []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:  "list2 has smaller values",
			list1: []int{4, 5, 6},
			list2: []int{1, 2, 3},
			want:  []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:  "alternating values",
			list1: []int{1, 3, 5},
			list2: []int{2, 4, 6},
			want:  []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:  "list1 is longer",
			list1: []int{1, 2, 3, 4, 5},
			list2: []int{1, 2},
			want:  []int{1, 1, 2, 2, 3, 4, 5},
		},
		{
			name:  "list2 is longer",
			list1: []int{1, 2},
			list2: []int{1, 2, 3, 4, 5},
			want:  []int{1, 1, 2, 2, 3, 4, 5},
		},
		{
			name:  "single node in each list",
			list1: []int{1},
			list2: []int{2},
			want:  []int{1, 2},
		},
		{
			name:  "single node - list1 smaller",
			list1: []int{1},
			list2: []int{2},
			want:  []int{1, 2},
		},
		{
			name:  "single node - list2 smaller",
			list1: []int{2},
			list2: []int{1},
			want:  []int{1, 2},
		},
		{
			name:  "duplicate values",
			list1: []int{1, 1, 1},
			list2: []int{1, 1, 1},
			want:  []int{1, 1, 1, 1, 1, 1},
		},
		{
			name:  "negative values",
			list1: []int{-5, -3, -1},
			list2: []int{-4, -2, 0},
			want:  []int{-5, -4, -3, -2, -1, 0},
		},
		{
			name:  "mixed positive and negative",
			list1: []int{-1, 0, 1},
			list2: []int{-2, 2},
			want:  []int{-2, -1, 0, 1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list1 := createList(tt.list1)
			list2 := createList(tt.list2)

			t.Logf("Input: list1 = %v, list2 = %v", tt.list1, tt.list2)

			got := mergeTwoLists(list1, list2)
			gotSlice := listToSlice(got)

			t.Logf("Output: %v", gotSlice)

			if !reflect.DeepEqual(gotSlice, tt.want) {
				t.Fatalf(
					"mergeTwoLists(%v, %v) = %v; want %v",
					tt.list1, tt.list2, gotSlice, tt.want,
				)
			}
		})
	}
}
