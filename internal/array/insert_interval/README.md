# Insert Interval

## Mô tả bài toán

Bạn được cho một mảng các intervals không chồng chéo `intervals` trong đó `intervals[i] = [starti, endi]` đại diện cho điểm đầu và điểm cuối của interval thứ i, và `intervals` được sắp xếp theo thứ tự tăng dần theo `starti`. Bạn cũng được cho một interval `newInterval = [start, end]` đại diện cho điểm đầu và điểm cuối của một interval khác.

Chèn `newInterval` vào `intervals` sao cho `intervals` vẫn được sắp xếp theo thứ tự tăng dần theo `starti` và `intervals` vẫn không có các intervals chồng chéo (gộp các intervals chồng chéo nếu cần thiết).

Trả về `intervals` sau khi chèn.

Lưu ý rằng bạn không cần sửa đổi `intervals` tại chỗ. Bạn có thể tạo một mảng mới và trả về nó.

**Ví dụ:**

### Example 1:
- Input: `intervals = [[1,3],[6,9]]`, `newInterval = [2,5]`
- Output: `[[1,5],[6,9]]`
- Giải thích: Interval [2,5] chồng chéo với [1,3], gộp thành [1,5].

### Example 2:
- Input: `intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]]`, `newInterval = [4,8]`
- Output: `[[1,2],[3,10],[12,16]]`
- Giải thích: Vì interval mới [4,8] chồng chéo với [3,5], [6,7], và [8,10], gộp tất cả thành [3,10].

**Ràng buộc:**
- `0 <= intervals.length <= 10^4`
- `intervals[i].length == 2`
- `0 <= starti <= endi <= 10^5`
- `intervals` được sắp xếp theo `starti` theo thứ tự tăng dần
- `newInterval.length == 2`
- `0 <= start <= end <= 10^5`

## Phân tích thuật toán

### Cách tiếp cận 1: Sử dụng Merge Intervals

**Ý tưởng:**
- Thêm `newInterval` vào mảng `intervals`
- Sắp xếp lại (nếu cần)
- Gọi hàm `merge()` từ bài Merge Intervals để gộp các intervals chồng chéo

**Độ phức tạp:**
- Thời gian: O(n log n) - do sắp xếp
- Không gian: O(n)

**Nhược điểm:** Không tận dụng được việc `intervals` đã được sắp xếp sẵn.

### Cách tiếp cận 2: Linear Scan (Tối ưu)

**Ý tưởng:**
- Vì `intervals` đã được sắp xếp, ta có thể duyệt một lần:
  1. Thêm tất cả intervals đứng trước `newInterval` (có `end < newInterval.start`)
  2. Gộp tất cả intervals chồng chéo với `newInterval`:
     - Interval chồng chéo khi: `interval.start <= newInterval.end`
     - Khi gộp: `start = min(tất cả starts)`, `end = max(tất cả ends)`
  3. Thêm tất cả intervals đứng sau interval đã gộp (có `start > mergedInterval.end`)

**Tại sao Linear Scan hoạt động?**
- Vì `intervals` đã được sắp xếp, các intervals chồng chéo với `newInterval` sẽ nằm liền kề nhau
- Chỉ cần một lần duyệt để:
  - Xác định các intervals đứng trước (không chồng chéo)
  - Gộp các intervals chồng chéo
  - Xác định các intervals đứng sau (không chồng chéo)

**Độ phức tạp:**
- Thời gian: O(n) - duyệt một lần
- Không gian: O(n) - cho mảng kết quả

**Ưu điểm:**
- Tận dụng được việc `intervals` đã được sắp xếp
- Không cần sắp xếp lại
- Hiệu quả hơn cách tiếp cận 1

## Giải thích chi tiết thuật toán

### Ví dụ 1: `intervals = [[1,3],[6,9]]`, `newInterval = [2,5]`

**Bước 1:** Thêm intervals đứng trước
- Xét [1, 3]: end (3) < newStart (2)? Không → dừng
- Không có interval nào đứng trước
- result = []

**Bước 2:** Gộp intervals chồng chéo
- Xét [1, 3]: start (1) <= newEnd (5)? Có → chồng chéo
  - newStart = min(1, 2) = 1
  - newEnd = max(3, 5) = 5
- Xét [6, 9]: start (6) <= newEnd (5)? Không → dừng
- Thêm [1, 5] vào result
- result = [[1, 5]]

**Bước 3:** Thêm intervals đứng sau
- Xét [6, 9]: start (6) > mergedEnd (5)? Có → thêm vào
- result = [[1, 5], [6, 9]]

**Kết quả:** `[[1,5],[6,9]]`

### Ví dụ 2: `intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]]`, `newInterval = [4,8]`

**Bước 1:** Thêm intervals đứng trước
- Xét [1, 2]: end (2) < newStart (4)? Có → thêm vào
- Xét [3, 5]: end (5) < newStart (4)? Không → dừng
- result = [[1, 2]]

**Bước 2:** Gộp intervals chồng chéo
- Xét [3, 5]: start (3) <= newEnd (8)? Có → chồng chéo
  - newStart = min(3, 4) = 3
  - newEnd = max(5, 8) = 8
- Xét [6, 7]: start (6) <= newEnd (8)? Có → chồng chéo
  - newStart = min(3, 6) = 3 (không đổi)
  - newEnd = max(8, 7) = 8 (không đổi)
- Xét [8, 10]: start (8) <= newEnd (8)? Có → chồng chéo
  - newStart = min(3, 8) = 3 (không đổi)
  - newEnd = max(8, 10) = 10
- Xét [12, 16]: start (12) <= newEnd (10)? Không → dừng
- Thêm [3, 10] vào result
- result = [[1, 2], [3, 10]]

**Bước 3:** Thêm intervals đứng sau
- Xét [12, 16]: start (12) > mergedEnd (10)? Có → thêm vào
- result = [[1, 2], [3, 10], [12, 16]]

**Kết quả:** `[[1,2],[3,10],[12,16]]`

## Giải thích code

### Hàm `insert`

```go
result := [][]int{}
i := 0
n := len(intervals)
newStart := newInterval[0]
newEnd := newInterval[1]
```

Khởi tạo mảng kết quả, biến đếm `i`, và lưu start/end của `newInterval`.

```go
for i < n && intervals[i][1] < newStart {
    result = append(result, intervals[i])
    i++
}
```

**Bước 1:** Thêm tất cả intervals đứng trước `newInterval`. Interval đứng trước khi `end < newInterval.start` (không chồng chéo và đứng trước).

```go
for i < n && intervals[i][0] <= newEnd {
    if intervals[i][0] < newStart {
        newStart = intervals[i][0]
    }
    if intervals[i][1] > newEnd {
        newEnd = intervals[i][1]
    }
    i++
}
```

**Bước 2:** Gộp tất cả intervals chồng chéo với `newInterval`. Interval chồng chéo khi `start <= newInterval.end`. Khi gộp:
- Cập nhật `newStart` = min của tất cả starts
- Cập nhật `newEnd` = max của tất cả ends

```go
result = append(result, []int{newStart, newEnd})
```

Thêm interval đã được gộp vào kết quả.

```go
for i < n {
    result = append(result, intervals[i])
    i++
}
```

**Bước 3:** Thêm tất cả intervals còn lại (đứng sau interval đã gộp).

## Phân tích độ phức tạp

### Độ phức tạp thời gian: O(n)

- Duyệt qua mảng một lần: O(n)
- Mỗi interval được xét đúng một lần
- Tổng: O(n)

### Độ phức tạp không gian: O(n)

- Mảng kết quả có thể chứa tối đa n+1 intervals: O(n)
- Không sử dụng thêm cấu trúc dữ liệu phụ
- Tổng: O(n)

## So sánh với các cách tiếp cận khác

| Cách tiếp cận | Thời gian | Không gian | Ghi chú |
|--------------|-----------|------------|---------|
| Merge Intervals | O(n log n) | O(n) | Không tận dụng được việc đã sắp xếp |
| Binary Search + Merge | O(n) | O(n) | Phức tạp hơn, không cần thiết |
| **Linear Scan** | **O(n)** | **O(n)** | **Tối ưu, đơn giản** |

## Follow-up Questions

### 1. Có thể tối ưu hơn O(n) không?

Trong trường hợp tổng quát, không. Phải duyệt qua ít nhất một lần để xác định các intervals chồng chéo.

Nếu cần xử lý nhiều insertions động, có thể dùng cấu trúc dữ liệu như Interval Tree để insert trong O(log n), nhưng phức tạp hơn nhiều.

### 2. Nếu intervals không được sắp xếp sẵn?

Cần sắp xếp trước: O(n log n), sau đó dùng thuật toán trên: O(n). Tổng: O(n log n).

Hoặc có thể dùng cách tiếp cận Merge Intervals: thêm `newInterval` vào mảng, sắp xếp, rồi merge.

### 3. Nếu cần insert nhiều intervals cùng lúc?

Có thể:
- Thêm tất cả intervals mới vào mảng
- Sắp xếp lại
- Gọi `merge()` - O(n log n)

Hoặc insert từng interval một - O(k * n) với k là số intervals mới.

### 4. Mối quan hệ với Merge Intervals?

Insert Interval là trường hợp đặc biệt của Merge Intervals:
- Input đã được sắp xếp
- Chỉ cần insert một interval
- Có thể tối ưu thành O(n) thay vì O(n log n)
