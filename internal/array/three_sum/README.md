# Three Sum

## Mô tả bài toán

Cho một mảng số nguyên `nums`, trả về tất cả các bộ ba `[nums[i], nums[j], nums[k]]` sao cho `i != j`, `i != k`, và `j != k`, và `nums[i] + nums[j] + nums[k] == 0`.

Lưu ý rằng tập nghiệm không được chứa các bộ ba trùng lặp.

**Ví dụ:**

### Example 1:
- Input: `nums = [-1,0,1,2,-1,-4]`
- Output: `[[-1,-1,2],[-1,0,1]]`
- Giải thích:
  - `nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0`
  - `nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0`
  - `nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0`
  - Các bộ ba phân biệt là `[-1,0,1]` và `[-1,-1,2]`
  - Lưu ý rằng thứ tự của output và thứ tự của các bộ ba không quan trọng

### Example 2:
- Input: `nums = [0,1,1]`
- Output: `[]`
- Giải thích: Bộ ba duy nhất có thể không có tổng bằng 0

### Example 3:
- Input: `nums = [0,0,0]`
- Output: `[[0,0,0]]`
- Giải thích: Bộ ba duy nhất có tổng bằng 0

**Ràng buộc:**
- `3 <= nums.length <= 3000`
- `-10^5 <= nums[i] <= 10^5`

## Phân tích thuật toán

### Cách tiếp cận 1: Brute Force (O(n³) time, O(1) space)

Kiểm tra tất cả các tổ hợp có thể có của ba số.

**Độ phức tạp:**
- Thời gian: O(n³) - có C(n,3) = n(n-1)(n-2)/6 tổ hợp
- Không gian: O(1) (không tính output array)

### Cách tiếp cận 2: Hash Set (O(n²) time, O(n) space)

Với mỗi cặp số đầu tiên, sử dụng hash set để tìm số thứ ba.

**Độ phức tạp:**
- Thời gian: O(n²)
- Không gian: O(n) cho hash set

### Cách tiếp cận 3: Two Pointers (O(n²) time, O(1) space) ⭐ Tối ưu

**Ý tưởng:**
1. Sắp xếp mảng để dễ dàng xử lý duplicate và sử dụng two pointers
2. Với mỗi phần tử đầu tiên, sử dụng two pointers để tìm hai số còn lại
3. Tránh duplicate bằng cách skip các phần tử giống nhau

**Các bước thực hiện:**
1. **Sắp xếp mảng**: Cho phép sử dụng two pointers và dễ dàng skip duplicate
2. **Duyệt qua phần tử đầu tiên**: Với mỗi `nums[i]`, tìm hai số `nums[left]` và `nums[right]` sao cho tổng bằng `-nums[i]`
3. **Two Pointers**: 
   - `left` bắt đầu từ `i+1`
   - `right` bắt đầu từ cuối mảng
   - Nếu tổng quá nhỏ, tăng `left`
   - Nếu tổng quá lớn, giảm `right`
   - Nếu tổng bằng target, thêm vào kết quả và skip duplicate
4. **Xử lý duplicate**: Skip các phần tử giống nhau để tránh duplicate triplets

**Độ phức tạp:**
- Thời gian: O(n²) - vòng lặp ngoài O(n), vòng lặp trong O(n)
- Không gian: O(1) - chỉ sử dụng các biến (không tính output array và không gian sắp xếp)

### Ví dụ minh họa

Với input: `nums = [-1,0,1,2,-1,-4]`

```
Bước 1: Sắp xếp mảng
  nums = [-4, -1, -1, 0, 1, 2]

Bước 2: Duyệt qua từng phần tử đầu tiên

i = 0, nums[i] = -4:
  target = 4
  left = 1, right = 5
  nums[left] + nums[right] = -1 + 2 = 1 < 4 → left++
  nums[left] + nums[right] = -1 + 2 = 1 < 4 → left++
  nums[left] + nums[right] = 0 + 2 = 2 < 4 → left++
  nums[left] + nums[right] = 1 + 2 = 3 < 4 → left++
  left >= right → dừng

i = 1, nums[i] = -1:
  target = 1
  left = 2, right = 5
  nums[left] + nums[right] = -1 + 2 = 1 == 1 ✓
    → Thêm [-1, -1, 2] vào kết quả
    → Skip duplicate: left++ đến 3, right-- đến 4
  nums[left] + nums[right] = 0 + 1 = 1 == 1 ✓
    → Thêm [-1, 0, 1] vào kết quả
    → left++, right-- → left >= right → dừng

i = 2, nums[i] = -1:
  Skip vì nums[2] == nums[1] (duplicate)

i = 3, nums[i] = 0:
  target = 0
  left = 4, right = 5
  nums[left] + nums[right] = 1 + 2 = 3 > 0 → right--
  left >= right → dừng

Kết quả: [[-1, -1, 2], [-1, 0, 1]]
```

## Giải thích code

### Cấu trúc hàm

```1:60:internal/array/three_sum/three_sum.go
package array

import "sort"

// threeSum tìm tất cả các bộ ba số có tổng bằng 0 trong mảng.
//
// Sử dụng kỹ thuật Two Pointers:
// 1. Sắp xếp mảng để dễ dàng xử lý duplicate và sử dụng two pointers
// 2. Với mỗi phần tử đầu tiên, sử dụng two pointers để tìm hai số còn lại
// 3. Tránh duplicate bằng cách skip các phần tử giống nhau
//
// Độ phức tạp: O(n²) thời gian, O(1) không gian (không tính output array)
func threeSum(nums []int) [][]int {
	n := len(nums)
	if n < 3 {
		return [][]int{}
	}

	// Sắp xếp mảng để dễ dàng xử lý duplicate và sử dụng two pointers
	sort.Ints(nums)
	result := [][]int{}

	// Duyệt qua từng phần tử làm phần tử đầu tiên
	for i := 0; i < n-2; i++ {
		// Skip duplicate cho phần tử đầu tiên
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// Two pointers: left bắt đầu từ i+1, right từ cuối mảng
		left := i + 1
		right := n - 1
		target := -nums[i] // Tổng của hai số còn lại phải bằng -nums[i]

		for left < right {
			sum := nums[left] + nums[right]

			if sum == target {
				// Tìm thấy một bộ ba hợp lệ
				result = append(result, []int{nums[i], nums[left], nums[right]})

				// Skip duplicate cho left pointer
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				// Skip duplicate cho right pointer
				for left < right && nums[right] == nums[right-1] {
					right--
				}

				// Di chuyển cả hai pointers
				left++
				right--
			} else if sum < target {
				// Tổng quá nhỏ, cần tăng left pointer
				left++
			} else {
				// Tổng quá lớn, cần giảm right pointer
				right--
			}
		}
	}

	return result
}
```

### Chi tiết từng phần

#### 1. Kiểm tra trường hợp biên (dòng 15-18)
```go
n := len(nums)
if n < 3 {
    return [][]int{}
}
```
- Nếu mảng có ít hơn 3 phần tử, không thể tạo bộ ba
- Trả về mảng rỗng ngay lập tức

#### 2. Sắp xếp mảng (dòng 20-21)
```go
sort.Ints(nums)
result := [][]int{}
```
- Sắp xếp mảng để:
  - Dễ dàng sử dụng two pointers
  - Dễ dàng skip duplicate (các phần tử giống nhau sẽ nằm cạnh nhau)
- Khởi tạo mảng kết quả

#### 3. Vòng lặp ngoài - Duyệt phần tử đầu tiên (dòng 23-28)
```go
for i := 0; i < n-2; i++ {
    if i > 0 && nums[i] == nums[i-1] {
        continue
    }
    // ...
}
```
- Duyệt qua từng phần tử làm phần tử đầu tiên của bộ ba
- `i < n-2`: Đảm bảo còn đủ 2 phần tử cho left và right
- Skip duplicate: Nếu `nums[i] == nums[i-1]`, bỏ qua vì đã xử lý

#### 4. Khởi tạo Two Pointers (dòng 30-32)
```go
left := i + 1
right := n - 1
target := -nums[i]
```
- `left`: Bắt đầu từ phần tử sau `i`
- `right`: Bắt đầu từ cuối mảng
- `target`: Tổng của `nums[left] + nums[right]` phải bằng `-nums[i]` để tổng ba số bằng 0

#### 5. Vòng lặp Two Pointers (dòng 34-54)
```go
for left < right {
    sum := nums[left] + nums[right]
    
    if sum == target {
        // Tìm thấy bộ ba hợp lệ
        result = append(result, []int{nums[i], nums[left], nums[right]})
        
        // Skip duplicate
        for left < right && nums[left] == nums[left+1] {
            left++
        }
        for left < right && nums[right] == nums[right-1] {
            right--
        }
        
        left++
        right--
    } else if sum < target {
        left++
    } else {
        right--
    }
}
```

**Giải thích:**
- `sum == target`: Tìm thấy bộ ba hợp lệ
  - Thêm vào kết quả
  - Skip duplicate cho cả left và right
  - Di chuyển cả hai pointers
- `sum < target`: Tổng quá nhỏ, cần số lớn hơn → tăng `left`
- `sum > target`: Tổng quá lớn, cần số nhỏ hơn → giảm `right`

**Tại sao cần skip duplicate?**

Để tránh các bộ ba trùng lặp. Ví dụ với `nums = [-1, -1, 0, 1]`:
- Nếu không skip, có thể tạo `[-1, 0, 1]` hai lần (với `nums[0]` và `nums[1]`)

### Độ phức tạp

- **Thời gian**: O(n²)
  - Sắp xếp: O(n log n)
  - Vòng lặp ngoài: O(n)
  - Vòng lặp trong (two pointers): O(n)
  - Tổng: O(n log n) + O(n²) = O(n²) (dominant term)

- **Không gian**: O(1)
  - Chỉ sử dụng các biến: `n`, `left`, `right`, `target`, `sum`
  - Không tính output array và không gian sắp xếp (thường là O(log n) cho quicksort)

### Tại sao thuật toán này đúng?

1. **Tính đầy đủ**: Với mỗi phần tử đầu tiên, ta kiểm tra tất cả các cặp có thể với two pointers, đảm bảo không bỏ sót bộ ba nào.

2. **Tính đúng đắn**: Khi tìm thấy `sum == target`, ta có `nums[i] + nums[left] + nums[right] = 0`.

3. **Tránh duplicate**: 
   - Skip duplicate cho phần tử đầu tiên: `if nums[i] == nums[i-1]`
   - Skip duplicate cho left và right sau khi tìm thấy bộ ba hợp lệ

4. **Tối ưu**: Two pointers giảm độ phức tạp từ O(n³) xuống O(n²).

## Test Cases

### Test Case 1: Example 1
```go
Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
```
Có hai bộ ba phân biệt có tổng bằng 0.

### Test Case 2: Example 2
```go
Input: nums = [0,1,1]
Output: []
```
Không có bộ ba nào có tổng bằng 0.

### Test Case 3: Example 3
```go
Input: nums = [0,0,0]
Output: [[0,0,0]]
```
Bộ ba duy nhất có tổng bằng 0.

### Test Case 4: Empty array
```go
Input: nums = []
Output: []
```
Mảng rỗng.

### Test Case 5: Less than 3 elements
```go
Input: nums = [1,2]
Output: []
```
Không đủ phần tử để tạo bộ ba.

### Test Case 6: All zeros
```go
Input: nums = [0,0,0,0]
Output: [[0,0,0]]
```
Chỉ có một bộ ba duy nhất (không duplicate).

### Test Case 7: Multiple solutions
```go
Input: nums = [-2,0,1,1,2]
Output: [[-2,0,2],[-2,1,1]]
```
Có nhiều bộ ba khác nhau.

## Chạy test

Để chạy các test case:

```bash
go test ./internal/array/three_sum/
```

Hoặc chạy với verbose mode để xem chi tiết:

```bash
go test -v ./internal/array/three_sum/
```

## Mở rộng

### 3Sum Closest (LeetCode 16)

Một biến thể của bài toán này là tìm bộ ba có tổng gần nhất với một giá trị target cho trước (không nhất thiết bằng 0). Giải pháp tương tự nhưng cần theo dõi tổng gần nhất.

### 4Sum (LeetCode 18)

Tìm tất cả các bộ bốn số có tổng bằng target. Có thể mở rộng từ 3Sum bằng cách thêm một vòng lặp ngoài nữa.

### K-Sum Problem

Tổng quát hóa cho K số. Có thể giải bằng đệ quy hoặc dynamic programming.

### Ứng dụng thực tế

Bài toán này có ứng dụng trong:
- Tìm các tổ hợp số trong tài chính
- Phân tích dữ liệu và thống kê
- Tối ưu hóa và tìm kiếm
- Xử lý hình ảnh và computer vision
