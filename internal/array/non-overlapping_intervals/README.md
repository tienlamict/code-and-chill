# Non-overlapping Intervals

## Mô tả bài toán

Cho một mảng các intervals `intervals` trong đó `intervals[i] = [starti, endi]`, trả về số lượng intervals tối thiểu cần xóa để các intervals còn lại không overlap (không chồng chéo).

**Lưu ý:** Các intervals chỉ chạm nhau tại một điểm được coi là không overlap. Ví dụ, `[1, 2]` và `[2, 3]` là không overlap.

**Ví dụ:**

### Example 1:
- Input: `intervals = [[1,2],[2,3],[3,4],[1,3]]`
- Output: `1`
- Giải thích: Có thể xóa `[1,3]` và các intervals còn lại sẽ không overlap.

```
Intervals:
[1,2]  |----|
[2,3]       |----|
[3,4]            |----|
[1,3]  |---------|  ← Xóa interval này

Sau khi xóa [1,3]:
[1,2]  |----|
[2,3]       |----|
[3,4]            |----|
→ Không còn overlap
```

### Example 2:
- Input: `intervals = [[1,2],[1,2],[1,2]]`
- Output: `2`
- Giải thích: Cần xóa hai intervals `[1,2]` để chỉ còn lại một interval không overlap.

```
Intervals:
[1,2]  |----|
[1,2]  |----|  ← Xóa
[1,2]  |----|  ← Xóa

Sau khi xóa 2 intervals:
[1,2]  |----|
→ Chỉ còn 1 interval
```

### Example 3:
- Input: `intervals = [[1,2],[2,3]]`
- Output: `0`
- Giải thích: Không cần xóa interval nào vì chúng đã không overlap.

```
Intervals:
[1,2]  |----|
[2,3]       |----|
→ Chỉ chạm nhau tại điểm 2, không overlap
```

**Ràng buộc:**
- `1 <= intervals.length <= 10^5`
- `intervals[i].length == 2`
- `-5 * 10^4 <= starti < endi <= 5 * 10^4`

## Phân tích thuật toán

### Cách tiếp cận 1: Brute Force với Dynamic Programming (O(2^n) time, O(n) space)

Thử tất cả các tổ hợp intervals có thể giữ lại và tìm tổ hợp có nhiều intervals nhất:

```go
func eraseOverlapIntervalsBruteForce(intervals [][]int) int {
    // Sắp xếp intervals theo start time
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    
    // DP: dp[i] = số intervals tối đa có thể giữ từ intervals[0..i]
    dp := make([]int, len(intervals))
    dp[0] = 1
    
    for i := 1; i < len(intervals); i++ {
        dp[i] = 1
        for j := 0; j < i; j++ {
            if intervals[j][1] <= intervals[i][0] {
                dp[i] = max(dp[i], dp[j] + 1)
            }
        }
    }
    
    maxKept := 0
    for _, count := range dp {
        maxKept = max(maxKept, count)
    }
    
    return len(intervals) - maxKept
}
```

**Độ phức tạp:**
- Thời gian: O(n²) - với n là số lượng intervals
- Không gian: O(n) - mảng DP

**Nhược điểm:** Vẫn chậm với n lên tới 10^5

### Cách tiếp cận 2: Greedy Algorithm với End Time Sorting (O(n log n) time, O(1) space) ⭐ Tối ưu

**Ý tưởng chính:**

Đây là bài toán **Interval Scheduling** kinh điển. Chiến lược tối ưu là:
- **Sắp xếp intervals theo end time** (kết thúc sớm nhất trước)
- **Luôn giữ lại interval có end time nhỏ nhất** khi có conflict
- Lý do: Interval kết thúc sớm hơn sẽ để lại nhiều không gian cho các intervals sau

**Thuật toán:**

1. **Sắp xếp intervals theo end time**: `intervals[i][1]`
2. **Giữ lại interval đầu tiên** (có end time nhỏ nhất)
3. **Duyệt qua các intervals còn lại**:
   - Nếu interval hiện tại overlap với interval đã giữ lại → xóa interval hiện tại
   - Nếu không overlap → giữ lại interval hiện tại và cập nhật end time

**Độ phức tạp:**
- Thời gian: O(n log n) - sắp xếp + O(n) duyệt = O(n log n)
- Không gian: O(1) - chỉ sử dụng thêm một vài biến

**Tại sao thuật toán này đúng?**

**Chứng minh bằng Greedy Choice Property:**

Giả sử có hai intervals A và B overlap, với `A.end < B.end`.

- **Nếu giữ A**: Có thể giữ thêm các intervals bắt đầu sau `A.end`
- **Nếu giữ B**: Chỉ có thể giữ các intervals bắt đầu sau `B.end` (≥ `A.end`)

Vì `A.end < B.end`, giữ A sẽ để lại nhiều không gian hơn → giữ A là lựa chọn tối ưu.

### So sánh các cách tiếp cận

| Cách tiếp cận | Thời gian | Không gian | Ưu điểm | Nhược điểm |
|--------------|-----------|------------|---------|------------|
| Brute Force | O(2^n) | O(n) | Đơn giản | Quá chậm |
| DP | O(n²) | O(n) | Đúng | Vẫn chậm với n lớn |
| Greedy (End Time) | O(n log n) | O(1) | Tối ưu | Cần hiểu greedy |

## Giải thích chi tiết thuật toán Greedy

### Ví dụ minh họa

Với input: `intervals = [[1,2],[2,3],[3,4],[1,3]]`

```
Bước 0: Sắp xếp theo end time
Ban đầu: [[1,2],[2,3],[3,4],[1,3]]
Sau sort: [[1,2],[2,3],[1,3],[3,4]]
          end: 2    end: 3    end: 3    end: 4

Bước 1: Giữ interval đầu tiên [1,2]
endTime = 2
removed = 0
Giữ lại: [1,2]

Bước 2: Xử lý [2,3]
intervals[1][0] = 2, endTime = 2
2 < 2? → KHÔNG (không overlap, chỉ chạm tại điểm 2)
→ Giữ lại [2,3]
endTime = 3
removed = 0
Giữ lại: [1,2], [2,3]

Bước 3: Xử lý [1,3]
intervals[2][0] = 1, endTime = 3
1 < 3? → CÓ (overlap!)
→ Xóa [1,3]
removed = 1
Giữ lại: [1,2], [2,3]

Bước 4: Xử lý [3,4]
intervals[3][0] = 3, endTime = 3
3 < 3? → KHÔNG (không overlap, chỉ chạm tại điểm 3)
→ Giữ lại [3,4]
endTime = 4
removed = 1
Giữ lại: [1,2], [2,3], [3,4]

Kết quả: removed = 1
```

Với input: `intervals = [[1,2],[1,2],[1,2]]`

```
Bước 0: Sắp xếp theo end time
Tất cả đều có end = 2, thứ tự không đổi
[[1,2],[1,2],[1,2]]

Bước 1: Giữ interval đầu tiên [1,2]
endTime = 2
removed = 0
Giữ lại: [1,2]

Bước 2: Xử lý [1,2] thứ hai
intervals[1][0] = 1, endTime = 2
1 < 2? → CÓ (overlap!)
→ Xóa [1,2] thứ hai
removed = 1
Giữ lại: [1,2] (đầu tiên)

Bước 3: Xử lý [1,2] thứ ba
intervals[2][0] = 1, endTime = 2
1 < 2? → CÓ (overlap!)
→ Xóa [1,2] thứ ba
removed = 2
Giữ lại: [1,2] (đầu tiên)

Kết quả: removed = 2
```

### Tại sao sắp xếp theo end time?

**Lý do 1: Greedy Choice Property**

Khi hai intervals overlap, interval có end time nhỏ hơn sẽ để lại nhiều không gian hơn cho các intervals sau.

**Ví dụ:**
```
Intervals: [1,3] và [2,4] overlap
- Giữ [1,3] (end=3): Có thể giữ intervals bắt đầu từ 3 trở đi
- Giữ [2,4] (end=4): Chỉ có thể giữ intervals bắt đầu từ 4 trở đi

→ Giữ [1,3] tốt hơn
```

**Lý do 2: Optimal Substructure**

Nếu ta đã chọn tối ưu cho các intervals trước đó, việc chọn interval có end time nhỏ nhất khi có conflict sẽ dẫn đến giải pháp tối ưu tổng thể.

## Giải thích code

### Hàm eraseOverlapIntervals

```9:50:internal/array/non-overlapping_intervals/non_overlapping_intervals.go
// eraseOverlapIntervals trả về số lượng intervals tối thiểu cần xóa để các intervals còn lại không overlap.
//
// Sử dụng Greedy Algorithm với Interval Scheduling:
// - Sắp xếp intervals theo end time (kết thúc sớm nhất trước)
// - Duyệt qua từng interval và giữ lại interval có end time nhỏ nhất khi có conflict
// - Khi hai intervals overlap, luôn xóa interval có end time lớn hơn
//   (vì giữ interval kết thúc sớm hơn sẽ để lại nhiều không gian cho các intervals sau)
//
// Lý do sắp xếp theo end time:
// - Interval kết thúc sớm hơn sẽ để lại nhiều không gian cho các intervals sau
// - Đây là chiến lược tối ưu để giữ lại nhiều intervals nhất có thể
//
// Độ phức tạp: O(n log n) thời gian, O(1) không gian
// với n là số lượng intervals
func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	// Sắp xếp intervals theo end time (intervals[i][1])
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	// Giữ lại interval đầu tiên (có end time nhỏ nhất)
	// endTime là end time của interval cuối cùng được giữ lại
	endTime := intervals[0][1]
	removed := 0

	// Duyệt qua các intervals còn lại
	for i := 1; i < len(intervals); i++ {
		// Nếu interval hiện tại overlap với interval đã giữ lại
		// (start của interval hiện tại < end của interval đã giữ lại)
		if intervals[i][0] < endTime {
			// Xóa interval hiện tại (vì nó có end time lớn hơn)
			removed++
		} else {
			// Giữ lại interval hiện tại và cập nhật endTime
			endTime = intervals[i][1]
		}
	}

	return removed
}
```

**Chi tiết từng phần:**

1. **Xử lý mảng rỗng (dòng 18-20):**
   ```go
   if len(intervals) == 0 {
       return 0
   }
   ```
   - Nếu không có interval nào, không cần xóa gì

2. **Sắp xếp theo end time (dòng 22-25):**
   ```go
   sort.Slice(intervals, func(i, j int) bool {
       return intervals[i][1] < intervals[j][1]
   })
   ```
   - Sắp xếp intervals theo `intervals[i][1]` (end time)
   - Intervals kết thúc sớm hơn sẽ đứng trước

3. **Khởi tạo (dòng 27-30):**
   ```go
   endTime := intervals[0][1]
   removed := 0
   ```
   - Giữ lại interval đầu tiên (có end time nhỏ nhất)
   - `endTime`: End time của interval cuối cùng được giữ lại
   - `removed`: Số lượng intervals đã xóa

4. **Duyệt và quyết định (dòng 32-42):**
   ```go
   for i := 1; i < len(intervals); i++ {
       if intervals[i][0] < endTime {
           removed++
       } else {
           endTime = intervals[i][1]
       }
   }
   ```
   - **Nếu overlap** (`intervals[i][0] < endTime`): Xóa interval hiện tại
   - **Nếu không overlap**: Giữ lại và cập nhật `endTime`

**Lưu ý quan trọng:**
- Điều kiện overlap: `intervals[i][0] < endTime` (không có `=`)
- Lý do: Intervals chỉ chạm nhau tại một điểm (ví dụ: `[1,2]` và `[2,3]`) được coi là không overlap

## Độ phức tạp

### Thời gian: O(n log n)

- **Sắp xếp intervals**: O(n log n)
- **Duyệt qua intervals**: O(n)
- **Tổng**: O(n log n)

### Không gian: O(1)

- Chỉ sử dụng thêm một vài biến (`endTime`, `removed`)
- Không sử dụng thêm cấu trúc dữ liệu nào
- Sắp xếp tại chỗ (in-place)

## Test Cases

### Test Case 1: Example 1
```go
Input: intervals = [[1,2],[2,3],[3,4],[1,3]]
Output: 1
```
Xóa `[1,3]` để các intervals còn lại không overlap.

### Test Case 2: Example 2
```go
Input: intervals = [[1,2],[1,2],[1,2]]
Output: 2
```
Xóa 2 intervals trùng lặp, chỉ giữ lại 1.

### Test Case 3: Example 3
```go
Input: intervals = [[1,2],[2,3]]
Output: 0
```
Các intervals đã không overlap, không cần xóa.

### Test Case 4: Hai intervals overlap
```go
Input: intervals = [[1,3],[2,4]]
Output: 1
```
Xóa một trong hai intervals overlap.

### Test Case 5: Intervals chạm nhau tại điểm
```go
Input: intervals = [[1,2],[2,3],[3,4]]
Output: 0
```
Các intervals chỉ chạm nhau tại điểm, không overlap.

### Test Case 6: Nested intervals
```go
Input: intervals = [[1,5],[2,3],[3,4]]
Output: 1
```
Interval `[1,5]` chứa các intervals khác, xóa nó.

### Test Case 7: Nhiều overlaps
```go
Input: intervals = [[1,2],[1,3],[1,4],[2,3]]
Output: 2
```
Xóa 2 intervals để chỉ giữ lại 1.

### Test Case 8: Số âm
```go
Input: intervals = [[-5,-2],[-3,0],[-4,-1]]
Output: 1
```
Thuật toán hoạt động với số âm.

## Chạy test

Để chạy các test case:

```bash
go test ./internal/array/non-overlapping_intervals/
```

Hoặc chạy với verbose mode để xem chi tiết:

```bash
go test -v ./internal/array/non-overlapping_intervals/
```

## Mở rộng

### Biến thể: Maximum Number of Non-Overlapping Intervals

Tìm số lượng intervals tối đa có thể giữ lại (thay vì số lượng cần xóa):

```go
func maxNonOverlappingIntervals(intervals [][]int) int {
    if len(intervals) == 0 {
        return 0
    }
    
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][1] < intervals[j][1]
    })
    
    count := 1
    endTime := intervals[0][1]
    
    for i := 1; i < len(intervals); i++ {
        if intervals[i][0] >= endTime {
            count++
            endTime = intervals[i][1]
        }
    }
    
    return count
}
```

### Biến thể: Merge Intervals

Thay vì xóa, merge các intervals overlap:

```go
func mergeIntervals(intervals [][]int) [][]int {
    // Sắp xếp và merge
    // ...
}
```

### Biến thể: Insert Interval

Chèn một interval mới vào danh sách intervals đã sắp xếp:

```go
func insertInterval(intervals [][]int, newInterval []int) [][]int {
    // Chèn và merge nếu cần
    // ...
}
```

### Ứng dụng thực tế

Thuật toán này được sử dụng trong:

- **Scheduling**: Lên lịch các công việc không overlap
- **Resource allocation**: Phân bổ tài nguyên không conflict
- **Calendar applications**: Tìm các sự kiện không overlap
- **Network routing**: Tìm các đường dẫn không conflict
- **Database transactions**: Quản lý các transaction không overlap

### Tips và Tricks

1. **Luôn sắp xếp theo end time**: Đây là chìa khóa của thuật toán
2. **Điều kiện overlap**: `start < end` (không có `=`) vì chạm tại điểm không tính là overlap
3. **Greedy choice**: Khi có conflict, luôn giữ interval có end time nhỏ hơn
4. **Edge cases**: Mảng rỗng, một interval, tất cả overlap, không overlap
5. **Số âm**: Thuật toán hoạt động với cả số âm

### So sánh với các cách tiếp cận khác

| Đặc điểm | DP | Greedy |
|----------|----|--------|
| Độ phức tạp thời gian | O(n²) | O(n log n) |
| Độ phức tạp không gian | O(n) | O(1) |
| Dễ hiểu | Trung bình | Dễ |
| Hiệu suất | Tốt | Tốt hơn |

### Lưu ý về điều kiện overlap

**Định nghĩa overlap:**
- Hai intervals `[a, b]` và `[c, d]` overlap nếu: `a < d` và `c < b`
- Hoặc đơn giản hơn: `max(a, c) < min(b, d)`

**Trong code:**
- Ta kiểm tra: `intervals[i][0] < endTime`
- Điều này tương đương với: Start của interval mới < End của interval đã giữ lại
- Nếu điều kiện đúng → overlap → xóa interval mới

**Ví dụ:**
- `[1,2]` và `[2,3]`: `2 < 2`? → KHÔNG → Không overlap ✓
- `[1,3]` và `[2,4]`: `2 < 3`? → CÓ → Overlap ✓

### Tối ưu hóa

1. **In-place sorting**: Sắp xếp tại chỗ để tiết kiệm không gian
2. **Early termination**: Nếu đã xóa quá nhiều, có thể dừng sớm (nhưng không áp dụng được ở đây vì cần duyệt hết)
3. **Custom sort**: Có thể tối ưu sort nếu biết phạm vi giá trị (nhưng với n lớn, quicksort vẫn tốt)
