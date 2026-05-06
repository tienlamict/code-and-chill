package binary_tree

// TreeNode định nghĩa một node trong binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// kthSmallest trả về giá trị nhỏ thứ k trong BST (1-indexed).
//
// Thuật toán: In-order traversal lặp (Iterative In-order)
// - In-order traversal của BST cho kết quả các node theo thứ tự tăng dần
// - Dùng stack để mô phỏng đệ quy, đi hết sang trái trước
// - Mỗi lần pop một node ra khỏi stack là ta đang thăm node theo thứ tự tăng dần
// - Khi đã thăm đúng k node thì trả về giá trị đó ngay (dừng sớm)
//
// Độ phức tạp: O(H + k) thời gian (H là chiều cao cây), O(H) không gian
func kthSmallest(root *TreeNode, k int) int {
	stack := []*TreeNode{}
	curr := root

	for curr != nil || len(stack) > 0 {
		// Đi hết sang trái, đẩy tất cả node vào stack
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		// Pop node nhỏ nhất chưa thăm
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		k--
		if k == 0 {
			return curr.Val
		}

		// Chuyển sang cây con phải
		curr = curr.Right
	}

	return -1 // không bao giờ xảy ra nếu k hợp lệ
}
