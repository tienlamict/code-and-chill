# 322. Coin Change

## Mô tả bài toán

Cho mảng số nguyên `coins` (các mệnh giá xu) và số nguyên `amount` (tổng tiền cần đổi).

Trả về **số đồng xu ít nhất** để tạo đúng tổng `amount`. Mỗi loại xu có thể dùng **vô hạn lần**. Nếu không thể tạo được tổng đó thì trả về **-1**.

## Ví dụ

**Ví dụ 1:**
- Input: `coins = [1,2,5]`, `amount = 11`
- Output: `3`
- Giải thích: 11 = 5 + 5 + 1 (3 xu).

**Ví dụ 2:**
- Input: `coins = [2]`, `amount = 3`
- Output: `-1`
- Giải thích: Chỉ có xu 2, không thể tạo 3.

**Ví dụ 3:**
- Input: `coins = [1]`, `amount = 0`
- Output: `0`
- Giải thích: Không cần xu nào.

## Ràng buộc

- `1 <= coins.length <= 12`
- `1 <= coins[i] <= 2^31 - 1`
- `0 <= amount <= 10^4`

## Các cách tiếp cận

### 1. Brute force (DFS / backtracking)

Thử mọi cách chọn xu: với mỗi tổng còn lại, thử từng loại xu rồi gọi đệ quy. Lấy cách dùng ít xu nhất.

- **Thời gian:** O(S^n) với S = amount, n = số loại xu (cây đệ quy rất lớn).
- **Không gian:** O(S) do stack đệ quy.
- Nhược điểm: Trùng lặp nhiều bài con (cùng một tổng con được tính lại nhiều lần).

### 2. Đệ quy có nhớ (memoization)

Giống brute force nhưng lưu kết quả cho mỗi tổng con (số xu ít nhất để tạo tổng đó). Mỗi giá trị chỉ tính một lần.

- **Thời gian:** O(amount * len(coins)).
- **Không gian:** O(amount) cho bảng nhớ + O(amount) stack.
- Ưu điểm: Dễ suy luận từ bài “thử từng xu”.

### 3. Quy hoạch động bottom-up (đã chọn)

Định nghĩa `dp[i]` = số xu ít nhất để tạo tổng `i` (`dp[0] = 0`). Tính lần lượt `dp[1], dp[2], ..., dp[amount]`. Với mỗi `i`, thử từng xu `c` (c ≤ i): `dp[i] = min(dp[i], 1 + dp[i-c])`.

- **Thời gian:** O(amount * len(coins)).
- **Không gian:** O(amount), không dùng stack đệ quy.
- Ưu điểm: Code đơn giản, chạy nhanh, dễ tối ưu bộ nhớ.

## Thuật toán được chọn: DP bottom-up

**Công thức:**
- `dp[0] = 0`.
- Với `i` từ 1 đến `amount`:  
  `dp[i] = min{ 1 + dp[i - c] : c ∈ coins, c ≤ i }`.  
  Nếu không có `c` nào thỏa mãn, coi `dp[i] = +∞` (trong code dùng `amount+1`).

**Ví dụ từng bước:** `coins = [1,2,5]`, `amount = 11`

| i   | Cách tạo (min xu) | dp[i] |
|-----|-------------------|-------|
| 0   | không dùng xu     | 0     |
| 1   | 1                 | 1     |
| 2   | 2 hoặc 1+1        | 1     |
| 3   | 2+1               | 2     |
| 4   | 2+2               | 2     |
| 5   | 5                 | 1     |
| ... | ...               | ...   |
| 11  | 5+5+1             | 3     |

Cuối cùng trả về `dp[11] = 3`. Nếu `dp[amount] > amount` (vẫn là “vô cùng”) thì trả về -1.

## Giải thích code

```go
if amount == 0 {
    return 0
}
dp := make([]int, amount+1)
for i := 1; i <= amount; i++ {
    dp[i] = amount + 1  // giá trị "vô cùng"
}
```

- `amount == 0`: không cần xu nào → trả về 0.
- `dp[i] = amount + 1`: tối đa cần `amount` xu (toàn xu 1), nên dùng `amount+1` để đánh dấu “chưa có cách”.

```go
for i := 1; i <= amount; i++ {
    for _, c := range coins {
        if c <= i {
            if onePlus := 1 + dp[i-c]; onePlus < dp[i] {
                dp[i] = onePlus
            }
        }
    }
}
```

- Với mỗi tổng `i`, thử từng xu `c`. Nếu `c <= i` thì có thể dùng 1 xu `c` và còn lại tổng `i-c` → số xu là `1 + dp[i-c]`. Cập nhật `dp[i]` nếu nhỏ hơn giá trị hiện tại.

```go
if dp[amount] > amount {
    return -1
}
return dp[amount]
```

- Nếu sau khi tính mà `dp[amount]` vẫn là `amount+1` thì không có cách tạo tổng `amount` → trả về -1.

## Độ phức tạp

- **Thời gian:** O(amount × len(coins)) — hai vòng lặp lồng nhau.
- **Không gian:** O(amount) — mảng `dp` có `amount+1` phần tử.

## Lưu ý

- Bài này **không dùng được greedy** (chọn xu lớn nhất có thể). Ví dụ `coins = [1,3,4]`, `amount = 6`: greedy cho 4+1+1 = 3 xu, trong khi 3+3 = 2 xu. DP luôn cho kết quả đúng.
