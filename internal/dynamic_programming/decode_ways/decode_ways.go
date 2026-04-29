package decode_ways

// numDecodings trả về số cách giải mã chuỗi số s theo bảng chữ cái A-Z (1-26).
//
// Thuật toán: Dynamic Programming (Bottom-up)
//
// Định nghĩa dp[i] = số cách giải mã chuỗi con s[0..i-1]
//
// Trạng thái ban đầu:
//   - dp[0] = 1 (chuỗi rỗng có 1 cách giải mã - base case)
//   - dp[1] = 1 nếu s[0] != '0', ngược lại dp[1] = 0
//
// Chuyển trạng thái tại mỗi vị trí i (1-indexed):
//  1. Lấy 1 chữ số: s[i-1]
//     - Nếu s[i-1] != '0' -> dp[i] += dp[i-1]
//  2. Lấy 2 chữ số: s[i-2..i-1]
//     - Nếu 10 <= twoDigit <= 26 -> dp[i] += dp[i-2]
//
// Tối ưu không gian: chỉ cần lưu 2 giá trị prev1 và prev2 thay vì toàn bộ mảng.
func numDecodings(s string) int {
	n := len(s)
	if n == 0 || s[0] == '0' {
		return 0
	}

	// prev2 = dp[i-2], prev1 = dp[i-1]
	prev2 := 1 // dp[0]
	prev1 := 1 // dp[1]: s[0] != '0' đã được kiểm tra ở trên

	for i := 2; i <= n; i++ {
		curr := 0

		// Giải mã 1 chữ số: s[i-1]
		oneDigit := s[i-1] - '0'
		if oneDigit >= 1 {
			curr += prev1
		}

		// Giải mã 2 chữ số: s[i-2..i-1]
		twoDigit := int(s[i-2]-'0')*10 + int(s[i-1]-'0')
		if twoDigit >= 10 && twoDigit <= 26 {
			curr += prev2
		}

		prev2 = prev1
		prev1 = curr
	}

	return prev1
}
