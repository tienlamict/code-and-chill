package linked_list

// ListNode định nghĩa một node trong linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// removeNthFromEnd xóa node thứ n từ cuối danh sách và trả về head của list mới.
//
// Sử dụng kỹ thuật Two Pointers:
// - Tạo dummy node để xử lý trường hợp xóa head
// - Di chuyển fast pointer n+1 bước trước
// - Sau đó di chuyển cả fast và slow cùng lúc cho đến khi fast đến cuối
// - Khi đó slow sẽ ở node trước node cần xóa
// - Xóa node bằng cách cập nhật slow.Next = slow.Next.Next
//
// Độ phức tạp: O(n) thời gian, O(1) không gian
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// Tạo dummy node để xử lý trường hợp xóa head
	dummy := &ListNode{Next: head}

	// Khởi tạo hai con trỏ
	slow := dummy
	fast := dummy

	// Di chuyển fast pointer n+1 bước trước
	// +1 để slow sẽ dừng ở node trước node cần xóa
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}

	// Di chuyển cả hai con trỏ cùng lúc cho đến khi fast đến cuối
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}

	// Xóa node bằng cách bỏ qua node thứ n từ cuối
	slow.Next = slow.Next.Next

	// Trả về head mới (có thể đã thay đổi nếu xóa head cũ)
	return dummy.Next
}

