# Combination Sum

## Mô tả bài toán

Cho một mảng các số nguyên phân biệt `candidates` và một số nguyên `target`, trả về danh sách tất cả các tổ hợp duy nhất của `candidates` sao cho tổng các số được chọn bằng `target`. Bạn có thể trả về các tổ hợp theo bất kỳ thứ tự nào.

Cùng một số có thể được chọn từ `candidates` không giới hạn số lần. Hai tổ hợp là duy nhất nếu tần suất của ít nhất một trong các số được chọn là khác nhau.

Các test case được tạo sao cho số lượng tổ hợp duy nhất có tổng bằng `target` nhỏ hơn 150 tổ hợp cho input đã cho.

**Ví dụ:**

### Example 1:
- Input: `candidates = [2,3,6,7]`, `target = 7`
- Output: `[[2,2,3],[7]]`
- Giải thích:
  - `2` và `3` là candidates, và `2 + 2 + 3 = 7`. Lưu ý rằng `2` có thể được sử dụng nhiều lần.
  - `7` là candidate, và `7 = 7`.
  - Đây là hai tổ hợp duy nhất.

### Example 2:
- Input: `candidates = [2,3,5]`, `target = 8`
- Output: `[[2,2,2,2],[2,3,3],[3,5]]`

### Example 3:
- Input: `candidates = [2]`, `target = 1`
- Output: `[]`

**Ràng buộc:**
- `1 <= candidates.length <= 30`
- `2 <= candidates[i] <= 40`
- Tất cả các phần tử của `candidates` là phân biệt
- `1 <= target <= 40`

## Phân tích thuật toán

### Cách tiếp cận: Backtracking

Đây là bài toán kinh điển sử dụng kỹ thuật **Backtracking** (quay lui).

**Ý tưởng:**
- Thử tất cả các tổ hợp có thể bằng cách đệ quy
- Mỗi số có thể được sử dụng nhiều lần
- Tránh duplicate bằng cách chỉ xét các số từ vị trí hiện tại trở đi
- Pruning: dừng sớm nếu tổng hiện tại vượt quá target

**Các bước thực hiện:**

1. **Sắp xếp mảng**: Để dễ dàng pruning và tránh duplicate
2. **Backtracking**: 
   - Với mỗi số từ vị trí `start` trở đi, thử thêm vào tổ hợp
   - Đệ quy với cùng số (có thể dùng lại) hoặc số tiếp theo
   - Nếu tổng bằng target, lưu tổ hợp vào kết quả
   - Nếu tổng vượt quá target, dừng nhánh đó (pruning)
3. **Backtrack**: Xóa số vừa thêm để thử số khác

### Ví dụ minh họa

Với input: `candidates = [2,3,6,7]`, `target = 7`

```
Sắp xếp: [2,3,6,7] (đã sắp xếp)

Backtracking tree:

                    []
            /        |        |        \
          2         3        6         7
         /|\       /|\       |         |
       2  3 6 7   3 6 7     6 7       7
      /|  | |     | |       |         |
    2 3  3 6 7   6 7       7         ✓ (7=7)
   /| |  |       |
  2 3 6 7       7
  |
  ✓ (2+2+3=7)

Kết quả: [[2,2,3], [7]]
```

Chi tiết từng bước:

```
Bước 1: Thử 2
  current = [2], remaining = 5
  Bước 1.1: Thử 2 (có thể dùng lại)
    current = [2,2], remaining = 3
    Bước 1.1.1: Thử 2
      current = [2,2,2], remaining = 1
      Bước 1.1.1.1: Thử 2 → 2 > 1, break (pruning)
      Bước 1.1.1.2: Thử 3 → 3 > 1, break
      Bước 1.1.1.3: Thử 6 → 6 > 1, break
      Bước 1.1.1.4: Thử 7 → 7 > 1, break
    Bước 1.1.2: Thử 3
      current = [2,2,3], remaining = 0 ✓
      → Thêm [2,2,3] vào kết quả
    Bước 1.1.3: Thử 6 → 6 > 3, break
    Bước 1.1.4: Thử 7 → 7 > 3, break
  Bước 1.2: Thử 3
    current = [2,3], remaining = 2
    Bước 1.2.1: Thử 3 → 3 > 2, break
    Bước 1.2.2: Thử 6 → 6 > 2, break
    Bước 1.2.3: Thử 7 → 7 > 2, break
  Bước 1.3: Thử 6 → 6 > 5, break
  Bước 1.4: Thử 7 → 7 > 5, break

Bước 2: Thử 3
  current = [3], remaining = 4
  Bước 2.1: Thử 3
    current = [3,3], remaining = 1
    Bước 2.1.1: Thử 3 → 3 > 1, break
    Bước 2.1.2: Thử 6 → 6 > 1, break
    Bước 2.1.3: Thử 7 → 7 > 1, break
  Bước 2.2: Thử 6 → 6 > 4, break
  Bước 2.3: Thử 7 → 7 > 4, break

Bước 3: Thử 6
  current = [6], remaining = 1
  Bước 3.1: Thử 6 → 6 > 1, break
  Bước 3.2: Thử 7 → 7 > 1, break

Bước 4: Thử 7
  current = [7], remaining = 0 ✓
  → Thêm [7] vào kết quả

Kết quả: [[2,2,3], [7]]
```

## Giải thích code

### Cấu trúc hàm

```1:52:internal/array/combination_sum/combination_sum.go
package array

import "sort"

// combinationSum tìm tất cả các tổ hợp duy nhất của candidates sao cho tổng bằng target.
//
// Sử dụng Backtracking:
// 1. Sắp xếp mảng để dễ dàng pruning và tránh duplicate
// 2. Với mỗi số, thử thêm vào tổ hợp hiện tại
// 3. Mỗi số có thể được sử dụng nhiều lần
// 4. Tránh duplicate bằng cách chỉ xét các số từ vị trí hiện tại trở đi
// 5. Pruning: nếu tổng hiện tại > target, dừng nhánh đó
//
// Độ phức tạp: O(2^target) thời gian trong trường hợp xấu nhất, O(target) không gian (cho call stack)
func combinationSum(candidates []int, target int) [][]int {
	// Sắp xếp mảng để dễ dàng pruning
	sort.Ints(candidates)

	result := [][]int{}
	current := []int{}

	// Backtracking function
	var backtrack func(int, int)
	backtrack = func(start int, remaining int) {
		// Base case: đã đạt target
		if remaining == 0 {
			// Tạo bản sao của current và thêm vào result
			combination := make([]int, len(current))
			copy(combination, current)
			result = append(result, combination)
			return
		}

		// Duyệt qua các số từ start đến cuối
		for i := start; i < len(candidates); i++ {
			// Pruning: nếu số hiện tại lớn hơn remaining, không cần xét tiếp
			// (vì mảng đã được sắp xếp, các số sau sẽ còn lớn hơn)
			if candidates[i] > remaining {
				break
			}

			// Thêm số hiện tại vào tổ hợp
			current = append(current, candidates[i])

			// Đệ quy: tiếp tục với số hiện tại (có thể dùng lại)
			// remaining giảm đi candidates[i]
			backtrack(i, remaining-candidates[i])

			// Backtrack: xóa số vừa thêm để thử số khác
			current = current[:len(current)-1]
		}
	}

	// Bắt đầu backtracking từ index 0
	backtrack(0, target)

	return result
}
```

### Chi tiết từng phần

#### 1. Sắp xếp mảng (dòng 15)
```go
sort.Ints(candidates)
```
- Sắp xếp mảng để:
  - Dễ dàng pruning: nếu `candidates[i] > remaining`, các số sau cũng sẽ lớn hơn
  - Tránh duplicate: chỉ xét các số từ `start` trở đi

#### 2. Khởi tạo (dòng 17-18)
```go
result := [][]int{}
current := []int{}
```
- `result`: Lưu tất cả các tổ hợp hợp lệ
- `current`: Tổ hợp hiện tại đang xây dựng

#### 3. Hàm backtrack (dòng 21-48)
```go
var backtrack func(int, int)
backtrack = func(start int, remaining int) {
    // Base case
    if remaining == 0 {
        combination := make([]int, len(current))
        copy(combination, current)
        result = append(result, combination)
        return
    }

    // Duyệt từ start đến cuối
    for i := start; i < len(candidates); i++ {
        if candidates[i] > remaining {
            break
        }
        current = append(current, candidates[i])
        backtrack(i, remaining-candidates[i])
        current = current[:len(current)-1]
    }
}
```

**Giải thích:**
- **Base case** (dòng 23-28): Nếu `remaining == 0`, đã tìm thấy một tổ hợp hợp lệ
  - Tạo bản sao của `current` (vì slice là reference type)
  - Thêm vào `result`
- **Vòng lặp** (dòng 30-45): Duyệt từ `start` đến cuối mảng
  - **Pruning** (dòng 32-34): Nếu `candidates[i] > remaining`, dừng vòng lặp
  - **Thêm số** (dòng 37): Thêm `candidates[i]` vào `current`
  - **Đệ quy** (dòng 40): Gọi `backtrack(i, remaining-candidates[i])`
    - `i` (không phải `i+1`): Cho phép dùng lại số hiện tại
    - `remaining-candidates[i]`: Giảm remaining đi giá trị số vừa thêm
  - **Backtrack** (dòng 43): Xóa số vừa thêm để thử số khác

#### 4. Bắt đầu backtracking (dòng 50)
```go
backtrack(0, target)
```
- Bắt đầu từ index 0 với remaining = target

### Độ phức tạp

- **Thời gian**: O(2^target) trong trường hợp xấu nhất
  - Mỗi số có thể được chọn hoặc không
  - Trong trường hợp xấu nhất, có thể có 2^target tổ hợp
  - Tuy nhiên, với pruning, thực tế nhanh hơn nhiều

- **Không gian**: O(target)
  - Call stack: tối đa target levels (mỗi level thêm một số)
  - `current`: tối đa target phần tử
  - `result`: phụ thuộc vào số lượng tổ hợp tìm được

### Tại sao thuật toán này đúng?

1. **Tính đầy đủ**: Backtracking thử tất cả các tổ hợp có thể.

2. **Tránh duplicate**: 
   - Chỉ xét các số từ `start` trở đi
   - Không quay lại các số đã xét trước đó
   - Đảm bảo các tổ hợp được tạo theo thứ tự không giảm

3. **Pruning hiệu quả**: 
   - Dừng sớm khi `candidates[i] > remaining`
   - Giảm đáng kể số lượng nhánh cần xét

4. **Xử lý số dùng lại**: 
   - Truyền `i` (không phải `i+1`) cho phép dùng lại số hiện tại
   - Đúng với yêu cầu "cùng một số có thể được chọn nhiều lần"

## Test Cases

### Test Case 1: Example 1
```go
Input: candidates = [2,3,6,7], target = 7
Output: [[2,2,3],[7]]
```
Hai tổ hợp duy nhất.

### Test Case 2: Example 2
```go
Input: candidates = [2,3,5], target = 8
Output: [[2,2,2,2],[2,3,3],[3,5]]
```
Ba tổ hợp với các độ dài khác nhau.

### Test Case 3: Example 3
```go
Input: candidates = [2], target = 1
Output: []
```
Không có tổ hợp nào.

### Test Case 4: Single candidate matches
```go
Input: candidates = [2], target = 2
Output: [[2]]
```
Chỉ cần một số.

### Test Case 5: Single candidate multiple times
```go
Input: candidates = [2], target = 8
Output: [[2,2,2,2]]
```
Cần dùng số nhiều lần.

### Test Case 6: No solution
```go
Input: candidates = [3,5,7], target = 2
Output: []
```
Tất cả số đều lớn hơn target.

## Chạy test

Để chạy các test case:

```bash
go test ./internal/array/combination_sum/
```

Hoặc chạy với verbose mode để xem chi tiết:

```bash
go test -v ./internal/array/combination_sum/
```

## Mở rộng

### Combination Sum II (LeetCode 40)

Mỗi số chỉ được sử dụng một lần. Cần xử lý duplicate trong mảng.

### Combination Sum III (LeetCode 216)

Tìm tất cả các tổ hợp của k số từ 1 đến 9 có tổng bằng n. Mỗi số chỉ được sử dụng một lần.

### Combination Sum IV (LeetCode 377)

Tìm số lượng các tổ hợp có thể (không phải tổ hợp cụ thể). Có thể giải bằng Dynamic Programming.

### Ứng dụng thực tế

Bài toán này có ứng dụng trong:
- **Tối ưu hóa**: Tìm các cách kết hợp tài nguyên
- **Tài chính**: Tìm các cách đầu tư với ngân sách cố định
- **Lập lịch**: Tìm các cách phân bổ thời gian
- **Game Development**: Tìm các cách kết hợp vật phẩm
- **Inventory Management**: Tìm các cách đóng gói hàng hóa
