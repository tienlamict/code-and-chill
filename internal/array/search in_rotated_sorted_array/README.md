# Search in Rotated Sorted Array

## Mô tả bài toán

Có một mảng số nguyên `nums` được sắp xếp theo thứ tự tăng dần (với các giá trị phân biệt).

Trước khi được truyền vào hàm, `nums` có thể bị rotate sang trái tại một chỉ số `k` không xác định (`1 <= k < nums.length`) sao cho mảng kết quả là `[nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]` (0-indexed). 

Ví dụ, `[0,1,2,4,5,6,7]` có thể bị rotate sang trái 3 vị trí và trở thành `[4,5,6,7,0,1,2]`.

Cho mảng `nums` sau khi có thể bị rotate và một số nguyên `target`, trả về chỉ số của `target` nếu nó có trong `nums`, hoặc `-1` nếu không có.

Bạn phải viết một thuật toán có độ phức tạp thời gian O(log n).

**Ví dụ:**

### Example 1:
- Input: `nums = [4,5,6,7,0,1,2]`, `target = 0`
- Output: `4`
- Giải thích: `target = 0` nằm ở vị trí index 4 trong mảng đã rotate.

```
Mảng gốc: [0,1,2,4,5,6,7]
Rotate tại k=3: [4,5,6,7,0,1,2]
                    ↑
                  index 4
```

### Example 2:
- Input: `nums = [4,5,6,7,0,1,2]`, `target = 3`
- Output: `-1`
- Giải thích: `target = 3` không có trong mảng.

### Example 3:
- Input: `nums = [1]`, `target = 0`
- Output: `-1`
- Giải thích: `target = 0` không có trong mảng chỉ có một phần tử.

**Ràng buộc:**
- `1 <= nums.length <= 5000`
- `-10^4 <= nums[i] <= 10^4`
- Tất cả các giá trị trong `nums` là duy nhất
- `nums` là một mảng tăng dần có thể bị rotate
- `-10^4 <= target <= 10^4`

## Phân tích thuật toán

### Cách tiếp cận 1: Linear Search (O(n) time, O(1) space)

Duyệt qua mảng từ đầu đến cuối để tìm `target`.

```go
func searchLinear(nums []int, target int) int {
    for i := 0; i < len(nums); i++ {
        if nums[i] == target {
            return i
        }
    }
    return -1
}
```

**Độ phức tạp:**
- Thời gian: O(n)
- Không gian: O(1)

**Nhược điểm:** Không đáp ứng yêu cầu O(log n)

### Cách tiếp cận 2: Binary Search với điều chỉnh (O(log n) time, O(1) space) ⭐ Tối ưu

**Ý tưởng chính:**

Mặc dù mảng đã bị rotate, tại mỗi bước của binary search, **một nửa của mảng luôn được sắp xếp**!

**Thuật toán:**

1. **Xác định nửa nào được sắp xếp:**
   - So sánh `nums[mid]` với `nums[left]`
   - Nếu `nums[left] <= nums[mid]`: Nửa trái `[left, mid]` được sắp xếp
   - Ngược lại: Nửa phải `[mid, right]` được sắp xếp

2. **Kiểm tra target có nằm trong nửa đã sắp xếp không:**
   - Nếu nửa trái được sắp xếp:
     - Kiểm tra: `nums[left] <= target < nums[mid]`
     - Nếu đúng → tìm ở nửa trái
     - Nếu sai → tìm ở nửa phải
   - Nếu nửa phải được sắp xếp:
     - Kiểm tra: `nums[mid] < target <= nums[right]`
     - Nếu đúng → tìm ở nửa phải
     - Nếu sai → tìm ở nửa trái

3. **Lặp lại cho đến khi tìm thấy hoặc không còn phần tử nào**

**Độ phức tạp:**
- Thời gian: O(log n) - binary search
- Không gian: O(1) - chỉ sử dụng thêm một vài biến

### Ví dụ minh họa

Với input: `nums = [4,5,6,7,0,1,2]`, `target = 0`

```
Bước 0: Khởi tạo
left = 0, right = 6
nums = [4, 5, 6, 7, 0, 1, 2]
        ↑              ↑
      left          right

Bước 1: mid = 3, nums[3] = 7
nums[left] = 4 <= nums[mid] = 7 → Nửa trái [0,3] được sắp xếp
nums[left] = 4 <= target = 0 < nums[mid] = 7? → KHÔNG
→ Tìm ở nửa phải
left = 4, right = 6

Bước 2: mid = 5, nums[5] = 1
nums[left] = 0 <= nums[mid] = 1 → Nửa trái [4,5] được sắp xếp
nums[left] = 0 <= target = 0 < nums[mid] = 1? → CÓ
→ Tìm ở nửa trái
left = 4, right = 4

Bước 3: mid = 4, nums[4] = 0
nums[4] == target → Tìm thấy!
Kết quả: 4
```

Với input: `nums = [4,5,6,7,0,1,2]`, `target = 3`

```
Bước 0: Khởi tạo
left = 0, right = 6

Bước 1: mid = 3, nums[3] = 7
Nửa trái được sắp xếp
nums[left] = 4 <= target = 3 < nums[mid] = 7? → KHÔNG
→ Tìm ở nửa phải
left = 4, right = 6

Bước 2: mid = 5, nums[5] = 1
Nửa trái được sắp xếp
nums[left] = 0 <= target = 3 < nums[mid] = 1? → KHÔNG
→ Tìm ở nửa phải
left = 6, right = 6

Bước 3: mid = 6, nums[6] = 2
nums[6] != target
left = 7 > right → Không tìm thấy
Kết quả: -1
```

## Giải thích code

### Hàm search

```9:50:internal/array/search in_rotated_sorted_array/search_in_rotated_sorted_array.go
// search tìm vị trí của target trong mảng đã được rotate.
//
// Sử dụng Binary Search với điều chỉnh:
// - Mặc dù mảng đã bị rotate, một nửa của mảng luôn được sắp xếp
// - Xác định nửa nào được sắp xếp bằng cách so sánh nums[mid] với nums[left]
// - Nếu nửa trái được sắp xếp:
//   - Kiểm tra xem target có nằm trong khoảng [left, mid] không
//   - Nếu có, tìm ở nửa trái; ngược lại, tìm ở nửa phải
// - Nếu nửa phải được sắp xếp:
//   - Kiểm tra xem target có nằm trong khoảng [mid, right] không
//   - Nếu có, tìm ở nửa phải; ngược lại, tìm ở nửa trái
//
// Độ phức tạp: O(log n) thời gian, O(1) không gian
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2

		// Tìm thấy target
		if nums[mid] == target {
			return mid
		}

		// Xác định nửa nào được sắp xếp
		if nums[left] <= nums[mid] {
			// Nửa trái [left, mid] được sắp xếp
			if nums[left] <= target && target < nums[mid] {
				// Target nằm trong nửa trái đã sắp xếp
				right = mid - 1
			} else {
				// Target nằm trong nửa phải (có thể bị rotate)
				left = mid + 1
			}
		} else {
			// Nửa phải [mid, right] được sắp xếp
			if nums[mid] < target && target <= nums[right] {
				// Target nằm trong nửa phải đã sắp xếp
				left = mid + 1
			} else {
				// Target nằm trong nửa trái (có thể bị rotate)
				right = mid - 1
			}
		}
	}

	// Không tìm thấy target
	return -1
}
```

### Chi tiết từng phần

#### 1. Khởi tạo (dòng 20-23)
```go
if len(nums) == 0 {
    return -1
}

left := 0
right := len(nums) - 1
```
- Xử lý trường hợp mảng rỗng
- Khởi tạo hai con trỏ cho binary search

#### 2. Vòng lặp binary search (dòng 25-48)
```go
for left <= right {
    mid := left + (right-left)/2
    
    if nums[mid] == target {
        return mid
    }
    // ...
}
```
- Tiếp tục tìm kiếm khi còn phần tử
- Tính `mid` an toàn để tránh overflow: `mid = left + (right-left)/2`
- Kiểm tra ngay nếu tìm thấy target

#### 3. Xác định nửa được sắp xếp (dòng 33-47)
```go
if nums[left] <= nums[mid] {
    // Nửa trái được sắp xếp
    if nums[left] <= target && target < nums[mid] {
        right = mid - 1
    } else {
        left = mid + 1
    }
} else {
    // Nửa phải được sắp xếp
    if nums[mid] < target && target <= nums[right] {
        left = mid + 1
    } else {
        right = mid - 1
    }
}
```

**Logic:**
- **Nếu `nums[left] <= nums[mid]`**: Nửa trái `[left, mid]` được sắp xếp
  - Nếu `target` nằm trong `[nums[left], nums[mid])` → tìm ở nửa trái
  - Ngược lại → tìm ở nửa phải
  
- **Nếu `nums[left] > nums[mid]`**: Nửa phải `[mid, right]` được sắp xếp
  - Nếu `target` nằm trong `(nums[mid], nums[right]]` → tìm ở nửa phải
  - Ngược lại → tìm ở nửa trái

**Lưu ý quan trọng:**
- Sử dụng `<=` và `<` một cách cẩn thận để tránh bỏ sót phần tử
- `target < nums[mid]` (không có `=`) vì đã kiểm tra `nums[mid] == target` ở trên
- `nums[mid] < target` (không có `=`) vì đã kiểm tra `nums[mid] == target` ở trên

### Độ phức tạp

- **Thời gian**: O(log n)
  - Mỗi lần lặp giảm một nửa không gian tìm kiếm
  - Tối đa log₂(n) lần lặp

- **Không gian**: O(1)
  - Chỉ sử dụng thêm một vài biến (`left`, `right`, `mid`)
  - Không sử dụng đệ quy hoặc cấu trúc dữ liệu phụ

### Tại sao thuật toán này đúng?

**Chứng minh:**

1. **Tính đúng đắn của việc xác định nửa được sắp xếp:**
   - Nếu `nums[left] <= nums[mid]`: 
     - Vì mảng gốc được sắp xếp và chỉ bị rotate một lần
     - Nếu phần tử đầu <= phần tử giữa, thì toàn bộ nửa trái phải được sắp xếp
   - Ngược lại, nửa phải được sắp xếp

2. **Tính đúng đắn của việc quyết định tìm kiếm:**
   - Nếu target nằm trong nửa đã sắp xếp, ta có thể áp dụng binary search thông thường
   - Nếu target không nằm trong nửa đã sắp xếp, nó phải nằm trong nửa còn lại (có thể bị rotate)

3. **Tính đầy đủ:**
   - Mỗi lần lặp loại bỏ ít nhất một nửa không gian tìm kiếm
   - Đảm bảo tìm thấy target nếu nó tồn tại, hoặc xác định không tồn tại

## Test Cases

### Test Case 1: Example 1
```go
Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4
```
Target nằm trong phần bị rotate.

### Test Case 2: Example 2
```go
Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1
```
Target không tồn tại trong mảng.

### Test Case 3: Example 3
```go
Input: nums = [1], target = 0
Output: -1
```
Mảng chỉ có một phần tử và target không khớp.

### Test Case 4: Target trong nửa trái đã sắp xếp
```go
Input: nums = [4,5,6,7,0,1,2], target = 5
Output: 1
```
Target nằm trong phần đã sắp xếp ở đầu mảng.

### Test Case 5: Target trong nửa phải đã sắp xếp
```go
Input: nums = [4,5,6,7,0,1,2], target = 1
Output: 5
```
Target nằm trong phần đã sắp xếp ở cuối mảng.

### Test Case 6: Mảng không bị rotate
```go
Input: nums = [1,2,3,4,5], target = 3
Output: 2
```
Mảng không bị rotate, hoạt động như binary search thông thường.

### Test Case 7: Rotate tại vị trí đầu
```go
Input: nums = [5,1,2,3,4], target = 1
Output: 1
```
Rotate tại vị trí 1 (chỉ có phần tử đầu bị đưa ra sau).

### Test Case 8: Rotate tại vị trí cuối
```go
Input: nums = [2,3,4,5,1], target = 1
Output: 4
```
Rotate tại vị trí cuối (chỉ có phần tử cuối bị đưa ra trước).

### Test Case 9: Hai phần tử
```go
Input: nums = [1,3], target = 3
Output: 1
```
Mảng chỉ có hai phần tử.

### Test Case 10: Target nhỏ hơn tất cả
```go
Input: nums = [4,5,6,7,0,1,2], target = -1
Output: -1
```
Target nhỏ hơn tất cả các phần tử trong mảng.

### Test Case 11: Target lớn hơn tất cả
```go
Input: nums = [4,5,6,7,0,1,2], target = 10
Output: -1
```
Target lớn hơn tất cả các phần tử trong mảng.

## Chạy test

Để chạy các test case:

```bash
go test ./internal/array/search\ in_rotated_sorted_array/
```

Hoặc chạy với verbose mode để xem chi tiết:

```bash
go test -v ./internal/array/search\ in_rotated_sorted_array/
```

## Mở rộng

### Biến thể: Search in Rotated Sorted Array II (có duplicate)

Nếu mảng có thể chứa các phần tử trùng lặp, thuật toán phức tạp hơn:

```go
func searchWithDuplicates(nums []int, target int) bool {
    left, right := 0, len(nums)-1
    
    for left <= right {
        mid := left + (right-left)/2
        
        if nums[mid] == target {
            return true
        }
        
        // Xử lý trường hợp nums[left] == nums[mid] == nums[right]
        if nums[left] == nums[mid] && nums[mid] == nums[right] {
            left++
            right--
            continue
        }
        
        // Logic tương tự nhưng cần xử lý duplicate
        // ...
    }
    
    return false
}
```

**Khác biệt:** Khi `nums[left] == nums[mid] == nums[right]`, không thể xác định nửa nào được sắp xếp, cần thu hẹp cả hai bên.

### Biến thể: Find Minimum in Rotated Sorted Array

Tìm phần tử nhỏ nhất trong mảng đã rotate:

```go
func findMin(nums []int) int {
    left, right := 0, len(nums)-1
    
    for left < right {
        mid := left + (right-left)/2
        
        if nums[mid] > nums[right] {
            // Minimum nằm ở nửa phải
            left = mid + 1
        } else {
            // Minimum nằm ở nửa trái (bao gồm mid)
            right = mid
        }
    }
    
    return nums[left]
}
```

### Ứng dụng thực tế

Thuật toán này được sử dụng trong:

- **Database indexing**: Tìm kiếm trong các index đã được rotate hoặc reorganized
- **Circular buffer**: Tìm kiếm trong buffer vòng tròn
- **Time-series data**: Tìm kiếm trong dữ liệu thời gian đã được rotate theo chu kỳ
- **Game development**: Tìm kiếm trong các mảng đã được shuffle hoặc rotate

### So sánh với các cách tiếp cận khác

| Cách tiếp cận | Thời gian | Không gian | Ưu điểm | Nhược điểm |
|--------------|-----------|------------|---------|------------|
| Linear Search | O(n) | O(1) | Đơn giản | Không đáp ứng yêu cầu |
| Binary Search | O(log n) | O(1) | Tối ưu | Logic phức tạp hơn |
| Find pivot rồi search | O(log n) | O(1) | Dễ hiểu | Cần 2 lần binary search |

### Tips và Tricks

1. **Luôn kiểm tra nửa nào được sắp xếp**: Đây là chìa khóa của thuật toán
2. **Sử dụng `<=` và `<` cẩn thận**: Tránh bỏ sót phần tử hoặc kiểm tra trùng lặp
3. **Tính `mid` an toàn**: Sử dụng `mid = left + (right-left)/2` thay vì `(left+right)/2` để tránh overflow
4. **Edge cases**: Mảng không rotate, rotate tại đầu/cuối, một/two phần tử
5. **Không có duplicate**: Đảm bảo logic đơn giản hơn (nếu có duplicate cần xử lý thêm)

### Lưu ý về điều kiện so sánh

**Tại sao `nums[left] <= nums[mid]` thay vì `nums[left] < nums[mid]`?**

- Khi `left == mid` (mảng có 1-2 phần tử), `nums[left] == nums[mid]`
- Trong trường hợp này, nửa trái được coi là "sắp xếp" (vì chỉ có 1 phần tử)
- Sử dụng `<=` để bao gồm trường hợp này

**Tại sao `target < nums[mid]` thay vì `target <= nums[mid]`?**

- Đã kiểm tra `nums[mid] == target` ở trên
- Nếu `target == nums[mid]`, đã return rồi
- Chỉ cần kiểm tra `target < nums[mid]` để quyết định tìm ở nửa trái

### Tối ưu hóa

Có thể viết lại code để rõ ràng hơn bằng cách tách logic:

```go
func searchOptimized(nums []int, target int) int {
    left, right := 0, len(nums)-1
    
    for left <= right {
        mid := left + (right-left)/2
        
        if nums[mid] == target {
            return mid
        }
        
        // Xác định nửa được sắp xếp
        leftSorted := nums[left] <= nums[mid]
        
        if leftSorted {
            // Nửa trái được sắp xếp
            if nums[left] <= target && target < nums[mid] {
                right = mid - 1
            } else {
                left = mid + 1
            }
        } else {
            // Nửa phải được sắp xếp
            if nums[mid] < target && target <= nums[right] {
                left = mid + 1
            } else {
                right = mid - 1
            }
        }
    }
    
    return -1
}
```

Cách này dễ đọc hơn nhưng logic tương tự.
