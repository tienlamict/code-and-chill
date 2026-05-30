package best_time_to_buy_and_sell_stock

// MaxProfit tìm lợi nhuận tối đa từ việc mua và bán cổ phiếu
// prices: mảng giá cổ phiếu theo từng ngày
// Trả về lợi nhuận tối đa có thể đạt được
func MaxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	// minPrice lưu giá thấp nhất đã gặp cho đến hiện tại
	minPrice := prices[0]
	// maxProfit lưu lợi nhuận tối đa tìm được
	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		// Nếu giá hiện tại nhỏ hơn minPrice, cập nhật minPrice
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else {
			// Nếu giá hiện tại lớn hơn minPrice, tính lợi nhuận tiềm năng
			// và cập nhật maxProfit nếu lợi nhuận này lớn hơn maxProfit hiện tại
			profit := prices[i] - minPrice
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}

	return maxProfit
}
