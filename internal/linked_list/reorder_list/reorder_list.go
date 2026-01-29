package linked_list

// ListNode định nghĩa một node trong linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// reorderList sắp xếp lại linked list theo pattern: L0 → Ln → L1 → Ln-1 → L2 → Ln-2 → ...
//
// Giải pháp 3 bước:
// 1. Tìm middle của list bằng slow/fast pointers
// 2. Reverse nửa sau của list
// 3. Merge hai nửa lại với nhau theo pattern
//
// Độ phức tạp: O(n) thời gian, O(1) không gian
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	// Bước 1: Tìm middle của list
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Bước 2: Reverse nửa sau của list
	secondHalf := slow.Next
	slow.Next = nil // Ngắt kết nối giữa hai nửa
	secondHalf = reverseList(secondHalf)

	// Bước 3: Merge hai nửa lại với nhau
	firstHalf := head
	for secondHalf != nil {
		// Lưu các node tiếp theo
		firstNext := firstHalf.Next
		secondNext := secondHalf.Next

		// Kết nối: firstHalf -> secondHalf -> firstNext
		firstHalf.Next = secondHalf
		secondHalf.Next = firstNext

		// Di chuyển đến node tiếp theo
		firstHalf = firstNext
		secondHalf = secondNext
	}
}

// reverseList đảo ngược một linked list
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}
