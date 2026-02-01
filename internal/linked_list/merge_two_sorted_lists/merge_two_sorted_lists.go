package linked_list

// ListNode định nghĩa một node trong linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// mergeTwoLists hợp nhất hai danh sách đã được sắp xếp thành một danh sách đã sắp xếp.
//
// Sử dụng kỹ thuật Two Pointers với dummy node:
// - Tạo dummy node để đơn giản hóa việc xử lý edge cases
// - So sánh từng node của hai list và chọn node nhỏ hơn
// - Nối node nhỏ hơn vào kết quả và di chuyển con trỏ tương ứng
// - Tiếp tục cho đến khi một trong hai list kết thúc
// - Nối phần còn lại của list chưa kết thúc vào kết quả
//
// Độ phức tạp: O(n + m) thời gian, O(1) không gian
// với n và m là độ dài của list1 và list2
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// Tạo dummy node để đơn giản hóa việc xử lý
	dummy := &ListNode{}
	current := dummy

	// Duyệt cả hai list đồng thời
	for list1 != nil && list2 != nil {
		// So sánh giá trị của hai node hiện tại
		if list1.Val <= list2.Val {
			// Chọn node từ list1 nếu giá trị nhỏ hơn hoặc bằng
			current.Next = list1
			list1 = list1.Next
		} else {
			// Chọn node từ list2 nếu giá trị nhỏ hơn
			current.Next = list2
			list2 = list2.Next
		}
		// Di chuyển con trỏ kết quả
		current = current.Next
	}

	// Nối phần còn lại của list chưa kết thúc
	// (một trong hai list sẽ là nil, nên chỉ một trong hai dòng này có hiệu lực)
	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}

	// Trả về head của list đã hợp nhất (bỏ qua dummy node)
	return dummy.Next
}
