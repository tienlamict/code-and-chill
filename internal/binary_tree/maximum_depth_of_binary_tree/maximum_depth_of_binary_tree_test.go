package binary_tree

import "testing"

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

func Test_MaxDepth(t *testing.T) {
	tests := []struct {
		name string
		root []interface{}
		want int
	}{
		{
			name: "example 1 - balanced tree",
			root: []interface{}{3, 9, 20, nil, nil, 15, 7},
			want: 3,
		},
		{
			name: "example 2 - right skewed",
			root: []interface{}{1, nil, 2},
			want: 2,
		},
		{
			name: "empty tree",
			root: []interface{}{},
			want: 0,
		},
		{
			name: "single node",
			root: []interface{}{1},
			want: 1,
		},
		{
			name: "left skewed tree depth 4",
			root: []interface{}{1, 2, nil, 3, nil, 4},
			want: 4,
		},
		{
			name: "perfect binary tree depth 4",
			root: []interface{}{1, 2, 3, 4, 5, 6, 7},
			want: 3,
		},
		{
			name: "negative values",
			root: []interface{}{-100, -50, -50, -25, -25, -25, -25},
			want: 3,
		},
		{
			name: "deep right skewed depth 5",
			root: []interface{}{1, nil, 2, nil, 3, nil, 4, nil, 5},
			want: 5,
		},
		{
			name: "two level tree",
			root: []interface{}{1, 2, 3},
			want: 2,
		},
		{
			name: "unbalanced - left deeper",
			root: []interface{}{1, 2, 3, 4, nil, nil, nil, 5},
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTree(tt.root)
			got := maxDepth(root)
			if got != tt.want {
				t.Errorf("maxDepth(%v) = %d; want %d", tt.root, got, tt.want)
			}
		})
	}
}
