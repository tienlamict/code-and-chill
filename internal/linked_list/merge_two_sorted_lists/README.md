# Merge Two Sorted Lists

## Mô tả bài toán

Cho hai danh sách liên kết đã được sắp xếp `list1` và `list2`, hợp nhất chúng thành một danh sách đã sắp xếp. Danh sách mới nên được tạo bằng cách nối các node của hai danh sách ban đầu.

Trả về head của danh sách đã hợp nhất.

**Ví dụ:**

### Example 1:
- Input: `list1 = [1,2,4]`, `list2 = [1,3,4]`
- Output: `[1,1,2,3,4,4]`
- Giải thích: Hợp nhất hai danh sách đã sắp xếp thành một danh sách đã sắp xếp.

```
list1: 1 -> 2 -> 4
list2: 1 -> 3 -> 4
Kết quả: 1 -> 1 -> 2 -> 3 -> 4 -> 4
```

### Example 2:
- Input: `list1 = []`, `list2 = []`
- Output: `[]`
- Giải thích: Cả hai danh sách đều rỗng, kết quả cũng rỗng.

### Example 3:
- Input: `list1 = []`, `list2 = [0]`
- Output: `[0]`
- Giải thích: Danh sách 1 rỗng, kết quả là danh sách 2.

```
list1: (rỗng)
list2: 0
Kết quả: 0
```

**Ràng buộc:**
- Số lượng nodes trong cả hai list nằm trong khoảng `[0, 50]`
- `-100 <= Node.val <= 100`
- Cả `list1` và `list2` đều được sắp xếp theo thứ tự không giảm (non-decreasing)

## Phân tích thuật toán

### Cách tiếp cận 1: Sử dụng mảng phụ (O(n+m) time, O(n+m) space)

1. Chuyển đổi cả hai linked list thành mảng
2. Hợp nhất hai mảng đã sắp xếp
3. Tạo linked list mới từ mảng đã hợp nhất

**Độ phức tạp:**
- Thời gian: O(n + m) - duyệt qua cả hai list và hợp nhất
- Không gian: O(n + m) - cần mảng phụ để lưu trữ

**Nhược điểm:** Sử dụng thêm không gian O(n+m), không tận dụng được cấu trúc linked list

### Cách tiếp cận 2: Two Pointers với Dummy Node (O(n+m) time, O(1) space) ⭐ Tối ưu

Sử dụng kỹ thuật **Two Pointers** với một dummy node để hợp nhất trực tiếp:

1. **Tạo dummy node**: Giúp đơn giản hóa việc xử lý edge cases (list rỗng, chỉ có một list)
2. **So sánh và chọn**: So sánh giá trị của node hiện tại trong hai list, chọn node nhỏ hơn
3. **Nối node**: Nối node đã chọn vào kết quả và di chuyển con trỏ tương ứng
4. **Lặp lại**: Tiếp tục cho đến khi một trong hai list kết thúc
5. **Nối phần còn lại**: Nối phần còn lại của list chưa kết thúc vào kết quả

**Độ phức tạp:**
- Thời gian: O(n + m) - duyệt qua tất cả nodes của cả hai list
- Không gian: O(1) - chỉ sử dụng thêm một vài con trỏ

**Ưu điểm:** 
- Không sử dụng thêm không gian
- Tận dụng được cấu trúc linked list
- Code đơn giản và dễ hiểu

### Tại sao cần dummy node?

Dummy node giúp:
- **Xử lý edge cases dễ dàng**: Khi một hoặc cả hai list rỗng
- **Đơn giản hóa code**: Không cần kiểm tra xem kết quả có rỗng hay không trước khi thêm node đầu tiên
- **Thống nhất logic**: Luôn có một node để nối vào (`current.Next`)

### Ví dụ minh họa

Với input: `list1 = [1,2,4]`, `list2 = [1,3,4]`

```
Bước 0: Khởi tạo
list1: 1 -> 2 -> 4
list2: 1 -> 3 -> 4
dummy -> nil
current = dummy

Bước 1: So sánh 1 và 1, chọn 1 từ list1 (hoặc list2, vì bằng nhau)
list1: 2 -> 4
list2: 1 -> 3 -> 4
dummy -> 1 -> nil
current = 1

Bước 2: So sánh 2 và 1, chọn 1 từ list2
list1: 2 -> 4
list2: 3 -> 4
dummy -> 1 -> 1 -> nil
current = 1 (node thứ hai)

Bước 3: So sánh 2 và 3, chọn 2 từ list1
list1: 4
list2: 3 -> 4
dummy -> 1 -> 1 -> 2 -> nil
current = 2

Bước 4: So sánh 4 và 3, chọn 3 từ list2
list1: 4
list2: 4
dummy -> 1 -> 1 -> 2 -> 3 -> nil
current = 3

Bước 5: So sánh 4 và 4, chọn 4 từ list1 (hoặc list2)
list1: nil
list2: 4
dummy -> 1 -> 1 -> 2 -> 3 -> 4 -> nil
current = 4

Bước 6: list1 đã kết thúc, nối phần còn lại của list2
list1: nil
list2: nil
dummy -> 1 -> 1 -> 2 -> 3 -> 4 -> 4 -> nil

Kết quả: [1,1,2,3,4,4]
```

Với input: `list1 = []`, `list2 = [0]`

```
Bước 0: Khởi tạo
list1: nil
list2: 0
dummy -> nil
current = dummy

Bước 1: list1 rỗng, nối phần còn lại của list2
list1: nil
list2: nil
dummy -> 0 -> nil

Kết quả: [0]
```

## Giải thích code

### Cấu trúc ListNode

```1:7:internal/linked_list/merge_two_sorted_lists/merge_two_sorted_lists.go
package linked_list

// ListNode định nghĩa một node trong linked list
type ListNode struct {
	Val  int
	Next *ListNode
}
```

- `Val`: Giá trị của node
- `Next`: Con trỏ đến node tiếp theo

### Hàm mergeTwoLists

```9:50:internal/linked_list/merge_two_sorted_lists/merge_two_sorted_lists.go
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
```

### Chi tiết từng phần

#### 1. Tạo dummy node (dòng 20-21)
```go
dummy := &ListNode{}
current := dummy
```
- Tạo một node giả để bắt đầu kết quả
- `current` trỏ đến node cuối cùng của kết quả hiện tại
- Giúp xử lý trường hợp cả hai list đều rỗng

#### 2. Vòng lặp hợp nhất (dòng 23-35)
```go
for list1 != nil && list2 != nil {
    if list1.Val <= list2.Val {
        current.Next = list1
        list1 = list1.Next
    } else {
        current.Next = list2
        list2 = list2.Next
    }
    current = current.Next
}
```
- Duyệt cả hai list đồng thời
- So sánh giá trị của node hiện tại trong hai list
- Chọn node nhỏ hơn (hoặc bằng) và nối vào kết quả
- Di chuyển con trỏ của list đã chọn và con trỏ kết quả

**Lưu ý:** Sử dụng `<=` để đảm bảo tính ổn định (stable) - nếu hai giá trị bằng nhau, ưu tiên node từ list1

#### 3. Nối phần còn lại (dòng 37-42)
```go
if list1 != nil {
    current.Next = list1
} else {
    current.Next = list2
}
```
- Sau khi một trong hai list kết thúc, nối phần còn lại của list kia
- Chỉ một trong hai điều kiện sẽ đúng (hoặc cả hai đều nil nếu cả hai list đều kết thúc cùng lúc)

#### 4. Trả về kết quả (dòng 44)
```go
return dummy.Next
```
- Trả về `dummy.Next` là head của list đã hợp nhất
- Bỏ qua dummy node vì nó chỉ là node phụ trợ

### Độ phức tạp

- **Thời gian**: O(n + m)
  - Duyệt qua tất cả nodes của cả hai list
  - Trong trường hợp xấu nhất, cần duyệt qua tất cả n + m nodes

- **Không gian**: O(1)
  - Chỉ sử dụng thêm một vài con trỏ (`dummy`, `current`)
  - Không sử dụng thêm cấu trúc dữ liệu nào
  - Không tạo node mới, chỉ nối lại các node hiện có

### Tại sao thuật toán này đúng?

**Chứng minh:**

1. **Tính đúng đắn của việc so sánh:**
   - Vì cả hai list đều đã được sắp xếp, node nhỏ nhất còn lại trong hai list chắc chắn là node nhỏ nhất tiếp theo trong kết quả
   - Bằng cách luôn chọn node nhỏ hơn, ta đảm bảo kết quả được sắp xếp đúng

2. **Tính đầy đủ:**
   - Vòng lặp xử lý tất cả nodes cho đến khi một trong hai list kết thúc
   - Phần còn lại được xử lý bằng cách nối trực tiếp (vì đã được sắp xếp)

3. **Xử lý edge cases:**
   - Nếu cả hai list rỗng: `dummy.Next` sẽ là `nil`, trả về `nil` ✓
   - Nếu một list rỗng: Nối trực tiếp list còn lại ✓
   - Nếu cả hai list có giá trị: Hợp nhất đúng thứ tự ✓

## Test Cases

### Test Case 1: Example 1
```go
Input: list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]
```
Hợp nhất hai list có độ dài bằng nhau với các giá trị xen kẽ.

### Test Case 2: Example 2
```go
Input: list1 = [], list2 = []
Output: []
```
Cả hai list đều rỗng.

### Test Case 3: Example 3
```go
Input: list1 = [], list2 = [0]
Output: [0]
```
Một list rỗng, kết quả là list còn lại.

### Test Case 4: List1 có giá trị nhỏ hơn
```go
Input: list1 = [1,2,3], list2 = [4,5,6]
Output: [1,2,3,4,5,6]
```
Tất cả giá trị trong list1 nhỏ hơn list2.

### Test Case 5: List2 có giá trị nhỏ hơn
```go
Input: list1 = [4,5,6], list2 = [1,2,3]
Output: [1,2,3,4,5,6]
```
Tất cả giá trị trong list2 nhỏ hơn list1.

### Test Case 6: Giá trị xen kẽ
```go
Input: list1 = [1,3,5], list2 = [2,4,6]
Output: [1,2,3,4,5,6]
```
Giá trị trong hai list xen kẽ nhau.

### Test Case 7: List1 dài hơn
```go
Input: list1 = [1,2,3,4,5], list2 = [1,2]
Output: [1,1,2,2,3,4,5]
```
List1 có nhiều nodes hơn list2.

### Test Case 8: List2 dài hơn
```go
Input: list1 = [1,2], list2 = [1,2,3,4,5]
Output: [1,1,2,2,3,4,5]
```
List2 có nhiều nodes hơn list1.

### Test Case 9: Mỗi list một node
```go
Input: list1 = [1], list2 = [2]
Output: [1,2]
```
Mỗi list chỉ có một node.

### Test Case 10: Giá trị trùng lặp
```go
Input: list1 = [1,1,1], list2 = [1,1,1]
Output: [1,1,1,1,1,1]
```
Cả hai list đều có giá trị trùng lặp.

### Test Case 11: Giá trị âm
```go
Input: list1 = [-5,-3,-1], list2 = [-4,-2,0]
Output: [-5,-4,-3,-2,-1,0]
```
Hợp nhất list có giá trị âm.

### Test Case 12: Giá trị dương và âm
```go
Input: list1 = [-1,0,1], list2 = [-2,2]
Output: [-2,-1,0,1,2]
```
Hợp nhất list có cả giá trị dương và âm.

## Chạy test

Để chạy các test case:

```bash
go test ./internal/linked_list/merge_two_sorted_lists/
```

Hoặc chạy với verbose mode để xem chi tiết:

```bash
go test -v ./internal/linked_list/merge_two_sorted_lists/
```

## Mở rộng

### Biến thể: Merge K Sorted Lists

Nếu bài toán yêu cầu hợp nhất K danh sách đã sắp xếp, có thể sử dụng:

1. **Cách tiếp cận 1: Merge từng cặp** (O(k * n) time)
   - Hợp nhất từng cặp list một
   - Độ phức tạp: O(k * n) với n là tổng số nodes

2. **Cách tiếp cận 2: Sử dụng Min Heap** (O(n * log k) time) ⭐ Tối ưu
   - Đưa head của mỗi list vào heap
   - Lấy node nhỏ nhất từ heap, thêm vào kết quả
   - Thêm node tiếp theo của list đó vào heap
   - Lặp lại cho đến khi heap rỗng

```go
func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 {
        return nil
    }
    
    // Sử dụng min heap để chọn node nhỏ nhất
    // Implementation chi tiết tùy thuộc vào cấu trúc heap
    // ...
}
```

### Ứng dụng thực tế

Thuật toán merge hai sorted lists được sử dụng trong:

- **Merge Sort cho linked list**: Chia list thành hai phần, sắp xếp từng phần, rồi hợp nhất
- **Hợp nhất kết quả từ nhiều nguồn đã sắp xếp**: Ví dụ hợp nhất kết quả tìm kiếm từ nhiều database
- **External sorting**: Khi dữ liệu quá lớn để load vào memory, merge các chunk đã sắp xếp
- **Priority queue**: Hợp nhất các queue đã được sắp xếp theo priority

### Tips và Tricks

1. **Luôn sử dụng dummy node** khi xây dựng linked list mới để đơn giản hóa code
2. **Xử lý phần còn lại**: Sau khi một list kết thúc, chỉ cần nối phần còn lại của list kia (không cần duyệt tiếp)
3. **So sánh với <=**: Sử dụng `<=` thay vì `<` để đảm bảo tính ổn định khi hai giá trị bằng nhau
4. **Không tạo node mới**: Chỉ nối lại các node hiện có để tiết kiệm memory
5. **Edge cases**: Luôn kiểm tra trường hợp một hoặc cả hai list rỗng

### So sánh với các cách tiếp cận khác

| Cách tiếp cận | Thời gian | Không gian | Ưu điểm | Nhược điểm |
|--------------|-----------|------------|---------|------------|
| Mảng phụ | O(n+m) | O(n+m) | Dễ hiểu | Tốn memory |
| Two Pointers | O(n+m) | O(1) | Tối ưu memory | Cần hiểu linked list |
| Recursive | O(n+m) | O(n+m) | Code ngắn gọn | Stack overflow với list dài |

### Lưu ý về Recursive Approach

Mặc dù có thể giải bằng đệ quy, nhưng cách này có nhược điểm:

```go
func mergeTwoListsRecursive(list1 *ListNode, list2 *ListNode) *ListNode {
    if list1 == nil {
        return list2
    }
    if list2 == nil {
        return list1
    }
    
    if list1.Val <= list2.Val {
        list1.Next = mergeTwoListsRecursive(list1.Next, list2)
        return list1
    } else {
        list2.Next = mergeTwoListsRecursive(list1, list2.Next)
        return list2
    }
}
```

**Nhược điểm:**
- Sử dụng O(n+m) không gian cho call stack
- Có thể gây stack overflow với list dài
- Khó debug hơn iterative approach

**Khi nào nên dùng:**
- List ngắn (đảm bảo không stack overflow)
- Code ngắn gọn quan trọng hơn performance
