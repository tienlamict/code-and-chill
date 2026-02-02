package linked_list

// ListNode định nghĩa một node trong linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// mergeKLists hợp nhất k linked lists đã được sắp xếp thành một linked list đã sắp xếp.
//
// Sử dụng kỹ thuật Divide and Conquer:
// 1. Chia mảng lists thành hai nửa
// 2. Đệ quy merge từng nửa
// 3. Merge hai kết quả lại với nhau
//
// Độ phức tạp: O(n log k) thời gian, O(log k) không gian (cho call stack)
// với n là tổng số nodes và k là số lượng lists
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	return mergeLists(lists, 0, len(lists)-1)
}

// mergeLists merge các lists từ index left đến right bằng divide and conquer
func mergeLists(lists []*ListNode, left, right int) *ListNode {
	// Base case: chỉ còn một list
	if left == right {
		return lists[left]
	}

	// Chia đôi
	mid := left + (right-left)/2

	// Đệ quy merge hai nửa
	leftList := mergeLists(lists, left, mid)
	rightList := mergeLists(lists, mid+1, right)

	// Merge hai kết quả
	return mergeTwoLists(leftList, rightList)
}

// mergeTwoLists hợp nhất hai linked lists đã được sắp xếp thành một linked list đã sắp xếp.
//
// Sử dụng kỹ thuật Two Pointers với dummy node:
// - Tạo dummy node để đơn giản hóa việc xử lý edge cases
// - So sánh từng node của hai list và chọn node nhỏ hơn
// - Nối node nhỏ hơn vào kết quả và di chuyển con trỏ tương ứng
// - Tiếp tục cho đến khi một trong hai list kết thúc
// - Nối phần còn lại của list chưa kết thúc vào kết quả
//
// Độ phức tạp: O(n + m) thời gian, O(1) không gian
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
	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}

	// Trả về head của list đã hợp nhất (bỏ qua dummy node)
	return dummy.Next
}
