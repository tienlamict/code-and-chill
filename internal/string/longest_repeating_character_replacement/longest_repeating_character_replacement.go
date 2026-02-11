package string

// characterReplacement trả về độ dài của chuỗi con dài nhất chứa cùng một ký tự
// sau khi thay đổi tối đa k ký tự.
//
// Sử dụng kỹ thuật Sliding Window:
// - Duy trì một cửa sổ trượt [left, right] đại diện cho chuỗi con hiện tại.
// - Trong cửa sổ, đếm tần suất của từng ký tự.
// - maxFreq là tần suất cao nhất của một ký tự trong cửa sổ hiện tại.
// - Số ký tự cần thay đổi = (độ dài cửa sổ) - maxFreq.
// - Nếu số ký tự cần thay đổi > k, thu hẹp cửa sổ từ bên trái.
// - Cập nhật kết quả là kích thước cửa sổ lớn nhất hợp lệ.
//
// Lưu ý quan trọng: maxFreq không cần giảm khi thu hẹp cửa sổ vì:
// - Ta chỉ quan tâm đến cửa sổ LỚN HƠN cửa sổ tốt nhất đã tìm được.
// - Kết quả chỉ cải thiện khi maxFreq tăng, nên không cần cập nhật khi giảm.
//
// Độ phức tạp: O(n) thời gian, O(1) không gian (mảng 26 ký tự)
func characterReplacement(s string, k int) int {
	// Mảng đếm tần suất của 26 ký tự in hoa (A-Z)
	count := [26]int{}

	left := 0    // Con trỏ trái của cửa sổ
	maxFreq := 0 // Tần suất cao nhất của một ký tự trong cửa sổ
	result := 0  // Kết quả: độ dài cửa sổ lớn nhất hợp lệ

	for right := 0; right < len(s); right++ {
		// Thêm ký tự bên phải vào cửa sổ
		count[s[right]-'A']++

		// Cập nhật tần suất cao nhất
		if count[s[right]-'A'] > maxFreq {
			maxFreq = count[s[right]-'A']
		}

		// Kiểm tra tính hợp lệ của cửa sổ:
		// Số ký tự cần thay đổi = (độ dài cửa sổ) - maxFreq
		// Nếu > k, cửa sổ không hợp lệ → thu hẹp từ bên trái
		windowLen := right - left + 1
		if windowLen-maxFreq > k {
			// Giảm tần suất của ký tự bên trái
			count[s[left]-'A']--
			// Di chuyển con trỏ trái sang phải
			left++
		}

		// Cập nhật kết quả
		if right-left+1 > result {
			result = right - left + 1
		}
	}

	return result
}
