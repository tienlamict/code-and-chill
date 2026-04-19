package hash_table

// Thuật toán: Dùng hash set để theo dõi các phần tử đã gặp.
// Duyệt qua từng phần tử, nếu đã tồn tại trong set thì trả về true,
// ngược lại thêm vào set. Nếu duyệt hết mà không có trùng, trả về false.
func containsDuplicate(nums []int) bool {
	seen := make(map[int]struct{})
	for _, n := range nums {
		if _, exists := seen[n]; exists {
			return true
		}
		seen[n] = struct{}{}
	}
	return false
}
