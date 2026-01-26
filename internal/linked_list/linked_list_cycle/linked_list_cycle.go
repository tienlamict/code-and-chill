package linked_list

// ListNode định nghĩa một node trong linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// hasCycle kiểm tra xem linked list có chứa cycle hay không.
//
// Sử dụng thuật toán Floyd's Cycle Detection (Two Pointers):
// - Sử dụng hai con trỏ: slow (di chuyển 1 bước) và fast (di chuyển 2 bước)
// - Nếu có cycle, fast sẽ đuổi kịp slow tại một điểm nào đó
// - Nếu không có cycle, fast sẽ đến null trước
//
// Độ phức tạp: O(n) thời gian, O(1) không gian
func hasCycle(head *ListNode) bool {
	// Trường hợp rỗng hoặc chỉ có 1 node
	if head == nil || head.Next == nil {
		return false
	}

	// Khởi tạo hai con trỏ
	slow := head      // Di chuyển 1 bước mỗi lần
	fast := head.Next // Di chuyển 2 bước mỗi lần

	// Duyệt cho đến khi fast gặp null (không có cycle)
	// hoặc fast đuổi kịp slow (có cycle)
	for fast != nil && fast.Next != nil {
		// Nếu hai con trỏ gặp nhau, có cycle
		if slow == fast {
			return true
		}

		// Di chuyển slow 1 bước
		slow = slow.Next
		// Di chuyển fast 2 bước
		fast = fast.Next.Next
	}

	// Fast đã đến null, không có cycle
	return false
}
