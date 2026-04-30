package hash_table

// topKFrequent trả về k phần tử xuất hiện nhiều nhất trong mảng nums.
//
// Sử dụng Bucket Sort (sắp xếp theo tần suất):
// - Đếm tần suất xuất hiện của mỗi phần tử bằng hash map
// - Tạo mảng bucket có độ dài n+1, bucket[i] chứa các phần tử xuất hiện đúng i lần
//   (vì tần suất tối đa là n - độ dài mảng)
// - Duyệt ngược từ bucket có tần suất cao nhất, thu thập đủ k phần tử
//
// Cách này đạt O(n) thời gian, tốt hơn yêu cầu O(n log n).
func topKFrequent(nums []int, k int) []int {
	// Bước 1: Đếm tần suất của từng phần tử
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}

	// Bước 2: Tạo bucket - bucket[i] chứa các số có tần suất = i
	// Tần suất tối đa là len(nums), nên cần len(nums)+1 bucket
	buckets := make([][]int, len(nums)+1)
	for num, count := range freq {
		buckets[count] = append(buckets[count], num)
	}

	// Bước 3: Duyệt ngược bucket từ tần suất cao đến thấp, lấy k phần tử
	result := make([]int, 0, k)
	for i := len(buckets) - 1; i >= 0 && len(result) < k; i-- {
		result = append(result, buckets[i]...)
	}

	return result[:k]
}
