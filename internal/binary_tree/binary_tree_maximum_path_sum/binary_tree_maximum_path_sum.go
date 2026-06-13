package binary_tree

import "math"

// TreeNode định nghĩa một node trong binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// maxPathSum trả về tổng đường đi lớn nhất trong cây nhị phân.
// Một đường đi không nhất thiết phải đi qua gốc.
//
// Thuật toán: Đệ quy (DFS - Post-order)
// - Với mỗi node, ta tính "giá trị đóng góp lớn nhất" (max gain) mà nó có thể mang lại cho node cha.
// - Giá trị đóng góp này chỉ lấy từ một trong hai nhánh (trái hoặc phải) cộng với giá trị chính nó.
// - Trong quá trình đệ quy, ta cập nhật biến toàn cục (hoặc biến tham chiếu) để lưu lại tổng đường đi
//   lớn nhất đi qua node hiện tại (như là đỉnh của đường đi đó).
//
// Độ phức tạp:
// - Thời gian: O(n) với n là số lượng node trong cây, vì mỗi node được thăm đúng một lần.
// - Không gian: O(h) với h là chiều cao của cây, do call stack của đệ quy.
func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32

	// helper là hàm đệ quy để tính toán gain lớn nhất từ mỗi node
	var helper func(*TreeNode) int
	helper = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		// Tính toán đóng góp từ nhánh trái và phải.
		// Nếu đóng góp âm, ta coi như không lấy (0).
		leftGain := max(0, helper(node.Left))
		rightGain := max(0, helper(node.Right))

		// Đường đi lớn nhất đi qua node hiện tại (node hiện tại là đỉnh)
		currentPathSum := node.Val + leftGain + rightGain

		// Cập nhật maxSum toàn cục
		if currentPathSum > maxSum {
			maxSum = currentPathSum
		}

		// Trả về giá trị đóng góp lớn nhất mà node này có thể cung cấp cho node cha
		// (Chỉ được chọn một trong hai nhánh)
		return node.Val + max(leftGain, rightGain)
	}

	helper(root)
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
