package valid_palindrome

// isPalindrome kiểm tra xem một chuỗi có phải là palindrome hay không sau khi đã:
// 1. Chuyển tất cả chữ hoa thành chữ thường.
// 2. Loại bỏ tất cả các ký tự không phải chữ cái và số (non-alphanumeric).
// Độ phức tạp thời gian: O(n), trong đó n là độ dài của chuỗi s.
// Độ phức tạp không gian: O(1), vì chúng ta sử dụng con trỏ hai đầu trực tiếp trên chuỗi.
func isPalindrome(s string) bool {
	// Sử dụng hai con trỏ: left bắt đầu từ đầu chuỗi, right bắt đầu từ cuối chuỗi.
	left, right := 0, len(s)-1

	for left < right {
		// Di chuyển left sang phải nếu gặp ký tự không phải chữ cái hoặc số.
		for left < right && !isAlphanumeric(s[left]) {
			left++
		}
		// Di chuyển right sang trái nếu gặp ký tự không phải chữ cái hoặc số.
		for left < right && !isAlphanumeric(s[right]) {
			right--
		}

		// So sánh hai ký tự sau khi đã chuyển về chữ thường.
		if toLower(s[left]) != toLower(s[right]) {
			return false
		}

		// Tiếp tục kiểm tra cặp ký tự tiếp theo.
		left++
		right--
	}

	return true
}

// isAlphanumeric kiểm tra xem một ký tự byte có phải là chữ cái hoặc số hay không.
func isAlphanumeric(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

// toLower chuyển một ký tự chữ hoa thành chữ thường.
func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + ('a' - 'A')
	}
	return c
}
