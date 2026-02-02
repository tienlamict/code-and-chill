# Merge k Sorted Lists

## Mô tả bài toán

Cho một mảng gồm `k` linked lists, mỗi linked list đã được sắp xếp theo thứ tự tăng dần.

Hợp nhất tất cả các linked lists thành một linked list đã được sắp xếp và trả về nó.

**Ví dụ:**

### Example 1:
- Input: `lists = [[1,4,5],[1,3,4],[2,6]]`
- Output: `[1,1,2,3,4,4,5,6]`
- Giải thích: Các linked lists là:
  ```
  [
    1->4->5,
    1->3->4,
    2->6
  ]
  ```
  Hợp nhất chúng thành một linked list đã sắp xếp:
  ```
  1->1->2->3->4->4->5->6
  ```

### Example 2:
- Input: `lists = []`
- Output: `[]`

### Example 3:
- Input: `lists = [[]]`
- Output: `[]`

**Ràng buộc:**
- `k == lists.length`
- `0 <= k <= 10^4`
- `0 <= lists[i].length <= 500`
- `-10^4 <= lists[i][j] <= 10^4`
- `lists[i]` được sắp xếp theo thứ tự tăng dần
- Tổng độ dài của tất cả `lists[i]` sẽ không vượt quá `10^4`

## Phân tích thuật toán

### Cách tiếp cận 1: Brute Force (O(n log n) time, O(n) space)

Thu thập tất cả các giá trị từ tất cả lists, sắp xếp chúng, sau đó tạo một linked list mới.

**Độ phức tạp:**
- Thời gian: O(n log n) - n là tổng số nodes
- Không gian: O(n) - lưu trữ tất cả giá trị

### Cách tiếp cận 2: Compare One by One (O(k * n) time, O(1) space)

Mỗi lần so sánh head của tất cả lists, chọn node nhỏ nhất, thêm vào kết quả.

**Độ phức tạp:**
- Thời gian: O(k * n) - với mỗi node, so sánh k lists
- Không gian: O(1) - không sử dụng thêm không gian

### Cách tiếp cận 3: Priority Queue / Min Heap (O(n log k) time, O(k) space)

Sử dụng min heap để luôn lấy node nhỏ nhất từ k lists.

**Độ phức tạp:**
- Thời gian: O(n log k) - mỗi node được push/pop từ heap
- Không gian: O(k) - heap chứa tối đa k nodes

### Cách tiếp cận 4: Divide and Conquer (O(n log k) time, O(log k) space) ⭐ Tối ưu

Chia mảng lists thành hai nửa, đệ quy merge từng nửa, sau đó merge hai kết quả lại.

**Ý tưởng:**
- Chia nhỏ vấn đề: Merge k lists = Merge 2 groups của k/2 lists
- Đệ quy: Tiếp tục chia cho đến khi còn 1 list
- Kết hợp: Merge 2 lists đã được merge

**Độ phức tạp:**
- Thời gian: O(n log k)
  - Có log k levels (chia đôi k lần)
  - Mỗi level merge tất cả n nodes
  - Tổng: O(n log k)
- Không gian: O(log k) - call stack cho đệ quy

### Ví dụ minh họa

Với input: `lists = [[1,4,5],[1,3,4],[2,6]]`

```
Level 0: [[1,4,5], [1,3,4], [2,6]]
         Chia thành 2 nhóm:
         Left: [[1,4,5], [1,3,4]]
         Right: [[2,6]]

Level 1 (Left):
  [[1,4,5], [1,3,4]]
  Chia thành: [1,4,5] và [1,3,4]
  Merge: [1,4,5] + [1,3,4] = [1,1,3,4,4,5]

Level 1 (Right):
  [[2,6]]
  Chỉ còn 1 list → return [2,6]

Level 0: Merge kết quả
  [1,1,3,4,4,5] + [2,6] = [1,1,2,3,4,4,5,6]

Kết quả: [1,1,2,3,4,4,5,6]
```

## Giải thích code

### Cấu trúc hàm chính

```1:20:internal/linked_list/merge_k_sorted_lists/merge_k_sorted_lists.go
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
```

### Hàm mergeLists (Divide and Conquer)

```22:36:internal/linked_list/merge_k_sorted_lists/merge_k_sorted_lists.go
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
```

**Giải thích:**
- **Base case** (dòng 25-27): Nếu chỉ còn một list, trả về list đó
- **Chia đôi** (dòng 29-30): Tính điểm giữa và chia mảng thành hai nửa
- **Đệ quy** (dòng 32-33): Đệ quy merge từng nửa
- **Kết hợp** (dòng 35): Merge hai kết quả bằng `mergeTwoLists`

### Hàm mergeTwoLists

```38:71:internal/linked_list/merge_k_sorted_lists/merge_k_sorted_lists.go
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
```

**Giải thích:**
- **Dummy node** (dòng 50-51): Tạo node giả để đơn giản hóa việc xử lý edge cases
- **Vòng lặp merge** (dòng 53-65): So sánh và chọn node nhỏ hơn từ hai lists
- **Nối phần còn lại** (dòng 67-71): Nối phần còn lại của list chưa kết thúc

### Độ phức tạp

- **Thời gian**: O(n log k)
  - Có log k levels trong cây đệ quy
  - Mỗi level merge tất cả n nodes
  - Tổng: O(n log k)

- **Không gian**: O(log k)
  - Call stack cho đệ quy: log k levels
  - Không sử dụng thêm cấu trúc dữ liệu nào

### Tại sao thuật toán này đúng?

1. **Base case đúng**: Một list đã được sắp xếp, không cần merge.

2. **Merge 2 lists đúng**: Hàm `mergeTwoLists` đảm bảo merge đúng thứ tự.

3. **Divide and Conquer đúng**: 
   - Chia nhỏ vấn đề thành các vấn đề con
   - Giải quyết từng vấn đề con
   - Kết hợp kết quả đúng

4. **Tính đầy đủ**: Tất cả nodes đều được xử lý qua các lần merge.

## So sánh các cách tiếp cận

| Cách tiếp cận | Thời gian | Không gian | Ghi chú |
|--------------|-----------|------------|---------|
| Brute Force | O(n log n) | O(n) | Đơn giản nhưng không tận dụng tính đã sắp xếp |
| Compare One by One | O(k * n) | O(1) | Chậm khi k lớn |
| Priority Queue | O(n log k) | O(k) | Cần implement heap |
| Divide and Conquer | O(n log k) | O(log k) | ⭐ Tối ưu, dễ hiểu |

## Test Cases

### Test Case 1: Example 1
```go
Input: lists = [[1,4,5],[1,3,4],[2,6]]
Output: [1,1,2,3,4,4,5,6]
```
Ba lists với độ dài khác nhau.

### Test Case 2: Example 2
```go
Input: lists = []
Output: []
```
Mảng rỗng.

### Test Case 3: Example 3
```go
Input: lists = [[]]
Output: []
```
Một list rỗng.

### Test Case 4: Single list
```go
Input: lists = [[1,2,3]]
Output: [1,2,3]
```
Chỉ có một list.

### Test Case 5: Two lists
```go
Input: lists = [[1,3,5],[2,4,6]]
Output: [1,2,3,4,5,6]
```
Hai lists đơn giản.

### Test Case 6: Lists with empty
```go
Input: lists = [[1,2],[],[3,4]]
Output: [1,2,3,4]
```
Có list rỗng trong mảng.

### Test Case 7: Different lengths
```go
Input: lists = [[1,5,9],[2,3],[4,6,7,8]]
Output: [1,2,3,4,5,6,7,8,9]
```
Lists có độ dài khác nhau.

## Chạy test

Để chạy các test case:

```bash
go test ./internal/linked_list/merge_k_sorted_lists/
```

Hoặc chạy với verbose mode để xem chi tiết:

```bash
go test -v ./internal/linked_list/merge_k_sorted_lists/
```

## Mở rộng

### Merge Two Sorted Lists (LeetCode 21)

Bài toán cơ bản hơn: merge 2 sorted lists. Đây là building block cho merge k lists.

### Merge Sorted Array (LeetCode 88)

Merge hai mảng đã sắp xếp. Tương tự nhưng với arrays thay vì linked lists.

### Ứng dụng thực tế

Bài toán này có ứng dụng trong:
- **External Sorting**: Merge nhiều file đã sắp xếp
- **Database**: Merge results từ nhiều sorted indexes
- **Distributed Systems**: Merge sorted data từ nhiều nodes
- **Stream Processing**: Merge sorted streams
- **Priority Scheduling**: Merge nhiều priority queues
