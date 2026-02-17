package array

// lengthOfLIS tìm độ dài của dãy con tăng dài nhất (Longest Increasing Subsequence).
//
// Thuật toán sử dụng Dynamic Programming với Binary Search để đạt độ phức tạp O(n log n):
//
// Ý tưởng chính:
// - Thay vì lưu độ dài LIS tại mỗi vị trí, ta duy trì một mảng tails[]
//   trong đó tails[i] = phần tử nhỏ nhất có thể kết thúc một dãy con tăng có độ dài i+1
//
// Cách hoạt động:
// 1. Duyệt qua từng phần tử trong nums
// 2. Với mỗi phần tử, tìm vị trí chèn phù hợp trong tails[] bằng binary search
// 3. Nếu phần tử lớn hơn tất cả phần tử trong tails[], thêm vào cuối (tăng độ dài LIS)
// 4. Nếu không, thay thế phần tử tại vị trí tìm được (giữ tails[] luôn có phần tử nhỏ nhất)
//
// Ví dụ: nums = [10,9,2,5,3,7,101,18]
// i=0: tails = [10]
// i=1: tails = [9]      (9 < 10, thay thế)
// i=2: tails = [2]      (2 < 9, thay thế)
// i=3: tails = [2,5]    (5 > 2, thêm vào)
// i=4: tails = [2,3]    (3 < 5, thay thế)
// i=5: tails = [2,3,7]  (7 > 3, thêm vào)
// i=6: tails = [2,3,7,101] (101 > 7, thêm vào)
// i=7: tails = [2,3,7,18]  (18 < 101, thay thế)
// Kết quả: len(tails) = 4
//
// Độ phức tạp: O(n log n) thời gian, O(n) không gian
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// tails[i] = phần tử nhỏ nhất có thể kết thúc một dãy con tăng có độ dài i+1
	tails := []int{nums[0]}

	// Duyệt qua từng phần tử trong nums
	for i := 1; i < len(nums); i++ {
		num := nums[i]

		// Nếu phần tử lớn hơn phần tử cuối cùng trong tails,
		// ta có thể mở rộng dãy con tăng dài nhất hiện tại
		if num > tails[len(tails)-1] {
			tails = append(tails, num)
		} else {
			// Tìm vị trí chèn phù hợp bằng binary search
			// Tìm phần tử nhỏ nhất trong tails[] mà >= num
			left, right := 0, len(tails)-1
			pos := right

			for left <= right {
				mid := left + (right-left)/2
				if tails[mid] >= num {
					pos = mid
					right = mid - 1
				} else {
					left = mid + 1
				}
			}

			// Thay thế phần tử tại vị trí pos để giữ tails[] có phần tử nhỏ nhất
			tails[pos] = num
		}
	}

	// Độ dài của tails chính là độ dài LIS
	return len(tails)
}

// lengthOfLISDP là cách giải bằng Dynamic Programming truyền thống O(n²).
// Giữ lại để tham khảo và so sánh.
//
// Thuật toán:
// 1. dp[i] = độ dài LIS kết thúc tại vị trí i
// 2. Với mỗi vị trí i, xét tất cả vị trí j < i
// 3. Nếu nums[j] < nums[i], có thể mở rộng dãy con từ j sang i
// 4. dp[i] = max(dp[j]) + 1 với mọi j < i và nums[j] < nums[i]
//
// Độ phức tạp: O(n²) thời gian, O(n) không gian
func lengthOfLISDP(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// dp[i] = độ dài LIS kết thúc tại vị trí i
	dp := make([]int, len(nums))
	
	// Khởi tạo: mỗi phần tử tự nó là một dãy con tăng có độ dài 1
	for i := range dp {
		dp[i] = 1
	}

	maxLen := 1

	// Với mỗi vị trí i, tìm độ dài LIS kết thúc tại i
	for i := 1; i < len(nums); i++ {
		// Xét tất cả các vị trí j trước i
		for j := 0; j < i; j++ {
			// Nếu nums[j] < nums[i], có thể mở rộng dãy con từ j sang i
			if nums[j] < nums[i] {
				// Cập nhật dp[i] nếu tìm thấy dãy con dài hơn
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
				}
			}
		}

		// Cập nhật độ dài LIS lớn nhất
		if dp[i] > maxLen {
			maxLen = dp[i]
		}
	}

	return maxLen
}
