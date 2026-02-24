package array

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
