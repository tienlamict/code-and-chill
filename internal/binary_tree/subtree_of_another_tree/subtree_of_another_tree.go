package binary_tree

// TreeNode định nghĩa một node trong binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// isSubtree kiểm tra xem subRoot có phải là subtree của root không.
//
// Thuật toán: DFS (Depth-First Search) + Tree Comparison
// - Duyệt tree root bằng DFS (pre-order traversal)
// - Với mỗi node trong root, kiểm tra xem subtree bắt đầu từ node đó
//   có giống với subRoot không bằng cách so sánh từng node
// - Nếu tìm thấy một subtree giống với subRoot, trả về true
//
// Độ phức tạp: O(m * n) thời gian, O(h) không gian (h là chiều cao của root)
// với m là số node trong root, n là số node trong subRoot
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	// Base case: Nếu subRoot là nil, nó là subtree của mọi tree
	if subRoot == nil {
		return true
	}

	// Base case: Nếu root là nil nhưng subRoot không phải nil, không thể có subtree
	if root == nil {
		return false
	}

	// Kiểm tra xem subtree bắt đầu từ node hiện tại có giống subRoot không
	if isSameTree(root, subRoot) {
		return true
	}

	// Đệ quy kiểm tra ở left subtree và right subtree
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

// isSameTree kiểm tra xem hai tree có giống nhau hoàn toàn không.
//
// Hai tree được coi là giống nhau nếu:
// - Cùng cấu trúc (cùng số node, cùng vị trí)
// - Cùng giá trị tại mỗi node
//
// Thuật toán: Đệ quy so sánh từng node
// - Nếu cả hai node đều nil → giống nhau
// - Nếu một trong hai node là nil → khác nhau
// - So sánh giá trị và đệ quy so sánh left và right subtree
//
// Độ phức tạp: O(min(m, n)) thời gian với m, n là số node của hai tree
func isSameTree(p *TreeNode, q *TreeNode) bool {
	// Base case: Cả hai đều nil → giống nhau
	if p == nil && q == nil {
		return true
	}

	// Base case: Một trong hai là nil → khác nhau
	if p == nil || q == nil {
		return false
	}

	// So sánh giá trị và đệ quy so sánh left và right subtree
	return p.Val == q.Val &&
		isSameTree(p.Left, q.Left) &&
		isSameTree(p.Right, q.Right)
}
