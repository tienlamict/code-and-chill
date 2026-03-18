package linked_list

// ListNode định nghĩa một node trong singly linked list (theo LeetCode).
type ListNode struct {
	Val  int
	Next *ListNode
}

// reverseList đảo ngược danh sách liên kết đơn (cách lặp — O(n) thời gian, O(1) không gian phụ).
//
// Thuật toán: duyệt từng node, gán Next của node hiện tại trỏ về prev,
// rồi tiến prev và curr lên phía trước.
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

// reverseListRecursive đảo ngược danh sách bằng đệ quy (Follow-up LeetCode).
//
// Ý tưởng: reverseListRecursive(head.Next) trả về head của list đã đảo từ node thứ 2 trở đi;
// sau đó head.Next.Next = head, head.Next = nil để nối head vào cuối.
// Độ phức tạp: O(n) thời gian, O(n) không gian stack đệ quy.
func reverseListRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseListRecursive(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
