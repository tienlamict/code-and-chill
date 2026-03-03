package binary_tree

import "testing"

// createTreeFromSlice tạo binary tree từ mảng các giá trị (level-order)
// Format: [1,2,3,null,null,4,5]
func createTreeFromSlice(vals []interface{}) *TreeNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}

	root := &TreeNode{Val: vals[0].(int)}
	queue := []*TreeNode{root}
	index := 1

	for len(queue) > 0 && index < len(vals) {
		node := queue[0]
		queue = queue[1:]

		// Left child
		if index < len(vals) && vals[index] != nil {
			node.Left = &TreeNode{Val: vals[index].(int)}
			queue = append(queue, node.Left)
		}
		index++

		// Right child
		if index < len(vals) && vals[index] != nil {
			node.Right = &TreeNode{Val: vals[index].(int)}
			queue = append(queue, node.Right)
		}
		index++
	}

	return root
}

func Test_IsSubtree(t *testing.T) {
	tests := []struct {
		name    string
		root    []interface{}
		subRoot []interface{}
		want    bool
	}{
		{
			name:    "example 1 - subtree exists",
			root:    []interface{}{3, 4, 5, 1, 2},
			subRoot: []interface{}{4, 1, 2},
			want:    true,
		},
		{
			name:    "example 2 - subtree does not exist",
			root:    []interface{}{3, 4, 5, 1, 2, nil, nil, nil, nil, 0},
			subRoot: []interface{}{4, 1, 2},
			want:    false,
		},
		{
			name:    "same tree",
			root:    []interface{}{1, 2, 3},
			subRoot: []interface{}{1, 2, 3},
			want:    true,
		},
		{
			name:    "single node - same value",
			root:    []interface{}{1},
			subRoot: []interface{}{1},
			want:    true,
		},
		{
			name:    "single node - different value",
			root:    []interface{}{1},
			subRoot: []interface{}{2},
			want:    false,
		},
		{
			name:    "subtree is root itself",
			root:    []interface{}{3, 4, 5, 1, 2},
			subRoot: []interface{}{3, 4, 5, 1, 2},
			want:    true,
		},
		{
			name:    "subtree in left branch",
			root:    []interface{}{3, 4, 5, 1, 2, nil, nil},
			subRoot: []interface{}{4, 1, 2},
			want:    true,
		},
		{
			name:    "subtree in right branch",
			root:    []interface{}{3, 4, 5, nil, nil, 1, 2},
			subRoot: []interface{}{5, 1, 2},
			want:    true,
		},
		{
			name:    "subtree is single node in left",
			root:    []interface{}{3, 4, 5},
			subRoot: []interface{}{4},
			want:    true,
		},
		{
			name:    "subtree is single node in right",
			root:    []interface{}{3, 4, 5},
			subRoot: []interface{}{5},
			want:    true,
		},
		{
			name:    "subtree not found - different structure",
			root:    []interface{}{3, 4, 5, 1, 2},
			subRoot: []interface{}{4, 1, nil},
			want:    false,
		},
		{
			name:    "subtree not found - different values",
			root:    []interface{}{3, 4, 5, 1, 2},
			subRoot: []interface{}{4, 1, 3},
			want:    false,
		},
		{
			name:    "empty subRoot",
			root:    []interface{}{1, 2, 3},
			subRoot: []interface{}{},
			want:    true, // Empty tree is subtree of any tree
		},
		{
			name:    "complex tree - subtree exists",
			root:    []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			subRoot: []interface{}{6, 12, 13},
			want:    true,
		},
		{
			name:    "complex tree - subtree not exists",
			root:    []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			subRoot: []interface{}{6, 12, 14},
			want:    false,
		},
		{
			name:    "left skewed tree - subtree exists",
			root:    []interface{}{1, 2, nil, 3, nil, 4, nil},
			subRoot: []interface{}{3, 4, nil},
			want:    true,
		},
		{
			name:    "negative values - subtree exists",
			root:    []interface{}{-1, -2, -3, -4, -5},
			subRoot: []interface{}{-2, -4, -5},
			want:    true,
		},
		{
			name:    "mixed positive and negative - subtree exists",
			root:    []interface{}{1, -2, 3, -4, 5},
			subRoot: []interface{}{-2, -4, 5},
			want:    true,
		},
		{
			name:    "zero values - subtree exists",
			root:    []interface{}{0, 0, 0, 0, 0},
			subRoot: []interface{}{0, 0, 0},
			want:    true,
		},
		{
			name:    "large values - subtree exists",
			root:    []interface{}{10000, 20000, 30000, 40000, 50000},
			subRoot: []interface{}{20000, 40000, 50000},
			want:    true,
		},
		{
			name:    "subtree with only left child",
			root:    []interface{}{1, 2, 3, 4, nil},
			subRoot: []interface{}{2, 4, nil},
			want:    true,
		},
		{
			name:    "subtree with only right child",
			root:    []interface{}{1, 2, 3, nil, 4},
			subRoot: []interface{}{2, nil, 4},
			want:    true,
		},
		{
			name:    "deep nested subtree",
			root:    []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31},
			subRoot: []interface{}{15, 30, 31},
			want:    true,
		},
		{
			name:    "subtree matches partial structure",
			root:    []interface{}{1, 2, 3, 4, 5},
			subRoot: []interface{}{2, 4},
			want:    false, // Partial match không được tính là subtree
		},
		{
			name:    "multiple occurrences - first match",
			root:    []interface{}{1, 1, 1, 1, 1},
			subRoot: []interface{}{1, 1, 1},
			want:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rootTree := createTreeFromSlice(tt.root)
			subRootTree := createTreeFromSlice(tt.subRoot)

			t.Logf("Root tree: %v", tt.root)
			t.Logf("SubRoot tree: %v", tt.subRoot)

			got := isSubtree(rootTree, subRootTree)

			t.Logf("Result: %v (expected: %v)", got, tt.want)

			if got != tt.want {
				t.Fatalf("isSubtree() = %v; want %v", got, tt.want)
			}
		})
	}
}

func Test_IsSameTree(t *testing.T) {
	tests := []struct {
		name string
		p    []interface{}
		q    []interface{}
		want bool
	}{
		{
			name: "same tree",
			p:    []interface{}{1, 2, 3},
			q:    []interface{}{1, 2, 3},
			want: true,
		},
		{
			name: "different values",
			p:    []interface{}{1, 2, 3},
			q:    []interface{}{1, 2, 4},
			want: false,
		},
		{
			name: "different structure",
			p:    []interface{}{1, 2, 3},
			q:    []interface{}{1, 2},
			want: false,
		},
		{
			name: "both empty",
			p:    []interface{}{},
			q:    []interface{}{},
			want: true,
		},
		{
			name: "one empty",
			p:    []interface{}{1},
			q:    []interface{}{},
			want: false,
		},
		{
			name: "single node same",
			p:    []interface{}{1},
			q:    []interface{}{1},
			want: true,
		},
		{
			name: "single node different",
			p:    []interface{}{1},
			q:    []interface{}{2},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pTree := createTreeFromSlice(tt.p)
			qTree := createTreeFromSlice(tt.q)

			got := isSameTree(pTree, qTree)

			if got != tt.want {
				t.Fatalf("isSameTree() = %v; want %v", got, tt.want)
			}
		})
	}
}

// Benchmark test
func Benchmark_IsSubtree(b *testing.B) {
	root := createTreeFromSlice([]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15})
	subRoot := createTreeFromSlice([]interface{}{6, 12, 13})
	for i := 0; i < b.N; i++ {
		isSubtree(root, subRoot)
	}
}
