# Reorder List

## Mô tả bài toán

Cho `head` của một singly linked list. List có thể được biểu diễn như sau:

```
L0 → L1 → … → Ln - 1 → Ln
```

Sắp xếp lại list theo dạng sau:

```
L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
```

Bạn không được thay đổi giá trị trong các node của list. Chỉ có thể thay đổi các node.

**Ví dụ:**

### Example 1:
- Input: `head = [1,2,3,4]`
- Output: `[1,4,2,3]`

```
Trước: 1 → 2 → 3 → 4
Sau:   1 → 4 → 2 → 3
```

### Example 2:
- Input: `head = [1,2,3,4,5]`
- Output: `[1,5,2,4,3]`

```
Trước: 1 → 2 → 3 → 4 → 5
Sau:   1 → 5 → 2 → 4 → 3
```

**Ràng buộc:**
- Số lượng nodes trong list nằm trong khoảng `[1, 5 * 10^4]`
- `1 <= Node.val <= 1000`

## Phân tích thuật toán

### Cách tiếp cận: Three-Step Solution

Giải pháp được chia thành 3 bước:

1. **Tìm middle của list**: Sử dụng slow/fast pointers để tìm điểm giữa
2. **Reverse nửa sau**: Đảo ngược nửa sau của list
3. **Merge hai nửa**: Kết hợp hai nửa theo pattern yêu cầu

### Các bước thực hiện

#### Bước 1: Tìm Middle của List

Sử dụng kỹ thuật **slow/fast pointers** (tương tự như Floyd's Cycle Detection):
- `slow`: Di chuyển 1 bước mỗi lần
- `fast`: Di chuyển 2 bước mỗi lần
- Khi `fast` đến cuối, `slow` sẽ ở giữa list

**Ví dụ:** Với list `[1,2,3,4,5]`
```
Ban đầu: slow = 1, fast = 1
Bước 1: slow = 2, fast = 3
Bước 2: slow = 3, fast = 5
fast.Next = nil → dừng
Middle = node 3
```

#### Bước 2: Reverse Nửa Sau

Đảo ngược nửa sau của list và ngắt kết nối giữa hai nửa.

**Ví dụ:** Với list `[1,2,3,4,5]`
```
Sau bước 1: [1,2,3] và [4,5]
Reverse [4,5] → [5,4]
Kết quả: [1,2,3] và [5,4] (đã ngắt kết nối)
```

#### Bước 3: Merge Hai Nửa

Kết hợp hai nửa theo pattern: L0 → Ln → L1 → Ln-1 → ...

**Ví dụ:** Với `[1,2,3]` và `[5,4]`
```
1 → 2 → 3
5 → 4

Merge:
1 → 5 → 2 → 4 → 3
```

### Ví dụ minh họa đầy đủ

Với input: `head = [1,2,3,4]`

```
Bước 1: Tìm middle
  List: 1 → 2 → 3 → 4
  slow = 1, fast = 1
  slow = 2, fast = 3
  fast.Next.Next = nil → dừng
  Middle = node 2
  First half: 1 → 2
  Second half: 3 → 4

Bước 2: Reverse nửa sau
  Second half: 3 → 4
  Reverse: 4 → 3
  Ngắt kết nối: First half = 1 → 2 (không còn trỏ đến 3)

Bước 3: Merge
  First half: 1 → 2
  Second half: 4 → 3
  
  Lặp 1:
    1 → 4 → 2
    3 (còn lại)
  
  Lặp 2:
    1 → 4 → 2 → 3
    (hết)

Kết quả: 1 → 4 → 2 → 3
```

Với input: `head = [1,2,3,4,5]`

```
Bước 1: Tìm middle
  List: 1 → 2 → 3 → 4 → 5
  slow = 1, fast = 1
  slow = 2, fast = 3
  slow = 3, fast = 5
  fast.Next = nil → dừng
  Middle = node 3
  First half: 1 → 2 → 3
  Second half: 4 → 5

Bước 2: Reverse nửa sau
  Second half: 4 → 5
  Reverse: 5 → 4
  Ngắt kết nối: First half = 1 → 2 → 3

Bước 3: Merge
  First half: 1 → 2 → 3
  Second half: 5 → 4
  
  Lặp 1:
    1 → 5 → 2 → 3
    4 (còn lại)
  
  Lặp 2:
    1 → 5 → 2 → 4 → 3
    (hết)

Kết quả: 1 → 5 → 2 → 4 → 3
```

## Giải thích code

### Cấu trúc hàm chính

```1:48:internal/linked_list/reorder_list/reorder_list.go
package linked_list

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
```

### Chi tiết từng phần

#### 1. Kiểm tra trường hợp biên (dòng 13-16)
```go
if head == nil || head.Next == nil {
    return
}
```
- Nếu list rỗng hoặc chỉ có 1 node, không cần sắp xếp lại
- Trả về ngay lập tức

#### 2. Bước 1: Tìm Middle (dòng 18-22)
```go
slow, fast := head, head
for fast.Next != nil && fast.Next.Next != nil {
    slow = slow.Next
    fast = fast.Next.Next
}
```

**Giải thích:**
- `slow` và `fast` đều bắt đầu từ `head`
- `slow` di chuyển 1 bước, `fast` di chuyển 2 bước mỗi lần
- Khi `fast` đến cuối, `slow` sẽ ở giữa list
- Điều kiện `fast.Next != nil && fast.Next.Next != nil` đảm bảo không gây lỗi null pointer

**Ví dụ:** Với `[1,2,3,4,5]`
```
Bước 0: slow=1, fast=1
Bước 1: slow=2, fast=3
Bước 2: slow=3, fast=5
fast.Next = nil → dừng
Middle = node 3
```

#### 3. Bước 2: Reverse Nửa Sau (dòng 24-26)
```go
secondHalf := slow.Next
slow.Next = nil // Ngắt kết nối giữa hai nửa
secondHalf = reverseList(secondHalf)
```

**Giải thích:**
- `slow.Next` là đầu của nửa sau
- `slow.Next = nil`: Ngắt kết nối giữa hai nửa để tránh cycle
- `reverseList`: Đảo ngược nửa sau

#### 4. Hàm reverseList (dòng 50-61)
```go
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
```

**Giải thích:**
- Sử dụng 3 con trỏ: `prev`, `curr`, `next`
- Đảo ngược từng node một cách tuần tự
- Trả về `prev` (node cuối cùng của list gốc, giờ là đầu của list đảo ngược)

**Ví dụ:** Reverse `[4,5]`
```
Bước 0: prev=nil, curr=4
Bước 1: next=5, 4.Next=nil, prev=4, curr=5
Bước 2: next=nil, 5.Next=4, prev=5, curr=nil
Kết quả: [5,4]
```

#### 5. Bước 3: Merge Hai Nửa (dòng 28-42)
```go
firstHalf := head
for secondHalf != nil {
    firstNext := firstHalf.Next
    secondNext := secondHalf.Next

    firstHalf.Next = secondHalf
    secondHalf.Next = firstNext

    firstHalf = firstNext
    secondHalf = secondNext
}
```

**Giải thích:**
- Duyệt qua nửa sau (ngắn hơn hoặc bằng nửa đầu)
- Với mỗi node trong nửa sau:
  - Lưu `firstNext` và `secondNext` để không mất kết nối
  - Kết nối: `firstHalf → secondHalf → firstNext`
  - Di chuyển cả hai con trỏ

**Ví dụ:** Merge `[1,2,3]` và `[5,4]`
```
Lặp 1:
  firstHalf=1, secondHalf=5
  firstNext=2, secondNext=4
  Kết nối: 1 → 5 → 2
  firstHalf=2, secondHalf=4

Lặp 2:
  firstHalf=2, secondHalf=4
  firstNext=3, secondNext=nil
  Kết nối: 2 → 4 → 3
  firstHalf=3, secondHalf=nil → dừng

Kết quả: 1 → 5 → 2 → 4 → 3
```

### Độ phức tạp

- **Thời gian**: O(n)
  - Tìm middle: O(n) - duyệt qua nửa list
  - Reverse nửa sau: O(n) - duyệt qua nửa list
  - Merge: O(n) - duyệt qua nửa sau (ngắn hơn)
  - Tổng: O(n)

- **Không gian**: O(1)
  - Chỉ sử dụng các con trỏ và biến
  - Không sử dụng thêm cấu trúc dữ liệu nào

### Tại sao thuật toán này đúng?

1. **Tìm middle chính xác**: Slow/fast pointers đảm bảo `slow` luôn ở giữa list khi `fast` đến cuối.

2. **Reverse đúng**: Thuật toán reverse chuẩn đảo ngược toàn bộ nửa sau.

3. **Merge đúng pattern**: 
   - Luôn lấy node đầu tiên từ nửa đầu
   - Sau đó lấy node đầu tiên từ nửa sau (đã reverse)
   - Lặp lại cho đến khi hết nửa sau
   - Pattern: L0 → Ln → L1 → Ln-1 → ...

4. **Xử lý edge cases**: 
   - List rỗng hoặc 1 node được xử lý ngay từ đầu
   - List có số lượng node chẵn/lẻ đều hoạt động đúng

## Test Cases

### Test Case 1: Example 1
```go
Input: [1,2,3,4]
Output: [1,4,2,3]
```
List có 4 nodes (chẵn).

### Test Case 2: Example 2
```go
Input: [1,2,3,4,5]
Output: [1,5,2,4,3]
```
List có 5 nodes (lẻ).

### Test Case 3: Single node
```go
Input: [1]
Output: [1]
```
Chỉ có 1 node, không cần sắp xếp lại.

### Test Case 4: Two nodes
```go
Input: [1,2]
Output: [1,2]
```
Có 2 nodes, pattern vẫn là L0 → L1.

### Test Case 5: Three nodes
```go
Input: [1,2,3]
Output: [1,3,2]
```
List có 3 nodes.

### Test Case 6: Six nodes
```go
Input: [1,2,3,4,5,6]
Output: [1,6,2,5,3,4]
```
List có số lượng node chẵn lớn hơn.

## Chạy test

Để chạy các test case:

```bash
go test ./internal/linked_list/reorder_list/
```

Hoặc chạy với verbose mode để xem chi tiết:

```bash
go test -v ./internal/linked_list/reorder_list/
```

## Mở rộng

### Reverse Linked List II (LeetCode 92)

Một biến thể là reverse một phần của linked list từ vị trí `left` đến `right`. Có thể sử dụng kỹ thuật tương tự nhưng chỉ reverse một phần.

### Reverse Nodes in k-Group (LeetCode 25)

Reverse linked list theo từng nhóm k node. Yêu cầu kỹ thuật reverse và merge phức tạp hơn.

### Ứng dụng thực tế

Bài toán này có ứng dụng trong:
- Xử lý và sắp xếp dữ liệu trong linked list
- Tối ưu hóa cấu trúc dữ liệu
- Thuật toán xử lý chuỗi và pattern matching
- Các bài toán liên quan đến palindrome trong linked list
