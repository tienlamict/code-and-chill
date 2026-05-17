package valid_anagram

// isAnagram kiểm tra xem t có phải là anagram của s hay không.
// Anagram là một từ hoặc cụm từ được hình thành bằng cách sắp xếp lại các chữ cái của một từ hoặc cụm từ khác,
// thông thường sử dụng tất cả các chữ cái ban đầu đúng một lần.
func isAnagram(s string, t string) bool {
	// Nếu độ dài hai chuỗi khác nhau, chúng không thể là anagram của nhau.
	if len(s) != len(t) {
		return false
	}

	// Sử dụng một mảng cố định kích thước 26 để đếm số lần xuất hiện của các chữ cái tiếng Anh thường.
	// Index 0 tương ứng với 'a', 1 tương ứng với 'b', ..., 25 tương ứng với 'z'.
	count := [26]int{}

	// Duyệt qua chuỗi s để tăng biến đếm và chuỗi t để giảm biến đếm.
	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
		count[t[i]-'a']--
	}

	// Nếu tất cả các phần tử trong mảng count đều bằng 0, t là anagram của s.
	for _, c := range count {
		if c != 0 {
			return false
		}
	}

	return true
}
