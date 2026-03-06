package bit_manipulation

// reverseBits đảo ngược các bit của một số nguyên 32-bit.
//
// Thuật toán: Bit Manipulation
// - Duyệt qua tất cả 32 bit của số đầu vào
// - Với mỗi bit, lấy bit cuối cùng (LSB) của n
// - Thêm bit đó vào kết quả (đã được dịch chuyển phù hợp)
// - Dịch chuyển n sang phải 1 bit và kết quả sang trái 1 bit
//
// Ví dụ với n = 43261596 (00000010100101000001111010011100):
// - Bit 0: lấy bit cuối (0), kết quả = 0, n >>= 1
// - Bit 1: lấy bit cuối (0), kết quả = 00, n >>= 1
// - ... tiếp tục cho đến bit 31
// - Kết quả: 964176192 (00111001011110000010100101000000)
//
// Độ phức tạp: O(1) thời gian (luôn duyệt 32 bit), O(1) không gian
func reverseBits(n uint32) uint32 {
	var result uint32

	// Duyệt qua tất cả 32 bit
	for i := 0; i < 32; i++ {
		// Lấy bit cuối cùng (LSB) của n bằng cách AND với 1
		bit := n & 1

		// Dịch chuyển kết quả sang trái 1 bit và thêm bit mới vào vị trí cuối
		result = (result << 1) | bit

		// Dịch chuyển n sang phải 1 bit để xử lý bit tiếp theo
		n >>= 1
	}

	return result
}
