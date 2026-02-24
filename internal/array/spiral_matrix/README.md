# Spiral Matrix

## Mô tả bài toán

Cho một ma trận `m x n`, trả về tất cả các phần tử của ma trận theo thứ tự xoắn ốc (spiral order).

**Thứ tự xoắn ốc** là cách duyệt ma trận bắt đầu từ góc trên bên trái, đi sang phải, xuống dưới, sang trái, lên trên, và tiếp tục theo mô hình xoắn ốc cho đến khi duyệt hết tất cả các phần tử.

**Ví dụ:**

### Example 1:
- Input: `matrix = [[1,2,3],[4,5,6],[7,8,9]]`
- Output: `[1,2,3,6,9,8,7,4,5]`
- Giải thích: Duyệt theo thứ tự: 1→2→3→6→9→8→7→4→5

```
Ma trận 3x3:
┌─────┬─────┬─────┐
│  1  │  2  │  3  │  → (phải)
├─────┼─────┼─────┤
│  4  │  5  │  6  │
├─────┼─────┼─────┤
│  7  │  8  │  9  │
└─────┴─────┴─────┘
     ↑ (trái)  ↓ (xuống)

Thứ tự: 1 → 2 → 3 → 6 → 9 → 8 → 7 → 4 → 5
```

### Example 2:
- Input: `matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]`
- Output: `[1,2,3,4,8,12,11,10,9,5,6,7]`
- Giải thích: Duyệt theo thứ tự: 1→2→3→4→8→12→11→10→9→5→6→7

```
Ma trận 3x4:
┌─────┬─────┬─────┬─────┐
│  1  │  2  │  3  │  4  │  → (phải)
├─────┼─────┼─────┼─────┤
│  5  │  6  │  7  │  8  │
├─────┼─────┼─────┼─────┤
│  9  │ 10  │ 11  │ 12  │
└─────┴─────┴─────┴─────┘
     ↑ (trái)  ↓ (xuống)

Thứ tự: 1 → 2 → 3 → 4 → 8 → 12 → 11 → 10 → 9 → 5 → 6 → 7
```

**Ràng buộc:**
- `m == matrix.length`
- `n == matrix[i].length`
- `1 <= m, n <= 10`
- `-100 <= matrix[i][j] <= 100`

## Phân tích thuật toán

### Cách tiếp cận 1: Boundary Tracking (Theo dõi biên) ⭐ Tối ưu nhất

**Ý tưởng chính:**

Sử dụng 4 biến để theo dõi các biên của ma trận:
- `top`: Hàng trên cùng chưa được duyệt
- `bottom`: Hàng dưới cùng chưa được duyệt
- `left`: Cột bên trái chưa được duyệt
- `right`: Cột bên phải chưa được duyệt

Duyệt ma trận theo 4 hướng tuần tự:
1. **Phải (Right)**: Từ `left` đến `right` ở hàng `top`
2. **Xuống (Down)**: Từ `top` đến `bottom` ở cột `right`
3. **Trái (Left)**: Từ `right` đến `left` ở hàng `bottom`
4. **Lên (Up)**: Từ `bottom` đến `top` ở cột `left`

Sau mỗi hướng, thu hẹp biên tương ứng và kiểm tra điều kiện dừng.

**Thuật toán:**

1. **Khởi tạo:**
   - `top = 0`, `bottom = m - 1`
   - `left = 0`, `right = n - 1`
   - `result = []` (mảng kết quả)

2. **Vòng lặp chính** (khi `top <= bottom && left <= right`):
   - **Bước 1 - Phải:** Duyệt từ `left` đến `right` ở hàng `top`, sau đó `top++`
   - **Bước 2 - Xuống:** Duyệt từ `top` đến `bottom` ở cột `right`, sau đó `right--`
   - **Kiểm tra:** Nếu `top > bottom`, dừng (không cần duyệt hàng dưới)
   - **Bước 3 - Trái:** Duyệt từ `right` đến `left` ở hàng `bottom`, sau đó `bottom--`
   - **Kiểm tra:** Nếu `left > right`, dừng (không cần duyệt cột trái)
   - **Bước 4 - Lên:** Duyệt từ `bottom` đến `top` ở cột `left`, sau đó `left++`

3. **Điều kiện dừng:**
   - Khi `top > bottom` hoặc `left > right`
   - Cần kiểm tra sau bước 2 và bước 3 để tránh duyệt lại các phần tử đã duyệt

**Độ phức tạp:**
- Thời gian: O(m × n) - duyệt qua tất cả các phần tử một lần
- Không gian: O(1) - chỉ sử dụng thêm các biến biên (không tính output array)

**Ưu điểm:**
- Đơn giản, dễ hiểu
- Hiệu quả về thời gian và không gian
- Không cần đệ quy hoặc cấu trúc dữ liệu phức tạp

**Nhược điểm:**
- Cần cẩn thận với điều kiện dừng để tránh duyệt lại phần tử

### Cách tiếp cận 2: Direction Array (Mảng hướng)

**Ý tưởng chính:**

Sử dụng mảng hướng `directions = [(0,1), (1,0), (0,-1), (-1,0)]` để đại diện cho 4 hướng: phải, xuống, trái, lên.

**Thuật toán:**

1. Khởi tạo ma trận `visited` để đánh dấu các phần tử đã duyệt
2. Bắt đầu từ `(0, 0)` với hướng ban đầu là phải
3. Duyệt theo hướng hiện tại cho đến khi gặp biên hoặc phần tử đã duyệt
4. Đổi hướng theo thứ tự: phải → xuống → trái → lên
5. Lặp lại cho đến khi duyệt hết tất cả phần tử

**Độ phức tạp:**
- Thời gian: O(m × n)
- Không gian: O(m × n) - cần ma trận `visited`

**Nhược điểm:**
- Tốn thêm không gian cho ma trận `visited`
- Phức tạp hơn cách tiếp cận 1

### Cách tiếp cận 3: Recursive (Đệ quy)

**Ý tưởng chính:**

Duyệt từng lớp (layer) của ma trận một cách đệ quy:
- Lớp ngoài cùng: hàng trên, cột phải, hàng dưới, cột trái
- Sau đó gọi đệ quy cho lớp bên trong

**Độ phức tạp:**
- Thời gian: O(m × n)
- Không gian: O(min(m, n)) - call stack

**Nhược điểm:**
- Phức tạp hơn, có thể gây stack overflow với ma trận lớn

## Ví dụ minh họa

### Ví dụ 1: Ma trận 3x3

Với input: `matrix = [[1,2,3],[4,5,6],[7,8,9]]`

```
Khởi tạo:
top = 0, bottom = 2
left = 0, right = 2

Lần lặp 1:
  Bước 1 - Phải: [1, 2, 3] (hàng 0, cột 0→2)
    top = 1
  Bước 2 - Xuống: [6, 9] (cột 2, hàng 1→2)
    right = 1
  Kiểm tra: top(1) <= bottom(2) ✓
  Bước 3 - Trái: [8, 7] (hàng 2, cột 1→0)
    bottom = 1
  Kiểm tra: left(0) <= right(1) ✓
  Bước 4 - Lên: [4] (cột 0, hàng 1→1)
    left = 1

Lần lặp 2:
  top = 1, bottom = 1, left = 1, right = 1
  Bước 1 - Phải: [5] (hàng 1, cột 1→1)
    top = 2
  Bước 2 - Xuống: (không có, vì top(2) > bottom(1))
  Dừng (top > bottom)

Kết quả: [1, 2, 3, 6, 9, 8, 7, 4, 5]
```

### Ví dụ 2: Ma trận 3x4

Với input: `matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]`

```
Khởi tạo:
top = 0, bottom = 2
left = 0, right = 3

Lần lặp 1:
  Bước 1 - Phải: [1, 2, 3, 4] (hàng 0, cột 0→3)
    top = 1
  Bước 2 - Xuống: [8, 12] (cột 3, hàng 1→2)
    right = 2
  Kiểm tra: top(1) <= bottom(2) ✓
  Bước 3 - Trái: [11, 10, 9] (hàng 2, cột 2→0)
    bottom = 1
  Kiểm tra: left(0) <= right(2) ✓
  Bước 4 - Lên: [5] (cột 0, hàng 1→1)
    left = 1

Lần lặp 2:
  top = 1, bottom = 1, left = 1, right = 2
  Bước 1 - Phải: [6, 7] (hàng 1, cột 1→2)
    top = 2
  Bước 2 - Xuống: (không có, vì top(2) > bottom(1))
  Dừng (top > bottom)

Kết quả: [1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7]
```

### Ví dụ 3: Ma trận 1 hàng

Với input: `matrix = [[1, 2, 3, 4]]`

```
Khởi tạo:
top = 0, bottom = 0
left = 0, right = 3

Lần lặp 1:
  Bước 1 - Phải: [1, 2, 3, 4] (hàng 0, cột 0→3)
    top = 1
  Bước 2 - Xuống: (không có, vì top(1) > bottom(0))
  Dừng (top > bottom)

Kết quả: [1, 2, 3, 4]
```

### Ví dụ 4: Ma trận 1 cột

Với input: `matrix = [[1], [2], [3], [4]]`

```
Khởi tạo:
top = 0, bottom = 3
left = 0, right = 0

Lần lặp 1:
  Bước 1 - Phải: [1] (hàng 0, cột 0→0)
    top = 1
  Bước 2 - Xuống: [2, 3, 4] (cột 0, hàng 1→3)
    right = -1
  Kiểm tra: top(1) <= bottom(3) ✓
  Bước 3 - Trái: (không có, vì right(-1) < left(0))
  Dừng (left > right)

Kết quả: [1, 2, 3, 4]
```

## Giải thích code

```3:78:internal/array/spiral_matrix/spiral_matrix.go
// spiralOrder trả về tất cả các phần tử của ma trận theo thứ tự xoắn ốc (spiral order).
//
// Thuật toán: Boundary Tracking (Theo dõi biên)
// - Sử dụng 4 biến để theo dõi các biên: top, bottom, left, right
// - Duyệt theo 4 hướng: phải -> xuống -> trái -> lên
// - Sau mỗi hướng, thu hẹp biên tương ứng
// - Dừng khi top > bottom hoặc left > right
//
// Độ phức tạp: O(m * n) thời gian, O(1) không gian (không tính output array)
func spiralOrder(matrix [][]int) []int {
	// Kiểm tra ma trận rỗng
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	// Lấy kích thước ma trận
	m := len(matrix)    // Số hàng
	n := len(matrix[0]) // Số cột

	// Khởi tạo mảng kết quả với kích thước m * n
	result := make([]int, 0, m*n)

	// Khởi tạo các biên
	top := 0         // Hàng trên cùng
	bottom := m - 1  // Hàng dưới cùng
	left := 0        // Cột bên trái
	right := n - 1   // Cột bên phải

	// Duyệt ma trận theo thứ tự xoắn ốc
	for top <= bottom && left <= right {
		// Bước 1: Duyệt từ trái sang phải ở hàng trên cùng
		for col := left; col <= right; col++ {
			result = append(result, matrix[top][col])
		}
		top++ // Thu hẹp biên trên

		// Bước 2: Duyệt từ trên xuống dưới ở cột bên phải
		for row := top; row <= bottom; row++ {
			result = append(result, matrix[row][right])
		}
		right-- // Thu hẹp biên phải

		// Kiểm tra điều kiện: nếu top > bottom, không cần duyệt hàng dưới
		if top > bottom {
			break
		}

		// Bước 3: Duyệt từ phải sang trái ở hàng dưới cùng
		for col := right; col >= left; col-- {
			result = append(result, matrix[bottom][col])
		}
		bottom-- // Thu hẹp biên dưới

		// Kiểm tra điều kiện: nếu left > right, không cần duyệt cột trái
		if left > right {
			break
		}

		// Bước 4: Duyệt từ dưới lên trên ở cột bên trái
		for row := bottom; row >= top; row-- {
			result = append(result, matrix[row][left])
		}
		left++ // Thu hẹp biên trái
	}

	return result
}
```

#### Chi tiết từng phần

**1. Kiểm tra ma trận rỗng (dòng 15-17)**
```go
if len(matrix) == 0 || len(matrix[0]) == 0 {
    return []int{}
}
```
- Xử lý trường hợp ma trận rỗng hoặc hàng rỗng
- Trả về mảng rỗng ngay lập tức

**2. Khởi tạo biên (dòng 24-27)**
```go
top := 0         // Hàng trên cùng
bottom := m - 1  // Hàng dưới cùng
left := 0        // Cột bên trái
right := n - 1   // Cột bên phải
```
- `top`, `bottom`: Giới hạn phạm vi hàng cần duyệt
- `left`, `right`: Giới hạn phạm vi cột cần duyệt
- Ban đầu, duyệt toàn bộ ma trận

**3. Bước 1 - Duyệt phải (dòng 30-33)**
```go
for col := left; col <= right; col++ {
    result = append(result, matrix[top][col])
}
top++ // Thu hẹp biên trên
```
- Duyệt từ `left` đến `right` ở hàng `top`
- Sau khi duyệt xong, tăng `top` để loại bỏ hàng đã duyệt

**4. Bước 2 - Duyệt xuống (dòng 35-38)**
```go
for row := top; row <= bottom; row++ {
    result = append(result, matrix[row][right])
}
right-- // Thu hẹp biên phải
```
- Duyệt từ `top` đến `bottom` ở cột `right`
- Sau khi duyệt xong, giảm `right` để loại bỏ cột đã duyệt
- Lưu ý: Bắt đầu từ `top` (không phải `top+1`) vì `top` đã được tăng ở bước 1

**5. Kiểm tra điều kiện sau bước 2 (dòng 40-42)**
```go
if top > bottom {
    break
}
```
- **Quan trọng:** Kiểm tra sau bước 2 để tránh duyệt lại hàng dưới
- Ví dụ: Ma trận 1 hàng, sau bước 1 `top = 1`, `bottom = 0`, không cần duyệt hàng dưới

**6. Bước 3 - Duyệt trái (dòng 44-47)**
```go
for col := right; col >= left; col-- {
    result = append(result, matrix[bottom][col])
}
bottom-- // Thu hẹp biên dưới
```
- Duyệt từ `right` đến `left` ở hàng `bottom` (ngược lại)
- Sau khi duyệt xong, giảm `bottom` để loại bỏ hàng đã duyệt

**7. Kiểm tra điều kiện sau bước 3 (dòng 49-51)**
```go
if left > right {
    break
}
```
- **Quan trọng:** Kiểm tra sau bước 3 để tránh duyệt lại cột trái
- Ví dụ: Ma trận 1 cột, sau bước 2 `right = -1`, `left = 0`, không cần duyệt cột trái

**8. Bước 4 - Duyệt lên (dòng 53-56)**
```go
for row := bottom; row >= top; row-- {
    result = append(result, matrix[row][left])
}
left++ // Thu hẹp biên trái
```
- Duyệt từ `bottom` đến `top` ở cột `left` (ngược lại)
- Sau khi duyệt xong, tăng `left` để loại bỏ cột đã duyệt

## Phân tích độ phức tạp

### Thời gian: O(m × n)

- Duyệt qua tất cả các phần tử của ma trận một lần duy nhất
- Mỗi phần tử được thêm vào kết quả trong thời gian O(1)
- Tổng: O(m × n)

**Phân tích chi tiết:**
- Vòng lặp ngoài: Tối đa `min(m, n) / 2` lần (mỗi lần loại bỏ 2 hàng hoặc 2 cột)
- Mỗi lần lặp: Duyệt O(m + n) phần tử (tổng các phần tử ở biên)
- Tổng: O(m × n) phần tử

### Không gian: O(1)

- Chỉ sử dụng thêm các biến: `top`, `bottom`, `left`, `right`, `m`, `n`
- Không sử dụng thêm cấu trúc dữ liệu nào (không tính output array)
- Không gian phụ: O(1)

**Lưu ý:** Nếu tính output array, không gian là O(m × n) - đây là yêu cầu của bài toán.

### So sánh các cách tiếp cận

| Phương pháp | Thời gian | Không gian | Ưu điểm | Nhược điểm |
|------------|-----------|------------|---------|------------|
| Boundary Tracking | O(m×n) | O(1) | Đơn giản, hiệu quả | Cần cẩn thận với điều kiện dừng |
| Direction Array | O(m×n) | O(m×n) | Dễ mở rộng | Tốn không gian cho visited |
| Recursive | O(m×n) | O(min(m,n)) | Elegant | Phức tạp, có thể stack overflow |

**Kết luận:** Boundary Tracking là giải pháp tối ưu nhất cho bài toán này.

## Tại sao cần kiểm tra điều kiện sau bước 2 và bước 3?

### Vấn đề

Sau khi duyệt phải và xuống, nếu không kiểm tra điều kiện, có thể duyệt lại các phần tử đã duyệt.

### Ví dụ minh họa

**Ma trận 1 hàng:** `[[1, 2, 3, 4]]`

```
Khởi tạo: top=0, bottom=0, left=0, right=3

Bước 1 - Phải: [1, 2, 3, 4]
  top = 1

Bước 2 - Xuống: (không có, vì top(1) > bottom(0))
  right = 2

Nếu không kiểm tra:
  Bước 3 - Trái: [3, 2] (hàng 0, cột 2→0) ❌ DUYỆT LẠI!
  
Nếu có kiểm tra:
  if top(1) > bottom(0): break ✓ DỪNG ĐÚNG
```

**Ma trận 1 cột:** `[[1], [2], [3], [4]]`

```
Khởi tạo: top=0, bottom=3, left=0, right=0

Bước 1 - Phải: [1]
  top = 1

Bước 2 - Xuống: [2, 3, 4]
  right = -1

Bước 3 - Trái: (không có, vì right(-1) < left(0))
  bottom = 2

Nếu không kiểm tra:
  Bước 4 - Lên: [3, 2] (cột 0, hàng 2→1) ❌ DUYỆT LẠI!
  
Nếu có kiểm tra:
  if left(0) > right(-1): break ✓ DỪNG ĐÚNG
```

### Kết luận

Kiểm tra điều kiện sau bước 2 và bước 3 là **bắt buộc** để tránh duyệt lại các phần tử đã duyệt, đặc biệt với ma trận hình chữ nhật (không vuông).

## Test Cases

### Test Case 1: Example 1
```go
Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [1,2,3,6,9,8,7,4,5]
```
Ma trận vuông 3x3.

### Test Case 2: Example 2
```go
Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
Output: [1,2,3,4,8,12,11,10,9,5,6,7]
```
Ma trận hình chữ nhật 3x4.

### Test Case 3: Single element
```go
Input: matrix = [[1]]
Output: [1]
```
Ma trận chỉ có một phần tử.

### Test Case 4: Single row
```go
Input: matrix = [[1,2,3,4]]
Output: [1,2,3,4]
```
Ma trận chỉ có một hàng.

### Test Case 5: Single column
```go
Input: matrix = [[1],[2],[3],[4]]
Output: [1,2,3,4]
```
Ma trận chỉ có một cột.

### Test Case 6: 2x2 matrix
```go
Input: matrix = [[1,2],[3,4]]
Output: [1,2,4,3]
```
Ma trận vuông nhỏ nhất.

### Test Case 7: Matrix with negative numbers
```go
Input: matrix = [[-1,-2,-3],[-4,-5,-6],[-7,-8,-9]]
Output: [-1,-2,-3,-6,-9,-8,-7,-4,-5]
```
Ma trận với số âm.

### Test Case 8: Matrix with zeros
```go
Input: matrix = [[0,0,0],[0,0,0],[0,0,0]]
Output: [0,0,0,0,0,0,0,0,0]
```
Ma trận toàn số 0.

### Test Case 9: Rectangular matrix 2x5
```go
Input: matrix = [[1,2,3,4,5],[6,7,8,9,10]]
Output: [1,2,3,4,5,10,9,8,7,6]
```
Ma trận hình chữ nhật rộng.

### Test Case 10: Rectangular matrix 5x2
```go
Input: matrix = [[1,2],[3,4],[5,6],[7,8],[9,10]]
Output: [1,2,4,6,8,10,9,7,5,3]
```
Ma trận hình chữ nhật cao.

## Chạy test

Để chạy các test case:

```bash
go test ./internal/array/spiral_matrix/
```

Hoặc chạy với verbose mode để xem chi tiết:

```bash
go test -v ./internal/array/spiral_matrix/
```

Chạy benchmark để kiểm tra hiệu suất:

```bash
go test -bench=. ./internal/array/spiral_matrix/
```

## Mở rộng

### Biến thể 1: Spiral Matrix II - Tạo ma trận xoắn ốc

Thay vì duyệt ma trận, tạo ma trận xoắn ốc từ các số 1 đến n²:

```go
func generateMatrix(n int) [][]int {
    matrix := make([][]int, n)
    for i := range matrix {
        matrix[i] = make([]int, n)
    }
    
    top, bottom := 0, n-1
    left, right := 0, n-1
    num := 1
    
    for top <= bottom && left <= right {
        // Tương tự spiralOrder nhưng gán giá trị thay vì đọc
        // ...
    }
    
    return matrix
}
```

### Biến thể 2: Spiral Matrix III - Bắt đầu từ vị trí bất kỳ

Duyệt ma trận xoắn ốc bắt đầu từ vị trí `(rStart, cStart)`:

```go
func spiralMatrixIII(rows int, cols int, rStart int, cStart int) [][]int {
    // Sử dụng direction array và mở rộng dần bán kính
    // ...
}
```

### Biến thể 3: Spiral Matrix IV - Với linked list

Duyệt ma trận và điền giá trị từ linked list:

```go
func spiralMatrix(m int, n int, head *ListNode) [][]int {
    // Kết hợp spiral order với linked list
    // ...
}
```

### Biến thể 4: Reverse Spiral Order

Duyệt ma trận theo thứ tự xoắn ốc ngược (từ trong ra ngoài):

```go
func reverseSpiralOrder(matrix [][]int) []int {
    // Bắt đầu từ giữa, mở rộng ra ngoài
    // ...
}
```

## Ứng dụng thực tế

Spiral Matrix được sử dụng trong:

1. **Xử lý ảnh:**
   - Quét ảnh theo thứ tự xoắn ốc
   - Nén ảnh (JPEG zigzag scan)
   - Image rotation

2. **Game development:**
   - Tạo hiệu ứng xoắn ốc
   - Pathfinding algorithms
   - Map generation

3. **Data visualization:**
   - Hiển thị dữ liệu theo thứ tự xoắn ốc
   - Heatmap visualization
   - Matrix traversal animations

4. **Printing:**
   - In ma trận theo thứ tự xoắn ốc
   - Format output

5. **Algorithm design:**
   - Pattern recognition
   - Matrix algorithms
   - Interview questions

## Tips và Tricks

1. **Kiểm tra điều kiện:** Luôn kiểm tra `top > bottom` và `left > right` sau bước 2 và bước 3
2. **Khởi tạo đúng:** Bắt đầu với `top=0, bottom=m-1, left=0, right=n-1`
3. **Thu hẹp biên:** Sau mỗi hướng, thu hẹp biên tương ứng ngay lập tức
4. **Xử lý ma trận rỗng:** Kiểm tra `len(matrix) == 0` hoặc `len(matrix[0]) == 0`
5. **Debug:** In ra `top, bottom, left, right` tại mỗi bước để hiểu thuật toán
6. **Tối ưu:** Sử dụng `make([]int, 0, m*n)` để pre-allocate capacity

## Lưu ý về Edge Cases

1. **Ma trận rỗng:** Trả về `[]`
2. **Ma trận 1 phần tử:** Trả về `[matrix[0][0]]`
3. **Ma trận 1 hàng:** Chỉ duyệt phải, không cần duyệt xuống/trái/lên
4. **Ma trận 1 cột:** Chỉ duyệt phải và xuống, không cần duyệt trái/lên
5. **Ma trận vuông:** Duyệt đầy đủ 4 hướng
6. **Ma trận hình chữ nhật:** Cần kiểm tra điều kiện cẩn thận

## Tối ưu hóa

### Pre-allocate capacity

```go
result := make([]int, 0, m*n)  // Pre-allocate capacity
```

Giúp giảm số lần reallocate khi `append`.

### Kiểm tra điều kiện sớm

```go
if top > bottom {
    break  // Dừng ngay khi không cần duyệt tiếp
}
```

Tránh duyệt không cần thiết.

### Sử dụng index thay vì append (nếu biết trước kích thước)

```go
result := make([]int, m*n)
idx := 0
// ...
result[idx] = matrix[top][col]
idx++
```

Nhanh hơn `append` một chút, nhưng code phức tạp hơn.

## Kết luận

**Spiral Matrix** là một bài toán kinh điển về duyệt ma trận, được giải quyết hiệu quả bằng phương pháp **Boundary Tracking** với độ phức tạp O(m × n) thời gian và O(1) không gian.

**Key takeaways:**
- Boundary Tracking là giải pháp tối ưu nhất
- Ý tưởng chính: Theo dõi 4 biên và thu hẹp sau mỗi hướng
- Quan trọng: Kiểm tra điều kiện sau bước 2 và bước 3
- Thuật toán có thể mở rộng cho nhiều biến thể khác

**Độ khó:** Medium - Cần hiểu cách duyệt ma trận và xử lý điều kiện biên cẩn thận

**Thời gian giải:** 20-40 phút (tùy vào kinh nghiệm với ma trận)
