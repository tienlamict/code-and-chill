package array

// findMin tìm phần tử nhỏ nhất trong một mảng đã được sắp xếp và xoay.
//
// Sử dụng Binary Search:
// - Mảng rotated có đặc điểm: một phần bên trái và một phần bên phải đều được sắp xếp
// - Phần tử nhỏ nhất nằm ở điểm "pivot" (điểm xoay)
// - So sánh nums[mid] với nums[right] để xác định phần nào chứa minimum
//   - Nếu nums[mid] < nums[right]: minimum nằm ở bên trái (bao gồm mid)
//   - Nếu nums[mid] > nums[right]: minimum nằm ở bên phải (sau mid)
//
// Độ phức tạp: O(log n) thời gian, O(1) không gian
func findMin(nums []int) int {
	left := 0
	right := len(nums) - 1

	// Binary search
	for left < right {
		mid := left + (right-left)/2

		// So sánh nums[mid] với nums[right]
		// Nếu nums[mid] < nums[right], phần bên phải đã được sắp xếp
		// → minimum nằm ở bên trái (bao gồm mid)
		if nums[mid] < nums[right] {
			right = mid
		} else {
			// Nếu nums[mid] > nums[right], phần bên trái đã được sắp xếp
			// → minimum nằm ở bên phải (sau mid)
			left = mid + 1
		}
	}

	// Khi left == right, ta đã tìm thấy minimum
	return nums[left]
}
