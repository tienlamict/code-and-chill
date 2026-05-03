package same_tree

// TreeNode định nghĩa một nút trong cây nhị phân.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// isSameTree kiểm tra xem hai cây nhị phân có giống nhau hay không.
// Hai cây được coi là giống nhau nếu chúng có cấu trúc giống hệt nhau và các nút có cùng giá trị.
func isSameTree(p *TreeNode, q *TreeNode) bool {
	// Nếu cả hai nút đều null, chúng giống nhau
	if p == nil && q == nil {
		return true
	}

	// Nếu một trong hai nút null hoặc giá trị của chúng khác nhau, chúng không giống nhau
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}

	// Đệ quy kiểm tra cây con bên trái và bên phải
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
