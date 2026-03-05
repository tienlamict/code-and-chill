# Maximum Subarray

## Mô tả bài toán

Cho một mảng số nguyên `nums`, tìm một subarray (mảng con liên tiếp) có tổng lớn nhất và trả về tổng đó.

**Subarray** là một dãy liên tiếp các phần tử trong mảng gốc.

**Ví dụ:**

### Example 1:
- Input: `nums = [-2,1,-3,4,-1,2,1,-5,4]`
- Output: `6`
- Giải thích: Subarray `[4,-1,2,1]` có tổng lớn nhất là `6`.

```
Một số subarray có thể:
[-2] = -2
[-2,1] = -1
[-2,1,-3] = -4
[1] = 1
[1,-3] = -2
[1,-3,4] = 2
[4] = 4
[4,-1] = 3
[4,-1,2] = 5
[4,-1,2,1] = 6 ✓ (lớn nhất)
[4,-1,2,1,-5] = 1
[-1,2,1] = 2
[2,1] = 3
...
```

### Example 2:
- Input: `nums = [1]`
- Output: `1`
- Giải thích: Subarray `[1]` có tổng lớn nhất là `1`.

### Example 3:
- Input: `nums = [5,4,-1,7,8]`
- Output: `23`
- Giải thích: Subarray `[5,4,-1,7,8]` (toàn bộ mảng) có tổng lớn nhất là `23`.

**Ràng buộc:**
- `1 <= nums.length <= 10^5`
- `-10^4 <= nums[i] <= 10^4`

**Follow-up:** Nếu bạn đã giải được bài toán với độ phức tạp O(n), hãy thử code một giải pháp khác sử dụng phương pháp **divide and conquer** (chia để trị), cách này tinh tế hơn.

## Phân tích thuật toán

### Cách tiếp cận 1: Brute Force (O(n²) time, O(1) space)

Duyệt qua tất cả các subarray có thể và tính tổng của chúng:

```go
func maxSubArrayBruteForce(nums []int) int {
    result := nums[0]
    for i := 0; i < len(nums); i++ {
        sum := 0
        for j := i; j < len(nums); j++ {
            sum += nums[j]
            if sum > result {
                result = sum
            }
        }
    }
    return result
}
```

**Độ phức tạp:**
- Thời gian: O(n²) - có n(n+1)/2 subarray
- Không gian: O(1)

**Nhược điểm:** Quá chậm với n lên tới 10^5

### Cách tiếp cận 2: Kadane's Algorithm - Dynamic Programming (O(n) time, O(1) space) ⭐ Tối ưu nhất

**Ý tưởng chính:**

Tại mỗi vị trí, ta có hai lựa chọn:
1. **Bắt đầu một subarray mới** từ vị trí hiện tại
2. **Mở rộng subarray hiện tại** bằng cách thêm phần tử hiện tại vào

Chọn lựa chọn nào cho tổng lớn hơn!

**Thuật toán:**

1. **Khởi tạo:**
   - `maxSum`: Tổng lớn nhất tổng thể (kết quả cuối cùng)
   - `currentSum`: Tổng của subarray kết thúc tại vị trí hiện tại

2. **Tại mỗi vị trí i:**
   - `currentSum = max(nums[i], currentSum + nums[i])`
     - Nếu `currentSum + nums[i] < nums[i]`, tức là `currentSum < 0`
     - Thì tốt hơn là bắt đầu subarray mới từ `nums[i]`
   - Cập nhật `maxSum = max(maxSum, currentSum)`

3. **Tại sao thuật toán này đúng?**
   - Nếu tổng hiện tại âm, việc thêm nó vào phần tử tiếp theo chỉ làm giảm tổng
   - Do đó, tốt hơn là bắt đầu lại từ phần tử hiện tại
   - Ta luôn theo dõi tổng lớn nhất gặp được

**Độ phức tạp:**
- Thời gian: O(n) - duyệt qua mảng một lần
- Không gian: O(1) - chỉ sử dụng thêm hai biến

### Cách tiếp cận 3: Divide and Conquer (O(n log n) time, O(log n) space)

**Ý tưởng chính:**

Chia mảng thành hai nửa. Subarray có tổng lớn nhất có thể nằm ở:
1. **Nửa trái** (left half)
2. **Nửa phải** (right half)
3. **Qua điểm giữa** (crossing middle) - bắt đầu từ bên trái mid, qua mid, đến bên phải mid

**Thuật toán:**

1. **Base case:** Nếu mảng chỉ có một phần tử, trả về phần tử đó

2. **Chia:** Chia mảng thành hai nửa tại điểm giữa `mid`

3. **Trị:**
   - Tìm tổng lớn nhất ở nửa trái (đệ quy)
   - Tìm tổng lớn nhất ở nửa phải (đệ quy)
   - Tìm tổng lớn nhất qua điểm giữa:
     - Tìm tổng lớn nhất từ mid về bên trái
     - Tìm tổng lớn nhất từ mid+1 về bên phải
     - Cộng hai tổng lại

4. **Kết hợp:** Trả về giá trị lớn nhất trong ba trường hợp

**Độ phức tạp:**
- Thời gian: O(n log n)
  - Mỗi lần chia đôi mảng: O(log n) levels
  - Mỗi level xử lý O(n) phần tử (tính crossing sum)
  - Tổng: O(n log n)
- Không gian: O(log n) - call stack của đệ quy

**So sánh:** Kadane's Algorithm nhanh hơn, nhưng Divide and Conquer là một cách tiếp cận thú vị và có thể mở rộng cho các bài toán phức tạp hơn.

## Ví dụ minh họa

### Ví dụ 1: Kadane's Algorithm

Với input: `nums = [-2,1,-3,4,-1,2,1,-5,4]`

```
Bước 0: Khởi tạo
i = 0, nums[0] = -2
currentSum = -2, maxSum = -2

Bước 1: i = 1, nums[1] = 1
currentSum = max(1, -2+1) = max(1, -1) = 1
maxSum = max(-2, 1) = 1

Bước 2: i = 2, nums[2] = -3
currentSum = max(-3, 1+(-3)) = max(-3, -2) = -2
maxSum = max(1, -2) = 1

Bước 3: i = 3, nums[3] = 4
currentSum = max(4, -2+4) = max(4, 2) = 4
maxSum = max(1, 4) = 4

Bước 4: i = 4, nums[4] = -1
currentSum = max(-1, 4+(-1)) = max(-1, 3) = 3
maxSum = max(4, 3) = 4

Bước 5: i = 5, nums[5] = 2
currentSum = max(2, 3+2) = max(2, 5) = 5
maxSum = max(4, 5) = 5

Bước 6: i = 6, nums[6] = 1
currentSum = max(1, 5+1) = max(1, 6) = 6
maxSum = max(5, 6) = 6

Bước 7: i = 7, nums[7] = -5
currentSum = max(-5, 6+(-5)) = max(-5, 1) = 1
maxSum = max(6, 1) = 6

Bước 8: i = 8, nums[8] = 4
currentSum = max(4, 1+4) = max(4, 5) = 5
maxSum = max(6, 5) = 6

Kết quả: 6 (từ subarray [4,-1,2,1])
```

### Ví dụ 2: Divide and Conquer

Với input: `nums = [-2,1,-3,4,-1,2,1,-5,4]`

```
Level 0: [−2,1,−3,4,−1,2,1,−5,4]
         mid = 4
         ├─ Left:  [−2,1,−3,4,−1]
         ├─ Right: [2,1,−5,4]
         └─ Cross: ?

Level 1 (Left): [−2,1,−3,4,−1]
                mid = 2
                ├─ Left:  [−2,1,−3]
                ├─ Right: [4,−1]
                └─ Cross: ?

Level 2 (Left-Left): [−2,1,−3]
                     mid = 1
                     ├─ Left:  [−2,1]
                     ├─ Right: [−3]
                     └─ Cross: ?

Level 3 (Left-Left-Left): [−2,1]
                          mid = 0
                          ├─ Left:  [−2] = -2
                          ├─ Right: [1] = 1
                          └─ Cross: -2 + 1 = -1
                          → max(-2, 1, -1) = 1

Level 3 (Left-Left-Right): [−3] = -3

Level 2 (Left-Left) Cross:
    From mid=1 going left:  max(1, 1+(-2)) = 1
    From mid=2 going right: -3
    Cross sum: 1 + (-3) = -2
    → max(1, -3, -2) = 1

Level 2 (Left-Right): [4,−1]
                      mid = 0
                      ├─ Left:  [4] = 4
                      ├─ Right: [−1] = -1
                      └─ Cross: 4 + (-1) = 3
                      → max(4, -1, 3) = 4

Level 1 (Left) Cross:
    From mid=2 going left:  max(-3, -3+1, -3+1+(-2)) = 1
    From mid=3 going right: max(4, 4+(-1)) = 4
    Cross sum: 1 + 4 = 5
    → max(1, 4, 5) = 5

Level 1 (Right): [2,1,−5,4]
                 mid = 1
                 ├─ Left:  [2,1]
                 ├─ Right: [−5,4]
                 └─ Cross: ?

Level 2 (Right-Left): [2,1]
                      ├─ Left:  [2] = 2
                      ├─ Right: [1] = 1
                      └─ Cross: 2 + 1 = 3
                      → max(2, 1, 3) = 3

Level 2 (Right-Right): [−5,4]
                       ├─ Left:  [−5] = -5
                       ├─ Right: [4] = 4
                       └─ Cross: -5 + 4 = -1
                       → max(-5, 4, -1) = 4

Level 1 (Right) Cross:
    From mid=1 going left:  max(1, 1+2) = 3
    From mid=2 going right: max(-5, -5+4) = -1
    Cross sum: 3 + (-1) = 2
    → max(3, 4, 2) = 4

Level 0 Cross:
    From mid=4 going left:  max(-1, -1+4, -1+4+(-3), -1+4+(-3)+1, -1+4+(-3)+1+(-2)) = 4
    From mid=5 going right: max(2, 2+1, 2+1+(-5), 2+1+(-5)+4) = 3
    Cross sum: 4 + 3 = 7

    Wait, let me recalculate the cross sum more carefully:
    From mid=4 going left:  
        i=4: sum = -1, leftSum = -1
        i=3: sum = -1+4 = 3, leftSum = 3
        i=2: sum = 3+(-3) = 0, leftSum = 3
        i=1: sum = 0+1 = 1, leftSum = 3
        i=0: sum = 1+(-2) = -1, leftSum = 3
        → leftSum = 3
    
    From mid=5 going right:
        i=5: sum = 2, rightSum = 2
        i=6: sum = 2+1 = 3, rightSum = 3
        i=7: sum = 3+(-5) = -2, rightSum = 3
        i=8: sum = -2+4 = 2, rightSum = 3
        → rightSum = 3
    
    Cross sum: 3 + 3 = 6

Final: max(5, 4, 6) = 6
```

## Giải thích code

### Cách 1: Kadane's Algorithm

```3:36:internal/array/maximum_subarray/maximum_subarray.go
// maxSubArray tìm tổng lớn nhất của một subarray trong mảng.
//
// Sử dụng Kadane's Algorithm - Dynamic Programming approach:
// - Tại mỗi vị trí, ta có hai lựa chọn:
//   1. Bắt đầu một subarray mới từ vị trí hiện tại
//   2. Mở rộng subarray hiện tại bằng cách thêm phần tử hiện tại
// - Chọn lựa chọn nào cho tổng lớn hơn
//
// Độ phức tạp: O(n) thời gian, O(1) không gian
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Khởi tạo với phần tử đầu tiên
	maxSum := nums[0]      // Tổng lớn nhất tổng thể
	currentSum := nums[0]  // Tổng của subarray kết thúc tại vị trí hiện tại

	// Duyệt qua các phần tử còn lại
	for i := 1; i < len(nums); i++ {
		// Quyết định: bắt đầu mới hay mở rộng subarray hiện tại?
		// Nếu currentSum + nums[i] < nums[i], tức là currentSum < 0
		// thì tốt hơn là bắt đầu subarray mới từ nums[i]
		currentSum = max(nums[i], currentSum+nums[i])

		// Cập nhật tổng lớn nhất tổng thể
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	return maxSum
}
```

#### Chi tiết từng phần

**1. Khởi tạo (dòng 17-18)**
```go
maxSum := nums[0]      // Tổng lớn nhất tổng thể
currentSum := nums[0]  // Tổng của subarray kết thúc tại vị trí hiện tại
```
- Khởi tạo với phần tử đầu tiên
- `maxSum`: Lưu kết quả cuối cùng - tổng lớn nhất tìm được
- `currentSum`: Tổng của subarray kết thúc tại vị trí hiện tại

**2. Quyết định: Bắt đầu mới hay mở rộng? (dòng 24)**
```go
currentSum = max(nums[i], currentSum+nums[i])
```
- **Hai lựa chọn:**
  1. Bắt đầu subarray mới từ `nums[i]`
  2. Mở rộng subarray hiện tại: `currentSum + nums[i]`
- **Khi nào bắt đầu mới?**
  - Khi `currentSum < 0`: Thêm số âm vào chỉ làm giảm tổng
  - Tốt hơn là bắt đầu lại từ `nums[i]`
- **Khi nào mở rộng?**
  - Khi `currentSum >= 0`: Thêm vào có thể tăng tổng
  - Mở rộng subarray hiện tại

**3. Cập nhật kết quả (dòng 27-29)**
```go
if currentSum > maxSum {
    maxSum = currentSum
}
```
- Theo dõi tổng lớn nhất tổng thể
- Cập nhật khi tìm thấy tổng lớn hơn

### Cách 2: Divide and Conquer

```38:52:internal/array/maximum_subarray/maximum_subarray.go
// maxSubArrayDivideConquer tìm tổng lớn nhất sử dụng phương pháp chia để trị.
//
// Ý tưởng:
// - Chia mảng thành hai nửa
// - Tổng lớn nhất có thể nằm ở:
//   1. Nửa trái
//   2. Nửa phải
//   3. Qua điểm giữa (cross middle)
// - Trả về giá trị lớn nhất trong ba trường hợp
//
// Độ phức tạp: O(n log n) thời gian, O(log n) không gian (call stack)
func maxSubArrayDivideConquer(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	return divideConquerHelper(nums, 0, len(nums)-1)
}
```

```54:76:internal/array/maximum_subarray/maximum_subarray.go
// divideConquerHelper hàm đệ quy để tính tổng lớn nhất trong đoạn [left, right]
func divideConquerHelper(nums []int, left, right int) int {
	// Base case: chỉ có một phần tử
	if left == right {
		return nums[left]
	}

	// Chia đôi mảng
	mid := left + (right-left)/2

	// Tìm tổng lớn nhất ở ba vị trí:
	// 1. Nửa trái
	leftMax := divideConquerHelper(nums, left, mid)

	// 2. Nửa phải
	rightMax := divideConquerHelper(nums, mid+1, right)

	// 3. Qua điểm giữa
	crossMax := maxCrossingSum(nums, left, mid, right)

	// Trả về giá trị lớn nhất trong ba trường hợp
	return max(max(leftMax, rightMax), crossMax)
}
```

#### Chi tiết từng phần

**1. Base case (dòng 57-59)**
```go
if left == right {
    return nums[left]
}
```
- Khi mảng chỉ có một phần tử, trả về phần tử đó

**2. Chia đôi mảng (dòng 62)**
```go
mid := left + (right-left)/2
```
- Tìm điểm giữa
- Sử dụng `left + (right-left)/2` thay vì `(left+right)/2` để tránh overflow

**3. Tìm max ở ba vị trí (dòng 65-73)**
```go
leftMax := divideConquerHelper(nums, left, mid)
rightMax := divideConquerHelper(nums, mid+1, right)
crossMax := maxCrossingSum(nums, left, mid, right)
```
- Đệ quy tìm max ở nửa trái
- Đệ quy tìm max ở nửa phải
- Tính max qua điểm giữa

**4. Hàm maxCrossingSum**

```78:106:internal/array/maximum_subarray/maximum_subarray.go
// maxCrossingSum tính tổng lớn nhất của subarray qua điểm giữa.
// Subarray này phải bao gồm nums[mid] và có thể mở rộng về cả hai phía.
func maxCrossingSum(nums []int, left, mid, right int) int {
	// Tìm tổng lớn nhất từ mid về bên trái
	leftSum := nums[mid]
	sum := 0
	for i := mid; i >= left; i-- {
		sum += nums[i]
		if sum > leftSum {
			leftSum = sum
		}
	}

	// Tìm tổng lớn nhất từ mid+1 về bên phải
	rightSum := nums[mid+1]
	sum = 0
	for i := mid + 1; i <= right; i++ {
		sum += nums[i]
		if sum > rightSum {
			rightSum = sum
		}
	}

	// Tổng qua điểm giữa = tổng bên trái + tổng bên phải
	return leftSum + rightSum
}
```

- **Tìm tổng lớn nhất từ mid về trái:**
  - Bắt đầu từ `mid`, duyệt về `left`
  - Tính tổng tích lũy và theo dõi tổng lớn nhất
- **Tìm tổng lớn nhất từ mid+1 về phải:**
  - Bắt đầu từ `mid+1`, duyệt về `right`
  - Tính tổng tích lũy và theo dõi tổng lớn nhất
- **Kết hợp:** `leftSum + rightSum`

### Hàm helper: max

```108:113:internal/array/maximum_subarray/maximum_subarray.go
// max trả về giá trị lớn hơn giữa a và b
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

## Phân tích độ phức tạp

### Kadane's Algorithm

**Thời gian: O(n)**
- Duyệt qua mảng một lần duy nhất
- Mỗi phần tử được xử lý trong thời gian O(1)
- Tổng: O(n)

**Không gian: O(1)**
- Chỉ sử dụng thêm hai biến: `maxSum`, `currentSum`
- Không sử dụng thêm cấu trúc dữ liệu nào
- Không gian phụ: O(1)

### Divide and Conquer

**Thời gian: O(n log n)**
- Mỗi lần chia đôi mảng: O(log n) levels trong cây đệ quy
- Mỗi level xử lý tổng cộng O(n) phần tử:
  - Tính crossing sum: O(n) cho mỗi lần gọi
  - Tổng tất cả các lần gọi ở mỗi level: O(n)
- Tổng: O(n) × O(log n) = O(n log n)

**Không gian: O(log n)**
- Call stack của đệ quy
- Độ sâu tối đa của cây đệ quy: O(log n)
- Mỗi lần gọi đệ quy chỉ sử dụng O(1) không gian
- Tổng: O(log n)

### So sánh

| Phương pháp | Thời gian | Không gian | Ưu điểm | Nhược điểm |
|------------|-----------|------------|---------|------------|
| Brute Force | O(n²) | O(1) | Dễ hiểu, dễ code | Quá chậm với n lớn |
| Kadane's Algorithm | O(n) | O(1) | Nhanh nhất, tối ưu nhất | Cần hiểu ý tưởng DP |
| Divide and Conquer | O(n log n) | O(log n) | Elegant, có thể song song hóa | Chậm hơn Kadane's |

**Kết luận:** Kadane's Algorithm là giải pháp tối ưu nhất cho bài toán này.

## Tại sao Kadane's Algorithm đúng?

### Chứng minh tính đúng đắn

**Định lý:** Kadane's Algorithm tìm được subarray có tổng lớn nhất.

**Chứng minh:**

1. **Quan sát 1:** Mọi subarray đều kết thúc tại một vị trí nào đó.
   - Subarray `[a, b, c]` kết thúc tại vị trí của `c`
   - Nếu ta tìm được tổng lớn nhất kết thúc tại mỗi vị trí, ta có thể tìm được tổng lớn nhất tổng thể

2. **Quan sát 2:** Tại vị trí `i`, subarray kết thúc tại `i` có hai khả năng:
   - **Khả năng 1:** Chỉ có `nums[i]` (bắt đầu subarray mới)
   - **Khả năng 2:** Mở rộng từ subarray kết thúc tại `i-1` (thêm `nums[i]` vào)

3. **Quyết định:** Chọn khả năng nào cho tổng lớn hơn
   - `currentSum[i] = max(nums[i], currentSum[i-1] + nums[i])`
   - Nếu `currentSum[i-1] < 0`, chọn `nums[i]` (bắt đầu mới)
   - Nếu `currentSum[i-1] >= 0`, chọn `currentSum[i-1] + nums[i]` (mở rộng)

4. **Tính đầy đủ:**
   - Ta xét tất cả các subarray kết thúc tại mỗi vị trí
   - Ta chọn giá trị lớn nhất trong tất cả các vị trí
   - Do đó, ta tìm được subarray có tổng lớn nhất

5. **Tính tối ưu:**
   - Tại mỗi vị trí, ta chọn quyết định tối ưu (max)
   - Quyết định tối ưu tại vị trí `i` phụ thuộc vào quyết định tối ưu tại vị trí `i-1`
   - Đây là tính chất **optimal substructure** của Dynamic Programming

**Kết luận:** Kadane's Algorithm đúng và tối ưu. ∎

## Follow-up: Divide and Conquer

Đề bài yêu cầu thử code giải pháp sử dụng **divide and conquer** nếu đã giải được bài toán với O(n).

### Tại sao Divide and Conquer?

1. **Học thuật:** Divide and Conquer là một kỹ thuật quan trọng trong thiết kế thuật toán
2. **Song song hóa:** Có thể song song hóa các phần đệ quy (parallel computing)
3. **Mở rộng:** Có thể áp dụng cho các bài toán phức tạp hơn
4. **Elegant:** Cách tiếp cận elegant và có cấu trúc rõ ràng

### So sánh với Kadane's Algorithm

**Kadane's Algorithm:**
- ✅ Nhanh hơn: O(n) vs O(n log n)
- ✅ Đơn giản hơn: Không cần đệ quy
- ✅ Ít memory hơn: O(1) vs O(log n)
- ❌ Khó song song hóa

**Divide and Conquer:**
- ✅ Có thể song song hóa: Tính left và right độc lập
- ✅ Elegant: Cấu trúc rõ ràng, dễ chứng minh đúng
- ✅ Mở rộng: Có thể áp dụng cho các bài toán phức tạp hơn
- ❌ Chậm hơn: O(n log n)
- ❌ Phức tạp hơn: Cần đệ quy và tính crossing sum

### Khi nào dùng Divide and Conquer?

- **Parallel computing:** Khi có nhiều CPU/cores và muốn song song hóa
- **Distributed systems:** Khi dữ liệu phân tán trên nhiều máy
- **Learning:** Khi muốn học và thực hành kỹ thuật Divide and Conquer
- **Interview:** Khi interviewer yêu cầu giải pháp Divide and Conquer

## Test Cases

### Test Case 1: Example 1
```go
Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
```
Subarray `[4,-1,2,1]` có tổng lớn nhất.

### Test Case 2: Example 2
```go
Input: nums = [1]
Output: 1
```
Mảng chỉ có một phần tử.

### Test Case 3: Example 3
```go
Input: nums = [5,4,-1,7,8]
Output: 23
```
Toàn bộ mảng có tổng lớn nhất.

### Test Case 4: All negative
```go
Input: nums = [-3,-2,-5,-1,-4]
Output: -1
```
Khi tất cả đều âm, chọn số âm nhỏ nhất (gần 0 nhất).

### Test Case 5: All positive
```go
Input: nums = [1,2,3,4,5]
Output: 15
```
Khi tất cả đều dương, tổng toàn bộ mảng là lớn nhất.

### Test Case 6: Zeros in array
```go
Input: nums = [-2,0,-1]
Output: 0
```
Số 0 có thể là kết quả khi tất cả số khác đều âm.

### Test Case 7: Complex case
```go
Input: nums = [8,-19,5,-4,20]
Output: 21
```
Subarray `[5,-4,20]` có tổng lớn nhất là 21.

## Chạy test

Để chạy các test case:

```bash
go test ./internal/array/maximum_subarray/
```

Hoặc chạy với verbose mode để xem chi tiết:

```bash
go test -v ./internal/array/maximum_subarray/
```

Chạy benchmark để so sánh hiệu suất:

```bash
go test -bench=. ./internal/array/maximum_subarray/
```

## Mở rộng

### Biến thể 1: Trả về chỉ số của subarray

Nếu cần trả về chỉ số bắt đầu và kết thúc của subarray:

```go
func maxSubArrayWithIndices(nums []int) (maxSum, start, end int) {
    maxSum = nums[0]
    currentSum := nums[0]
    start, end = 0, 0
    tempStart := 0
    
    for i := 1; i < len(nums); i++ {
        if nums[i] > currentSum + nums[i] {
            currentSum = nums[i]
            tempStart = i  // Bắt đầu subarray mới
        } else {
            currentSum = currentSum + nums[i]
        }
        
        if currentSum > maxSum {
            maxSum = currentSum
            start = tempStart
            end = i
        }
    }
    
    return maxSum, start, end
}
```

### Biến thể 2: Maximum Product Subarray

Tương tự nhưng với tích thay vì tổng (đã có trong folder `maximum_product_subarray`):

```go
func maxProduct(nums []int) int {
    maxProd := nums[0]
    minProd := nums[0]
    result := nums[0]
    
    for i := 1; i < len(nums); i++ {
        if nums[i] < 0 {
            maxProd, minProd = minProd, maxProd
        }
        maxProd = max(nums[i], maxProd * nums[i])
        minProd = min(nums[i], minProd * nums[i])
        result = max(result, maxProd)
    }
    
    return result
}
```

**Khác biệt:**
- Với tổng: Chỉ cần theo dõi một giá trị (max)
- Với tích: Cần theo dõi cả max và min (vì số âm)

### Biến thể 3: Maximum Subarray Sum với ràng buộc độ dài

Tìm tổng lớn nhất của subarray có độ dài từ `minLen` đến `maxLen`:

```go
func maxSubArrayWithLength(nums []int, minLen, maxLen int) int {
    // Sử dụng sliding window hoặc prefix sum
    // ...
}
```

### Biến thể 4: Circular Array

Mảng được xem như circular (phần tử cuối nối với phần tử đầu):

```go
func maxSubArrayCircular(nums []int) int {
    // Hai trường hợp:
    // 1. Subarray không wrap around: Kadane's Algorithm thông thường
    // 2. Subarray wrap around: Total sum - minimum subarray sum
    // Trả về max của hai trường hợp
    // ...
}
```

## Ứng dụng thực tế

Kadane's Algorithm được sử dụng trong:

1. **Phân tích tài chính:**
   - Tìm khoảng thời gian có lợi nhuận tích lũy lớn nhất
   - Phân tích xu hướng giá cổ phiếu

2. **Xử lý tín hiệu:**
   - Tìm đoạn tín hiệu có biên độ tích lũy lớn nhất
   - Phát hiện anomaly trong time series

3. **Machine Learning:**
   - Feature selection
   - Tối ưu hóa loss function

4. **Genomics:**
   - Tìm đoạn DNA có GC content cao nhất
   - Phân tích sequence alignment

5. **Image Processing:**
   - Tìm vùng ảnh có độ sáng tích lũy lớn nhất
   - Object detection

## Tips và Tricks

1. **Khởi tạo đúng:** Bắt đầu với `nums[0]`, không phải 0 hoặc -∞
2. **Xử lý mảng toàn số âm:** Thuật toán tự động chọn số âm nhỏ nhất (gần 0 nhất)
3. **Tối ưu code:** Có thể viết gọn hơn bằng cách inline các phép toán
4. **Debug:** In ra `currentSum` và `maxSum` tại mỗi bước để hiểu thuật toán
5. **Mở rộng:** Dễ dàng mở rộng để trả về chỉ số của subarray

## Lưu ý về Edge Cases

1. **Mảng chỉ có một phần tử:** Trả về phần tử đó
2. **Mảng toàn số âm:** Trả về số âm nhỏ nhất (gần 0 nhất)
3. **Mảng toàn số dương:** Trả về tổng toàn bộ mảng
4. **Có số 0:** Số 0 có thể là kết quả nếu tất cả số khác đều âm
5. **Overflow:** Với Go, `int` thường là 64-bit nên ít khi overflow

## Tối ưu hóa

### Viết gọn hơn

```go
func maxSubArrayCompact(nums []int) int {
    maxSum, currentSum := nums[0], nums[0]
    for i := 1; i < len(nums); i++ {
        currentSum = max(nums[i], currentSum+nums[i])
        maxSum = max(maxSum, currentSum)
    }
    return maxSum
}
```

### Sử dụng built-in max (Go 1.21+)

```go
import "cmp"

func maxSubArrayBuiltin(nums []int) int {
    maxSum, currentSum := nums[0], nums[0]
    for i := 1; i < len(nums); i++ {
        currentSum = max(nums[i], currentSum+nums[i])
        maxSum = max(maxSum, currentSum)
    }
    return maxSum
}
```

## Kết luận

**Maximum Subarray** là một bài toán kinh điển trong Dynamic Programming và được giải quyết tối ưu bởi **Kadane's Algorithm** với độ phức tạp O(n) thời gian và O(1) không gian.

**Key takeaways:**
- Kadane's Algorithm là giải pháp tối ưu nhất
- Ý tưởng chính: Tại mỗi vị trí, quyết định bắt đầu mới hay mở rộng subarray hiện tại
- Divide and Conquer là một cách tiếp cận thú vị nhưng chậm hơn
- Thuật toán có thể mở rộng cho nhiều biến thể khác

**Độ khó:** Medium - Cần hiểu ý tưởng DP nhưng code đơn giản

**Thời gian giải:** 15-30 phút (nếu đã biết Kadane's Algorithm)
