package bit_manipulation

// hammingWeight trả về số bit 1 (số lần set bit) trong biểu diễn nhị phân của n (Hamming weight).
//
// Thuật toán: Brian Kernighan - xóa lần lượt bit 1 ngoài cùng bên phải.
// - Công thức n & (n-1): luôn xóa đúng một bit 1 ngoài cùng bên phải.
// - Lặp và đếm số lần cho đến khi n = 0.
//
// Ví dụ n = 11 (binary 1011):
// - 11 & 10 = 1011 & 1010 = 1010, count = 1
// - 10 & 9  = 1010 & 1001 = 1000, count = 2
// - 8 & 7   = 1000 & 0111 = 0000, count = 3 → return 3
//
// Độ phức tạp: O(k) thời gian với k = số bit 1, O(1) không gian.
func hammingWeight(n int) int {
	count := 0
	for n != 0 {
		// n & (n-1) xóa bit 1 ngoài cùng bên phải
		n &= n - 1
		count++
	}
	return count
}
