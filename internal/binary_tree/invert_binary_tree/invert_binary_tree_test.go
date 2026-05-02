package binary_tree

import (
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

// treeToSlice chuyển đổi binary tree thành mảng level-order để so sánh và log
func treeToSlice(root *TreeNode) []interface{} {
	if root == nil {
		return []interface{}{}
	}

	var res []interface{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node != nil {
			res = append(res, node.Val)
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		} else {
			res = append(res, nil)
		}
	}

	// Cắt tỉa các giá trị nil ở cuối mảng để mảng gọn hơn
	for i := len(res) - 1; i >= 0; i-- {
		if res[i] == nil {
			res = res[:i]
		} else {
			break
		}
	}

	return res
}

// isSameTree kiểm tra xem hai cây có giống hệt nhau không
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func Test_InvertTree(t *testing.T) {
	tests := []struct {
		name string
		root []interface{}
		want []interface{}
	}{
		{
			name: "example 1",
			root: []interface{}{4, 2, 7, 1, 3, 6, 9},
			want: []interface{}{4, 7, 2, 9, 6, 3, 1},
		},
		{
			name: "example 2",
			root: []interface{}{2, 1, 3},
			want: []interface{}{2, 3, 1},
		},
		{
			name: "example 3 - empty tree",
			root: []interface{}{},
			want: []interface{}{},
		},
		{
			name: "single node",
			root: []interface{}{1},
			want: []interface{}{1},
		},
		{
			name: "tree with null nodes",
			root: []interface{}{1, 2, nil, 3, nil},
			want: []interface{}{1, nil, 2, nil, 3},
		},
		{
			name: "negative values",
			root: []interface{}{-1, -2, -3, -4, -5},
			want: []interface{}{-1, -3, -2, nil, nil, -5, -4},
		},
		{
			name: "asymmetric tree",
			root: []interface{}{1, 2, 3, nil, 4},
			want: []interface{}{1, 3, 2, nil, nil, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTree(tt.root)
			expected := buildTree(tt.want)
			got := invertTree(root)

			if !isSameTree(got, expected) {
				t.Errorf("invertTree(%v) = %v; want %v", tt.root, treeToSlice(got), tt.want)
			}
		})
	}
}
