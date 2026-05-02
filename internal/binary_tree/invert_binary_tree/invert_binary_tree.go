package binary_tree

// TreeNode định nghĩa một node trong binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// invertTree đảo ngược một binary tree.
//
// Thuật toán: Đệ quy (DFS - Pre-order)
// - Với mỗi node, ta tráo đổi node con bên trái và node con bên phải.
// - Sau đó đệ quy thực hiện tương tự cho node con bên trái và node con bên phải đã được tráo đổi.
// - Base case: Nếu node là nil, trả về nil.
//
// Độ phức tạp:
// - Thời gian: O(n) với n là số lượng node trong cây, vì ta phải duyệt qua mỗi node một lần.
// - Không gian: O(h) với h là chiều cao của cây (do call stack của đệ quy).
//   Trong trường hợp xấu nhất (cây lệch), h = n. Trong trường hợp tốt nhất (cây cân bằng), h = log(n).
func invertTree(root *TreeNode) *TreeNode {
	// Base case: Nếu cây trống, không có gì để đảo ngược
	if root == nil {
		return nil
	}

	// Tráo đổi node con bên trái và bên phải
	root.Left, root.Right = root.Right, root.Left

	// Tiếp tục đệ quy đảo ngược các cây con trái và phải
	invertTree(root.Left)
	invertTree(root.Right)

	return root
}
