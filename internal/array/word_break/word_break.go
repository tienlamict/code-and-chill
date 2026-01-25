package array

// wordBreak kiểm tra xem chuỗi s có thể được phân tách thành các từ trong wordDict hay không.
//
// Sử dụng Dynamic Programming với memoization:
// - dp[i] = true nếu s[0:i] có thể được phân tách thành các từ trong wordDict
// - Với mỗi vị trí i, kiểm tra tất cả các từ trong wordDict:
//   Nếu s[j:i] là một từ trong wordDict và dp[j] = true, thì dp[i] = true
func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	
	// Tạo map để kiểm tra từ nhanh hơn (O(1) thay vì O(n))
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}
	
	// dp[i] = true nếu s[0:i] có thể được phân tách thành các từ trong wordDict
	dp := make([]bool, n+1)
	dp[0] = true // Chuỗi rỗng luôn có thể phân tách được
	
	// Duyệt qua từng vị trí trong chuỗi
	for i := 1; i <= n; i++ {
		// Kiểm tra tất cả các vị trí j trước đó
		for j := 0; j < i; j++ {
			// Nếu s[0:j] có thể phân tách được (dp[j] = true)
			// và s[j:i] là một từ trong wordDict
			// thì s[0:i] cũng có thể phân tách được
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
				break // Đã tìm thấy cách phân tách, không cần kiểm tra tiếp
			}
		}
	}
	
	return dp[n]
}
