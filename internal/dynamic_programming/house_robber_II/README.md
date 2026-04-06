# 213. House Robber II

## Mô tả bài toán

Bạn là một tên trộm chuyên nghiệp đang lên kế hoạch trộm các ngôi nhà dọc theo một con phố. Mỗi ngôi nhà có một lượng tiền nhất định. Tất cả các ngôi nhà tại đây được sắp xếp theo một **vòng tròn**. Điều đó có nghĩa là ngôi nhà đầu tiên là hàng xóm của ngôi nhà cuối cùng.

Trong khi đó, các ngôi nhà liền kề có hệ thống an ninh kết nối với nhau, và nó sẽ tự động liên lạc với cảnh sát nếu hai ngôi nhà liền kề bị đột nhập trong cùng một đêm.

Cho một mảng số nguyên `nums` đại diện cho số tiền của mỗi ngôi nhà, hãy trả về số tiền tối đa bạn có thể trộm được đêm nay mà không báo động cảnh sát.

### Ví dụ 1:
- **Đầu vào:** `nums = [2,3,2]`
- **Đầu ra:** `3`
- **Giải thích:** Bạn không thể trộm nhà thứ 1 (tiền = 2) và sau đó trộm nhà thứ 3 (tiền = 2), vì chúng là hàng xóm của nhau.

### Ví dụ 2:
- **Đầu vào:** `nums = [1,2,3,1]`
- **Đầu ra:** `4`
- **Giải thích:** Trộm nhà 1 (tiền = 1) và sau đó trộm nhà 3 (tiền = 3). Tổng tiền = 1 + 3 = 4.

### Ví dụ 3:
- **Đầu vào:** `nums = [1,2,3]`
- **Đầu ra:** `3`

### Ràng buộc:
- `1 <= nums.length <= 100`
- `0 <= nums[i] <= 1000`

---

## Phân tích cách tiếp cận

### 1. Sự khác biệt so với House Robber I
Trong House Robber I, các ngôi nhà nằm trên một đường thẳng. Ở bài này, các ngôi nhà nằm trên một vòng tròn. Điểm khác biệt mấu chốt là: **Nhà đầu tiên và nhà cuối cùng là hàng xóm của nhau.**

Điều này dẫn đến 3 khả năng:
1. Trộm nhà đầu tiên, nhưng không thể trộm nhà cuối cùng.
2. Trộm nhà cuối cùng, nhưng không thể trộm nhà đầu tiên.
3. Không trộm cả hai nhà này.

Thực tế, trường hợp 3 đã nằm trong trường hợp 1 hoặc 2 (khi chúng ta bỏ qua việc trộm một trong hai nhà).

### 2. Thuật toán tối ưu: Chia để trị + Quy hoạch động
Chúng ta có thể chia bài toán vòng tròn thành hai bài toán đường thẳng:
- **Trường hợp A:** Xét dãy nhà từ chỉ số `0` đến `n-2` (loại bỏ nhà cuối cùng).
- **Trường hợp B:** Xét dãy nhà từ chỉ số `1` đến `n-1` (loại bỏ nhà đầu tiên).

Kết quả cuối cùng sẽ là giá trị lớn nhất của hai trường hợp trên: `max(robLinear(nums[0:n-1]), robLinear(nums[1:n]))`.

#### Công thức Quy hoạch động (cho dãy thẳng):
- Gọi `dp[i]` là số tiền tối đa thu được khi xét đến nhà thứ `i`.
- Tại nhà `i`, ta có 2 lựa chọn:
    - Trộm nhà `i`: `nums[i] + dp[i-2]`
    - Không trộm nhà `i`: `dp[i-1]`
- Công thức: `dp[i] = max(dp[i-1], nums[i] + dp[i-2])`

#### Tối ưu không gian:
Vì `dp[i]` chỉ phụ thuộc vào `dp[i-1]` và `dp[i-2]`, ta chỉ cần 2 biến để lưu trữ giá trị thay vì cả một mảng.

---

## Giải thích Code

```go
func rob(nums []int) int {
    n := len(nums)
    if n == 1 {
        return nums[0]
    }

    // Chia làm 2 trường hợp: 
    // 1. Không xét nhà cuối
    // 2. Không xét nhà đầu
    case1 := robLinear(nums[:n-1])
    case2 := robLinear(nums[1:])

    return max(case1, case2)
}
```

Hàm trợ giúp giải quyết bài toán đường thẳng:
```go
func robLinear(nums []int) int {
    prev2, prev1 := 0, 0
    for _, num := range nums {
        // Tương đương: dp[i] = max(dp[i-1], num + dp[i-2])
        curr := max(prev1, num + prev2)
        prev2 = prev1
        prev1 = curr
    }
    return prev1
}
```

---

## Độ phức tạp

- **Thời gian:** $O(n)$
    - Chúng ta duyệt qua mảng 2 lần (mỗi lần gần như toàn bộ mảng).
- **Không gian:** $O(1)$
    - Chỉ sử dụng một vài biến hỗ trợ, không sử dụng mảng phụ thuộc vào kích thước đầu vào.
