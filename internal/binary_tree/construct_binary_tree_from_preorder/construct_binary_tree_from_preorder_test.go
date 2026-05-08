package construct_binary_tree_from_preorder

import (
	"testing"
)

// treeToSlice chuyển cây nhị phân thành slice theo level-order (BFS), dùng để so sánh kết quả.
func treeToSlice(root *TreeNode) []interface{} {
	if root == nil {
		return []interface{}{}
	}

	result := []interface{}{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			result = append(result, nil)
		} else {
			result = append(result, node.Val)
			queue = append(queue, node.Left, node.Right)
		}
	}

	// Loại bỏ các nil ở cuối
	for len(result) > 0 && result[len(result)-1] == nil {
		result = result[:len(result)-1]
	}

	return result
}

// sliceEqual so sánh hai slice interface{}.
func sliceEqual(a, b []interface{}) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestBuildTree(t *testing.T) {
	tests := []struct {
		name     string
		preorder []int
		inorder  []int
		expected []interface{}
	}{
		{
			name:     "Example 1",
			preorder: []int{3, 9, 20, 15, 7},
			inorder:  []int{9, 3, 15, 20, 7},
			expected: []interface{}{3, 9, 20, nil, nil, 15, 7},
		},
		{
			name:     "Example 2 - single node negative",
			preorder: []int{-1},
			inorder:  []int{-1},
			expected: []interface{}{-1},
		},
		{
			name:     "Single node",
			preorder: []int{1},
			inorder:  []int{1},
			expected: []interface{}{1},
		},
		{
			name:     "Left skewed tree",
			preorder: []int{3, 2, 1},
			inorder:  []int{1, 2, 3},
			expected: []interface{}{3, 2, nil, 1},
		},
		{
			name:     "Right skewed tree",
			preorder: []int{1, 2, 3},
			inorder:  []int{1, 2, 3},
			expected: []interface{}{1, nil, 2, nil, 3},
		},
		{
			name:     "Perfect binary tree",
			preorder: []int{1, 2, 4, 5, 3, 6, 7},
			inorder:  []int{4, 2, 5, 1, 6, 3, 7},
			expected: []interface{}{1, 2, 3, 4, 5, 6, 7},
		},
		{
			name:     "Root only right child",
			preorder: []int{1, 2},
			inorder:  []int{1, 2},
			expected: []interface{}{1, nil, 2},
		},
		{
			name:     "Root only left child",
			preorder: []int{1, 2},
			inorder:  []int{2, 1},
			expected: []interface{}{1, 2},
		},
		{
			name:     "Negative values",
			preorder: []int{-3, -9, -20},
			inorder:  []int{-9, -3, -20},
			expected: []interface{}{-3, -9, -20},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildTree(tt.preorder, tt.inorder)
			gotSlice := treeToSlice(got)
			if !sliceEqual(gotSlice, tt.expected) {
				t.Errorf("buildTree() = %v, want %v", gotSlice, tt.expected)
			}
		})
	}
}
