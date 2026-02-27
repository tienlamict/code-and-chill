package array

import "sort"

// merge gộp tất cả các intervals chồng chéo và trả về mảng các intervals không chồng chéo
// bao phủ toàn bộ các intervals trong input.
//
// Mỗi interval có dạng [start, end]. Hai intervals chồng chéo nếu interval này kết thúc
// tại hoặc sau khi interval kia bắt đầu (ví dụ: [1,4] và [4,5] chồng chéo).
//
// Thuật toán Greedy:
// 1. Sắp xếp intervals theo start time (đầu mút trái)
// 2. Duyệt qua từng interval đã sắp xếp
// 3. Nếu interval hiện tại chồng chéo với interval cuối trong kết quả (start <= end của interval cuối),
//    gộp bằng cách mở rộng end của interval cuối
// 4. Nếu không chồng chéo, thêm interval hiện tại vào kết quả
//
// Độ phức tạp: O(n log n) thời gian (do sắp xếp), O(n) không gian (cho kết quả)
func merge(intervals [][]int) [][]int {
	// Xử lý trường hợp mảng rỗng
	if len(intervals) == 0 {
		return [][]int{}
	}

	// Bước 1: Sắp xếp intervals theo start time (intervals[i][0])
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// Kết quả: bắt đầu với interval đầu tiên
	result := [][]int{{intervals[0][0], intervals[0][1]}}

	// Bước 2: Duyệt qua các intervals còn lại
	for i := 1; i < len(intervals); i++ {
		last := result[len(result)-1]
		currentStart := intervals[i][0]
		currentEnd := intervals[i][1]

		// Bước 3: Kiểm tra chồng chéo
		// Hai intervals chồng chéo khi: start của interval hiện tại <= end của interval cuối
		if currentStart <= last[1] {
			// Chồng chéo: gộp bằng cách mở rộng end của interval cuối
			// End mới = max(end của interval cuối, end của interval hiện tại)
			if currentEnd > last[1] {
				last[1] = currentEnd
			}
		} else {
			// Không chồng chéo: thêm interval hiện tại vào kết quả
			result = append(result, []int{currentStart, currentEnd})
		}
	}

	return result
}
