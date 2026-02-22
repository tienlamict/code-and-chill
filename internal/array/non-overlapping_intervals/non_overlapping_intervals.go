package array

import "sort"

// eraseOverlapIntervals trả về số lượng intervals tối thiểu cần xóa để các intervals còn lại không overlap.
//
// Sử dụng Greedy Algorithm với Interval Scheduling:
// - Sắp xếp intervals theo end time (kết thúc sớm nhất trước)
// - Duyệt qua từng interval và giữ lại interval có end time nhỏ nhất khi có conflict
// - Khi hai intervals overlap, luôn xóa interval có end time lớn hơn
//   (vì giữ interval kết thúc sớm hơn sẽ để lại nhiều không gian cho các intervals sau)
//
// Lý do sắp xếp theo end time:
// - Interval kết thúc sớm hơn sẽ để lại nhiều không gian cho các intervals sau
// - Đây là chiến lược tối ưu để giữ lại nhiều intervals nhất có thể
//
// Độ phức tạp: O(n log n) thời gian, O(1) không gian
// với n là số lượng intervals
func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	// Sắp xếp intervals theo end time (intervals[i][1])
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	// Giữ lại interval đầu tiên (có end time nhỏ nhất)
	// endTime là end time của interval cuối cùng được giữ lại
	endTime := intervals[0][1]
	removed := 0

	// Duyệt qua các intervals còn lại
	for i := 1; i < len(intervals); i++ {
		// Nếu interval hiện tại overlap với interval đã giữ lại
		// (start của interval hiện tại < end của interval đã giữ lại)
		if intervals[i][0] < endTime {
			// Xóa interval hiện tại (vì nó có end time lớn hơn)
			removed++
		} else {
			// Giữ lại interval hiện tại và cập nhật endTime
			endTime = intervals[i][1]
		}
	}

	return removed
}
