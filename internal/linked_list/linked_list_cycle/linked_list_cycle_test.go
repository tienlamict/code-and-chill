package linked_list

import "testing"

// createCycleList tạo một linked list với cycle tại vị trí pos
// Nếu pos = -1, tạo linked list không có cycle
func createCycleList(vals []int, pos int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	// Tạo các nodes
	nodes := make([]*ListNode, len(vals))
	for i, val := range vals {
		nodes[i] = &ListNode{Val: val}
	}

	// Kết nối các nodes
	for i := 0; i < len(vals)-1; i++ {
		nodes[i].Next = nodes[i+1]
	}

	// Tạo cycle nếu pos >= 0
	if pos >= 0 && pos < len(vals) {
		nodes[len(vals)-1].Next = nodes[pos]
	}

	return nodes[0]
}

func Test_HasCycle(t *testing.T) {

	tests := []struct {
		name string
		vals []int
		pos  int
		want bool
	}{
		{
			name: "example 1 - cycle at index 1",
			vals: []int{3, 2, 0, -4},
			pos:  1,
			want: true,
		},
		{
			name: "example 2 - cycle at index 0",
			vals: []int{1, 2},
			pos:  0,
			want: true,
		},
		{
			name: "example 3 - no cycle",
			vals: []int{1},
			pos:  -1,
			want: false,
		},
		{
			name: "empty list",
			vals: []int{},
			pos:  -1,
			want: false,
		},
		{
			name: "two nodes no cycle",
			vals: []int{1, 2},
			pos:  -1,
			want: false,
		},
		{
			name: "two nodes with cycle",
			vals: []int{1, 2},
			pos:  0,
			want: true,
		},
		{
			name: "three nodes cycle at tail",
			vals: []int{1, 2, 3},
			pos:  2,
			want: true,
		},
		{
			name: "three nodes cycle at head",
			vals: []int{1, 2, 3},
			pos:  0,
			want: true,
		},
		{
			name: "long list no cycle",
			vals: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			pos:  -1,
			want: false,
		},
		{
			name: "long list with cycle",
			vals: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			pos:  5,
			want: true,
		},
		{
			name: "single node self cycle",
			vals: []int{1},
			pos:  0,
			want: true,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			head := createCycleList(tt.vals, tt.pos)

			t.Logf("Input: vals = %v, pos = %d", tt.vals, tt.pos)

			got := hasCycle(head)

			t.Logf("Output result: %v", got)

			if got != tt.want {

				t.Fatalf(
					"hasCycle(%v, pos=%d) = %v; want %v",
					tt.vals, tt.pos, got, tt.want,
				)

			}

		})

	}

}
