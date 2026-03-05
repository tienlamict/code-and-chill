package dynamic_programming

// uniquePaths tính số đường đi duy nhất từ góc trên-trái (0,0) đến góc dưới-phải (m-1, n-1)
// trên lưới m x n, chỉ được di chuyển xuống hoặc sang phải.
//
// Sử dụng Dynamic Programming:
// 1. dp[i][j] = số cách đến ô (i, j)
// 2. Công thức: dp[i][j] = dp[i-1][j] + dp[i][j-1]
//    (có thể đến từ ô trên hoặc ô trái)
// 3. Base case: dp[0][0] = 1, hàng đầu và cột đầu đều là 1
// 4. Tối ưu không gian: chỉ cần mảng 1D vì chỉ dùng hàng trước đó
//
// Độ phức tạp: O(m*n) thời gian, O(n) không gian
func uniquePaths(m int, n int) int {
	// Edge case: nếu chỉ có 1 hàng hoặc 1 cột, chỉ có 1 đường đi
	if m == 1 || n == 1 {
		return 1
	}

	// Tối ưu không gian: chỉ dùng mảng 1D để lưu hàng hiện tại
	// dp[j] đại diện cho số cách đến ô ở hàng hiện tại, cột j
	dp := make([]int, n)

	// Khởi tạo: hàng đầu tiên, tất cả các ô đều có 1 cách (chỉ đi sang phải)
	for j := 0; j < n; j++ {
		dp[j] = 1
	}

	// Duyệt từ hàng thứ 2 đến hàng cuối
	for i := 1; i < m; i++ {
		// Cột đầu tiên luôn là 1 (chỉ đi xuống)
		// Cập nhật các cột còn lại: dp[j] = dp[j] (từ trên xuống) + dp[j-1] (từ trái sang)
		for j := 1; j < n; j++ {
			dp[j] = dp[j] + dp[j-1]
		}
	}

	// Kết quả là số cách đến ô cuối cùng
	return dp[n-1]
}
