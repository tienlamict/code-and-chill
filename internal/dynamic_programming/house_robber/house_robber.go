package house_robber

// rob trả về tổng tiền tối đa có thể trộm được mà không bị báo cảnh sát.
// Không được trộm hai nhà liền kề (hệ thống an ninh sẽ báo).
//
// Thuật toán: Quy hoạch động.
// - dp[i] = tổng tiền tối đa khi xét đến nhà thứ i (0-indexed).
// - Tại nhà i: có hai lựa chọn
//   1. Trộm nhà i → lấy nums[i] + dp[i-2] (vì không được trộm nhà i-1).
//   2. Bỏ qua nhà i → lấy dp[i-1].
// - Công thức: dp[i] = max(dp[i-1], nums[i] + dp[i-2])
// - Base case: dp[0] = nums[0], dp[1] = max(nums[0], nums[1])
//
// Tối ưu không gian: chỉ cần hai biến prev2, prev1 vì mỗi bước chỉ dùng hai giá trị trước.
//
// Độ phức tạp: O(n) thời gian, O(1) không gian.
func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	// prev2 = dp[i-2], prev1 = dp[i-1]
	prev2 := nums[0]
	prev1 := max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		// dp[i] = max(trộm nhà i, bỏ nhà i)
		curr := max(prev1, nums[i]+prev2)
		prev2, prev1 = prev1, curr
	}
	return prev1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
