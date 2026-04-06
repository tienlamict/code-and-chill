package house_robber_II

// rob trả về tổng tiền tối đa có thể trộm được mà không bị báo cảnh sát.
// Tại nơi này, các ngôi nhà được sắp xếp theo một vòng tròn.
// Điều đó có nghĩa là ngôi nhà đầu tiên là hàng xóm của ngôi nhà cuối cùng.
// Do đó, ta không thể trộm cả hai nhà này cùng lúc.
//
// Thuật toán: Quy hoạch động.
// Vì các ngôi nhà tạo thành một vòng tròn, chúng ta có thể chia bài toán thành 2 trường hợp:
// 1. Trộm các nhà từ chỉ số 0 đến n-2 (không bao gồm nhà cuối cùng).
// 2. Trộm các nhà từ chỉ số 1 đến n-1 (không bao gồm nhà đầu tiên).
// Sau đó, trả về giá trị lớn nhất của 2 trường hợp này.
// Mỗi trường hợp đều là bài toán House Robber I truyền thống.
//
// Độ phức tạp thời gian: O(n) vì chúng ta quét mảng 2 lần.
// Độ phức tạp không gian: O(1) bằng cách sử dụng các biến thay vì mảng DP.
func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}

	// Trường hợp 1: Trộm từ nhà đầu đến nhà áp chót (0 đến n-2)
	case1 := robLinear(nums[:n-1])
	// Trường hợp 2: Trộm từ nhà thứ hai đến nhà cuối (1 đến n-1)
	case2 := robLinear(nums[1:])

	return max(case1, case2)
}

// robLinear giải quyết bài toán House Robber cho một dãy nhà thẳng.
func robLinear(nums []int) int {
	prev2, prev1 := 0, 0
	for _, num := range nums {
		// Công thức: dp[i] = max(dp[i-1], num + dp[i-2])
		curr := max(prev1, num+prev2)
		prev2 = prev1
		prev1 = curr
	}
	return prev1
}
