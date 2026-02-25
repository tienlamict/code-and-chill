package array

// canJump kiểm tra xem có thể nhảy từ vị trí đầu tiên đến vị trí cuối cùng của mảng hay không.
//
// Mỗi phần tử nums[i] đại diện cho độ dài bước nhảy tối đa từ vị trí i.
// Trả về true nếu có thể đến được vị trí cuối cùng, false nếu không.
//
// Thuật toán Greedy:
// 1. Duyệt qua mảng từ trái sang phải
// 2. Tại mỗi vị trí, cập nhật vị trí xa nhất có thể đạt được (maxReach)
// 3. Nếu tại bất kỳ vị trí nào, vị trí hiện tại vượt quá maxReach, nghĩa là không thể đến được vị trí đó
// 4. Kiểm tra xem maxReach có đạt được vị trí cuối cùng hay không
//
// Độ phức tạp: O(n) thời gian, O(1) không gian
func canJump(nums []int) bool {
	// Vị trí xa nhất có thể đạt được từ các vị trí đã duyệt
	maxReach := 0

	// Duyệt qua từng vị trí trong mảng
	for i := 0; i < len(nums); i++ {
		// Nếu vị trí hiện tại vượt quá maxReach, nghĩa là không thể đến được vị trí này
		// Do đó không thể đến được vị trí cuối cùng
		if i > maxReach {
			return false
		}

		// Cập nhật vị trí xa nhất có thể đạt được
		// Từ vị trí i, có thể nhảy tối đa i + nums[i] bước
		currentReach := i + nums[i]
		if currentReach > maxReach {
			maxReach = currentReach
		}

		// Nếu đã đạt được vị trí cuối cùng hoặc vượt quá, trả về true
		if maxReach >= len(nums)-1 {
			return true
		}
	}

	// Kiểm tra lại sau khi duyệt xong (trường hợp đặc biệt)
	return maxReach >= len(nums)-1
}
