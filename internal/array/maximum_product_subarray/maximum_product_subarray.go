package array

// maxProduct tìm tích lớn nhất của một subarray trong mảng.
//
// Sử dụng Dynamic Programming với hai biến:
// - maxProd: Tích lớn nhất kết thúc tại vị trí hiện tại
// - minProd: Tích nhỏ nhất kết thúc tại vị trí hiện tại
//
// Lý do cần cả max và min:
// - Khi gặp số âm, tích nhỏ nhất có thể trở thành tích lớn nhất
// - Ví dụ: minProd = -10, nums[i] = -2 => maxProd = 20
//
// Độ phức tạp: O(n) thời gian, O(1) không gian
func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Khởi tạo với phần tử đầu tiên
	maxProd := nums[0]
	minProd := nums[0]
	result := nums[0]

	// Duyệt qua các phần tử còn lại
	for i := 1; i < len(nums); i++ {
		num := nums[i]

		// Nếu số hiện tại là số âm, swap max và min
		// vì tích với số âm sẽ đảo ngược thứ tự
		if num < 0 {
			maxProd, minProd = minProd, maxProd
		}

		// Cập nhật max và min product kết thúc tại vị trí i
		// Có hai lựa chọn:
		// 1. Bắt đầu subarray mới từ vị trí i (num)
		// 2. Mở rộng subarray hiện tại (maxProd * num hoặc minProd * num)
		maxProd = max(num, maxProd*num)
		minProd = min(num, minProd*num)

		// Cập nhật kết quả tổng thể
		if maxProd > result {
			result = maxProd
		}
	}

	return result
}

// max trả về giá trị lớn hơn giữa a và b
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// min trả về giá trị nhỏ hơn giữa a và b
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
