package array

// search tìm vị trí của target trong mảng đã được rotate.
//
// Sử dụng Binary Search với điều chỉnh:
// - Mặc dù mảng đã bị rotate, một nửa của mảng luôn được sắp xếp
// - Xác định nửa nào được sắp xếp bằng cách so sánh nums[mid] với nums[left]
// - Nếu nửa trái được sắp xếp:
//   - Kiểm tra xem target có nằm trong khoảng [left, mid] không
//   - Nếu có, tìm ở nửa trái; ngược lại, tìm ở nửa phải
// - Nếu nửa phải được sắp xếp:
//   - Kiểm tra xem target có nằm trong khoảng [mid, right] không
//   - Nếu có, tìm ở nửa phải; ngược lại, tìm ở nửa trái
//
// Độ phức tạp: O(log n) thời gian, O(1) không gian
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2

		// Tìm thấy target
		if nums[mid] == target {
			return mid
		}

		// Xác định nửa nào được sắp xếp
		if nums[left] <= nums[mid] {
			// Nửa trái [left, mid] được sắp xếp
			if nums[left] <= target && target < nums[mid] {
				// Target nằm trong nửa trái đã sắp xếp
				right = mid - 1
			} else {
				// Target nằm trong nửa phải (có thể bị rotate)
				left = mid + 1
			}
		} else {
			// Nửa phải [mid, right] được sắp xếp
			if nums[mid] < target && target <= nums[right] {
				// Target nằm trong nửa phải đã sắp xếp
				left = mid + 1
			} else {
				// Target nằm trong nửa trái (có thể bị rotate)
				right = mid - 1
			}
		}
	}

	// Không tìm thấy target
	return -1
}
