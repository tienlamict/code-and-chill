# Set Matrix Zeroes (LeetCode 73)

## Mô tả bài toán

Cho một ma trận số nguyên `matrix` kích thước `m x n`. Nếu một phần tử bằng `0` thì **cả hàng** và **cả cột** chứa phần tử đó phải được đặt về `0`.

Yêu cầu: thực hiện **in-place** (không được tạo ma trận mới).

### Ví dụ 1

- Input: `matrix = [[1,1,1],[1,0,1],[1,1,1]]`
- Output: `[[1,0,1],[0,0,0],[1,0,1]]`

### Ví dụ 2

- Input: `matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]`
- Output: `[[0,0,0,0],[0,4,5,0],[0,3,1,0]]`

### Ràng buộc

- `m == matrix.length`
- `n == matrix[0].length`
- `1 <= m, n <= 200`
- `-2^31 <= matrix[i][j] <= 2^31 - 1`

### Follow-up

- O(mn) extra space là quá tệ.
- O(m + n) extra space tốt hơn nhưng vẫn chưa tối ưu.
- Hãy giải bằng **O(1) extra space**.

---

## Các cách tiếp cận

### 1. Dùng ma trận phụ (O(mn) space)

**Ý tưởng**: Tạo một ma trận mới `copy`, duyệt toàn bộ `matrix`. Khi gặp `0` ở `(i, j)` thì set hàng `i` và cột `j` trong `copy` về 0. Cuối cùng sao chép `copy` vào `matrix`.

- Thời gian: O(mn)
- Không gian: O(mn) (ma trận phụ)
- **Nhược điểm**: Tốn quá nhiều bộ nhớ, không đáp ứng follow-up.

### 2. Dùng hai mảng đánh dấu hàng & cột (O(m + n) space)

**Ý tưởng**:

- Dùng:
  - `rows[m]`: `rows[i] = true` nếu hàng `i` phải set về 0.
  - `cols[n]`: `cols[j] = true` nếu cột `j` phải set về 0.
- Bước 1: Duyệt ma trận, khi gặp `matrix[i][j] == 0`:
  - Đánh dấu `rows[i] = true`, `cols[j] = true`.
- Bước 2: Duyệt lại ma trận:
  - Nếu `rows[i] == true` hoặc `cols[j] == true` thì set `matrix[i][j] = 0`.

**Độ phức tạp**:

- Thời gian: O(mn)
- Không gian: O(m + n)

**Ưu điểm**: Dễ hiểu, code đơn giản.

**Nhược điểm**: Vẫn cần thêm O(m + n) bộ nhớ, chưa đạt O(1).

### 3. Dùng hàng 0 & cột 0 làm marker (O(1) extra space) ⭐

Đây là cách tối ưu và được triển khai trong code.

**Ý tưởng chính**:

- Thay vì tạo mảng `rows` và `cols` riêng, ta dùng luôn:
  - **Hàng đầu tiên** (`matrix[0][j]`)
  - **Cột đầu tiên** (`matrix[i][0]`)
  để lưu trạng thái đánh dấu.
- Khi gặp `matrix[i][j] == 0` (với `i > 0`, `j > 0`):
  - Đặt `matrix[i][0] = 0` → hàng `i` sẽ thành 0.
  - Đặt `matrix[0][j] = 0` → cột `j` sẽ thành 0.

**Vấn đề**: Hàng 0 và cột 0 vừa là dữ liệu thật, vừa là nơi đánh dấu → cần biết ban đầu chúng có chứa 0 hay không.

Giải quyết bằng 2 biến cờ:

- `firstRowZero`: hàng đầu có chứa 0 ban đầu không.
- `firstColZero`: cột đầu có chứa 0 ban đầu không.

**Thuật toán chi tiết**:

1. Nếu ma trận rỗng → return.
2. Kiểm tra hàng 0:
   - Nếu có phần tử 0 → `firstRowZero = true`.
3. Kiểm tra cột 0:
   - Nếu có phần tử 0 → `firstColZero = true`.
4. Duyệt các ô còn lại (`i` từ 1..m-1, `j` từ 1..n-1):
   - Nếu `matrix[i][j] == 0`:
     - Đặt `matrix[i][0] = 0`.
     - Đặt `matrix[0][j] = 0`.
5. Duyệt lại các ô còn lại:
   - Nếu `matrix[i][0] == 0` **hoặc** `matrix[0][j] == 0` → set `matrix[i][j] = 0`.
6. Cuối cùng:
   - Nếu `firstRowZero == true` → set cả hàng 0 về 0.
   - Nếu `firstColZero == true` → set cả cột 0 về 0.

**Độ phức tạp**:

- Thời gian: O(mn) (duyệt ma trận vài lần, nhưng cùng thứ tự).
- Không gian: O(1) extra space (chỉ dùng vài biến cờ).

---

## Giải thích code chính

```12:78:internal/hash_table/set_matrix_zeroes/set_matrix_zeroes.go
func setZeroes(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}

	m, n := len(matrix), len(matrix[0])

	firstRowZero := false
	firstColZero := false

	// Kiểm tra xem hàng đầu có phần tử 0 hay không.
	for j := 0; j < n; j++ {
		if matrix[0][j] == 0 {
			firstRowZero = true
			break
		}
	}

	// Kiểm tra xem cột đầu có phần tử 0 hay không.
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			firstColZero = true
			break
		}
	}

	// Dùng hàng 0 và cột 0 làm marker cho các hàng/cột còn lại.
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	// Set các ô (trừ hàng 0, cột 0) dựa trên marker.
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	// Nếu hàng đầu ban đầu có 0, set cả hàng 0 thành 0.
	if firstRowZero {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}

	// Nếu cột đầu ban đầu có 0, set cả cột 0 thành 0.
	if firstColZero {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}
```

---

## Độ phức tạp

- **Thời gian**: O(mn)
  - Duyệt ma trận tối đa 3 lần, vẫn là O(mn).
- **Không gian**: O(1) extra space
  - Không dùng thêm ma trận/phần tử phụ tỉ lệ với m hoặc n.

---

## Test

Các test case được viết theo dạng table-driven trong file:

- `set_matrix_zeroes_test.go`

Các case chính:

- Hai ví dụ từ đề bài.
- Ma trận 1x1 (0 và không 0).
- Ma trận không có số 0.
- Ma trận toàn số 0.
- 0 ở hàng đầu, cột đầu, cả hai.
- Giá trị âm và 0 trộn lẫn.

Chạy test:

```bash
go test ./internal/hash_table/set_matrix_zeroes/
```

