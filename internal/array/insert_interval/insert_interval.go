package array

// insert chèn một interval mới vào mảng intervals đã được sắp xếp và không chồng chéo,
// gộp các intervals chồng chéo nếu cần thiết, và trả về mảng intervals sau khi chèn.
//
// intervals đã được sắp xếp theo start time và không chồng chéo.
// newInterval = [start, end] là interval cần chèn.
//
// Thuật toán Greedy:
// 1. Thêm tất cả intervals có end < newInterval.start (đứng trước newInterval)
// 2. Gộp tất cả intervals chồng chéo với newInterval:
//    - Interval chồng chéo khi: interval.end >= newInterval.start và interval.start <= newInterval.end
//    - Khi gộp: start = min(start của tất cả intervals chồng chéo), end = max(end của tất cả intervals chồng chéo)
// 3. Thêm tất cả intervals có start > mergedInterval.end (đứng sau newInterval)
//
// Độ phức tạp: O(n) thời gian, O(n) không gian
func insert(intervals [][]int, newInterval []int) [][]int {
	result := [][]int{}
	i := 0
	n := len(intervals)
	newStart := newInterval[0]
	newEnd := newInterval[1]

	// Bước 1: Thêm tất cả intervals đứng trước newInterval
	// (intervals có end < newInterval.start)
	for i < n && intervals[i][1] < newStart {
		result = append(result, intervals[i])
		i++
	}

	// Bước 2: Gộp tất cả intervals chồng chéo với newInterval
	// Interval chồng chéo khi: interval.start <= newInterval.end
	// Trong khi gộp, cập nhật newStart và newEnd để bao phủ tất cả intervals chồng chéo
	for i < n && intervals[i][0] <= newEnd {
		// Cập nhật start và end của interval được gộp
		if intervals[i][0] < newStart {
			newStart = intervals[i][0]
		}
		if intervals[i][1] > newEnd {
			newEnd = intervals[i][1]
		}
		i++
	}

	// Thêm interval đã được gộp vào kết quả
	result = append(result, []int{newStart, newEnd})

	// Bước 3: Thêm tất cả intervals đứng sau newInterval
	// (intervals có start > mergedInterval.end)
	for i < n {
		result = append(result, intervals[i])
		i++
	}

	return result
}
