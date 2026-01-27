package hash_table

// missingNumber tìm số còn thiếu trong mảng chứa n số phân biệt trong khoảng [0, n].
//
// Sử dụng công thức tổng:
// - Tổng của các số từ 0 đến n: n * (n + 1) / 2
// - Tính tổng các số trong mảng
// - Số còn thiếu = Tổng lý thuyết - Tổng thực tế
//
// Độ phức tạp: O(n) thời gian, O(1) không gian
func missingNumber(nums []int) int {
	n := len(nums)
	
	// Tổng lý thuyết của các số từ 0 đến n
	expectedSum := n * (n + 1) / 2
	
	// Tính tổng thực tế của các số trong mảng
	actualSum := 0
	for _, num := range nums {
		actualSum += num
	}
	
	// Số còn thiếu = Tổng lý thuyết - Tổng thực tế
	return expectedSum - actualSum
}
