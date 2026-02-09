package data_stream

import "container/heap"

// MedianFinder tìm median từ data stream sử dụng hai heaps.
//
// Sử dụng kỹ thuật Two Heaps:
// - Max heap (small): Lưu nửa nhỏ hơn hoặc bằng của dữ liệu, phần tử lớn nhất ở top
// - Min heap (large): Lưu nửa lớn hơn của dữ liệu, phần tử nhỏ nhất ở top
// - Đảm bảo: len(small) >= len(large) và len(small) - len(large) <= 1
// - Median = top của small (nếu số lượng lẻ) hoặc (top small + top large) / 2 (nếu số lượng chẵn)
//
// Độ phức tạp:
// - AddNum: O(log n) thời gian
// - FindMedian: O(1) thời gian
// - Không gian: O(n)
type MedianFinder struct {
	small *MaxHeap // Nửa nhỏ hơn hoặc bằng (max heap)
	large *MinHeap // Nửa lớn hơn (min heap)
}

// Constructor khởi tạo đối tượng MedianFinder
func Constructor() MedianFinder {
	return MedianFinder{
		small: &MaxHeap{},
		large: &MinHeap{},
	}
}

// AddNum thêm một số vào data stream
func (mf *MedianFinder) AddNum(num int) {
	// Thêm vào small heap trước
	heap.Push(mf.small, num)

	// Đảm bảo tất cả phần tử trong small <= tất cả phần tử trong large
	if mf.small.Len() > 0 && mf.large.Len() > 0 && mf.small.Peek() > mf.large.Peek() {
		// Di chuyển phần tử lớn nhất từ small sang large
		val := heap.Pop(mf.small).(int)
		heap.Push(mf.large, val)
	}

	// Đảm bảo cân bằng: len(small) >= len(large) và len(small) - len(large) <= 1
	if mf.small.Len() > mf.large.Len()+1 {
		val := heap.Pop(mf.small).(int)
		heap.Push(mf.large, val)
	} else if mf.large.Len() > mf.small.Len() {
		val := heap.Pop(mf.large).(int)
		heap.Push(mf.small, val)
	}
}

// FindMedian trả về median của tất cả các số đã thêm
func (mf *MedianFinder) FindMedian() float64 {
	if mf.small.Len() > mf.large.Len() {
		// Số lượng lẻ: median là phần tử ở giữa (top của small)
		return float64(mf.small.Peek())
	}
	// Số lượng chẵn: median là trung bình của hai phần tử giữa
	return float64(mf.small.Peek()+mf.large.Peek()) / 2.0
}

// MaxHeap là max heap (phần tử lớn nhất ở top)
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] } // Đảo dấu để có max heap
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h MaxHeap) Peek() int {
	return h[0]
}

// MinHeap là min heap (phần tử nhỏ nhất ở top)
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h MinHeap) Peek() int {
	return h[0]
}
