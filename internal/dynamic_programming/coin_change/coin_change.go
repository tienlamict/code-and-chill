package dynamic_programming

// coinChange trả về số đồng xu ít nhất để tạo ra amount (tổng tiền).
// Mỗi loại xu có thể dùng vô hạn lần. Không thể tạo được thì trả về -1.
//
// Thuật toán: Quy hoạch động (bottom-up).
// - dp[i] = số xu ít nhất để tạo tổng i (dp[0] = 0).
// - Với mỗi tổng i từ 1 đến amount: thử từng loại xu c, nếu c <= i thì
//   dp[i] = min(dp[i], 1 + dp[i-c]).
// - Khởi tạo dp[i] = amount+1 (vô cùng) vì tối đa cần không quá amount xu (xu 1).
//
// Ví dụ coins = [1,2,5], amount = 11:
// dp[0]=0, dp[1]=1, dp[2]=1, ..., dp[5]=1, ..., dp[10]=2, dp[11]=3 → return 3.
//
// Độ phức tạp: O(amount * len(coins)) thời gian, O(amount) không gian.
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	// dp[i] = số xu ít nhất để tạo tổng i
	// Khởi tạo amount+1 thay cho "vô cùng" (tối đa cần amount xu nếu chỉ dùng xu 1)
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}
	// Lần lượt tính dp[1], dp[2], ..., dp[amount]
	for i := 1; i <= amount; i++ {
		for _, c := range coins {
			if c <= i {
				// Chọn xu c: số xu = 1 + số xu cần cho tổng i-c
				if onePlus := 1 + dp[i-c]; onePlus < dp[i] {
					dp[i] = onePlus
				}
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}
