# Maximum Product Subarray

## Mô tả bài toán

Cho một mảng số nguyên `nums`, tìm một subarray có tích lớn nhất và trả về tích đó.

Các test cases được tạo sao cho kết quả sẽ nằm trong phạm vi số nguyên 32-bit.

**Lưu ý:** Tích của một mảng với một phần tử duy nhất là giá trị của phần tử đó.

**Ví dụ:**

### Example 1:
- Input: `nums = [2,3,-2,4]`
- Output: `6`
- Giải thích: `[2,3]` có tích lớn nhất là `6`.

```
Subarray có thể:
[2] = 2
[2,3] = 6 ✓ (lớn nhất)
[2,3,-2] = -12
[2,3,-2,4] = -48
[3] = 3
[3,-2] = -6
[3,-2,4] = -24
[-2] = -2
[-2,4] = -8
[4] = 4
```

### Example 2:
- Input: `nums = [-2,0,-1]`
- Output: `0`
- Giải thích: Kết quả không thể là `2`, vì `[-2,-1]` không phải là một subarray (có `0` ở giữa).

```
Subarray có thể:
[-2] = -2
[-2,0] = 0 ✓ (lớn nhất)
[-2,0,-1] = 0 ✓
[0] = 0 ✓
[0,-1] = 0 ✓
[-1] = -1
```

**Ràng buộc:**
- `1 <= nums.length <= 2 * 10^4`
- `-10 <= nums[i] <= 10`
- Tích của bất kỳ subarray nào của `nums` đều đảm bảo nằm trong phạm vi số nguyên 32-bit

## Phân tích thuật toán

### Cách tiếp cận 1: Brute Force (O(n²) time, O(1) space)

Duyệt qua tất cả các subarray có thể và tính tích của chúng:

```go
func maxProductBruteForce(nums []int) int {
    result := nums[0]
    for i := 0; i < len(nums); i++ {
        product := 1
        for j := i; j < len(nums); j++ {
            product *= nums[j]
            if product > result {
                result = product
            }
        }
    }
    return result
}
```

**Độ phức tạp:**
- Thời gian: O(n²) - có n(n+1)/2 subarray
- Không gian: O(1)

**Nhược điểm:** Quá chậm với n lên tới 2 * 10^4

### Cách tiếp cận 2: Dynamic Programming với Max và Min (O(n) time, O(1) space) ⭐ Tối ưu

**Ý tưởng chính:**

Khi làm việc với tích, số âm có thể biến tích nhỏ nhất thành tích lớn nhất!

Ví dụ:
- `minProd = -10`, `nums[i] = -2` → `maxProd = 20` (từ tích nhỏ nhất)
- `maxProd = 5`, `nums[i] = -2` → `minProd = -10` (từ tích lớn nhất)

**Thuật toán:**

1. **Theo dõi hai giá trị tại mỗi vị trí:**
   - `maxProd`: Tích lớn nhất của subarray kết thúc tại vị trí hiện tại
   - `minProd`: Tích nhỏ nhất của subarray kết thúc tại vị trí hiện tại

2. **Tại mỗi vị trí i:**
   - Nếu `nums[i] < 0`: Swap `maxProd` và `minProd` (vì nhân với số âm sẽ đảo ngược thứ tự)
   - Cập nhật:
     - `maxProd = max(nums[i], maxProd * nums[i])`
     - `minProd = min(nums[i], minProd * nums[i])`
   - Cập nhật kết quả tổng thể: `result = max(result, maxProd)`

3. **Tại sao cần swap khi gặp số âm?**
   - Khi nhân với số âm, số lớn hơn sẽ trở thành số nhỏ hơn và ngược lại
   - Swap trước khi nhân để đảm bảo logic đúng

**Độ phức tạp:**
- Thời gian: O(n) - duyệt qua mảng một lần
- Không gian: O(1) - chỉ sử dụng thêm một vài biến

### Ví dụ minh họa

Với input: `nums = [2,3,-2,4]`

```
Bước 0: Khởi tạo
i = 0, nums[0] = 2
maxProd = 2, minProd = 2, result = 2

Bước 1: i = 1, nums[1] = 3 (dương)
maxProd = max(3, 2*3) = max(3, 6) = 6
minProd = min(3, 2*3) = min(3, 6) = 3
result = max(2, 6) = 6

Bước 2: i = 2, nums[2] = -2 (âm)
Swap: maxProd = 3, minProd = 6 (tạm thời)
maxProd = max(-2, 3*(-2)) = max(-2, -6) = -2
minProd = min(-2, 6*(-2)) = min(-2, -12) = -12
result = max(6, -2) = 6

Bước 3: i = 3, nums[3] = 4 (dương)
maxProd = max(4, -2*4) = max(4, -8) = 4
minProd = min(4, -12*4) = min(4, -48) = -48
result = max(6, 4) = 6

Kết quả: 6 (từ subarray [2,3])
```

Với input: `nums = [-2,0,-1]`

```
Bước 0: Khởi tạo
i = 0, nums[0] = -2
maxProd = -2, minProd = -2, result = -2

Bước 1: i = 1, nums[1] = 0
maxProd = max(0, -2*0) = max(0, 0) = 0
minProd = min(0, -2*0) = min(0, 0) = 0
result = max(-2, 0) = 0

Bước 2: i = 2, nums[2] = -1 (âm)
Swap: maxProd = 0, minProd = 0 (không thay đổi vì cả hai đều 0)
maxProd = max(-1, 0*(-1)) = max(-1, 0) = 0
minProd = min(-1, 0*(-1)) = min(-1, 0) = -1
result = max(0, 0) = 0

Kết quả: 0
```

## Giải thích code

### Hàm maxProduct

```9:50:internal/array/maximum_product_subarray/maximum_product_subarray.go
// maxProduct tìm tích lớn nhất của một subarray trong mảng.
//
// Sử dụng Dynamic Programming với hai biến:
// - maxProd: Tích lớn nhất kết thúc tại vị trí hiện tại
// - minProd: Tích nhỏ nhất kết thúc tại vị trí hiện tại
//
// Lý do cần cả max và min:
// - Khi gặp số âm, tích nhỏ nhất có thể trở thành tích lớn nhất
// - Ví dụ: minProd = -10, nums[i] = -2 => maxProd = 20
//
// Độ phức tạp: O(n) thời gian, O(1) không gian
func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Khởi tạo với phần tử đầu tiên
	maxProd := nums[0]
	minProd := nums[0]
	result := nums[0]

	// Duyệt qua các phần tử còn lại
	for i := 1; i < len(nums); i++ {
		num := nums[i]

		// Nếu số hiện tại là số âm, swap max và min
		// vì tích với số âm sẽ đảo ngược thứ tự
		if num < 0 {
			maxProd, minProd = minProd, maxProd
		}

		// Cập nhật max và min product kết thúc tại vị trí i
		// Có hai lựa chọn:
		// 1. Bắt đầu subarray mới từ vị trí i (num)
		// 2. Mở rộng subarray hiện tại (maxProd * num hoặc minProd * num)
		maxProd = max(num, maxProd*num)
		minProd = min(num, minProd*num)

		// Cập nhật kết quả tổng thể
		if maxProd > result {
			result = maxProd
		}
	}

	return result
}
```

### Chi tiết từng phần

#### 1. Khởi tạo (dòng 18-21)
```go
maxProd := nums[0]
minProd := nums[0]
result := nums[0]
```
- Khởi tạo với phần tử đầu tiên
- Subarray chỉ có một phần tử: tích chính là giá trị phần tử đó

#### 2. Swap khi gặp số âm (dòng 27-30)
```go
if num < 0 {
    maxProd, minProd = minProd, maxProd
}
```
- Khi nhân với số âm, số lớn hơn sẽ trở thành số nhỏ hơn
- Swap trước khi nhân để đảm bảo logic đúng
- Ví dụ: `maxProd = 5`, `minProd = -10`, `num = -2`
  - Sau swap: `maxProd = -10`, `minProd = 5`
  - Sau nhân: `maxProd = 20`, `minProd = -10` ✓

#### 3. Cập nhật max và min (dòng 32-36)
```go
maxProd = max(num, maxProd*num)
minProd = min(num, minProd*num)
```
- **Hai lựa chọn:**
  1. Bắt đầu subarray mới từ vị trí i: `num`
  2. Mở rộng subarray hiện tại: `maxProd * num` hoặc `minProd * num`
- Chọn giá trị lớn nhất/nhỏ nhất giữa hai lựa chọn

#### 4. Cập nhật kết quả (dòng 38-41)
```go
if maxProd > result {
    result = maxProd
}
```
- Theo dõi tích lớn nhất tổng thể
- Cập nhật khi tìm thấy tích lớn hơn

### Hàm helper: max và min

```52:66:internal/array/maximum_product_subarray/maximum_product_subarray.go
// max trả về giá trị lớn hơn giữa a và b
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// min trả về giá trị nhỏ hơn giữa a và b
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```

### Độ phức tạp

- **Thời gian**: O(n)
  - Duyệt qua mảng một lần duy nhất
  - Mỗi phần tử được xử lý trong thời gian O(1)

- **Không gian**: O(1)
  - Chỉ sử dụng thêm một vài biến (`maxProd`, `minProd`, `result`)
  - Không sử dụng thêm cấu trúc dữ liệu nào

### Tại sao thuật toán này đúng?

**Chứng minh:**

1. **Tính đúng đắn của việc theo dõi max và min:**
   - Tại mỗi vị trí i, ta cần tìm tích lớn nhất/nhỏ nhất của subarray kết thúc tại i
   - Có hai lựa chọn: bắt đầu mới hoặc mở rộng từ vị trí trước
   - Chọn giá trị lớn nhất/nhỏ nhất giữa hai lựa chọn

2. **Tính đúng đắn của việc swap:**
   - Khi nhân với số âm: `a > b` → `a * (-1) < b * (-1)`
   - Swap trước khi nhân để đảm bảo sau khi nhân, `maxProd` vẫn là lớn nhất và `minProd` vẫn là nhỏ nhất

3. **Tính đầy đủ:**
   - Xét tất cả các subarray kết thúc tại mỗi vị trí
   - Subarray kết thúc tại i có thể là:
     - Chỉ có `nums[i]`
     - Mở rộng từ subarray kết thúc tại i-1
   - Ta xét cả hai trường hợp và chọn giá trị tối ưu

## Test Cases

### Test Case 1: Example 1
```go
Input: nums = [2,3,-2,4]
Output: 6
```
Subarray `[2,3]` có tích lớn nhất.

### Test Case 2: Example 2
```go
Input: nums = [-2,0,-1]
Output: 0
```
Subarray chứa `0` có tích lớn nhất.

### Test Case 3: Single element
```go
Input: nums = [5]
Output: 5
```
Mảng chỉ có một phần tử.

### Test Case 4: Single negative element
```go
Input: nums = [-5]
Output: -5
```
Mảng chỉ có một phần tử âm.

### Test Case 5: All positive
```go
Input: nums = [1,2,3,4]
Output: 24
```
Tích của toàn bộ mảng là lớn nhất.

### Test Case 6: All negative - even count
```go
Input: nums = [-1,-2,-3,-4]
Output: 24
```
Số lượng số âm chẵn, tích của toàn bộ mảng là dương và lớn nhất.

### Test Case 7: All negative - odd count
```go
Input: nums = [-1,-2,-3]
Output: 6
```
Số lượng số âm lẻ, tích của hai số đầu là lớn nhất.

### Test Case 8: Two negatives make positive
```go
Input: nums = [-2,3,-4]
Output: 24
```
Tích của toàn bộ mảng (có hai số âm) là lớn nhất.

### Test Case 9: Zero resets product
```go
Input: nums = [2,0,3,4]
Output: 12
```
Số `0` chia mảng thành hai phần, phần sau có tích lớn nhất.

### Test Case 10: Complex case
```go
Input: nums = [2,-5,-2,-4,3]
Output: 24
```
Subarray `[-5,-2,-4]` có tích `-40`, nhưng khi nhân với `3` ở sau sẽ thành `-120`.
Subarray `[-2,-4,3]` có tích `24` là lớn nhất.

## Chạy test

Để chạy các test case:

```bash
go test ./internal/array/maximum_product_subarray/
```

Hoặc chạy với verbose mode để xem chi tiết:

```bash
go test -v ./internal/array/maximum_product_subarray/
```

## Mở rộng

### Biến thể: Maximum Sum Subarray (Kadane's Algorithm)

Bài toán tương tự nhưng với tổng thay vì tích:

```go
func maxSubArray(nums []int) int {
    maxSum := nums[0]
    currentSum := nums[0]
    
    for i := 1; i < len(nums); i++ {
        currentSum = max(nums[i], currentSum + nums[i])
        maxSum = max(maxSum, currentSum)
    }
    
    return maxSum
}
```

**Khác biệt:**
- Với tổng, chỉ cần theo dõi một giá trị (max)
- Với tích, cần theo dõi cả max và min (vì số âm)

### Biến thể: Maximum Product với K phần tử

Tìm tích lớn nhất của subarray có đúng K phần tử:

```go
func maxProductK(nums []int, k int) int {
    // Sliding window với kích thước cố định k
    // ...
}
```

### Ứng dụng thực tế

Thuật toán này được sử dụng trong:

- **Phân tích tài chính**: Tìm khoảng thời gian có lợi nhuận tích lũy lớn nhất
- **Xử lý tín hiệu**: Tìm đoạn tín hiệu có tích lớn nhất
- **Machine Learning**: Tối ưu hóa các hàm mục tiêu liên quan đến tích
- **Game development**: Tính toán điểm số hoặc multiplier tích lũy

### So sánh với các cách tiếp cận khác

| Cách tiếp cận | Thời gian | Không gian | Ưu điểm | Nhược điểm |
|--------------|-----------|------------|---------|------------|
| Brute Force | O(n²) | O(1) | Dễ hiểu | Quá chậm |
| DP với Max/Min | O(n) | O(1) | Tối ưu | Cần hiểu logic swap |
| DP với mảng | O(n) | O(n) | Dễ debug | Tốn memory |

### Tips và Tricks

1. **Luôn theo dõi cả max và min**: Số âm có thể biến min thành max
2. **Swap khi gặp số âm**: Đảm bảo logic đúng khi nhân với số âm
3. **Xử lý số 0**: Số 0 reset tích về 0, có thể bắt đầu subarray mới
4. **Edge case**: Mảng chỉ có một phần tử, mảng toàn số âm
5. **Khởi tạo đúng**: Bắt đầu với phần tử đầu tiên, không phải 0 hoặc 1

### Lưu ý về Overflow

Mặc dù đề bài đảm bảo kết quả nằm trong phạm vi 32-bit, nhưng trong quá trình tính toán có thể xảy ra overflow tạm thời. Trong Go, kiểu `int` thường là 64-bit trên các hệ thống hiện đại, nên không có vấn đề. Tuy nhiên, nếu cần xử lý với 32-bit integer, có thể cần kiểm tra overflow.

### Tối ưu hóa

Có thể viết lại code không cần swap bằng cách tính cả hai giá trị và chọn:

```go
func maxProductOptimized(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    
    maxProd := nums[0]
    minProd := nums[0]
    result := nums[0]
    
    for i := 1; i < len(nums); i++ {
        num := nums[i]
        
        // Tính cả hai khả năng
        tempMax := max(num, max(maxProd*num, minProd*num))
        tempMin := min(num, min(maxProd*num, minProd*num))
        
        maxProd = tempMax
        minProd = tempMin
        result = max(result, maxProd)
    }
    
    return result
}
```

Cách này không cần swap nhưng phải tính nhiều hơn một chút. Cả hai cách đều có độ phức tạp O(n) và O(1) không gian.
