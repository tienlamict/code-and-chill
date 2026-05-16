package product_of_array_except_self

// ProductExceptSelf tính toán mảng kết quả sao cho mỗi phần tử tại index i
// là tích của tất cả các phần tử trong mảng nums ngoại trừ nums[i].
// Thuật toán đạt độ phức tạp thời gian O(n) và không gian O(1) (không tính mảng kết quả).
func ProductExceptSelf(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}

	// Khởi tạo mảng kết quả với kích thước n
	res := make([]int, n)

	// Bước 1: Tính tích các phần tử bên trái (Prefix Product)
	// res[i] sẽ chứa tích của tất cả các phần tử từ index 0 đến i-1
	res[0] = 1
	for i := 1; i < n; i++ {
		res[i] = res[i-1] * nums[i-1]
	}

	// Bước 2: Tính tích các phần tử bên phải (Suffix Product) và nhân vào kết quả
	// suffixProduct lưu trữ tích của các phần tử từ index n-1 xuống i+1
	suffixProduct := 1
	for i := n - 1; i >= 0; i-- {
		// Kết quả cuối cùng tại i = (tích bên trái) * (tích bên phải)
		res[i] = res[i] * suffixProduct
		// Cập nhật suffixProduct cho bước tiếp theo (phần tử bên trái tiếp theo)
		suffixProduct *= nums[i]
	}

	return res
}
