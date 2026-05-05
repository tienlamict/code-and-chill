package binary_tree

// TreeNode định nghĩa một node trong binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// levelOrder trả về danh sách các giá trị node theo từng tầng (level order traversal).
//
// Thuật toán: BFS (Breadth-First Search) sử dụng queue
// - Đẩy root vào queue.
// - Với mỗi tầng, xác định số lượng node trong queue (= số node ở tầng hiện tại).
// - Lần lượt lấy từng node ra, ghi nhận giá trị, rồi đẩy node con trái/phải vào queue.
// - Lặp lại cho đến khi queue rỗng.
//
// Độ phức tạp:
// - Thời gian: O(n) vì mỗi node được xử lý đúng một lần.
// - Không gian: O(n) cho queue (tối đa chứa toàn bộ một tầng, tầng cuối có thể có n/2 node).
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var result [][]int
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		level := make([]int, 0, levelSize)

		// Xử lý tất cả node trong tầng hiện tại
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			level = append(level, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, level)
	}

	return result
}
