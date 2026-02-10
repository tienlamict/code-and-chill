package binary_tree

import (
	"strconv"
	"strings"
)

// TreeNode định nghĩa một node trong binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Codec cung cấp các phương thức serialize và deserialize
type Codec struct{}

// Constructor khởi tạo Codec mới
func Constructor() Codec {
	return Codec{}
}

// serialize chuyển đổi binary tree thành chuỗi string.
//
// Sử dụng Level-Order Traversal (BFS):
// - Duyệt tree theo từng level từ trên xuống dưới
// - Với mỗi node, thêm giá trị vào chuỗi (hoặc "null" nếu node là nil)
// - Sử dụng dấu phẩy để phân tách các giá trị
// - Format: "1,2,3,null,null,4,5"
//
// Độ phức tạp: O(n) thời gian, O(n) không gian
func (c *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	var result []string
	queue := []*TreeNode{root}

	// Duyệt tree theo level-order
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			result = append(result, "null")
		} else {
			result = append(result, strconv.Itoa(node.Val))
			// Thêm cả left và right vào queue (kể cả khi nil)
			// để giữ được cấu trúc tree
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}
	}

	// Loại bỏ các "null" thừa ở cuối
	// (các null này không cần thiết vì không có node nào sau chúng)
	for len(result) > 0 && result[len(result)-1] == "null" {
		result = result[:len(result)-1]
	}

	return strings.Join(result, ",")
}

// deserialize chuyển đổi chuỗi string thành binary tree.
//
// Sử dụng Level-Order Traversal (BFS) ngược lại:
// - Parse chuỗi thành mảng các giá trị
// - Sử dụng queue để xây dựng tree từ trên xuống dưới
// - Với mỗi node không null, tạo node mới và thêm children vào queue
//
// Độ phức tạp: O(n) thời gian, O(n) không gian
func (c *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}

	// Parse chuỗi thành mảng các giá trị
	values := strings.Split(data, ",")
	if len(values) == 0 {
		return nil
	}

	// Parse giá trị đầu tiên làm root
	rootVal, err := strconv.Atoi(values[0])
	if err != nil {
		return nil
	}
	root := &TreeNode{Val: rootVal}

	queue := []*TreeNode{root}
	index := 1

	// Xây dựng tree theo level-order
	for len(queue) > 0 && index < len(values) {
		node := queue[0]
		queue = queue[1:]

		// Xử lý left child
		if index < len(values) && values[index] != "null" {
			leftVal, err := strconv.Atoi(values[index])
			if err == nil {
				node.Left = &TreeNode{Val: leftVal}
				queue = append(queue, node.Left)
			}
		}
		index++

		// Xử lý right child
		if index < len(values) && values[index] != "null" {
			rightVal, err := strconv.Atoi(values[index])
			if err == nil {
				node.Right = &TreeNode{Val: rightVal}
				queue = append(queue, node.Right)
			}
		}
		index++
	}

	return root
}
