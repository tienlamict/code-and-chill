package construct_binary_tree_from_preorder

// TreeNode định nghĩa một nút trong cây nhị phân.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// buildTree xây dựng cây nhị phân từ mảng preorder và inorder traversal.
//
// Thuật toán (Divide & Conquer + HashMap):
//  1. Phần tử đầu tiên của preorder luôn là gốc (root).
//  2. Tìm vị trí root trong inorder → chia inorder thành cây con trái và phải.
//  3. Số lượng nút cây con trái = (vị trí root trong inorder) - inStart.
//  4. Đệ quy xây dựng cây con trái và phải với các chỉ số tương ứng.
//  5. Dùng HashMap lưu index của từng giá trị trong inorder để tra cứu O(1).
func buildTree(preorder []int, inorder []int) *TreeNode {
	inorderIdx := make(map[int]int, len(inorder))
	for i, v := range inorder {
		inorderIdx[v] = i
	}

	var build func(preStart, inStart, inEnd int) *TreeNode
	build = func(preStart, inStart, inEnd int) *TreeNode {
		if inStart > inEnd {
			return nil
		}

		rootVal := preorder[preStart]
		root := &TreeNode{Val: rootVal}

		mid := inorderIdx[rootVal]
		leftSize := mid - inStart

		root.Left = build(preStart+1, inStart, mid-1)
		root.Right = build(preStart+1+leftSize, mid+1, inEnd)

		return root
	}

	return build(0, 0, len(inorder)-1)
}
