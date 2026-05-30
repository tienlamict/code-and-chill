# Best Time to Buy and Sell Stock (Thời điểm tốt nhất để mua và bán cổ phiếu)

## Mô tả bài toán

Cho một mảng `prices` trong đó `prices[i]` là giá của một cổ phiếu nhất định vào ngày thứ `i`.

Bạn muốn tối đa hóa lợi nhuận của mình bằng cách chọn **một ngày duy nhất** để mua một cổ phiếu và chọn **một ngày khác trong tương lai** để bán cổ phiếu đó.

Trả về lợi nhuận tối đa bạn có thể đạt được từ giao dịch này. Nếu bạn không thể đạt được bất kỳ lợi nhuận nào, hãy trả về `0`.

### Ví dụ 1:
- **Đầu vào:** `prices = [7,1,5,3,6,4]`
- **Đầu ra:** `5`
- **Giải thích:** Mua vào ngày 2 (giá = 1) và bán vào ngày 5 (giá = 6), lợi nhuận = 6 - 1 = 5. Lưu ý rằng việc mua vào ngày 2 và bán vào ngày 1 là không được phép vì bạn phải mua trước khi bán.

### Ví dụ 2:
- **Đầu vào:** `prices = [7,6,4,3,1]`
- **Đầu ra:** `0`
- **Giải thích:** Trong trường hợp này, không có giao dịch nào được thực hiện và lợi nhuận tối đa = 0.

### Ràng buộc:
- `1 <= prices.length <= 10^5`
- `0 <= prices[i] <= 10^4`

---

## Phân tích các cách tiếp cận

### 1. Brute Force (Vét cạn)
- **Ý tưởng:** Thử tất cả các cặp (mua, bán) có thể có, trong đó ngày bán phải sau ngày mua.
- **Độ phức tạp thời gian:** O(n²) - Duyệt qua tất cả các cặp.
- **Độ phức tạp không gian:** O(1).
- **Đánh giá:** Không hiệu quả với mảng có kích thước lên đến 10^5.

### 2. One Pass (Duyệt một lần) - Tối ưu
- **Ý tưởng:** Khi chúng ta duyệt qua mảng, chúng ta chỉ cần quan tâm đến hai điều:
    1. Giá thấp nhất mà chúng ta đã thấy cho đến nay (`minPrice`).
    2. Lợi nhuận tối đa mà chúng ta có thể đạt được nếu chúng ta bán vào ngày hôm nay (`currentPrice - minPrice`).
- **Độ phức tạp thời gian:** O(n) - Duyệt qua mảng đúng một lần.
- **Độ phức tạp không gian:** O(1) - Chỉ sử dụng hai biến bổ sung.

---

## Giải thích thuật toán tối ưu

Giả sử chúng ta có mảng giá: `[7, 1, 5, 3, 6, 4]`

1. **Ngày 1:** Giá = 7. `minPrice = 7`, `maxProfit = 0`.
2. **Ngày 2:** Giá = 1. Giá hiện tại (1) < `minPrice` (7) -> Cập nhật `minPrice = 1`. `maxProfit` vẫn là 0.
3. **Ngày 3:** Giá = 5. Giá hiện tại (5) > `minPrice` (1) -> Lợi nhuận tiềm năng = 5 - 1 = 4. `maxProfit = max(0, 4) = 4`.
4. **Ngày 4:** Giá = 3. Giá hiện tại (3) > `minPrice` (1) -> Lợi nhuận tiềm năng = 3 - 1 = 2. `maxProfit = max(4, 2) = 4`.
5. **Ngày 5:** Giá = 6. Giá hiện tại (6) > `minPrice` (1) -> Lợi nhuận tiềm năng = 6 - 1 = 5. `maxProfit = max(4, 5) = 5`.
6. **Ngày 6:** Giá = 4. Giá hiện tại (4) > `minPrice` (1) -> Lợi nhuận tiềm năng = 4 - 1 = 3. `maxProfit = max(5, 3) = 5`.

**Kết quả cuối cùng:** 5.

---

## Giải thích Code

```go
func MaxProfit(prices []int) int {
    if len(prices) == 0 {
        return 0
    }

    // Khởi tạo giá thấp nhất là giá của ngày đầu tiên
    minPrice := prices[0]
    maxProfit := 0

    for i := 1; i < len(prices); i++ {
        // Nếu tìm thấy giá thấp hơn, chúng ta cập nhật minPrice
        // Vì để tối đa lợi nhuận, chúng ta muốn mua ở giá thấp nhất có thể
        if prices[i] < minPrice {
            minPrice = prices[i]
        } else {
            // Nếu giá hiện tại cao hơn minPrice, tính lợi nhuận nếu bán hôm nay
            profit := prices[i] - minPrice
            // Cập nhật lợi nhuận tối đa nếu lợi nhuận hôm nay lớn hơn
            if profit > maxProfit {
                maxProfit = profit
            }
        }
    }

    return maxProfit
}
```

---

## Độ phức tạp

- **Thời gian:** `O(n)` - Trong đó `n` là số lượng phần tử trong mảng `prices`. Chúng ta chỉ duyệt qua mảng một lần.
- **Không gian:** `O(1)` - Chúng ta chỉ sử dụng hai biến `minPrice` và `maxProfit` bất kể kích thước của mảng đầu vào.
