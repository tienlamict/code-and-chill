package linked_list

import (
	"reflect"
	"testing"
)

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

func listToSlice(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	var out []int
	for head != nil {
		out = append(out, head.Val)
		head = head.Next
	}
	return out
}

func Test_reverseList(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want []int
	}{
		{name: "example 1", vals: []int{1, 2, 3, 4, 5}, want: []int{5, 4, 3, 2, 1}},
		{name: "example 2", vals: []int{1, 2}, want: []int{2, 1}},
		{name: "example 3 empty", vals: []int{}, want: []int{}},
		{name: "single node", vals: []int{42}, want: []int{42}},
		{name: "two same values", vals: []int{0, 0}, want: []int{0, 0}},
		{name: "negative values", vals: []int{-1, -2, -3}, want: []int{-3, -2, -1}},
		{name: "min constraint value", vals: []int{-5000, 5000}, want: []int{5000, -5000}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := createList(tt.vals)
			got := reverseList(head)
			if !reflect.DeepEqual(listToSlice(got), tt.want) {
				t.Fatalf("reverseList(...) = %v; want %v", listToSlice(got), tt.want)
			}
		})
	}
}

func Test_reverseListRecursive_matches_iterative(t *testing.T) {
	tests := [][]int{
		{},
		{1},
		{1, 2},
		{1, 2, 3, 4, 5},
		{9, 8, 7},
	}
	for _, vals := range tests {
		h1 := createList(vals)
		h2 := createList(vals)
		iter := reverseList(h1)
		rec := reverseListRecursive(h2)
		if !reflect.DeepEqual(listToSlice(iter), listToSlice(rec)) {
			t.Fatalf("vals=%v iterative %v != recursive %v", vals, listToSlice(iter), listToSlice(rec))
		}
	}
}
