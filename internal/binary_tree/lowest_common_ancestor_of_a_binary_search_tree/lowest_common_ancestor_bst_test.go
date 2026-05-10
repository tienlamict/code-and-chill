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

// findNode tìm một node có giá trị val trong cây.
func findNode(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	left := findNode(root.Left, val)
	if left != nil {
		return left
	}
	return findNode(root.Right, val)
}

func TestLowestCommonAncestor(t *testing.T) {
	tests := []struct {
		name     string
		rootVals []interface{}
		pVal     int
		qVal     int
		wantVal  int
	}{
		{
			name:     "Example 1",
			rootVals: []interface{}{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5},
			pVal:     2,
			qVal:     8,
			wantVal:  6,
		},
		{
			name:     "Example 2",
			rootVals: []interface{}{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5},
			pVal:     2,
			qVal:     4,
			wantVal:  2,
		},
		{
			name:     "Example 3",
			rootVals: []interface{}{2, 1},
			pVal:     2,
			qVal:     1,
			wantVal:  2,
		},
		{
			name:     "p and q are in the right subtree",
			rootVals: []interface{}{6, 2, 8, 0, 4, 7, 9},
			pVal:     7,
			qVal:     9,
			wantVal:  8,
		},
		{
			name:     "p and q are in the left subtree",
			rootVals: []interface{}{6, 2, 8, 0, 4, 7, 9},
			pVal:     0,
			qVal:     4,
			wantVal:  2,
		},
		{
			name:     "Negative values",
			rootVals: []interface{}{-5, -8, -2, -10, -6, -4, -1},
			pVal:     -10,
			qVal:     -6,
			wantVal:  -8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTree(tt.rootVals)
			p := findNode(root, tt.pVal)
			q := findNode(root, tt.qVal)
			got := lowestCommonAncestor(root, p, q)
			if got == nil || got.Val != tt.wantVal {
				gotVal := "nil"
				if got != nil {
					gotVal = string(rune(got.Val)) // wait, this is wrong for int
				}
				_ = gotVal // fix below
				if got == nil {
					t.Errorf("lowestCommonAncestor() = nil; want node with val %v", tt.wantVal)
				} else {
					t.Errorf("lowestCommonAncestor() = %v; want %v", got.Val, tt.wantVal)
				}
			}
		})
	}
}
