# Merge Intervals

## Mô tả bài toán

Cho một mảng các intervals trong đó `intervals[i] = [starti, endi]`, gộp tất cả các intervals chồng chéo và trả về mảng các intervals không chồng chéo bao phủ toàn bộ các intervals trong input.

**Ví dụ:**

### Example 1:
- Input: `intervals = [[1,3],[2,6],[8,10],[15,18]]`
- Output: `[[1,6],[8,10],[15,18]]`
- Giải thích: Vì intervals [1,3] và [2,6] chồng chéo, gộp chúng thành [1,6].

### Example 2:
- Input: `intervals = [[1,4],[4,5]]`
- Output: `[[1,5]]`
- Giải thích: Intervals [1,4] và [4,5] được coi là chồng chéo (chạm nhau tại điểm 4).

### Example 3:
- Input: `intervals = [[4,7],[1,4]]`
- Output: `[[1,7]]`
- Giải thích: Intervals [1,4] và [4,7] được coi là chồng chéo.

**Ràng buộc:**
- `1 <= intervals.length <= 10^4`
- `intervals[i].length == 2`
- `0 <= starti <= endi <= 10^4`

## Phân tích thuật toán

### Cách tiếp cận 1: Brute Force

**Ý tưởng:**
- Với mỗi interval, kiểm tra xem nó có chồng chéo với bất kỳ interval nào khác không
- Nếu có, gộp chúng lại và lặp lại cho đến khi không còn chồng chéo
- Lặp lại quá trình cho đến khi không còn thay đổi

**Độ phức tạp:**
- Thời gian: O(n²) hoặc cao hơn - cần nhiều lần duyệt
- Không gian: O(n)

**Nhược điểm:** Không hiệu quả, khó implement đúng.

### Cách tiếp cận 2: Sort + Greedy (Tối ưu)

**Ý tưởng:**
- Sắp xếp intervals theo start time (đầu mút trái)
- Sau khi sắp xếp, các intervals chồng chéo sẽ nằm liền kề nhau
- Duyệt qua từng interval: nếu chồng chéo với interval cuối trong kết quả thì gộp, không thì thêm mới

**Định nghĩa chồng chéo:**
- Hai intervals [a, b] và [c, d] chồng chéo khi: `c <= b` (khi đã sắp xếp theo start, ta có a <= c)
- Khi gộp: interval mới = [a, max(b, d)]

**Tại sao sắp xếp theo start hoạt động?**
- Sau khi sắp xếp, interval có start nhỏ hơn luôn đứng trước
- Khi xét interval hiện tại, chỉ cần so sánh với interval cuối trong kết quả
- Nếu current.start <= last.end → chồng chéo, gộp
- Nếu current.start > last.end → không chồng chéo, thêm mới

**Độ phức tạp:**
- Thời gian: O(n log n) - chủ yếu do sắp xếp
- Không gian: O(n) - cho mảng kết quả (hoặc O(log n) cho call stack của sort nếu sort in-place)

## Giải thích chi tiết thuật toán

### Ví dụ 1: `intervals = [[1,3],[2,6],[8,10],[15,18]]`

**Bước 1:** Sắp xếp (đã sắp xếp sẵn)
```
[1,3], [2,6], [8,10], [15,18]
```

**Bước 2:** Khởi tạo result = [[1, 3]]

**Bước 3:** Xét [2, 6]
- last = [1, 3], currentStart = 2, currentEnd = 6
- 2 <= 3 → chồng chéo
- max(3, 6) = 6 → cập nhật last = [1, 6]
- result = [[1, 6]]

**Bước 4:** Xét [8, 10]
- last = [1, 6], currentStart = 8
- 8 > 6 → không chồng chéo
- Thêm [8, 10] vào result
- result = [[1, 6], [8, 10]]

**Bước 5:** Xét [15, 18]
- last = [8, 10], currentStart = 15
- 15 > 10 → không chồng chéo
- Thêm [15, 18] vào result
- result = [[1, 6], [8, 10], [15, 18]]

**Kết quả:** `[[1,6],[8,10],[15,18]]`

### Ví dụ 2: `intervals = [[4,7],[1,4]]` (chưa sắp xếp)

**Bước 1:** Sắp xếp theo start
```
[1, 4], [4, 7]
```

**Bước 2:** Khởi tạo result = [[1, 4]]

**Bước 3:** Xét [4, 7]
- last = [1, 4], currentStart = 4, currentEnd = 7
- 4 <= 4 → chồng chéo (intervals chạm nhau cũng được gộp)
- max(4, 7) = 7 → cập nhật last = [1, 7]
- result = [[1, 7]]

**Kết quả:** `[[1, 7]]`

## Giải thích code

### Hàm `merge`

```go
if len(intervals) == 0 {
    return [][]int{}
}
```

Xử lý trường hợp mảng rỗng (theo constraints không xảy ra nhưng an toàn).

```go
sort.Slice(intervals, func(i, j int) bool {
    return intervals[i][0] < intervals[j][0]
})
```

Sắp xếp intervals theo start time (phần tử đầu tiên của mỗi interval).

```go
result := [][]int{{intervals[0][0], intervals[0][1]}}
```

Khởi tạo kết quả với interval đầu tiên. Cần tạo bản sao để tránh tham chiếu đến slice gốc.

```go
if currentStart <= last[1] {
    if currentEnd > last[1] {
        last[1] = currentEnd
    }
} else {
    result = append(result, []int{currentStart, currentEnd})
}
```

Điều kiện gộp: `currentStart <= last[1]` nghĩa là interval hiện tại bắt đầu trước hoặc tại thời điểm interval cuối kết thúc → chồng chéo. Khi gộp, mở rộng end nếu interval hiện tại kéo dài hơn.

## Phân tích độ phức tạp

### Độ phức tạp thời gian: O(n log n)

- Sắp xếp: O(n log n)
- Duyệt và gộp: O(n)
- Tổng: O(n log n)

### Độ phức tạp không gian: O(n)

- Kết quả có thể chứa tối đa n intervals: O(n)
- Sort trong Go có thể sử dụng O(log n) cho call stack
- Tổng: O(n)

## So sánh với các cách tiếp cận khác

| Cách tiếp cận | Thời gian | Không gian | Ghi chú |
|--------------|-----------|------------|---------|
| Brute Force | O(n²) hoặc cao hơn | O(n) | Phức tạp, dễ sai |
| **Sort + Greedy** | **O(n log n)** | **O(n)** | **Tối ưu, dễ hiểu** |

## Follow-up Questions

### 1. Có thể tối ưu hơn O(n log n) không?

Trong trường hợp tổng quát, không. Phải so sánh các intervals với nhau, và việc so sánh n phần tử cần ít nhất O(n log n) trong mô hình so sánh.

Nếu biết trước range của start/end (0 đến 10^4), có thể dùng Counting Sort để đạt O(n + k) với k là range, nhưng phức tạp hơn.

### 2. Nếu cần xử lý intervals được thêm động (online)?

Có thể dùng cấu trúc dữ liệu như Interval Tree hoặc Segment Tree để thêm interval mới và merge trong O(log n) mỗi thao tác.

### 3. Insert Interval (LeetCode 57)?

Bài toán mở rộng: cho một mảng intervals đã sắp xếp và không chồng chéo, thêm interval mới và merge nếu cần. Có thể:
- Thêm interval mới vào mảng
- Gọi merge() - O(n)
- Hoặc dùng binary search để tìm vị trí chèn và merge - O(n) nhưng có thể tối ưu hơn trong một số trường hợp
