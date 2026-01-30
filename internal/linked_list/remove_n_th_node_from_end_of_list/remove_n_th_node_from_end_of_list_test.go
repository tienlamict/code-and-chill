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

func Test_RemoveNthFromEnd(t *testing.T) {
	tests := []struct {
		name  string
		vals  []int
		n     int
		want  []int
	}{
		{
			name: "example 1 - remove 2nd from end",
			vals: []int{1, 2, 3, 4, 5},
			n:    2,
			want: []int{1, 2, 3, 5},
		},
		{
			name: "example 2 - remove head (single node)",
			vals: []int{1},
			n:    1,
			want: []int{},
		},
		{
			name: "example 3 - remove last node",
			vals: []int{1, 2},
			n:    1,
			want: []int{1},
		},
		{
			name: "remove head (multiple nodes)",
			vals: []int{1, 2, 3, 4, 5},
			n:    5,
			want: []int{2, 3, 4, 5},
		},
		{
			name: "remove middle node",
			vals: []int{1, 2, 3, 4, 5},
			n:    3,
			want: []int{1, 2, 4, 5},
		},
		{
			name: "remove second from end (3 nodes)",
			vals: []int{1, 2, 3},
			n:    2,
			want: []int{1, 3},
		},
		{
			name: "remove first from end (3 nodes)",
			vals: []int{1, 2, 3},
			n:    1,
			want: []int{1, 2},
		},
		{
			name: "remove third from end (3 nodes)",
			vals: []int{1, 2, 3},
			n:    3,
			want: []int{2, 3},
		},
		{
			name: "long list - remove 4th from end",
			vals: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			n:    4,
			want: []int{1, 2, 3, 4, 5, 6, 8, 9, 10},
		},
		{
			name: "two nodes - remove head",
			vals: []int{1, 2},
			n:    2,
			want: []int{2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := createList(tt.vals)

			t.Logf("Input: list = %v, n = %d", tt.vals, tt.n)

			got := removeNthFromEnd(head, tt.n)
			gotSlice := listToSlice(got)

			t.Logf("Output: %v", gotSlice)

			if !reflect.DeepEqual(gotSlice, tt.want) {
				t.Fatalf(
					"removeNthFromEnd(%v, n=%d) = %v; want %v",
					tt.vals, tt.n, gotSlice, tt.want,
				)
			}
		})
	}
}

