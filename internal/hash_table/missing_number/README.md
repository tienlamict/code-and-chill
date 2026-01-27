# Missing Number

## Mô tả bài toán

Cho một mảng `nums` chứa `n` số phân biệt trong khoảng `[0, n]`, trả về số duy nhất trong khoảng đó còn thiếu trong mảng.

**Ví dụ:**

### Example 1:
- Input: `nums = [3,0,1]`
- Output: `2`
- Giải thích: `n = 3` vì có 3 số, nên tất cả các số nằm trong khoảng `[0,3]`. Số `2` là số còn thiếu trong khoảng vì nó không xuất hiện trong `nums`.

### Example 2:
- Input: `nums = [0,1]`
- Output: `2`
- Giải thích: `n = 2` vì có 2 số, nên tất cả các số nằm trong khoảng `[0,2]`. Số `2` là số còn thiếu trong khoảng vì nó không xuất hiện trong `nums`.

### Example 3:
- Input: `nums = [9,6,4,2,3,5,7,0,1]`
- Output: `8`
- Giải thích: `n = 9` vì có 9 số, nên tất cả các số nằm trong khoảng `[0,9]`. Số `8` là số còn thiếu trong khoảng vì nó không xuất hiện trong `nums`.

**Ràng buộc:**
- `n == nums.length`
- `1 <= n <= 10^4`
- `0 <= nums[i] <= n`
- Tất cả các số trong `nums` là duy nhất

**Follow up:** Bạn có thể implement một giải pháp chỉ sử dụng O(1) độ phức tạp không gian bổ sung và O(n) độ phức tạp thời gian không?

## Phân tích thuật toán

### Cách tiếp cận 1: Hash Set (O(n) time, O(n) space)

Tạo một hash set chứa tất cả các số trong mảng, sau đó duyệt từ 0 đến n để tìm số không có trong set.

**Độ phức tạp:**
- Thời gian: O(n)
- Không gian: O(n)

### Cách tiếp cận 2: Sorting (O(n log n) time, O(1) space)

Sắp xếp mảng, sau đó duyệt để tìm số còn thiếu.

**Độ phức tạp:**
- Thời gian: O(n log n)
- Không gian: O(1) (nếu không tính không gian sắp xếp)

### Cách tiếp cận 3: Sum Formula (O(n) time, O(1) space) ⭐ Tối ưu

Sử dụng công thức tổng của dãy số từ 0 đến n:
- Tổng lý thuyết: `n * (n + 1) / 2`
- Tính tổng thực tế của các số trong mảng
- Số còn thiếu = Tổng lý thuyết - Tổng thực tế

**Tại sao cách này hoạt động?**

Nếu mảng chứa đầy đủ các số từ 0 đến n, tổng của chúng sẽ bằng `n * (n + 1) / 2`. Nếu thiếu một số, tổng thực tế sẽ nhỏ hơn tổng lý thuyết đúng bằng số còn thiếu.

**Độ phức tạp:**
- Thời gian: O(n) - duyệt qua mảng một lần
- Không gian: O(1) - chỉ sử dụng các biến

### Cách tiếp cận 4: XOR (O(n) time, O(1) space)

Sử dụng tính chất của phép XOR:
- `a XOR a = 0`
- `a XOR 0 = a`
- XOR có tính giao hoán và kết hợp

Ý tưởng: XOR tất cả các số từ 0 đến n với tất cả các số trong mảng. Số còn thiếu sẽ là kết quả cuối cùng.

**Độ phức tạp:**
- Thời gian: O(n)
- Không gian: O(1)

## Giải thích code

### Cấu trúc hàm

```1:22:internal/hash_table/missing_number/missing_number.go
package hash_table

// missingNumber tìm số còn thiếu trong mảng chứa n số phân biệt trong khoảng [0, n].
//
// Sử dụng công thức tổng:
// - Tổng của các số từ 0 đến n: n * (n + 1) / 2
// - Tính tổng các số trong mảng
// - Số còn thiếu = Tổng lý thuyết - Tổng thực tế
//
// Độ phức tạp: O(n) thời gian, O(1) không gian
func missingNumber(nums []int) int {
	n := len(nums)
	
	// Tổng lý thuyết của các số từ 0 đến n
	expectedSum := n * (n + 1) / 2
	
	// Tính tổng thực tế của các số trong mảng
	actualSum := 0
	for _, num := range nums {
		actualSum += num
	}
	
	// Số còn thiếu = Tổng lý thuyết - Tổng thực tế
	return expectedSum - actualSum
}
```

### Chi tiết từng phần

#### 1. Xác định độ dài mảng (dòng 11)
```go
n := len(nums)
```
- `n` là số lượng phần tử trong mảng
- Cũng là giá trị lớn nhất trong khoảng `[0, n]`

#### 2. Tính tổng lý thuyết (dòng 13)
```go
expectedSum := n * (n + 1) / 2
```
- Công thức tổng của dãy số từ 0 đến n: `0 + 1 + 2 + ... + n = n * (n + 1) / 2`
- Đây là tổng nếu mảng chứa đầy đủ các số từ 0 đến n

**Ví dụ:** Với n = 3:
```
Tổng = 0 + 1 + 2 + 3 = 6
Hoặc: 3 * (3 + 1) / 2 = 3 * 4 / 2 = 6
```

#### 3. Tính tổng thực tế (dòng 15-18)
```go
actualSum := 0
for _, num := range nums {
    actualSum += num
}
```
- Duyệt qua mảng và tính tổng tất cả các số có trong mảng
- Đây là tổng thực tế của các số hiện có

**Ví dụ:** Với `nums = [3, 0, 1]`:
```
actualSum = 3 + 0 + 1 = 4
```

#### 4. Tìm số còn thiếu (dòng 20)
```go
return expectedSum - actualSum
```
- Số còn thiếu chính là hiệu giữa tổng lý thuyết và tổng thực tế

**Ví dụ:** Với `nums = [3, 0, 1]`:
```
expectedSum = 6
actualSum = 4
missing = 6 - 4 = 2 ✓
```

### Ví dụ minh họa

Với input: `nums = [3, 0, 1]`

```
Bước 1: Xác định n
  n = len(nums) = 3

Bước 2: Tính tổng lý thuyết
  expectedSum = 3 * (3 + 1) / 2 = 3 * 4 / 2 = 6

Bước 3: Tính tổng thực tế
  actualSum = 3 + 0 + 1 = 4

Bước 4: Tìm số còn thiếu
  missing = 6 - 4 = 2

Kết quả: 2
```

Với input: `nums = [0, 1]`

```
Bước 1: Xác định n
  n = len(nums) = 2

Bước 2: Tính tổng lý thuyết
  expectedSum = 2 * (2 + 1) / 2 = 2 * 3 / 2 = 3

Bước 3: Tính tổng thực tế
  actualSum = 0 + 1 = 1

Bước 4: Tìm số còn thiếu
  missing = 3 - 1 = 2

Kết quả: 2
```

Với input: `nums = [9, 6, 4, 2, 3, 5, 7, 0, 1]`

```
Bước 1: Xác định n
  n = len(nums) = 9

Bước 2: Tính tổng lý thuyết
  expectedSum = 9 * (9 + 1) / 2 = 9 * 10 / 2 = 45

Bước 3: Tính tổng thực tế
  actualSum = 9 + 6 + 4 + 2 + 3 + 5 + 7 + 0 + 1 = 37

Bước 4: Tìm số còn thiếu
  missing = 45 - 37 = 8

Kết quả: 8
```

### Độ phức tạp

- **Thời gian**: O(n)
  - Duyệt qua mảng một lần để tính tổng: O(n)
  - Các phép tính khác: O(1)
  - Tổng: O(n)

- **Không gian**: O(1)
  - Chỉ sử dụng các biến: `n`, `expectedSum`, `actualSum`
  - Không sử dụng thêm cấu trúc dữ liệu nào

### Tại sao công thức tổng hoạt động?

**Chứng minh toán học:**

Tổng của dãy số từ 0 đến n:
```
S = 0 + 1 + 2 + ... + n
```

Có thể viết lại:
```
S = n + (n-1) + (n-2) + ... + 1 + 0
```

Cộng hai phương trình:
```
2S = (0+n) + (1+(n-1)) + (2+(n-2)) + ... + (n+0)
2S = n + n + n + ... + n  (n+1 lần)
2S = n * (n + 1)
S = n * (n + 1) / 2
```

Nếu mảng thiếu một số `x`, tổng thực tế sẽ là:
```
actualSum = expectedSum - x
```

Do đó:
```
x = expectedSum - actualSum
```

## Các cách tiếp cận khác

### Cách tiếp cận: XOR

```go
func missingNumberXOR(nums []int) int {
    result := len(nums) // Bắt đầu với n
    
    for i, num := range nums {
        result ^= i ^ num
    }
    
    return result
}
```

**Giải thích:**
- XOR tất cả các chỉ số từ 0 đến n-1 với tất cả các giá trị trong mảng
- Số còn thiếu sẽ là kết quả cuối cùng

**Ví dụ:** Với `nums = [3, 0, 1]`:
```
result = 3
i=0: result = 3 ^ 0 ^ 3 = 0
i=1: result = 0 ^ 1 ^ 0 = 1
i=2: result = 1 ^ 2 ^ 1 = 2
Kết quả: 2
```

Cách này cũng có độ phức tạp O(n) time và O(1) space, nhưng công thức tổng dễ hiểu hơn.

## Test Cases

### Test Case 1: Example 1
```go
Input: nums = [3,0,1]
Output: 2
```
Thiếu số 2 trong khoảng [0,3].

### Test Case 2: Example 2
```go
Input: nums = [0,1]
Output: 2
```
Thiếu số 2 trong khoảng [0,2].

### Test Case 3: Example 3
```go
Input: nums = [9,6,4,2,3,5,7,0,1]
Output: 8
```
Thiếu số 8 trong khoảng [0,9].

### Test Case 4: Missing 0
```go
Input: nums = [1,2]
Output: 0
```
Thiếu số 0 trong khoảng [0,2].

### Test Case 5: Missing last number
```go
Input: nums = [0,1,2,3]
Output: 4
```
Thiếu số cuối cùng trong khoảng [0,4].

### Test Case 6: Single element missing 0
```go
Input: nums = [1]
Output: 0
```
Chỉ có 1 phần tử, thiếu số 0.

### Test Case 7: Single element missing 1
```go
Input: nums = [0]
Output: 1
```
Chỉ có 1 phần tử, thiếu số 1.

## Chạy test

Để chạy các test case:

```bash
go test ./internal/hash_table/missing_number/
```

Hoặc chạy với verbose mode để xem chi tiết:

```bash
go test -v ./internal/hash_table/missing_number/
```

## Mở rộng

### Missing Number II

Một biến thể của bài toán này là tìm **hai số còn thiếu** trong mảng chứa n-2 số phân biệt trong khoảng [0, n]. Giải pháp sẽ phức tạp hơn, có thể sử dụng:
- Tổng và tích của các số
- XOR với phân chia bit
- Hash set

### Ứng dụng thực tế

Bài toán này có ứng dụng trong:
- Kiểm tra tính toàn vẹn dữ liệu
- Tìm lỗi trong chuỗi số liệu
- Phát hiện phần tử thiếu trong database
- Kiểm tra tính liên tục của dãy số
