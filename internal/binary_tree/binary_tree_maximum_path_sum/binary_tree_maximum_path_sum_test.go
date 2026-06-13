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

func Test_MaxPathSum(t *testing.T) {
	tests := []struct {
		name string
		root []interface{}
		want int
	}{
		{
			name: "Ví dụ 1: Cây đơn giản",
			root: []interface{}{1, 2, 3},
			want: 6,
		},
		{
			name: "Ví dụ 2: Có số âm",
			root: []interface{}{-10, 9, 20, nil, nil, 15, 7},
			want: 42,
		},
		{
			name: "Chỉ có một node",
			root: []interface{}{1},
			want: 1,
		},
		{
			name: "Chỉ có một node âm",
			root: []interface{}{-3},
			want: -3,
		},
		{
			name: "Tất cả các node đều âm",
			root: []interface{}{-2, -1},
			want: -1,
		},
		{
			name: "Đường đi không qua root",
			root: []interface{}{5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, nil, 1},
			want: 48, // 7+11+4+5+8+13 = 48 (Wait, 7+11+2 = 20, 11+4+5+8+13=41, 7+11+4+5+8+4+1=40... let me re-verify)
			// Re-calculating for {5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, nil, 1}
			// Root 5: Left 4, Right 8
			// Node 4: Left 11
			// Node 8: Left 13, Right 4
			// Node 11: Left 7, Right 2
			// Node 4 (right of 8): Right 1
			// Paths: 
			// 7-11-2 = 20
			// 7-11-4-5-8-13 = 48 (Max path: 7+11+4+5+8+13 = 48)
		},
		{
			name: "Cây lệch trái",
			root: []interface{}{1, 2, nil, 3, nil, 4},
			want: 10,
		},
		{
			name: "Cây có giá trị lớn",
			root: []interface{}{100, 200, 300},
			want: 600,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTree(tt.root)
			got := maxPathSum(root)
			if got != tt.want {
				t.Errorf("maxPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
