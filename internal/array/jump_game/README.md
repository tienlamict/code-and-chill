# Jump Game

## Mô tả bài toán

Bạn được cho một mảng số nguyên `nums`. Bạn bắt đầu ở vị trí đầu tiên của mảng, và mỗi phần tử trong mảng đại diện cho độ dài bước nhảy tối đa tại vị trí đó.

Trả về `true` nếu bạn có thể đến được vị trí cuối cùng, hoặc `false` nếu không thể.

**Ví dụ:**

### Example 1:
- Input: `nums = [2,3,1,1,4]`
- Output: `true`
- Giải thích: Nhảy 1 bước từ index 0 đến 1, sau đó nhảy 3 bước đến index cuối cùng.

### Example 2:
- Input: `nums = [3,2,1,0,4]`
- Output: `false`
- Giải thích: Bạn sẽ luôn đến được index 3 bất kể làm gì. Độ dài bước nhảy tối đa của nó là 0, khiến không thể đến được index cuối cùng.

**Ràng buộc:**
- `1 <= nums.length <= 10^4`
- `0 <= nums[i] <= 10^5`

## Phân tích thuật toán

### Cách tiếp cận 1: Dynamic Programming (Bottom-up)

**Ý tưởng:**
- Tạo mảng `dp[]` trong đó `dp[i]` đại diện cho việc có thể đến được vị trí `i` hay không
- Khởi tạo: `dp[0] = true` (bắt đầu ở vị trí 0)
- Với mỗi vị trí `i`, nếu `dp[i] = true`, đánh dấu tất cả các vị trí có thể nhảy từ `i` là `true`
- Trả về `dp[n-1]`

**Độ phức tạp:**
- Thời gian: O(n²) - với mỗi vị trí, có thể cần cập nhật nhiều vị trí khác
- Không gian: O(n) - mảng `dp[]`

**Nhược điểm:** Không tối ưu, có thể tối ưu hơn bằng cách sử dụng Greedy.

### Cách tiếp cận 2: Greedy (Tối ưu)

**Ý tưởng:**
- Thay vì lưu trữ tất cả các vị trí có thể đến được, chỉ cần theo dõi vị trí xa nhất có thể đạt được
- Duyệt qua mảng từ trái sang phải
- Tại mỗi vị trí `i`, cập nhật vị trí xa nhất có thể đạt được: `maxReach = max(maxReach, i + nums[i])`
- Nếu tại bất kỳ vị trí nào, `i > maxReach`, nghĩa là không thể đến được vị trí `i`, do đó không thể đến được vị trí cuối cùng
- Nếu `maxReach >= n-1`, có thể đến được vị trí cuối cùng

**Tại sao Greedy hoạt động?**
- Nếu có thể đến được vị trí `i`, thì có thể đến được tất cả các vị trí từ 0 đến `i`
- Chỉ cần quan tâm đến vị trí xa nhất có thể đạt được, không cần quan tâm đến cách đến đó
- Nếu tại vị trí `i`, có thể nhảy đến `i + nums[i]`, và `i + nums[i] > maxReach`, thì cập nhật `maxReach`
- Nếu `i > maxReach`, nghĩa là không có cách nào đến được vị trí `i`, do đó không thể đến được vị trí cuối cùng

**Độ phức tạp:**
- Thời gian: O(n) - duyệt qua mảng một lần
- Không gian: O(1) - chỉ sử dụng một biến `maxReach`

**Ưu điểm:**
- Tối ưu về thời gian và không gian
- Dễ hiểu và dễ implement
- Không cần lưu trữ thêm mảng

## Giải thích chi tiết thuật toán Greedy

### Ví dụ 1: `nums = [2,3,1,1,4]`

```
Vị trí:  0  1  2  3  4
nums:    2  3  1  1  4
```

**Bước 1:** i = 0, nums[0] = 2
- maxReach = 0
- i (0) <= maxReach (0) ✓
- currentReach = 0 + 2 = 2
- maxReach = max(0, 2) = 2
- maxReach (2) < n-1 (4), tiếp tục

**Bước 2:** i = 1, nums[1] = 3
- maxReach = 2
- i (1) <= maxReach (2) ✓
- currentReach = 1 + 3 = 4
- maxReach = max(2, 4) = 4
- maxReach (4) >= n-1 (4) ✓
- **Trả về true**

### Ví dụ 2: `nums = [3,2,1,0,4]`

```
Vị trí:  0  1  2  3  4
nums:    3  2  1  0  4
```

**Bước 1:** i = 0, nums[0] = 3
- maxReach = 0
- i (0) <= maxReach (0) ✓
- currentReach = 0 + 3 = 3
- maxReach = max(0, 3) = 3
- maxReach (3) < n-1 (4), tiếp tục

**Bước 2:** i = 1, nums[1] = 2
- maxReach = 3
- i (1) <= maxReach (3) ✓
- currentReach = 1 + 2 = 3
- maxReach = max(3, 3) = 3
- maxReach (3) < n-1 (4), tiếp tục

**Bước 3:** i = 2, nums[2] = 1
- maxReach = 3
- i (2) <= maxReach (3) ✓
- currentReach = 2 + 1 = 3
- maxReach = max(3, 3) = 3
- maxReach (3) < n-1 (4), tiếp tục

**Bước 4:** i = 3, nums[3] = 0
- maxReach = 3
- i (3) <= maxReach (3) ✓
- currentReach = 3 + 0 = 3
- maxReach = max(3, 3) = 3
- maxReach (3) < n-1 (4), tiếp tục

**Bước 5:** i = 4, nums[4] = 4
- maxReach = 3
- i (4) > maxReach (3) ✗
- **Trả về false**

## Giải thích code

### Hàm `canJump`

```go
func canJump(nums []int) bool {
	maxReach := 0
```

Khởi tạo biến `maxReach` để lưu vị trí xa nhất có thể đạt được. Ban đầu là 0 vì bắt đầu ở vị trí 0.

```go
	for i := 0; i < len(nums); i++ {
		if i > maxReach {
			return false
		}
```

Duyệt qua từng vị trí trong mảng. Nếu vị trí hiện tại `i` vượt quá `maxReach`, nghĩa là không thể đến được vị trí này, do đó không thể đến được vị trí cuối cùng. Trả về `false` ngay lập tức.

```go
		currentReach := i + nums[i]
		if currentReach > maxReach {
			maxReach = currentReach
		}
```

Tính vị trí xa nhất có thể đạt được từ vị trí hiện tại: `i + nums[i]`. Nếu vị trí này xa hơn `maxReach`, cập nhật `maxReach`.

```go
		if maxReach >= len(nums)-1 {
			return true
		}
	}
```

Nếu `maxReach` đã đạt được hoặc vượt quá vị trí cuối cùng (index `len(nums)-1`), trả về `true` ngay lập tức vì đã chắc chắn có thể đến được vị trí cuối cùng.

```go
	return maxReach >= len(nums)-1
}
```

Sau khi duyệt xong, kiểm tra lại một lần nữa (trường hợp đặc biệt) và trả về kết quả.

## Phân tích độ phức tạp

### Độ phức tạp thời gian: O(n)

- Duyệt qua mảng một lần: O(n)
- Mỗi lần lặp thực hiện các phép toán O(1): so sánh, cập nhật biến
- Tổng cộng: O(n)

### Độ phức tạp không gian: O(1)

- Chỉ sử dụng một biến `maxReach`: O(1)
- Không sử dụng thêm mảng hoặc cấu trúc dữ liệu phụ
- Tổng cộng: O(1)

## So sánh với các cách tiếp cận khác

| Cách tiếp cận | Thời gian | Không gian | Ghi chú |
|--------------|-----------|------------|---------|
| Brute Force (DFS) | O(2^n) | O(n) | Thử tất cả các cách nhảy có thể |
| Dynamic Programming | O(n²) | O(n) | Lưu trữ kết quả cho từng vị trí |
| **Greedy** | **O(n)** | **O(1)** | **Tối ưu nhất** |

## Follow-up Questions

### 1. Có thể tối ưu hơn không?

Không, thuật toán Greedy đã tối ưu:
- Thời gian: O(n) - phải duyệt qua ít nhất một lần
- Không gian: O(1) - chỉ cần một biến

### 2. Nếu cần tìm số bước nhảy tối thiểu thì sao?

Đây là bài toán **Jump Game II** (LeetCode 45). Cần sử dụng Greedy với cách tiếp cận khác:
- Theo dõi số bước nhảy hiện tại và vị trí xa nhất có thể đạt được với số bước đó
- Khi đến giới hạn của số bước hiện tại, tăng số bước và cập nhật giới hạn mới

### 3. Nếu có thể nhảy cả về phía trước và phía sau?

Đây là biến thể khác của bài toán. Cần sử dụng BFS hoặc DFS để tìm đường đi, vì có thể có chu trình.

## Kết luận

Bài toán Jump Game là một ví dụ điển hình về việc sử dụng thuật toán Greedy để giải quyết bài toán tối ưu. Thay vì lưu trữ tất cả các trạng thái có thể, chỉ cần theo dõi vị trí xa nhất có thể đạt được, giúp giảm độ phức tạp từ O(n²) xuống O(n) về thời gian và từ O(n) xuống O(1) về không gian.
