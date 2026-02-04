# Find Minimum in Rotated Sorted Array

## Mô tả bài toán

Giả sử một mảng có độ dài `n` được sắp xếp theo thứ tự tăng dần được xoay từ 1 đến n lần. Ví dụ, mảng `nums = [0,1,2,4,5,6,7]` có thể trở thành:

- `[4,5,6,7,0,1,2]` nếu được xoay 4 lần
- `[0,1,2,4,5,6,7]` nếu được xoay 7 lần

Lưu ý rằng xoay một mảng `[a[0], a[1], a[2], ..., a[n-1]]` 1 lần sẽ cho kết quả là mảng `[a[n-1], a[0], a[1], a[2], ..., a[n-2]]`.

Cho một mảng `nums` đã được sắp xếp và xoay với các phần tử duy nhất, trả về phần tử nhỏ nhất của mảng này.

Bạn phải viết một thuật toán chạy trong thời gian O(log n).

**Ví dụ:**

### Example 1:
- Input: `nums = [3,4,5,1,2]`
- Output: `1`
- Giải thích: Mảng gốc là `[1,2,3,4,5]` được xoay 3 lần.

### Example 2:
- Input: `nums = [4,5,6,7,0,1,2]`
- Output: `0`
- Giải thích: Mảng gốc là `[0,1,2,4,5,6,7]` và được xoay 4 lần.

### Example 3:
- Input: `nums = [11,13,15,17]`
- Output: `11`
- Giải thích: Mảng gốc là `[11,13,15,17]` và được xoay 4 lần (hoặc không xoay).

**Ràng buộc:**
- `n == nums.length`
- `1 <= n <= 5000`
- `-5000 <= nums[i] <= 5000`
- Tất cả các số nguyên trong `nums` là duy nhất
- `nums` được sắp xếp và xoay từ 1 đến n lần

## Phân tích thuật toán

### Đặc điểm của Rotated Sorted Array

Một rotated sorted array có đặc điểm:
- Có một điểm "pivot" (điểm xoay) chia mảng thành hai phần đã được sắp xếp
- Phần tử nhỏ nhất nằm tại điểm pivot
- Tất cả phần tử bên trái pivot đều lớn hơn tất cả phần tử bên phải pivot

**Ví dụ:** `[4,5,6,7,0,1,2]`
- Pivot tại index 4 (giá trị 0)
- Bên trái: `[4,5,6,7]` (đã sắp xếp, tất cả > 0)
- Bên phải: `[0,1,2]` (đã sắp xếp, tất cả < 4)

### Cách tiếp cận: Binary Search

Thay vì duyệt tuyến tính O(n), ta sử dụng Binary Search O(log n).

**Ý tưởng:**
- So sánh `nums[mid]` với `nums[right]` để xác định phần nào chứa minimum
- Nếu `nums[mid] < nums[right]`: Phần bên phải đã được sắp xếp → minimum nằm ở bên trái (bao gồm mid)
- Nếu `nums[mid] > nums[right]`: Phần bên trái đã được sắp xếp → minimum nằm ở bên phải (sau mid)

**Tại sao so sánh với `nums[right]` thay vì `nums[left]`?**

- So sánh với `nums[right]` đơn giản hơn và tránh edge cases
- Khi `nums[mid] < nums[right]`, ta biết chắc chắn phần bên phải đã được sắp xếp
- Khi `nums[mid] > nums[right]`, ta biết chắc chắn phần bên trái đã được sắp xếp

### Các bước thực hiện

1. **Khởi tạo**: `left = 0`, `right = len(nums) - 1`
2. **Binary Search**: Trong khi `left < right`:
   - Tính `mid = left + (right - left) / 2`
   - So sánh `nums[mid]` với `nums[right]`:
     - Nếu `nums[mid] < nums[right]`: `right = mid` (minimum ở bên trái)
     - Nếu `nums[mid] > nums[right]`: `left = mid + 1` (minimum ở bên phải)
3. **Kết quả**: Khi `left == right`, `nums[left]` là minimum

### Ví dụ minh họa

Với input: `nums = [4,5,6,7,0,1,2]`

```
Bước 0: left = 0, right = 6
  nums[left] = 4, nums[right] = 2

Bước 1: mid = 3, nums[mid] = 7, nums[right] = 2
  nums[mid] (7) > nums[right] (2)
  → Phần bên trái đã được sắp xếp, minimum ở bên phải
  → left = mid + 1 = 4

Bước 2: left = 4, right = 6
  mid = 5, nums[mid] = 1, nums[right] = 2
  nums[mid] (1) < nums[right] (2)
  → Phần bên phải đã được sắp xếp, minimum ở bên trái (bao gồm mid)
  → right = mid = 5

Bước 3: left = 4, right = 5
  mid = 4, nums[mid] = 0, nums[right] = 1
  nums[mid] (0) < nums[right] (1)
  → right = mid = 4

Bước 4: left = 4, right = 4
  left == right → dừng
  Kết quả: nums[4] = 0
```

Với input: `nums = [3,4,5,1,2]`

```
Bước 0: left = 0, right = 4
  nums[left] = 3, nums[right] = 2

Bước 1: mid = 2, nums[mid] = 5, nums[right] = 2
  nums[mid] (5) > nums[right] (2)
  → left = mid + 1 = 3

Bước 2: left = 3, right = 4
  mid = 3, nums[mid] = 1, nums[right] = 2
  nums[mid] (1) < nums[right] (2)
  → right = mid = 3

Bước 3: left = 3, right = 3
  left == right → dừng
  Kết quả: nums[3] = 1
```

Với input: `nums = [11,13,15,17]` (không xoay)

```
Bước 0: left = 0, right = 3
  nums[left] = 11, nums[right] = 17

Bước 1: mid = 1, nums[mid] = 13, nums[right] = 17
  nums[mid] (13) < nums[right] (17)
  → right = mid = 1

Bước 2: left = 0, right = 1
  mid = 0, nums[mid] = 11, nums[right] = 13
  nums[mid] (11) < nums[right] (13)
  → right = mid = 0

Bước 3: left = 0, right = 0
  left == right → dừng
  Kết quả: nums[0] = 11
```

## Giải thích code

### Cấu trúc hàm

```1:32:internal/array/minimum_rotated_sorted_array/minimum_rotated_sorted_array.go
package array

// findMin tìm phần tử nhỏ nhất trong một mảng đã được sắp xếp và xoay.
//
// Sử dụng Binary Search:
// - Mảng rotated có đặc điểm: một phần bên trái và một phần bên phải đều được sắp xếp
// - Phần tử nhỏ nhất nằm ở điểm "pivot" (điểm xoay)
// - So sánh nums[mid] với nums[right] để xác định phần nào chứa minimum
//   - Nếu nums[mid] < nums[right]: minimum nằm ở bên trái (bao gồm mid)
//   - Nếu nums[mid] > nums[right]: minimum nằm ở bên phải (sau mid)
//
// Độ phức tạp: O(log n) thời gian, O(1) không gian
func findMin(nums []int) int {
	left := 0
	right := len(nums) - 1

	// Binary search
	for left < right {
		mid := left + (right-left)/2

		// So sánh nums[mid] với nums[right]
		// Nếu nums[mid] < nums[right], phần bên phải đã được sắp xếp
		// → minimum nằm ở bên trái (bao gồm mid)
		if nums[mid] < nums[right] {
			right = mid
		} else {
			// Nếu nums[mid] > nums[right], phần bên trái đã được sắp xếp
			// → minimum nằm ở bên phải (sau mid)
			left = mid + 1
		}
	}

	// Khi left == right, ta đã tìm thấy minimum
	return nums[left]
}
```

### Chi tiết từng phần

#### 1. Khởi tạo (dòng 14-15)
```go
left := 0
right := len(nums) - 1
```
- `left`: Chỉ số bắt đầu của vùng tìm kiếm
- `right`: Chỉ số kết thúc của vùng tìm kiếm

#### 2. Binary Search Loop (dòng 17-28)
```go
for left < right {
    mid := left + (right-left)/2
    
    if nums[mid] < nums[right] {
        right = mid
    } else {
        left = mid + 1
    }
}
```

**Giải thích:**
- `mid := left + (right-left)/2`: Tính điểm giữa (tránh overflow)
- `nums[mid] < nums[right]`:
  - Phần bên phải từ mid đến right đã được sắp xếp
  - Minimum không thể nằm sau mid (vì nums[mid] < nums[right])
  - Minimum có thể là nums[mid] hoặc nằm ở bên trái
  - → `right = mid` (bao gồm mid trong vùng tìm kiếm)
- `nums[mid] > nums[right]`:
  - Phần bên trái từ left đến mid đã được sắp xếp
  - Minimum không thể nằm ở left đến mid (vì nums[mid] > nums[right])
  - Minimum phải nằm sau mid
  - → `left = mid + 1` (loại bỏ mid và phần bên trái)

#### 3. Trả về kết quả (dòng 31)
```go
return nums[left]
```
- Khi `left == right`, ta đã tìm thấy minimum
- `nums[left]` (hoặc `nums[right]`) là phần tử nhỏ nhất

### Độ phức tạp

- **Thời gian**: O(log n)
  - Mỗi lần lặp giảm một nửa vùng tìm kiếm
  - Tối đa log n lần lặp

- **Không gian**: O(1)
  - Chỉ sử dụng các biến: `left`, `right`, `mid`
  - Không sử dụng thêm cấu trúc dữ liệu nào

### Tại sao thuật toán này đúng?

1. **Invariant**: Minimum luôn nằm trong khoảng [left, right]

2. **Giảm vùng tìm kiếm**:
   - Khi `nums[mid] < nums[right]`: Minimum không thể nằm sau mid
   - Khi `nums[mid] > nums[right]`: Minimum không thể nằm ở left đến mid

3. **Termination**: Khi `left == right`, chỉ còn một phần tử → đó là minimum

4. **Edge cases**:
   - Mảng không xoay: Vẫn hoạt động đúng (minimum ở đầu)
   - Mảng xoay hoàn toàn: Vẫn hoạt động đúng (minimum ở cuối)

## Test Cases

### Test Case 1: Example 1
```go
Input: nums = [3,4,5,1,2]
Output: 1
```
Minimum ở giữa mảng.

### Test Case 2: Example 2
```go
Input: nums = [4,5,6,7,0,1,2]
Output: 0
```
Minimum ở giữa mảng.

### Test Case 3: Example 3
```go
Input: nums = [11,13,15,17]
Output: 11
```
Mảng không xoay, minimum ở đầu.

### Test Case 4: Single element
```go
Input: nums = [1]
Output: 1
```
Chỉ có một phần tử.

### Test Case 5: Two elements rotated
```go
Input: nums = [2,1]
Output: 1
```
Hai phần tử đã xoay.

### Test Case 6: Minimum at start
```go
Input: nums = [1,2,3,4,5]
Output: 1
```
Mảng không xoay.

### Test Case 7: Minimum at end
```go
Input: nums = [2,3,4,5,1]
Output: 1
```
Xoay một lần.

### Test Case 8: Negative numbers
```go
Input: nums = [-1,0,1,2,-2]
Output: -2
```
Có số âm.

## Chạy test

Để chạy các test case:

```bash
go test ./internal/array/minimum_rotated_sorted_array/
```

Hoặc chạy với verbose mode để xem chi tiết:

```bash
go test -v ./internal/array/minimum_rotated_sorted_array/
```

## Mở rộng

### Search in Rotated Sorted Array (LeetCode 33)

Tìm một phần tử cụ thể trong rotated sorted array. Tương tự nhưng phức tạp hơn vì cần xử lý nhiều trường hợp.

### Find Minimum in Rotated Sorted Array II (LeetCode 154)

Tương tự nhưng mảng có thể chứa duplicates. Yêu cầu xử lý edge case khi `nums[mid] == nums[right]`.

### Ứng dụng thực tế

Bài toán này có ứng dụng trong:
- **Database Indexing**: Tìm giá trị nhỏ nhất trong rotated index
- **Circular Buffer**: Tìm điểm bắt đầu trong buffer xoay vòng
- **Time Series Analysis**: Tìm điểm thấp nhất trong dữ liệu xoay vòng
- **Game Development**: Tìm góc quay nhỏ nhất
- **Signal Processing**: Tìm phase offset nhỏ nhất
