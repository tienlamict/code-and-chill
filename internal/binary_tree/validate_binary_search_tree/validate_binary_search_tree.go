package binary_tree

import "math"

// TreeNode định nghĩa một node trong binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// isValidBST kiểm tra xem một binary tree có phải là BST hợp lệ hay không.
//
// Thuật toán: Recursive với giới hạn min/max (Valid Range Approach)
// - Mỗi node phải nằm trong khoảng hợp lệ (min, max)
// - Khi đi sang trái: cập nhật max = node.Val (node con trái phải < node hiện tại)
// - Khi đi sang phải: cập nhật min = node.Val (node con phải phải > node hiện tại)
// - Khởi đầu với khoảng (-∞, +∞), sử dụng int64 để tránh overflow với biên int32
//
// Độ phức tạp: O(n) thời gian, O(h) không gian (h là chiều cao cây)
func isValidBST(root *TreeNode) bool {
	return validate(root, math.MinInt64, math.MaxInt64)
}

// validate kiểm tra đệ quy mỗi node có nằm trong khoảng (min, max) không.
func validate(node *TreeNode, min, max int64) bool {
	if node == nil {
		return true
	}

	val := int64(node.Val)

	// Node hiện tại phải nằm trong khoảng (min, max) không bao gồm biên
	if val <= min || val >= max {
		return false
	}

	// Đi sang trái: max mới là val; đi sang phải: min mới là val
	return validate(node.Left, min, val) && validate(node.Right, val, max)
}
