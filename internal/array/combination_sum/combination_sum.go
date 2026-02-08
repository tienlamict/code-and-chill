package array

import "sort"

// combinationSum tìm tất cả các tổ hợp duy nhất của candidates sao cho tổng bằng target.
//
// Sử dụng Backtracking:
// 1. Sắp xếp mảng để dễ dàng pruning và tránh duplicate
// 2. Với mỗi số, thử thêm vào tổ hợp hiện tại
// 3. Mỗi số có thể được sử dụng nhiều lần
// 4. Tránh duplicate bằng cách chỉ xét các số từ vị trí hiện tại trở đi
// 5. Pruning: nếu tổng hiện tại > target, dừng nhánh đó
//
// Độ phức tạp: O(2^target) thời gian trong trường hợp xấu nhất, O(target) không gian (cho call stack)
func combinationSum(candidates []int, target int) [][]int {
	// Sắp xếp mảng để dễ dàng pruning
	sort.Ints(candidates)

	result := [][]int{}
	current := []int{}

	// Backtracking function
	var backtrack func(int, int)
	backtrack = func(start int, remaining int) {
		// Base case: đã đạt target
		if remaining == 0 {
			// Tạo bản sao của current và thêm vào result
			combination := make([]int, len(current))
			copy(combination, current)
			result = append(result, combination)
			return
		}

		// Duyệt qua các số từ start đến cuối
		for i := start; i < len(candidates); i++ {
			// Pruning: nếu số hiện tại lớn hơn remaining, không cần xét tiếp
			// (vì mảng đã được sắp xếp, các số sau sẽ còn lớn hơn)
			if candidates[i] > remaining {
				break
			}

			// Thêm số hiện tại vào tổ hợp
			current = append(current, candidates[i])

			// Đệ quy: tiếp tục với số hiện tại (có thể dùng lại)
			// remaining giảm đi candidates[i]
			backtrack(i, remaining-candidates[i])

			// Backtrack: xóa số vừa thêm để thử số khác
			current = current[:len(current)-1]
		}
	}

	// Bắt đầu backtracking từ index 0
	backtrack(0, target)

	return result
}
