package binary_tree

import (
	"reflect"
	"testing"
)

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

// treeToSlice chuyển đổi binary tree thành mảng (level-order)
func treeToSlice(root *TreeNode) []interface{} {
	if root == nil {
		return []interface{}{}
	}

	var result []interface{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			result = append(result, nil)
		} else {
			result = append(result, node.Val)
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}
	}

	// Loại bỏ các nil thừa ở cuối
	for len(result) > 0 && result[len(result)-1] == nil {
		result = result[:len(result)-1]
	}

	return result
}

// treesEqual kiểm tra xem hai tree có giống nhau không
func treesEqual(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	return t1.Val == t2.Val &&
		treesEqual(t1.Left, t2.Left) &&
		treesEqual(t1.Right, t2.Right)
}

func Test_SerializeDeserialize(t *testing.T) {
	tests := []struct {
		name string
		vals []interface{}
	}{
		{
			name: "example 1",
			vals: []interface{}{1, 2, 3, nil, nil, 4, 5},
		},
		{
			name: "example 2 - empty tree",
			vals: []interface{}{},
		},
		{
			name: "single node",
			vals: []interface{}{1},
		},
		{
			name: "only left child",
			vals: []interface{}{1, 2, nil},
		},
		{
			name: "only right child",
			vals: []interface{}{1, nil, 2},
		},
		{
			name: "complete binary tree",
			vals: []interface{}{1, 2, 3, 4, 5, 6, 7},
		},
		{
			name: "left skewed tree",
			vals: []interface{}{1, 2, nil, 3, nil, 4, nil},
		},
		{
			name: "right skewed tree",
			vals: []interface{}{1, nil, 2, nil, 3, nil, 4},
		},
		{
			name: "negative values",
			vals: []interface{}{-1, -2, -3, nil, nil, -4, -5},
		},
		{
			name: "mixed positive and negative",
			vals: []interface{}{1, -2, 3, nil, nil, 4, -5},
		},
		{
			name: "large values",
			vals: []interface{}{1000, 2000, 3000},
		},
		{
			name: "zero values",
			vals: []interface{}{0, 0, 0},
		},
		{
			name: "complex tree 1",
			vals: []interface{}{1, 2, 3, nil, nil, 4, 5, 6, 7},
		},
		{
			name: "complex tree 2",
			vals: []interface{}{5, 2, 3, nil, nil, 2, 4, 3, 1},
		},
		{
			name: "tree with many nulls",
			vals: []interface{}{1, nil, 2, nil, nil, nil, 3},
		},
		{
			name: "unbalanced tree",
			vals: []interface{}{1, 2, nil, 3, nil, 4, nil, 5, nil},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Tạo tree từ input
			originalTree := createTreeFromSlice(tt.vals)

			t.Logf("Original tree: %v", treeToSlice(originalTree))

			// Serialize
			codec := Constructor()
			serialized := codec.serialize(originalTree)
			t.Logf("Serialized: %s", serialized)

			// Deserialize
			deserializedTree := codec.deserialize(serialized)
			t.Logf("Deserialized tree: %v", treeToSlice(deserializedTree))

			// Kiểm tra xem tree sau khi deserialize có giống tree ban đầu không
			if !treesEqual(originalTree, deserializedTree) {
				t.Fatalf(
					"Tree mismatch!\nOriginal: %v\nDeserialized: %v",
					treeToSlice(originalTree),
					treeToSlice(deserializedTree),
				)
			}
		})
	}
}

func Test_Serialize(t *testing.T) {
	tests := []struct {
		name     string
		vals     []interface{}
		expected string
	}{
		{
			name:     "example 1",
			vals:     []interface{}{1, 2, 3, nil, nil, 4, 5},
			expected: "1,2,3,null,null,4,5",
		},
		{
			name:     "empty tree",
			vals:     []interface{}{},
			expected: "",
		},
		{
			name:     "single node",
			vals:     []interface{}{1},
			expected: "1",
		},
		{
			name:     "only left child",
			vals:     []interface{}{1, 2, nil},
			expected: "1,2",
		},
		{
			name:     "only right child",
			vals:     []interface{}{1, nil, 2},
			expected: "1,null,2",
		},
		{
			name:     "complete binary tree",
			vals:     []interface{}{1, 2, 3, 4, 5, 6, 7},
			expected: "1,2,3,4,5,6,7",
		},
	}

	codec := Constructor()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := createTreeFromSlice(tt.vals)
			got := codec.serialize(tree)

			if got != tt.expected {
				t.Fatalf("serialize(%v) = %s; want %s", tt.vals, got, tt.expected)
			}
		})
	}
}

func Test_Deserialize(t *testing.T) {
	tests := []struct {
		name     string
		data     string
		expected []interface{}
	}{
		{
			name:     "example 1",
			data:     "1,2,3,null,null,4,5",
			expected: []interface{}{1, 2, 3, nil, nil, 4, 5},
		},
		{
			name:     "empty string",
			data:     "",
			expected: []interface{}{},
		},
		{
			name:     "single node",
			data:     "1",
			expected: []interface{}{1},
		},
		{
			name:     "only left child",
			data:     "1,2",
			expected: []interface{}{1, 2},
		},
		{
			name:     "only right child",
			data:     "1,null,2",
			expected: []interface{}{1, nil, 2},
		},
		{
			name:     "complete binary tree",
			data:     "1,2,3,4,5,6,7",
			expected: []interface{}{1, 2, 3, 4, 5, 6, 7},
		},
	}

	codec := Constructor()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := codec.deserialize(tt.data)
			got := treeToSlice(tree)

			if !reflect.DeepEqual(got, tt.expected) {
				t.Fatalf("deserialize(%s) = %v; want %v", tt.data, got, tt.expected)
			}
		})
	}
}
