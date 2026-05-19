package sum_of_two_integers

// getSum tính tổng của hai số nguyên a và b mà không sử dụng toán tử + hoặc -
// Thuật toán: Sử dụng các phép toán bitwise:
// - XOR (^) để tính tổng các bit mà không màng tới nhớ (carry).
// - AND (&) và dịch trái (<<) để tính toán bit nhớ (carry).
// Lặp lại quá trình cho đến khi không còn bit nhớ.
func getSum(a int, b int) int {
	for b != 0 {
		// carry chứa các bit sẽ được nhớ lên hàng tiếp theo
		carry := (a & b) << 1
		// a chứa tổng của các bit hiện tại mà chưa có nhớ
		a = a ^ b
		// b trở thành carry để tiếp tục cộng vào a ở vòng lặp sau
		b = carry
	}
	return a
}
