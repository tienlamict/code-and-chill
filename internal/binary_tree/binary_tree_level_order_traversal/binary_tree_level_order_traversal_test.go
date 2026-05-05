package binary_tree

import (
	"reflect"
	"testing"
)

// buildTree tạo binary tree từ mảng level-order (nil = vắng mặt)
func buildTree(vals []interface{}) *TreeNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}

	root := &TreeNode{Val: vals[0].(int)}
	queue := []*TreeNode{root}
	index := 1

	for len(queue) > 0 && index < len(vals) {
		node := queue[0]
		queue = queue[1:]

		if index < len(vals) && vals[index] != nil {
			node.Left = &TreeNode{Val: vals[index].(int)}
			queue = append(queue, node.Left)
		}
		index++

		if index < len(vals) && vals[index] != nil {
			node.Right = &TreeNode{Val: vals[index].(int)}
			queue = append(queue, node.Right)
		}
		index++
	}

	return root
}

func Test_LevelOrder(t *testing.T) {
	tests := []struct {
		name string
		root []interface{}
		want [][]int
	}{
		{
			name: "example 1 - balanced tree",
			root: []interface{}{3, 9, 20, nil, nil, 15, 7},
			want: [][]int{{3}, {9, 20}, {15, 7}},
		},
		{
			name: "example 2 - single node",
			root: []interface{}{1},
			want: [][]int{{1}},
		},
		{
			name: "example 3 - empty tree",
			root: []interface{}{},
			want: [][]int{},
		},
		{
			name: "left-skewed tree",
			root: []interface{}{1, 2, nil, 3, nil},
			want: [][]int{{1}, {2}, {3}},
		},
		{
			name: "right-skewed tree",
			root: []interface{}{1, nil, 2, nil, 3},
			want: [][]int{{1}, {2}, {3}},
		},
		{
			name: "negative values",
			root: []interface{}{-1, -2, -3},
			want: [][]int{{-1}, {-2, -3}},
		},
		{
			name: "complete binary tree",
			root: []interface{}{1, 2, 3, 4, 5, 6, 7},
			want: [][]int{{1}, {2, 3}, {4, 5, 6, 7}},
		},
		{
			name: "tree with duplicate values",
			root: []interface{}{1, 1, 1, 1, 1},
			want: [][]int{{1}, {1, 1}, {1, 1}},
		},
		{
			name: "boundary values - max",
			root: []interface{}{1000, -1000, 1000},
			want: [][]int{{1000}, {-1000, 1000}},
		},
		{
			name: "two levels asymmetric",
			root: []interface{}{1, 2, 3, 4, nil, nil, 5},
			want: [][]int{{1}, {2, 3}, {4, 5}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTree(tt.root)
			got := levelOrder(root)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder(%v) = %v; want %v", tt.root, got, tt.want)
			}
		})
	}
}
