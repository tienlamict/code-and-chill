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

func Test_KthSmallest(t *testing.T) {
	tests := []struct {
		name string
		root []interface{}
		k    int
		want int
	}{
		// Ví dụ từ đề bài
		{
			name: "example 1 - [3,1,4,null,2] k=1",
			root: []interface{}{3, 1, 4, nil, 2},
			k:    1,
			want: 1,
		},
		{
			name: "example 2 - [5,3,6,2,4,null,null,1] k=3",
			root: []interface{}{5, 3, 6, 2, 4, nil, nil, 1},
			k:    3,
			want: 3,
		},

		// Edge cases
		{
			name: "single node k=1",
			root: []interface{}{1},
			k:    1,
			want: 1,
		},
		{
			name: "two nodes left child k=1",
			root: []interface{}{2, 1, nil},
			k:    1,
			want: 1,
		},
		{
			name: "two nodes left child k=2",
			root: []interface{}{2, 1, nil},
			k:    2,
			want: 2,
		},
		{
			name: "two nodes right child k=1",
			root: []interface{}{1, nil, 2},
			k:    1,
			want: 1,
		},
		{
			name: "two nodes right child k=2",
			root: []interface{}{1, nil, 2},
			k:    2,
			want: 2,
		},

		// k = n (phần tử lớn nhất)
		{
			name: "k equals n - return max",
			root: []interface{}{3, 1, 4, nil, 2},
			k:    4,
			want: 4,
		},
		{
			name: "k equals n in example 2",
			root: []interface{}{5, 3, 6, 2, 4, nil, nil, 1},
			k:    6,
			want: 6,
		},

		// Cây lệch trái (left-skewed)
		{
			name: "left-skewed BST k=1",
			root: []interface{}{5, 4, nil, 3, nil, 2, nil, 1, nil},
			k:    1,
			want: 1,
		},
		{
			name: "left-skewed BST k=3",
			root: []interface{}{5, 4, nil, 3, nil, 2, nil},
			k:    3,
			want: 4,
		},

		// Cây lệch phải (right-skewed)
		{
			name: "right-skewed BST k=1",
			root: []interface{}{1, nil, 2, nil, 3, nil, 4},
			k:    1,
			want: 1,
		},
		{
			name: "right-skewed BST k=4",
			root: []interface{}{1, nil, 2, nil, 3, nil, 4},
			k:    4,
			want: 4,
		},

		// Cây hoàn chỉnh (complete BST)
		{
			name: "complete BST [4,2,6,1,3,5,7] k=1",
			root: []interface{}{4, 2, 6, 1, 3, 5, 7},
			k:    1,
			want: 1,
		},
		{
			name: "complete BST [4,2,6,1,3,5,7] k=4",
			root: []interface{}{4, 2, 6, 1, 3, 5, 7},
			k:    4,
			want: 4,
		},
		{
			name: "complete BST [4,2,6,1,3,5,7] k=7",
			root: []interface{}{4, 2, 6, 1, 3, 5, 7},
			k:    7,
			want: 7,
		},

		// Giá trị lớn
		{
			name: "large values BST",
			root: []interface{}{5000, 2500, 7500, 1000, 4000, 6000, 9000},
			k:    3,
			want: 4000,
		},

		// Giá trị 0
		{
			name: "tree with zero value k=1",
			root: []interface{}{1, 0, 2},
			k:    1,
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTree(tt.root)
			got := kthSmallest(root, tt.k)
			if got != tt.want {
				t.Errorf("kthSmallest(%v, %d) = %d; want %d", tt.root, tt.k, got, tt.want)
			}
		})
	}
}

func Benchmark_KthSmallest(b *testing.B) {
	root := buildTree([]interface{}{4, 2, 6, 1, 3, 5, 7})
	for i := 0; i < b.N; i++ {
		kthSmallest(root, 3)
	}
}
