package hash_table

// setZeroes đặt toàn bộ hàng và cột thành 0 nếu một phần tử trong đó bằng 0.
//
// Yêu cầu follow-up: làm in-place với O(1) extra space.
// Ý tưởng:
// - Dùng hàng đầu (row 0) và cột đầu (col 0) làm nơi đánh dấu (marker).
// - Khi gặp phần tử matrix[i][j] == 0 (với i > 0, j > 0), ta:
//   + Đặt matrix[i][0] = 0 để đánh dấu cả hàng i sẽ thành 0.
//   + Đặt matrix[0][j] = 0 để đánh dấu cả cột j sẽ thành 0.
// - Cần hai cờ riêng:
//   + firstRowZero: hàng 0 có chứa 0 ban đầu không.
//   + firstColZero: cột 0 có chứa 0 ban đầu không.
// - Sau khi đánh dấu, duyệt lại từ i=1..m-1, j=1..n-1:
//   + Nếu matrix[i][0] == 0 hoặc matrix[0][j] == 0 thì đặt matrix[i][j] = 0.
// - Cuối cùng, xử lý riêng hàng 0 và cột 0 dựa trên hai cờ firstRowZero, firstColZero.
//
// Độ phức tạp:
// - Thời gian: O(m * n) vì chỉ duyệt ma trận vài lần.
// - Không gian: O(1) extra space, chỉ dùng vài biến cờ, không dùng ma trận phụ.
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

