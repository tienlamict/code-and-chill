package array

import "sort"

// threeSum tìm tất cả các bộ ba số có tổng bằng 0 trong mảng.
//
// Sử dụng kỹ thuật Two Pointers:
// 1. Sắp xếp mảng để dễ dàng xử lý duplicate và sử dụng two pointers
// 2. Với mỗi phần tử đầu tiên, sử dụng two pointers để tìm hai số còn lại
// 3. Tránh duplicate bằng cách skip các phần tử giống nhau
//
// Độ phức tạp: O(n²) thời gian, O(1) không gian (không tính output array)
func threeSum(nums []int) [][]int {
	n := len(nums)
	if n < 3 {
		return [][]int{}
	}

	// Sắp xếp mảng để dễ dàng xử lý duplicate và sử dụng two pointers
	sort.Ints(nums)
	result := [][]int{}

	// Duyệt qua từng phần tử làm phần tử đầu tiên
	for i := 0; i < n-2; i++ {
		// Skip duplicate cho phần tử đầu tiên
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// Two pointers: left bắt đầu từ i+1, right từ cuối mảng
		left := i + 1
		right := n - 1
		target := -nums[i] // Tổng của hai số còn lại phải bằng -nums[i]

		for left < right {
			sum := nums[left] + nums[right]

			if sum == target {
				// Tìm thấy một bộ ba hợp lệ
				result = append(result, []int{nums[i], nums[left], nums[right]})

				// Skip duplicate cho left pointer
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				// Skip duplicate cho right pointer
				for left < right && nums[right] == nums[right-1] {
					right--
				}

				// Di chuyển cả hai pointers
				left++
				right--
			} else if sum < target {
				// Tổng quá nhỏ, cần tăng left pointer
				left++
			} else {
				// Tổng quá lớn, cần giảm right pointer
				right--
			}
		}
	}

	return result
}
