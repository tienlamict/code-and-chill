# Remove Nth Node From End of List

## Mô tả bài toán

Cho `head` là đầu của một linked list, xóa node thứ `n` từ cuối danh sách và trả về head của list mới.

**Ví dụ:**

### Example 1:
- Input: `head = [1,2,3,4,5]`, `n = 2`
- Output: `[1,2,3,5]`
- Giải thích: Xóa node thứ 2 từ cuối (node có giá trị 4).

```
Trước: 1 -> 2 -> 3 -> 4 -> 5
Sau:   1 -> 2 -> 3 -> 5
```

### Example 2:
- Input: `head = [1]`, `n = 1`
- Output: `[]`
- Giải thích: Xóa node duy nhất trong list.

```
Trước: 1
Sau:   (rỗng)
```

### Example 3:
- Input: `head = [1,2]`, `n = 1`
- Output: `[1]`
- Giải thích: Xóa node cuối cùng (node có giá trị 2).

```
Trước: 1 -> 2
Sau:   1
```

**Ràng buộc:**
- Số lượng nodes trong list là `sz`
- `1 <= sz <= 30`
- `0 <= Node.val <= 100`
- `1 <= n <= sz`

**Follow up:** Bạn có thể giải quyết bài toán này trong một lần duyệt (one pass) không?

## Phân tích thuật toán

### Cách tiếp cận 1: Two Pass (O(n) time, O(1) space)

1. Duyệt list một lần để đếm tổng số nodes
2. Tính vị trí của node cần xóa từ đầu: `position = length - n`
3. Duyệt lại list đến vị trí `position - 1` và xóa node

**Độ phức tạp:**
- Thời gian: O(n) - duyệt list 2 lần
- Không gian: O(1)

**Nhược điểm:** Cần duyệt list 2 lần

### Cách tiếp cận 2: Two Pointers - One Pass (O(n) time, O(1) space) ⭐ Tối ưu

Sử dụng kỹ thuật **Two Pointers** với một dummy node:

1. **Tạo dummy node**: Giúp xử lý trường hợp xóa head một cách dễ dàng
2. **Di chuyển fast pointer n+1 bước trước**: Đảm bảo khi fast đến cuối, slow sẽ ở node trước node cần xóa
3. **Di chuyển cả hai pointers cùng lúc**: Cho đến khi fast đến cuối list
4. **Xóa node**: Cập nhật `slow.Next = slow.Next.Next`

**Độ phức tạp:**
- Thời gian: O(n) - chỉ duyệt list 1 lần
- Không gian: O(1) - chỉ sử dụng thêm một vài con trỏ

**Ưu điểm:** Chỉ cần duyệt list 1 lần, đáp ứng yêu cầu follow up

### Tại sao cần dummy node?

Khi xóa node trong linked list, ta cần truy cập node **trước** node cần xóa. Tuy nhiên, nếu node cần xóa là head, không có node nào trước nó. Dummy node giải quyết vấn đề này bằng cách:

- Đặt dummy node trước head: `dummy -> head -> ...`
- Bây giờ luôn có node trước node cần xóa (kể cả head)
- Trả về `dummy.Next` là head mới

### Ví dụ minh họa

Với input: `head = [1,2,3,4,5]`, `n = 2`

```
Linked list ban đầu:
1 -> 2 -> 3 -> 4 -> 5

Bước 1: Tạo dummy node
dummy -> 1 -> 2 -> 3 -> 4 -> 5
slow = dummy
fast = dummy

Bước 2: Di chuyển fast n+1 = 3 bước trước
dummy -> 1 -> 2 -> 3 -> 4 -> 5
slow    fast
        (sau 3 bước)

Bước 3: Di chuyển cả hai cùng lúc cho đến khi fast đến cuối
dummy -> 1 -> 2 -> 3 -> 4 -> 5
         slow         fast
         (sau khi fast đến cuối)

Bước 4: Xóa node (slow.Next = slow.Next.Next)
dummy -> 1 -> 2 -> 3 -> 5
         slow         

Kết quả: [1,2,3,5]
```

Với input: `head = [1]`, `n = 1` (xóa head)

```
Linked list ban đầu:
1

Bước 1: Tạo dummy node
dummy -> 1
slow = dummy
fast = dummy

Bước 2: Di chuyển fast n+1 = 2 bước trước
dummy -> 1 -> nil
slow    fast
        (sau 2 bước, fast = nil)

Bước 3: fast đã là nil, không cần di chuyển thêm
slow vẫn ở dummy

Bước 4: Xóa node (slow.Next = slow.Next.Next = nil)
dummy -> nil

Kết quả: [] (rỗng)
```

## Giải thích code

### Cấu trúc ListNode

```1:7:internal/linked_list/remove_n_th_node_from_end_of_list/remove_n_th_node_from_end_of_list.go
package linked_list

// ListNode định nghĩa một node trong linked list
type ListNode struct {
	Val  int
	Next *ListNode
}
```

- `Val`: Giá trị của node
- `Next`: Con trỏ đến node tiếp theo

### Hàm removeNthFromEnd

```9:44:internal/linked_list/remove_n_th_node_from_end_of_list/remove_n_th_node_from_end_of_list.go
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
```

### Chi tiết từng phần

#### 1. Tạo dummy node (dòng 20)
```go
dummy := &ListNode{Next: head}
```
- Tạo một node giả trước head
- Giúp xử lý trường hợp xóa head một cách thống nhất
- `dummy.Next` trỏ đến head thực sự

#### 2. Khởi tạo hai con trỏ (dòng 22-24)
```go
slow := dummy
fast := dummy
```
- Cả hai con trỏ bắt đầu từ dummy node
- `slow`: Sẽ dừng ở node trước node cần xóa
- `fast`: Sẽ đi trước để đánh dấu khoảng cách

#### 3. Di chuyển fast pointer n+1 bước (dòng 26-29)
```go
for i := 0; i <= n; i++ {
    fast = fast.Next
}
```
- Di chuyển fast `n+1` bước (không phải `n` bước)
- Lý do: Khi fast đến cuối, slow sẽ ở node trước node cần xóa
- Ví dụ: nếu `n=2`, fast đi 3 bước, khi fast đến cuối, slow sẽ ở node trước node thứ 2 từ cuối

#### 4. Di chuyển cả hai con trỏ cùng lúc (dòng 31-35)
```go
for fast != nil {
    slow = slow.Next
    fast = fast.Next
}
```
- Di chuyển cả hai con trỏ cùng tốc độ (1 bước mỗi lần)
- Dừng khi `fast == nil` (đã đến cuối list)
- Lúc này `slow` đang ở node trước node cần xóa

#### 5. Xóa node (dòng 37)
```go
slow.Next = slow.Next.Next
```
- Bỏ qua node cần xóa bằng cách cập nhật con trỏ
- Node bị xóa sẽ được garbage collector tự động dọn dẹp

#### 6. Trả về head mới (dòng 39)
```go
return dummy.Next
```
- Trả về `dummy.Next` là head mới
- Nếu xóa head cũ, `dummy.Next` sẽ trỏ đến node mới
- Nếu không xóa head, `dummy.Next` vẫn là head cũ

### Độ phức tạp

- **Thời gian**: O(n)
  - Duyệt qua list một lần duy nhất
  - Trong trường hợp xấu nhất, cần duyệt qua tất cả n nodes

- **Không gian**: O(1)
  - Chỉ sử dụng thêm một vài con trỏ (`dummy`, `slow`, `fast`)
  - Không sử dụng thêm cấu trúc dữ liệu nào

### Tại sao thuật toán này đúng?

**Chứng minh:**

Giả sử:
- Tổng số nodes: `L`
- Node cần xóa là node thứ `n` từ cuối
- Vị trí của node cần xóa từ đầu: `L - n + 1`
- Node trước node cần xóa: `L - n`

**Bước 1:** Fast pointer di chuyển `n+1` bước từ dummy
- Fast ở vị trí: `n+1` (từ dummy)

**Bước 2:** Di chuyển cả hai cùng lúc cho đến khi fast đến cuối
- Khoảng cách giữa fast và slow: `n+1` (không đổi)
- Khi fast ở cuối (vị trí `L+1` từ dummy), slow ở vị trí: `(L+1) - (n+1) = L - n`
- Đây chính là node trước node cần xóa! ✓

## Test Cases

### Test Case 1: Example 1
```go
Input: head = [1,2,3,4,5], n = 2
Output: [1,2,3,5]
```
Xóa node thứ 2 từ cuối (node có giá trị 4).

### Test Case 2: Example 2
```go
Input: head = [1], n = 1
Output: []
```
Xóa node duy nhất, list trở thành rỗng.

### Test Case 3: Example 3
```go
Input: head = [1,2], n = 1
Output: [1]
```
Xóa node cuối cùng.

### Test Case 4: Xóa head (multiple nodes)
```go
Input: head = [1,2,3,4,5], n = 5
Output: [2,3,4,5]
```
Xóa node đầu tiên (head).

### Test Case 5: Xóa node ở giữa
```go
Input: head = [1,2,3,4,5], n = 3
Output: [1,2,4,5]
```
Xóa node ở giữa (node có giá trị 3).

### Test Case 6: List dài
```go
Input: head = [1,2,3,4,5,6,7,8,9,10], n = 4
Output: [1,2,3,4,5,6,8,9,10]
```
Xóa node thứ 4 từ cuối trong list dài.

### Test Case 7: Hai nodes - xóa head
```go
Input: head = [1,2], n = 2
Output: [2]
```
Xóa head trong list có 2 nodes.

### Test Case 8: Ba nodes - các trường hợp
```go
Input: head = [1,2,3], n = 1  // Xóa cuối
Output: [1,2]

Input: head = [1,2,3], n = 2  // Xóa giữa
Output: [1,3]

Input: head = [1,2,3], n = 3  // Xóa đầu
Output: [2,3]
```

## Chạy test

Để chạy các test case:

```bash
go test ./internal/linked_list/remove_n_th_node_from_end_of_list/
```

Hoặc chạy với verbose mode để xem chi tiết:

```bash
go test -v ./internal/linked_list/remove_n_th_node_from_end_of_list/
```

## Mở rộng

### Biến thể: Xóa node thứ k từ đầu

Nếu bài toán yêu cầu xóa node thứ `k` từ **đầu** (không phải từ cuối), giải pháp sẽ đơn giản hơn:

```go
func removeNthFromStart(head *ListNode, k int) *ListNode {
    dummy := &ListNode{Next: head}
    current := dummy
    
    // Di chuyển đến node trước node cần xóa
    for i := 1; i < k; i++ {
        current = current.Next
    }
    
    // Xóa node
    current.Next = current.Next.Next
    
    return dummy.Next
}
```

### Ứng dụng thực tế

Thuật toán Two Pointers được sử dụng trong nhiều bài toán linked list:
- Tìm middle node
- Phát hiện cycle
- Tìm điểm giao nhau của hai list
- Reverse linked list
- Merge sorted lists

### Tips và Tricks

1. **Luôn sử dụng dummy node** khi cần xóa node trong linked list để xử lý edge cases dễ dàng hơn
2. **Khoảng cách giữa hai pointers** là chìa khóa: với bài toán "từ cuối", cần giữ khoảng cách `n+1`
3. **Kiểm tra null pointer**: Luôn đảm bảo `slow.Next` không phải `nil` trước khi truy cập `slow.Next.Next`
4. **One pass vs Two pass**: Nếu không có yêu cầu one pass, two pass đơn giản hơn và dễ hiểu hơn

