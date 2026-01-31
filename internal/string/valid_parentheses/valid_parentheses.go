package string

// isValid kiểm tra xem chuỗi chứa các dấu ngoặc có hợp lệ hay không.
//
// Sử dụng Stack để kiểm tra tính hợp lệ:
// - Khi gặp dấu ngoặc mở '(', '{', '[', push vào stack
// - Khi gặp dấu ngoặc đóng ')', '}', ']', kiểm tra xem có match với dấu ngoặc mở trên đỉnh stack không
// - Nếu match, pop khỏi stack; nếu không match hoặc stack rỗng, return false
// - Cuối cùng, stack phải rỗng để chuỗi hợp lệ
//
// Độ phức tạp: O(n) thời gian, O(n) không gian
func isValid(s string) bool {
	// Stack để lưu các dấu ngoặc mở
	stack := []rune{}

	// Map để ánh xạ dấu ngoặc đóng với dấu ngoặc mở tương ứng
	closingToOpening := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	// Duyệt qua từng ký tự trong chuỗi
	for _, char := range s {
		// Nếu là dấu ngoặc đóng
		if opening, isClosing := closingToOpening[char]; isClosing {
			// Kiểm tra xem stack có rỗng không
			if len(stack) == 0 {
				return false
			}

			// Lấy dấu ngoặc mở trên đỉnh stack
			top := stack[len(stack)-1]

			// Nếu không match, chuỗi không hợp lệ
			if top != opening {
				return false
			}

			// Pop khỏi stack (match thành công)
			stack = stack[:len(stack)-1]
		} else {
			// Nếu là dấu ngoặc mở, push vào stack
			stack = append(stack, char)
		}
	}

	// Stack phải rỗng để chuỗi hợp lệ
	return len(stack) == 0
}
