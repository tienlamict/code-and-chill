# Find Median from Data Stream

## Mô tả bài toán

Median là giá trị ở giữa trong một danh sách số nguyên đã được sắp xếp. Nếu kích thước của danh sách là chẵn, không có giá trị ở giữa, và median là trung bình của hai giá trị ở giữa.

Ví dụ:
- Với `arr = [2,3,4]`, median là `3`
- Với `arr = [2,3]`, median là `(2 + 3) / 2 = 2.5`

Implement class `MedianFinder`:

- `MedianFinder()` khởi tạo đối tượng MedianFinder
- `void addNum(int num)` thêm số nguyên `num` từ data stream vào cấu trúc dữ liệu
- `double findMedian()` trả về median của tất cả các phần tử cho đến nay. Câu trả lời trong phạm vi 10^-5 của câu trả lời thực tế sẽ được chấp nhận.

**Ví dụ:**

### Example 1:
- Input: 
  ```
  ["MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"]
  [[], [1], [2], [], [3], []]
  ```
- Output: `[null, null, null, 1.5, null, 2.0]`
- Giải thích:
  ```
  MedianFinder medianFinder = new MedianFinder();
  medianFinder.addNum(1);    // arr = [1]
  medianFinder.addNum(2);    // arr = [1, 2]
  medianFinder.findMedian(); // return 1.5 (i.e., (1 + 2) / 2)
  medianFinder.addNum(3);    // arr = [1, 2, 3]
  medianFinder.findMedian(); // return 2.0
  ```

**Ràng buộc:**
- `-10^5 <= num <= 10^5`
- Sẽ có ít nhất một phần tử trong cấu trúc dữ liệu trước khi gọi `findMedian`
- Tối đa `5 * 10^4` lần gọi sẽ được thực hiện cho `addNum` và `findMedian`

**Follow up:**
- Nếu tất cả các số nguyên từ stream nằm trong khoảng `[0, 100]`, bạn sẽ tối ưu giải pháp như thế nào?
- Nếu 99% các số nguyên từ stream nằm trong khoảng `[0, 100]`, bạn sẽ tối ưu giải pháp như thế nào?

## Phân tích thuật toán

### Cách tiếp cận 1: Sắp xếp mỗi lần (O(n log n) time per query)

Mỗi lần gọi `findMedian()`, sắp xếp lại tất cả các số và tìm median.

**Độ phức tạp:**
- `addNum`: O(1) thời gian
- `findMedian`: O(n log n) thời gian

### Cách tiếp cận 2: Insertion Sort (O(n) time per addNum)

Giữ mảng luôn được sắp xếp bằng cách chèn số mới vào đúng vị trí.

**Độ phức tạp:**
- `addNum`: O(n) thời gian
- `findMedian`: O(1) thời gian

### Cách tiếp cận 3: Two Heaps (O(log n) time per addNum, O(1) time per findMedian) ⭐ Tối ưu

Sử dụng hai heaps để duy trì hai nửa của dữ liệu:
- **Max Heap (small)**: Lưu nửa nhỏ hơn hoặc bằng, phần tử lớn nhất ở top
- **Min Heap (large)**: Lưu nửa lớn hơn, phần tử nhỏ nhất ở top

**Ý tưởng:**
- Chia dữ liệu thành hai nửa bằng nhau (hoặc chênh lệch 1)
- Nửa nhỏ hơn trong max heap, nửa lớn hơn trong min heap
- Median = top của max heap (nếu số lượng lẻ) hoặc (top max heap + top min heap) / 2 (nếu số lượng chẵn)

**Điều kiện duy trì:**
1. `len(small) >= len(large)` và `len(small) - len(large) <= 1`
2. Tất cả phần tử trong `small` <= tất cả phần tử trong `large`

**Độ phức tạp:**
- `addNum`: O(log n) thời gian - push/pop từ heap
- `findMedian`: O(1) thời gian - chỉ cần peek top của heap
- Không gian: O(n) - lưu trữ tất cả các số

### Ví dụ minh họa

Với các operations: `addNum(1)`, `addNum(2)`, `findMedian()`, `addNum(3)`, `findMedian()`

```
Bước 1: addNum(1)
  small (max heap): [1]
  large (min heap): []
  len(small) = 1, len(large) = 0 ✓

Bước 2: addNum(2)
  Thêm 2 vào small: small = [2, 1] (max heap)
  Kiểm tra: small.top() = 2, large.top() = không có → OK
  Cân bằng: len(small) = 2, len(large) = 0
  Di chuyển 2 từ small sang large:
    small = [1]
    large = [2]
  len(small) = 1, len(large) = 1 ✓

Bước 3: findMedian()
  len(small) = 1, len(large) = 1 (chẵn)
  median = (small.top() + large.top()) / 2 = (1 + 2) / 2 = 1.5

Bước 4: addNum(3)
  Thêm 3 vào small: small = [3, 1] (max heap)
  Kiểm tra: small.top() = 3, large.top() = 2
  3 > 2 → Di chuyển 3 từ small sang large:
    small = [1]
    large = [2, 3] (min heap)
  Cân bằng: len(small) = 1, len(large) = 2
  Di chuyển 2 từ large sang small:
    small = [2, 1] (max heap)
    large = [3]
  len(small) = 2, len(large) = 1 ✓

Bước 5: findMedian()
  len(small) = 2, len(large) = 1 (lẻ, small lớn hơn)
  median = small.top() = 2.0
```

## Giải thích code

### Cấu trúc MedianFinder

```1:30:internal/data_steam/find_median_from_data_stream/find_median_from_data_stream.go
package data_stream

import "container/heap"

// MedianFinder tìm median từ data stream sử dụng hai heaps.
//
// Sử dụng kỹ thuật Two Heaps:
// - Max heap (small): Lưu nửa nhỏ hơn hoặc bằng của dữ liệu, phần tử lớn nhất ở top
// - Min heap (large): Lưu nửa lớn hơn của dữ liệu, phần tử nhỏ nhất ở top
// - Đảm bảo: len(small) >= len(large) và len(small) - len(large) <= 1
// - Median = top của small (nếu số lượng lẻ) hoặc (top small + top large) / 2 (nếu số lượng chẵn)
//
// Độ phức tạp:
// - AddNum: O(log n) thời gian
// - FindMedian: O(1) thời gian
// - Không gian: O(n)
type MedianFinder struct {
	small *MaxHeap // Nửa nhỏ hơn hoặc bằng (max heap)
	large *MinHeap // Nửa lớn hơn (min heap)
}

// Constructor khởi tạo đối tượng MedianFinder
func Constructor() MedianFinder {
	return MedianFinder{
		small: &MaxHeap{},
		large: &MinHeap{},
	}
}
```

### Hàm AddNum

```32:50:internal/data_steam/find_median_from_data_stream/find_median_from_data_stream.go
// AddNum thêm một số vào data stream
func (mf *MedianFinder) AddNum(num int) {
	// Thêm vào small heap trước
	heap.Push(mf.small, num)

	// Đảm bảo tất cả phần tử trong small <= tất cả phần tử trong large
	if mf.small.Len() > 0 && mf.large.Len() > 0 && mf.small.Peek() > mf.large.Peek() {
		// Di chuyển phần tử lớn nhất từ small sang large
		val := heap.Pop(mf.small).(int)
		heap.Push(mf.large, val)
	}

	// Đảm bảo cân bằng: len(small) >= len(large) và len(small) - len(large) <= 1
	if mf.small.Len() > mf.large.Len()+1 {
		val := heap.Pop(mf.small).(int)
		heap.Push(mf.large, val)
	} else if mf.large.Len() > mf.small.Len() {
		val := heap.Pop(mf.large).(int)
		heap.Push(mf.small, val)
	}
}
```

**Giải thích:**
1. **Thêm vào small** (dòng 35): Thêm số mới vào max heap `small`
2. **Kiểm tra thứ tự** (dòng 37-41): Nếu top của `small` > top của `large`, di chuyển phần tử lớn nhất từ `small` sang `large`
3. **Cân bằng kích thước** (dòng 43-48):
   - Nếu `small` lớn hơn `large` quá 1 phần tử, di chuyển phần tử lớn nhất từ `small` sang `large`
   - Nếu `large` lớn hơn `small`, di chuyển phần tử nhỏ nhất từ `large` sang `small`

### Hàm FindMedian

```52:60:internal/data_steam/find_median_from_data_stream/find_median_from_data_stream.go
// FindMedian trả về median của tất cả các số đã thêm
func (mf *MedianFinder) FindMedian() float64 {
	if mf.small.Len() > mf.large.Len() {
		// Số lượng lẻ: median là phần tử ở giữa (top của small)
		return float64(mf.small.Peek())
	}
	// Số lượng chẵn: median là trung bình của hai phần tử giữa
	return float64(mf.small.Peek()+mf.large.Peek()) / 2.0
}
```

**Giải thích:**
- Nếu số lượng lẻ: `len(small) > len(large)` → median = top của `small`
- Nếu số lượng chẵn: `len(small) == len(large)` → median = (top `small` + top `large`) / 2

### MaxHeap và MinHeap

```62:95:internal/data_steam/find_median_from_data_stream/find_median_from_data_stream.go
// MaxHeap là max heap (phần tử lớn nhất ở top)
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] } // Đảo dấu để có max heap
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h MaxHeap) Peek() int {
	return h[0]
}

// MinHeap là min heap (phần tử nhỏ nhất ở top)
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h MinHeap) Peek() int {
	return h[0]
}
```

**Giải thích:**
- **MaxHeap**: Implement heap interface với `Less(i, j) = h[i] > h[j]` để có max heap
- **MinHeap**: Implement heap interface với `Less(i, j) = h[i] < h[j]` để có min heap
- **Peek()**: Trả về phần tử ở top mà không xóa nó

### Độ phức tạp

- **AddNum**: O(log n)
  - Push vào heap: O(log n)
  - Pop từ heap: O(log n)
  - Tối đa 2 lần push/pop: O(log n)

- **FindMedian**: O(1)
  - Chỉ cần peek top của heap: O(1)

- **Không gian**: O(n)
  - Lưu trữ tất cả các số trong hai heaps

### Tại sao thuật toán này đúng?

1. **Invariant**: 
   - Tất cả phần tử trong `small` <= tất cả phần tử trong `large`
   - `len(small) >= len(large)` và `len(small) - len(large) <= 1`

2. **Median chính xác**:
   - Nếu số lượng lẻ: Phần tử ở giữa là top của `small` (vì `small` lớn hơn 1 phần tử)
   - Nếu số lượng chẵn: Hai phần tử ở giữa là top của `small` và `large`

3. **Cân bằng tự động**: Sau mỗi `addNum`, các điều kiện được đảm bảo

## Test Cases

### Test Case 1: Example 1
```go
Operations: ["MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"]
Values: [nil, 1, 2, nil, 3, nil]
Expected: [nil, nil, nil, 1.5, nil, 2.0]
```

### Test Case 2: Single number
```go
Operations: ["MedianFinder", "addNum", "findMedian"]
Values: [nil, 5, nil]
Expected: [nil, nil, 5.0]
```

### Test Case 3: Even count
```go
Operations: ["MedianFinder", "addNum", "addNum", "findMedian"]
Values: [nil, 1, 2, nil]
Expected: [nil, nil, nil, 1.5]
```

### Test Case 4: Odd count
```go
Operations: ["MedianFinder", "addNum", "addNum", "addNum", "findMedian"]
Values: [nil, 1, 2, 3, nil]
Expected: [nil, nil, nil, nil, 2.0]
```

## Chạy test

Để chạy các test case:

```bash
go test ./internal/data_steam/find_median_from_data_stream/
```

Hoặc chạy với verbose mode để xem chi tiết:

```bash
go test -v ./internal/data_steam/find_median_from_data_stream/
```

## Follow-up Questions

### Follow-up 1: Numbers trong khoảng [0, 100]

Nếu tất cả số đều trong khoảng [0, 100], có thể sử dụng **Counting Sort**:
- Tạo mảng `count[101]` để đếm số lần xuất hiện của mỗi số
- `addNum`: O(1) - chỉ cần tăng count
- `findMedian`: O(100) - duyệt qua mảng để tìm median

**Độ phức tạp:**
- `addNum`: O(1)
- `findMedian`: O(100) = O(1) (hằng số)

### Follow-up 2: 99% số trong khoảng [0, 100]

Kết hợp hai cách:
- Sử dụng counting array cho số trong [0, 100]
- Sử dụng two heaps cho số ngoài [0, 100]
- Khi tìm median, xử lý cả hai phần

**Độ phức tạp:**
- `addNum`: O(1) cho số trong [0, 100], O(log n) cho số ngoài
- `findMedian`: O(100) = O(1) trong hầu hết trường hợp

## Mở rộng

### Sliding Window Median (LeetCode 480)

Tìm median trong một sliding window kích thước k. Yêu cầu kỹ thuật tương tự nhưng cần xử lý việc xóa phần tử cũ.

### Ứng dụng thực tế

Bài toán này có ứng dụng trong:
- **Real-time Analytics**: Tính toán median trong stream dữ liệu
- **Financial Systems**: Theo dõi median giá cổ phiếu
- **Network Monitoring**: Tính toán median latency
- **Database Systems**: Tối ưu hóa truy vấn median
- **Statistics**: Tính toán thống kê real-time
