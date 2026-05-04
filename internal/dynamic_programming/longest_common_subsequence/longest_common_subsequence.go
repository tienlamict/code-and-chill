package longest_common_subsequence

// longestCommonSubsequence trả về độ dài dãy con chung dài nhất của text1 và text2.
//
// Thuật toán: Dynamic Programming (Bottom-up, 2D)
//
// Định nghĩa dp[i][j] = độ dài LCS của text1[0..i-1] và text2[0..j-1]
//
// Trạng thái ban đầu:
//   - dp[0][j] = 0 với mọi j (text1 rỗng)
//   - dp[i][0] = 0 với mọi i (text2 rỗng)
//
// Chuyển trạng thái:
//   - Nếu text1[i-1] == text2[j-1]: dp[i][j] = dp[i-1][j-1] + 1
//   - Ngược lại:                    dp[i][j] = max(dp[i-1][j], dp[i][j-1])
func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)

	// dp[i][j] = LCS của text1[0..i-1] và text2[0..j-1]
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				// Ký tự khớp: mở rộng LCS trước đó thêm 1
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				// Ký tự không khớp: lấy max bỏ ký tự cuối của text1 hoặc text2
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[m][n]
}
