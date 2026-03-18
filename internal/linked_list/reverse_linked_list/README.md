# Reverse Linked List (LeetCode 206)

## Mô tả bài toán

Cho `head` của một **singly linked list**, đảo ngược list và trả về head của list sau khi đảo.

## Ví dụ

- Input: `[1,2,3,4,5]` → Output: `[5,4,3,2,1]`
- Input: `[1,2]` → Output: `[2,1]`
- Input: `[]` → Output: `[]`

## Ràng buộc

- Số node: `[0, 5000]`
- `-5000 <= Node.val <= 5000`

## Follow-up

Có thể đảo **lặp (iterative)** hoặc **đệ quy (recursive)** — trong code có cả hai.

---

## Cách tiếp cận

### 1. Iterative — O(n) time, O(1) extra space ⭐

Dùng ba con trỏ `prev`, `curr`, `next`:

1. `next = curr.Next` lưu phần còn lại.
2. `curr.Next = prev` đảo hướng.
3. `prev = curr`, `curr = next`.

Kết thúc khi `curr == nil`, `prev` là head mới.

### 2. Recursive — O(n) time, O(n) stack

- Base: `head == nil` hoặc `head.Next == nil` → trả về `head`.
- Gọi `newHead = reverseListRecursive(head.Next)`.
- `head.Next.Next = head`, `head.Next = nil`.

---

## Độ phức tạp

| Cách       | Thời gian | Không gian phụ   |
|-----------|-----------|------------------|
| Iterative | O(n)      | O(1)             |
| Recursive | O(n)      | O(n) stack       |

---

## Code (trích)

Xem `reverse_linked_list.go`: `reverseList` (lặp) và `reverseListRecursive` (đệ quy).

## Chạy test

```bash
go test ./internal/linked_list/reverse_linked_list/
```
