# 338. Counting Bits (Đếm số bit 1)

## Mô tả bài toán
Cho một số nguyên `n`, hãy trả về một mảng `ans` có độ dài `n + 1` sao cho với mỗi `i` (`0 <= i <= n`), `ans[i]` là số lượng số **1** trong biểu diễn nhị phân của `i`.

## Ví dụ
### Ví dụ 1:
- **Input:** `n = 2`
- **Output:** `[0,1,1]`
- **Giải thích:**
  - 0 --> 0 (0 bit 1)
  - 1 --> 1 (1 bit 1)
  - 2 --> 10 (1 bit 1)

### Ví dụ 2:
- **Input:** `n = 5`
- **Output:** `[0,1,1,2,1,2]`
- **Giải thích:**
  - 0 --> 0 (0 bit 1)
  - 1 --> 1 (1 bit 1)
  - 2 --> 10 (1 bit 1)
  - 3 --> 11 (2 bit 1)
  - 4 --> 100 (1 bit 1)
  - 5 --> 101 (2 bit 1)

## Ràng buộc
- `0 <= n <= 10^5`

## Phân tích các cách tiếp cận

### 1. Brute Force (Duyệt qua từng số và đếm bit)
- **Cách làm:** Duyệt từ 0 đến n, với mỗi số sử dụng vòng lặp hoặc hàm có sẵn để đếm số lượng bit 1.
- **Độ phức tạp:**
  - Thời gian: $O(n \log n)$ vì mỗi số nguyên `i` có tối đa $\log i$ bit.
  - Không gian: $O(1)$ (không tính mảng kết quả).

### 2. Quy hoạch động (Dynamic Programming) - Tối ưu
Có nhiều công thức quy hoạch động để giải bài này trong $O(n)$.
- **Công thức:** $P(i) = P(i >> 1) + (i \& 1)$
- **Giải thích:** 
  - `i >> 1` là số `i` sau khi bỏ đi bit cuối cùng.
  - `i & 1` là giá trị của bit cuối cùng (0 hoặc 1).
  - Số lượng bit 1 của `i` bằng số lượng bit 1 của số `i/2` cộng thêm 1 nếu bit cuối cùng của `i` là 1.
- **Độ phức tạp:**
  - Thời gian: $O(n)$ (duyệt 1 vòng duy nhất).
  - Không gian: $O(1)$ (không tính mảng kết quả).

## Chi tiết thuật toán được chọn
Chúng ta chọn thuật toán Quy hoạch động vì nó đáp ứng được yêu cầu Follow-up (thời gian tuyến tính $O(n)$ và không dùng hàm có sẵn).

### Ví dụ từng bước (với n = 5):
1. Khởi tạo `ans = [0, 0, 0, 0, 0, 0]`
2. `i = 1`: `ans[1] = ans[1 >> 1] + (1 & 1) = ans[0] + 1 = 0 + 1 = 1`. `ans = [0, 1, 0, 0, 0, 0]`
3. `i = 2`: `ans[2] = ans[2 >> 1] + (2 & 1) = ans[1] + 0 = 1 + 0 = 1`. `ans = [0, 1, 1, 0, 0, 0]`
4. `i = 3`: `ans[3] = ans[3 >> 1] + (3 & 1) = ans[1] + 1 = 1 + 1 = 2`. `ans = [0, 1, 1, 2, 0, 0]`
5. `i = 4`: `ans[4] = ans[4 >> 1] + (4 & 1) = ans[2] + 0 = 1 + 0 = 1`. `ans = [0, 1, 1, 2, 1, 0]`
6. `i = 5`: `ans[5] = ans[5 >> 1] + (5 & 1) = ans[2] + 1 = 1 + 1 = 2`. `ans = [0, 1, 1, 2, 1, 2]`

## Giải thích code
```go
func countBits(n int) []int {
	ans := make([]int, n+1) // Tạo mảng kết quả n+1 phần tử
	for i := 1; i <= n; i++ {
		// Dùng kết quả đã tính trước đó của i/2 (i >> 1)
		// và cộng thêm bit cuối cùng của i (i & 1)
		ans[i] = ans[i>>1] + (i & 1)
	}
	return ans
}
```

## Phân tích độ phức tạp
- **Thời gian:** $O(n)$, vì chúng ta chỉ duyệt qua các số từ 1 đến n đúng một lần.
- **Không gian:** $O(1)$, nếu không tính mảng kết quả `ans`. Nếu tính cả mảng kết quả thì là $O(n)$.

## Trả lời Follow-up
- **Linear time O(n) and single pass?** Có, thuật toán trên thực hiện trong $O(n)$ và chỉ một vòng lặp duy nhất.
- **Without built-in function?** Có, chúng ta chỉ sử dụng các phép toán bit cơ bản như `>>` và `&`.
