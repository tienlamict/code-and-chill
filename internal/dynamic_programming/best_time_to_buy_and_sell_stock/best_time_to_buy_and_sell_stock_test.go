package best_time_to_buy_and_sell_stock

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name     string
		prices   []int
		expected int
	}{
		{
			name:     "Ví dụ 1: Có lợi nhuận",
			prices:   []int{7, 1, 5, 3, 6, 4},
			expected: 5,
		},
		{
			name:     "Ví dụ 2: Không có lợi nhuận (giá giảm dần)",
			prices:   []int{7, 6, 4, 3, 1},
			expected: 0,
		},
		{
			name:     "Mảng rỗng",
			prices:   []int{},
			expected: 0,
		},
		{
			name:     "Một phần tử",
			prices:   []int{1},
			expected: 0,
		},
		{
			name:     "Hai phần tử tăng dần",
			prices:   []int{1, 5},
			expected: 4,
		},
		{
			name:     "Hai phần tử giảm dần",
			prices:   []int{5, 1},
			expected: 0,
		},
		{
			name:     "Giá bằng nhau",
			prices:   []int{3, 3, 3, 3},
			expected: 0,
		},
		{
			name:     "Lợi nhuận ở cuối",
			prices:   []int{2, 1, 2, 0, 1},
			expected: 1,
		},
		{
			name:     "Giá trị lớn",
			prices:   []int{10000, 0, 10000},
			expected: 10000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := MaxProfit(tt.prices)
			if actual != tt.expected {
				t.Errorf("MaxProfit(%v) = %d; expected %d", tt.prices, actual, tt.expected)
			}
		})
	}
}
