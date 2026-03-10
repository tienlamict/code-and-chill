# 70. Climbing Stairs

## Mô tả bài toán

Bạn đang leo cầu thang. Cần leo **n** bậc để lên đỉnh.

Mỗi lần có thể bước **1** hoặc **2** bậc. Hỏi có **bao nhiêu cách khác nhau** để leo lên đỉnh?

## Ví dụ

**Ví dụ 1:**
- Input: `n = 2`
- Output: `2`
- Giải thích: Hai cách — (1+1) hoặc (2).

**Ví dụ 2:**
- Input: `n = 3`
- Output: `3`
- Giải thích: Ba cách — (1+1+1), (1+2), (2+1).

## Ràng buộc

- `1 <= n <= 45`

## Các cách tiếp cận

### 1. Brute force (đệ quy)

Với mỗi bậc, thử bước 1 hoặc 2 rồi gọi đệ quy. Đếm số cách khi đến đúng n bậc.

- **Thời gian:** O(2^n) — cây đệ quy nhị phân.
- **Không gian:** O(n) — độ sâu stack.
- Nhược điểm: Tính trùng lặp nhiều bài con.

### 2. Đệ quy có nhớ (memoization)

Lưu số cách leo đến bậc `i` vào bảng. Mỗi giá trị chỉ tính một lần.

- **Thời gian:** O(n).
- **Không gian:** O(n).
- Ưu điểm: Tránh tính lại.

### 3. Quy hoạch động bottom-up (đã chọn)

Định nghĩa `dp[i]` = số cách leo đến bậc `i`. Công thức: `dp[i] = dp[i-1] + dp[i-2]` (bước 1 từ i-1 hoặc bước 2 từ i-2). Chỉ cần 2 biến nên tối ưu không gian O(1).

- **Thời gian:** O(n).
- **Không gian:** O(1).
- Ưu điểm: Code gọn, không dùng mảng.

## Thuật toán được chọn: DP với O(1) không gian

**Công thức:**
- `dp[1] = 1`, `dp[2] = 2`
- `dp[i] = dp[i-1] + dp[i-2]` với i ≥ 3

Đây chính là dãy Fibonacci (bắt đầu từ 1, 2 thay vì 1, 1).

**Ví dụ từng bước:** n = 5

| Bậc i | Cách leo (minh họa) | dp[i] |
|-------|---------------------|-------|
| 1     | 1                   | 1     |
| 2     | 1+1, 2              | 2     |
| 3     | 1+1+1, 1+2, 2+1     | 3     |
| 4     | ...                 | 5     |
| 5     | ...                 | 8     |

Kết quả: **8** cách.

## Giải thích code

```go
if n <= 2 {
    return n
}
prev, curr := 1, 2
```

- `n <= 2`: trả về n (1 cách với n=1, 2 cách với n=2).
- `prev` = dp[i-2], `curr` = dp[i-1] khi đang tính dp[i].

```go
for i := 3; i <= n; i++ {
    prev, curr = curr, prev+curr
}
return curr
```

- Mỗi vòng: `prev` = giá trị cũ của `curr`, `curr` = `prev + curr` (dp[i] = dp[i-1] + dp[i-2]).
- Sau vòng lặp, `curr` là dp[n].

## Độ phức tạp

- **Thời gian:** O(n) — một vòng lặp từ 3 đến n.
- **Không gian:** O(1) — chỉ dùng 2 biến.
