package array

// rotate xoay ma trận n x n 90 độ theo chiều kim đồng hồ, in-place.
//
// Thuật toán: Transpose + Reverse từng hàng
//
// Ý tưởng:
// 1. Transpose: hoán đổi matrix[i][j] với matrix[j][i] (với i < j)
//    - Chuyển hàng thành cột, cột thành hàng
// 2. Reverse mỗi hàng: đảo ngược thứ tự phần tử trong từng hàng
//
// Ví dụ với matrix 3x3:
// Ban đầu:     Transpose:    Reverse mỗi hàng:
// 1 2 3       1 4 7         7 4 1
// 4 5 6   ->  2 5 8    ->   8 5 2
// 7 8 9       3 6 9         9 6 3
//
// Độ phức tạp: O(n²) thời gian, O(1) không gian (in-place)
func rotate(matrix [][]int) {
	n := len(matrix)
	if n <= 1 {
		return
	}

	// Bước 1: Transpose - hoán đổi phần tử qua đường chéo chính
	// Chỉ cần xét i < j để tránh hoán đổi 2 lần
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// Bước 2: Reverse mỗi hàng - đảo ngược thứ tự phần tử trong hàng
	for i := 0; i < n; i++ {
		left, right := 0, n-1
		for left < right {
			matrix[i][left], matrix[i][right] = matrix[i][right], matrix[i][left]
			left++
			right--
		}
	}
}
