# Linked List Cycle

## Mô tả bài toán

Cho `head`, là đầu của một linked list, xác định xem linked list có chứa cycle (chu kỳ) hay không.

Có một cycle trong linked list nếu có một node nào đó trong list có thể được truy cập lại bằng cách liên tục theo con trỏ `next`. Trong nội bộ, `pos` được sử dụng để biểu thị chỉ số của node mà con trỏ `next` của tail trỏ đến. Lưu ý rằng `pos` không được truyền vào như một tham số.

Trả về `true` nếu linked list có cycle. Ngược lại, trả về `false`.

**Ví dụ:**

### Example 1:
- Input: `head = [3,2,0,-4]`, `pos = 1`
- Output: `true`
- Giải thích: Có một cycle trong linked list, nơi tail kết nối với node thứ 1 (0-indexed).

```
3 -> 2 -> 0 -> -4
     ^          |
     |__________|
```

### Example 2:
- Input: `head = [1,2]`, `pos = 0`
- Output: `true`
- Giải thích: Có một cycle trong linked list, nơi tail kết nối với node thứ 0.

```
1 -> 2
^    |
|____|
```

### Example 3:
- Input: `head = [1]`, `pos = -1`
- Output: `false`
- Giải thích: Không có cycle trong linked list.

```
1 -> nil
```

**Ràng buộc:**
- Số lượng nodes trong list nằm trong khoảng `[0, 10^4]`
- `-10^5 <= Node.val <= 10^5`
- `pos` là -1 hoặc một chỉ số hợp lệ trong linked-list

**Follow up:** Bạn có thể giải quyết bài toán này với O(1) (tức là hằng số) bộ nhớ không?

## Phân tích thuật toán

### Cách tiếp cận 1: Hash Set (O(n) time, O(n) space)

Duyệt qua linked list và lưu trữ mỗi node đã gặp vào một hash set. Nếu gặp một node đã có trong set, có cycle.

**Độ phức tạp:**
- Thời gian: O(n)
- Không gian: O(n)

### Cách tiếp cận 2: Floyd's Cycle Detection - Two Pointers (O(n) time, O(1) space) ⭐ Tối ưu

Sử dụng hai con trỏ di chuyển với tốc độ khác nhau:
- **Slow pointer**: Di chuyển 1 bước mỗi lần
- **Fast pointer**: Di chuyển 2 bước mỗi lần

**Nguyên lý:**
- Nếu không có cycle: Fast pointer sẽ đến `null` trước
- Nếu có cycle: Fast pointer sẽ đuổi kịp slow pointer tại một điểm nào đó trong cycle

**Tại sao thuật toán này hoạt động?**

Hãy tưởng tượng hai người chạy trên một đường tròn:
- Người chạy nhanh (fast) chạy với tốc độ gấp đôi người chạy chậm (slow)
- Nếu có cycle (đường tròn), người chạy nhanh sẽ đuổi kịp người chạy chậm
- Nếu không có cycle (đường thẳng), người chạy nhanh sẽ đến đích trước

**Độ phức tạp:**
- Thời gian: O(n) - trong trường hợp xấu nhất, fast pointer sẽ duyệt qua tối đa 2n nodes
- Không gian: O(1) - chỉ sử dụng hai con trỏ

### Ví dụ minh họa

Với input: `head = [3,2,0,-4]`, `pos = 1`

```
Linked list:
3 -> 2 -> 0 -> -4
     ^          |
     |__________|

Bước 0:
  slow = 3 (head)
  fast = 2 (head.Next)

Bước 1:
  slow = 2 (slow.Next)
  fast = 0 (fast.Next.Next)
  slow != fast, tiếp tục

Bước 2:
  slow = 0 (slow.Next)
  fast = 2 (fast.Next.Next) - quay lại node 2 do cycle
  slow != fast, tiếp tục

Bước 3:
  slow = -4 (slow.Next)
  fast = -4 (fast.Next.Next)
  slow == fast ✓ → Có cycle!
```

Với input: `head = [1]`, `pos = -1`

```
Linked list:
1 -> nil

Bước 0:
  slow = 1 (head)
  fast = nil (head.Next)

Kiểm tra: fast == nil → Không có cycle
```

## Giải thích code

### Cấu trúc ListNode

```1:6:internal/linked_list/linked_list_cycle/linked_list_cycle.go
package linked_list

// ListNode định nghĩa một node trong linked list
type ListNode struct {
	Val  int
	Next *ListNode
}
```

- `Val`: Giá trị của node
- `Next`: Con trỏ đến node tiếp theo

### Hàm hasCycle

```8:44:internal/linked_list/linked_list_cycle/linked_list_cycle.go
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
```

### Chi tiết từng phần

#### 1. Kiểm tra trường hợp biên (dòng 16-19)
```go
if head == nil || head.Next == nil {
    return false
}
```
- Nếu list rỗng hoặc chỉ có 1 node, không thể có cycle
- Trả về `false` ngay lập tức

#### 2. Khởi tạo hai con trỏ (dòng 21-23)
```go
slow := head      // Di chuyển 1 bước mỗi lần
fast := head.Next // Di chuyển 2 bước mỗi lần
```
- `slow`: Bắt đầu từ `head`, di chuyển 1 node mỗi lần
- `fast`: Bắt đầu từ `head.Next`, di chuyển 2 nodes mỗi lần
- Bắt đầu từ `head.Next` để tránh trường hợp `slow == fast` ngay từ đầu

#### 3. Vòng lặp chính (dòng 25-36)
```go
for fast != nil && fast.Next != nil {
    if slow == fast {
        return true
    }
    slow = slow.Next
    fast = fast.Next.Next
}
```

**Giải thích:**
- Điều kiện `fast != nil && fast.Next != nil`: Đảm bảo `fast.Next.Next` không gây lỗi null pointer
- `if slow == fast`: Nếu hai con trỏ gặp nhau, có cycle → trả về `true`
- `slow = slow.Next`: Di chuyển slow 1 bước
- `fast = fast.Next.Next`: Di chuyển fast 2 bước

#### 4. Trả về kết quả (dòng 39)
```go
return false
```
- Nếu vòng lặp kết thúc (fast đã đến null), không có cycle

### Độ phức tạp

- **Thời gian**: O(n)
  - Trong trường hợp xấu nhất, fast pointer sẽ duyệt qua tối đa 2n nodes
  - Khi có cycle, fast sẽ đuổi kịp slow sau tối đa n bước
  - Khi không có cycle, fast sẽ đến null sau n/2 bước

- **Không gian**: O(1)
  - Chỉ sử dụng hai con trỏ `slow` và `fast`
  - Không sử dụng thêm cấu trúc dữ liệu nào

### Tại sao thuật toán này đúng?

**Chứng minh toán học:**

Giả sử:
- Khoảng cách từ head đến điểm bắt đầu cycle: `a`
- Độ dài của cycle: `c`
- Khoảng cách từ điểm bắt đầu cycle đến điểm gặp nhau: `b`

Khi slow và fast gặp nhau:
- Slow đã đi được: `a + b` bước
- Fast đã đi được: `a + b + k*c` bước (với k là số vòng đã chạy)

Vì fast đi nhanh gấp đôi slow:
```
2(a + b) = a + b + k*c
a + b = k*c
```

Điều này có nghĩa là khoảng cách từ head đến điểm gặp nhau bằng bội số của độ dài cycle, chứng tỏ có cycle.

## Test Cases

### Test Case 1: Example 1
```go
Input: head = [3,2,0,-4], pos = 1
Output: true
```
Cycle tại node thứ 1 (giá trị 2).

### Test Case 2: Example 2
```go
Input: head = [1,2], pos = 0
Output: true
```
Cycle tại node thứ 0 (giá trị 1).

### Test Case 3: Example 3
```go
Input: head = [1], pos = -1
Output: false
```
Không có cycle.

### Test Case 4: Empty list
```go
Input: head = [], pos = -1
Output: false
```
List rỗng.

### Test Case 5: Two nodes no cycle
```go
Input: head = [1,2], pos = -1
Output: false
```
Hai nodes không có cycle.

### Test Case 6: Single node self cycle
```go
Input: head = [1], pos = 0
Output: true
```
Node tự trỏ về chính nó.

### Test Case 7: Long list with cycle
```go
Input: head = [1,2,3,4,5,6,7,8,9,10], pos = 5
Output: true
```
List dài với cycle tại node thứ 5.

## Chạy test

Để chạy các test case:

```bash
go test ./internal/linked_list/linked_list_cycle/
```

Hoặc chạy với verbose mode để xem chi tiết:

```bash
go test -v ./internal/linked_list/linked_list_cycle/
```

## Mở rộng

### Linked List Cycle II (LeetCode 142)

Một biến thể của bài toán này là tìm **điểm bắt đầu của cycle** (nếu có). Giải pháp vẫn sử dụng Two Pointers nhưng cần thêm một bước:
1. Sử dụng Floyd's Cycle Detection để tìm điểm gặp nhau
2. Đặt một con trỏ về head, một con trỏ tại điểm gặp nhau
3. Di chuyển cả hai cùng tốc độ (1 bước) cho đến khi gặp nhau
4. Điểm gặp nhau lần này chính là điểm bắt đầu của cycle

### Ứng dụng thực tế

Thuật toán Floyd's Cycle Detection được sử dụng trong:
- Phát hiện cycle trong linked list
- Phát hiện cycle trong đồ thị
- Thuật toán Pollard's rho để phân tích thừa số
- Phát hiện lặp vô hạn trong các hàm đệ quy
