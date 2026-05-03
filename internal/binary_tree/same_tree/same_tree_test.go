package same_tree

import (
	"testing"
)

// buildTree helper function để tạo cây nhị phân từ slice (level-order traversal)
func buildTree(nodes []interface{}) *TreeNode {
	if len(nodes) == 0 || nodes[0] == nil {
		return nil
	}

	root := &TreeNode{Val: nodes[0].(int)}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(nodes) {
		curr := queue[0]
		queue = queue[1:]

		// Left child
		if i < len(nodes) && nodes[i] != nil {
			curr.Left = &TreeNode{Val: nodes[i].(int)}
			queue = append(queue, curr.Left)
		}
		i++

		// Right child
		if i < len(nodes) && nodes[i] != nil {
			curr.Right = &TreeNode{Val: nodes[i].(int)}
			queue = append(queue, curr.Right)
		}
		i++
	}

	return root
}

func TestIsSameTree(t *testing.T) {
	tests := []struct {
		name     string
		p        []interface{}
		q        []interface{}
		expected bool
	}{
		{
			name:     "Example 1",
			p:        []interface{}{1, 2, 3},
			q:        []interface{}{1, 2, 3},
			expected: true,
		},
		{
			name:     "Example 2",
			p:        []interface{}{1, 2},
			q:        []interface{}{1, nil, 2},
			expected: false,
		},
		{
			name:     "Example 3",
			p:        []interface{}{1, 2, 1},
			q:        []interface{}{1, 1, 2},
			expected: false,
		},
		{
			name:     "Both empty",
			p:        []interface{}{},
			q:        []interface{}{},
			expected: true,
		},
		{
			name:     "One empty",
			p:        []interface{}{1},
			q:        []interface{}{},
			expected: false,
		},
		{
			name:     "Different values",
			p:        []interface{}{1, 2, 3},
			q:        []interface{}{1, 2, 4},
			expected: false,
		},
		{
			name:     "Different structure",
			p:        []interface{}{1, 2, nil, 3},
			q:        []interface{}{1, 2, 3},
			expected: false,
		},
		{
			name:     "Deep trees same",
			p:        []interface{}{1, 2, 3, 4, 5, 6, 7},
			q:        []interface{}{1, 2, 3, 4, 5, 6, 7},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := buildTree(tt.p)
			q := buildTree(tt.q)
			got := isSameTree(p, q)
			if got != tt.expected {
				t.Errorf("isSameTree() = %v, want %v", got, tt.expected)
			}
		})
	}
}
