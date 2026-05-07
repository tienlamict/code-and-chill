package binary_tree

// TreeNode định nghĩa một node trong binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// maxDepth trả về độ sâu lớn nhất của binary tree.
//
// Thuật toán: Đệ quy (DFS - Post-order)
// - Độ sâu của cây = 1 + max(độ sâu cây con trái, độ sâu cây con phải)
// - Base case: Nếu node là nil, độ sâu = 0.
//
// Độ phức tạp:
// - Thời gian: O(n) vì phải duyệt qua toàn bộ n node.
// - Không gian: O(h) với h là chiều cao cây (call stack đệ quy).
//   Trường hợp xấu nhất (cây lệch): O(n). Cây cân bằng: O(log n).
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}
