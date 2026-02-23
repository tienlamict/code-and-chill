package array

// maxSubArray tìm tổng lớn nhất của một subarray trong mảng.
//
// Sử dụng Kadane's Algorithm - Dynamic Programming approach:
// - Tại mỗi vị trí, ta có hai lựa chọn:
//   1. Bắt đầu một subarray mới từ vị trí hiện tại
//   2. Mở rộng subarray hiện tại bằng cách thêm phần tử hiện tại
// - Chọn lựa chọn nào cho tổng lớn hơn
//
// Độ phức tạp: O(n) thời gian, O(1) không gian
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Khởi tạo với phần tử đầu tiên
	maxSum := nums[0]      // Tổng lớn nhất tổng thể
	currentSum := nums[0]  // Tổng của subarray kết thúc tại vị trí hiện tại

	// Duyệt qua các phần tử còn lại
	for i := 1; i < len(nums); i++ {
		// Quyết định: bắt đầu mới hay mở rộng subarray hiện tại?
		// Nếu currentSum + nums[i] < nums[i], tức là currentSum < 0
		// thì tốt hơn là bắt đầu subarray mới từ nums[i]
		currentSum = max(nums[i], currentSum+nums[i])

		// Cập nhật tổng lớn nhất tổng thể
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	return maxSum
}

// maxSubArrayDivideConquer tìm tổng lớn nhất sử dụng phương pháp chia để trị.
//
// Ý tưởng:
// - Chia mảng thành hai nửa
// - Tổng lớn nhất có thể nằm ở:
//   1. Nửa trái
//   2. Nửa phải
//   3. Qua điểm giữa (cross middle)
// - Trả về giá trị lớn nhất trong ba trường hợp
//
// Độ phức tạp: O(n log n) thời gian, O(log n) không gian (call stack)
func maxSubArrayDivideConquer(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	return divideConquerHelper(nums, 0, len(nums)-1)
}

// divideConquerHelper hàm đệ quy để tính tổng lớn nhất trong đoạn [left, right]
func divideConquerHelper(nums []int, left, right int) int {
	// Base case: chỉ có một phần tử
	if left == right {
		return nums[left]
	}

	// Chia đôi mảng
	mid := left + (right-left)/2

	// Tìm tổng lớn nhất ở ba vị trí:
	// 1. Nửa trái
	leftMax := divideConquerHelper(nums, left, mid)

	// 2. Nửa phải
	rightMax := divideConquerHelper(nums, mid+1, right)

	// 3. Qua điểm giữa
	crossMax := maxCrossingSum(nums, left, mid, right)

	// Trả về giá trị lớn nhất trong ba trường hợp
	return max(max(leftMax, rightMax), crossMax)
}

// maxCrossingSum tính tổng lớn nhất của subarray qua điểm giữa.
// Subarray này phải bao gồm nums[mid] và có thể mở rộng về cả hai phía.
func maxCrossingSum(nums []int, left, mid, right int) int {
	// Tìm tổng lớn nhất từ mid về bên trái
	leftSum := nums[mid]
	sum := 0
	for i := mid; i >= left; i-- {
		sum += nums[i]
		if sum > leftSum {
			leftSum = sum
		}
	}

	// Tìm tổng lớn nhất từ mid+1 về bên phải
	rightSum := nums[mid+1]
	sum = 0
	for i := mid + 1; i <= right; i++ {
		sum += nums[i]
		if sum > rightSum {
			rightSum = sum
		}
	}

	// Tổng qua điểm giữa = tổng bên trái + tổng bên phải
	return leftSum + rightSum
}

// max trả về giá trị lớn hơn giữa a và b
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
