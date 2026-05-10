package binary_tree

// TreeNode định nghĩa một node trong binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// lowestCommonAncestor tìm node tổ tiên chung thấp nhất của hai node p và q trong BST.
//
// Thuật toán: Dựa trên tính chất của Binary Search Tree (BST)
// - Trong BST, mọi node ở cây con trái đều nhỏ hơn node gốc, và mọi node ở cây con phải đều lớn hơn node gốc.
// - Nếu cả p và q đều nhỏ hơn root, LCA phải nằm ở bên trái.
// - Nếu cả p và q đều lớn hơn root, LCA phải nằm ở bên phải.
// - Nếu p và q nằm ở hai phía khác nhau (một nhỏ hơn, một lớn hơn hoặc bằng root), thì root chính là LCA.
//
// Độ phức tạp:
// - Thời gian: O(h) với h là chiều cao của cây. Trong trường hợp xấu nhất là O(n).
// - Không gian: O(1) nếu dùng vòng lặp, O(h) nếu dùng đệ quy (do call stack).
// Ở đây ta sử dụng vòng lặp để tối ưu bộ nhớ.
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	for root != nil {
		if p.Val < root.Val && q.Val < root.Val {
			// Cả p và q đều nằm bên trái
			root = root.Left
		} else if p.Val > root.Val && q.Val > root.Val {
			// Cả p và q đều nằm bên phải
			root = root.Right
		} else {
			// p và q nằm ở hai phía hoặc một trong hai là root
			return root
		}
	}
	return nil
}
