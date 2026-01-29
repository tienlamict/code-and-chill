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

func Test_ReorderList(t *testing.T) {

	tests := []struct {
		name string
		vals []int
		want []int
	}{
		{
			name: "example 1",
			vals: []int{1, 2, 3, 4},
			want: []int{1, 4, 2, 3},
		},
		{
			name: "example 2",
			vals: []int{1, 2, 3, 4, 5},
			want: []int{1, 5, 2, 4, 3},
		},
		{
			name: "single node",
			vals: []int{1},
			want: []int{1},
		},
		{
			name: "two nodes",
			vals: []int{1, 2},
			want: []int{1, 2},
		},
		{
			name: "three nodes",
			vals: []int{1, 2, 3},
			want: []int{1, 3, 2},
		},
		{
			name: "six nodes",
			vals: []int{1, 2, 3, 4, 5, 6},
			want: []int{1, 6, 2, 5, 3, 4},
		},
		{
			name: "seven nodes",
			vals: []int{1, 2, 3, 4, 5, 6, 7},
			want: []int{1, 7, 2, 6, 3, 5, 4},
		},
		{
			name: "eight nodes",
			vals: []int{1, 2, 3, 4, 5, 6, 7, 8},
			want: []int{1, 8, 2, 7, 3, 6, 4, 5},
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			head := createList(tt.vals)

			t.Logf("Input: %v", tt.vals)

			reorderList(head)

			got := listToSlice(head)

			t.Logf("Output: %v", got)
			t.Logf("Expected: %v", tt.want)

			if len(got) != len(tt.want) {
				t.Fatalf(
					"reorderList(%v) length = %d; want %d",
					tt.vals, len(got), len(tt.want),
				)
			}

			for i := 0; i < len(got); i++ {
				if got[i] != tt.want[i] {
					t.Fatalf(
						"reorderList(%v) = %v; want %v (mismatch at index %d: got %d, want %d)",
						tt.vals, got, tt.want, i, got[i], tt.want[i],
					)
				}
			}

		})

	}

}
