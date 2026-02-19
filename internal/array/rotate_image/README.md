# 48. Rotate Image

## Mô tả bài toán

Cho một ma trận 2D `n x n` đại diện cho một hình ảnh, hãy xoay hình ảnh đó 90 độ theo chiều kim đồng hồ.

**Yêu cầu quan trọng**: Bạn phải xoay hình ảnh **in-place**, tức là phải sửa đổi trực tiếp ma trận đầu vào. KHÔNG được cấp phát thêm ma trận 2D mới để thực hiện phép xoay.

## Ví dụ minh họa

### Ví dụ 1:
```
Input:  matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [[7,4,1],[8,5,2],[9,6,3]]

Trước:          Sau khi xoay 90° CW:
1 2 3           7 4 1
4 5 6     ->    8 5 2
7 8 9           9 6 3
```

### Ví dụ 2:
```
Input:  matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
Output: [[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]
```

## Ràng buộc

- `n == matrix.length == matrix[i].length` (ma trận vuông)
- `1 <= n <= 20`
- `-1000 <= matrix[i][j] <= 1000`

## Phân tích các cách tiếp cận

### Cách 1: Cấp phát ma trận mới (Không được phép)

**Ý tưởng**: Tạo ma trận mới và copy các phần tử theo quy tắc xoay 90°.

**Quy tắc**: `new[i][j] = old[n-1-j][i]`

**Độ phức tạp**:
- Thời gian: O(n²)
- Không gian: O(n²)

**Nhận xét**: Vi phạm yêu cầu in-place của bài toán.

---

### Cách 2: Hoán đổi từng nhóm 4 phần tử (Layer-by-layer)

**Ý tưởng**: Xoay ma trận theo từng "lớp" từ ngoài vào trong. Với mỗi phần tử ở góc trên-trái của lớp, hoán đổi 4 phần tử tương ứng (trên-trái, trên-phải, dưới-phải, dưới-trái).

**Công thức hoán đổi** cho phần tử tại (row, col) trong lớp có kích thước `n`:
- (row, col) ↔ (col, n-1-row)
- (col, n-1-row) ↔ (n-1-row, n-1-col)
- (n-1-row, n-1-col) ↔ (n-1-col, row)
- (n-1-col, row) ↔ (row, col)

**Độ phức tạp**:
- Thời gian: O(n²) - mỗi phần tử được xử lý đúng 1 lần
- Không gian: O(1) - chỉ dùng biến tạm

**Ưu điểm**: In-place, O(1) không gian

**Nhược điểm**: Logic phức tạp hơn, dễ nhầm lẫn với chỉ số

---

### Cách 3: Transpose + Reverse từng hàng ⭐

**Ý tưởng**: Phép xoay 90° CW có thể phân tích thành 2 bước đơn giản:
1. **Transpose**: Hoán đổi matrix[i][j] với matrix[j][i] (với i < j)
2. **Reverse**: Đảo ngược thứ tự phần tử trong mỗi hàng

**Chứng minh**:
- Transpose chuyển hàng thành cột: cột j của ma trận gốc trở thành hàng j của ma trận sau transpose
- Reverse mỗi hàng tương đương với việc "lật" cột từ dưới lên trên

**Ví dụ từng bước với matrix 3x3**:

```
Bước 0 - Ban đầu:
1 2 3
4 5 6
7 8 9

Bước 1 - Transpose (hoán đổi qua đường chéo chính):
- (0,1)↔(1,0): 2↔4
- (0,2)↔(2,0): 3↔7
- (1,2)↔(2,1): 6↔8

Kết quả:
1 4 7
2 5 8
3 6 9

Bước 2 - Reverse mỗi hàng:
- Hàng 0: [1,4,7] -> [7,4,1]
- Hàng 1: [2,5,8] -> [8,5,2]
- Hàng 2: [3,6,9] -> [9,6,3]

Kết quả cuối:
7 4 1
8 5 2
9 6 3  ✓ (đúng với output mong đợi)
```

**Độ phức tạp**:
- Thời gian: O(n²) - transpose O(n²), reverse O(n²)
- Không gian: O(1) - in-place, chỉ dùng biến tạm khi swap

**Ưu điểm**: Đơn giản, dễ hiểu, dễ implement, dễ kiểm chứng

**Nhược điểm**: Không có

---

## Giải thích code

### Hàm chính: `rotate`

```go
func rotate(matrix [][]int) {
	n := len(matrix)
	if n <= 1 {
		return
	}
```

Kiểm tra trường hợp ma trận rỗng hoặc 1x1 (xoay 90° không thay đổi gì).

```go
	// Bước 1: Transpose - hoán đổi phần tử qua đường chéo chính
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
```

Transpose: chỉ cần xét `j > i` để tránh hoán đổi 2 lần (nếu xét cả i và j thì mỗi cặp sẽ bị swap 2 lần, trả về trạng thái cũ).

```go
	// Bước 2: Reverse mỗi hàng
	for i := 0; i < n; i++ {
		left, right := 0, n-1
		for left < right {
			matrix[i][left], matrix[i][right] = matrix[i][right], matrix[i][left]
			left++
			right--
		}
	}
}
```

Reverse từng hàng bằng two pointers: swap phần tử đầu và cuối, di chuyển vào trong cho đến khi gặp nhau.

## Phân tích độ phức tạp

**Thời gian**: O(n²)
- Transpose: duyệt n(n-1)/2 cặp phần tử → O(n²)
- Reverse: mỗi hàng n phần tử, n hàng → O(n²)
- Tổng: O(n²)

**Không gian**: O(1)
- Không cấp phát thêm cấu trúc dữ liệu
- Chỉ dùng biến tạm khi swap (Go swap in-place không cần biến phụ)

## Kết luận

Bài toán Rotate Image có thể giải bằng cách kết hợp **Transpose** và **Reverse** - hai thao tác đơn giản, dễ hiểu. Cách tiếp cận này đạt O(n²) thời gian và O(1) không gian, đáp ứng đầy đủ yêu cầu in-place của bài toán.
