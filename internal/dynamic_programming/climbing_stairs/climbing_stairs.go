package dynamic_programming

// climbStairs trả về số cách khác nhau để leo lên n bậc thang.
// Mỗi lần có thể bước 1 hoặc 2 bậc.
//
// Thuật toán: Quy hoạch động (tương tự dãy Fibonacci).
// - dp[i] = số cách leo đến bậc i.
// - Công thức: dp[i] = dp[i-1] + dp[i-2]
//   (đến bậc i từ bậc i-1 bước 1, hoặc từ bậc i-2 bước 2).
// - Base case: dp[1]=1, dp[2]=2.
//
// Tối ưu không gian: chỉ cần 2 biến prev, curr vì mỗi bước chỉ dùng 2 giá trị trước.
//
// Ví dụ n=3: 1+1+1, 1+2, 2+1 → 3 cách.
//
// Độ phức tạp: O(n) thời gian, O(1) không gian.
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	prev, curr := 1, 2
	for i := 3; i <= n; i++ {
		prev, curr = curr, prev+curr
	}
	return curr
}
