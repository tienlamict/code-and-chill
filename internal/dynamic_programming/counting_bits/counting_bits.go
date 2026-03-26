package counting_bits

// CountBits nhận một số nguyên n và trả về một mảng ans có độ dài n + 1
// sao cho với mỗi i (0 <= i <= n), ans[i] là số lượng bit 1 trong biểu diễn nhị phân của i.
// Thuật toán sử dụng Quy hoạch động (Dynamic Programming) để đạt độ phức tạp O(n).
func countBits(n int) []int {
	// Khởi tạo mảng kết quả với độ dài n + 1, mặc định các giá trị là 0.
	ans := make([]int, n+1)

	// Duyệt qua từng số từ 1 đến n.
	// Với số 0, ans[0] đã là 0 (đúng).
	for i := 1; i <= n; i++ {
		// Công thức Quy hoạch động:
		// Số lượng bit 1 của i = (số lượng bit 1 của i dịch phải 1 bit) + (bit cuối cùng của i).
		// i >> 1 tương đương với i / 2.
		// i & 1 tương đương với i % 2.
		// Ví dụ: i = 5 (101)
		// i >> 1 = 2 (10), ans[2] = 1
		// i & 1 = 1
		// ans[5] = ans[2] + 1 = 1 + 1 = 2 (đúng)
		ans[i] = ans[i>>1] + (i & 1)
	}

	return ans
}
