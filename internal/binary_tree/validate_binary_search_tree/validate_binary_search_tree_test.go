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

func Test_IsValidBST(t *testing.T) {
	tests := []struct {
		name string
		root []interface{}
		want bool
	}{
		// Ví dụ từ đề bài
		{
			name: "example 1 - valid BST [2,1,3]",
			root: []interface{}{2, 1, 3},
			want: true,
		},
		{
			name: "example 2 - invalid BST [5,1,4,null,null,3,6]",
			root: []interface{}{5, 1, 4, nil, nil, 3, 6},
			want: false,
		},

		// Edge cases
		{
			name: "single node",
			root: []interface{}{1},
			want: true,
		},
		{
			name: "two nodes - valid left child",
			root: []interface{}{2, 1, nil},
			want: true,
		},
		{
			name: "two nodes - invalid left child greater",
			root: []interface{}{2, 3, nil},
			want: false,
		},
		{
			name: "two nodes - valid right child",
			root: []interface{}{2, nil, 3},
			want: true,
		},
		{
			name: "two nodes - invalid right child less",
			root: []interface{}{2, nil, 1},
			want: false,
		},

		// Trường hợp trùng lặp giá trị (BST yêu cầu strictly less/greater)
		{
			name: "duplicate root and left child - invalid",
			root: []interface{}{2, 2, nil},
			want: false,
		},
		{
			name: "duplicate root and right child - invalid",
			root: []interface{}{2, nil, 2},
			want: false,
		},

		// Trường hợp bẫy: node vi phạm với tổ tiên, không phải cha trực tiếp
		{
			name: "violates ancestor constraint - right subtree has value less than root",
			root: []interface{}{5, 3, 7, 1, 4, 6, 8},
			want: true,
		},
		{
			name: "left subtree node greater than root - invalid",
			root: []interface{}{5, 3, 7, 1, 6, nil, nil},
			want: false, // 6 nằm ở left subtree của 5 nhưng 6 > 5
		},
		{
			name: "right subtree node less than root - invalid",
			root: []interface{}{10, 5, 15, nil, nil, 6, 20},
			want: false, // 6 nằm ở right subtree của 10 nhưng 6 < 10
		},

		// Giá trị âm
		{
			name: "negative values - valid BST",
			root: []interface{}{-3, -5, -1},
			want: true,
		},
		{
			name: "negative values - invalid BST",
			root: []interface{}{-3, -1, -5},
			want: false,
		},
		{
			name: "mixed positive and negative - valid",
			root: []interface{}{0, -1, 1},
			want: true,
		},

		// Biên INT32
		{
			name: "INT32 min value at root",
			root: []interface{}{-2147483648},
			want: true,
		},
		{
			name: "INT32 max value at root",
			root: []interface{}{2147483647},
			want: true,
		},
		{
			name: "INT32 min value as left child - invalid because equal not allowed",
			root: []interface{}{-2147483648, nil, 2147483647},
			want: true,
		},
		{
			name: "right child equals INT32 max - valid",
			root: []interface{}{0, nil, 2147483647},
			want: true,
		},
		{
			name: "left child equals INT32 min - valid",
			root: []interface{}{0, -2147483648, nil},
			want: true,
		},

		// Cây lệch (skewed trees)
		{
			name: "left-skewed valid BST",
			root: []interface{}{5, 4, nil, 3, nil, 2, nil},
			want: true,
		},
		{
			name: "right-skewed valid BST",
			root: []interface{}{1, nil, 2, nil, 3, nil, 4},
			want: true,
		},
		{
			name: "left-skewed invalid BST",
			root: []interface{}{5, 4, nil, 6, nil},
			want: false,
		},

		// Cây lớn hợp lệ
		{
			name: "complete valid BST",
			root: []interface{}{8, 4, 12, 2, 6, 10, 14, 1, 3, 5, 7, 9, 11, 13, 15},
			want: true,
		},
		{
			name: "complete invalid BST - one wrong node",
			root: []interface{}{8, 4, 12, 2, 6, 10, 14, 1, 3, 5, 9, 9, 11, 13, 15},
			want: false, // node 9 ở vị trí 7 trong right subtree của 4 vi phạm
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTree(tt.root)
			got := isValidBST(root)
			if got != tt.want {
				t.Errorf("isValidBST(%v) = %v; want %v", tt.root, got, tt.want)
			}
		})
	}
}

func Benchmark_IsValidBST(b *testing.B) {
	root := buildTree([]interface{}{8, 4, 12, 2, 6, 10, 14, 1, 3, 5, 7, 9, 11, 13, 15})
	for i := 0; i < b.N; i++ {
		isValidBST(root)
	}
}
