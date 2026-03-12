# 198. House Robber

## Mô tả bài toán

Bạn là tên trộm chuyên nghiệp và dự định trộm các nhà dọc một con phố. Mỗi nhà có một lượng tiền nhất định. Ràng buộc duy nhất: **hai nhà liền kề có hệ thống an ninh nối với nhau** — nếu trộm cả hai nhà cạnh nhau trong cùng đêm thì cảnh sát sẽ được báo.

Cho mảng số nguyên `nums` biểu thị số tiền ở mỗi nhà, hãy trả về **tổng tiền tối đa** có thể trộm trong một đêm mà không kích hoạt cảnh sát.

## Ví dụ

**Ví dụ 1**

- **Input:** `nums = [1, 2, 3, 1]`
- **Output:** `4`
- **Giải thích:** Trộm nhà 1 (tiền = 1) và nhà 3 (tiền = 3). Tổng = 1 + 3 = 4.

**Ví dụ 2**

- **Input:** `nums = [2, 7, 9, 3, 1]`
- **Output:** `12`
- **Giải thích:** Trộm nhà 1 (2), nhà 3 (9), nhà 5 (1). Tổng = 2 + 9 + 1 = 12.

## Ràng buộc

- `1 <= nums.length <= 100`
- `0 <= nums[i] <= 400`

---

## Các cách tiếp cận

### 1. Brute force (đệ quy / thử mọi cách)

- Với mỗi nhà, thử **trộm** hoặc **bỏ qua**, đảm bảo không trộm hai nhà liền kề.
- Độ phức tạp thời gian: **O(2^n)** vì mỗi nhà có 2 lựa chọn.
- Không khả thi với n lớn.

### 2. Đệ quy có nhớ (Memoization)

- Gọi `f(i)` = tổng tiền tối đa khi xét từ nhà 0 đến nhà `i`.
- `f(i) = max(f(i-1), nums[i] + f(i-2))`: không trộm nhà `i` hoặc trộm nhà `i` (khi đó bỏ nhà `i-1`).
- Lưu kết quả đã tính để tránh gọi đệ quy trùng.
- Thời gian: **O(n)**, không gian: **O(n)** cho bảng nhớ + stack.

### 3. Quy hoạch động (DP) dạng bảng

- Định nghĩa `dp[i]` = tổng tiền tối đa khi xét từ nhà 0 đến nhà `i`.
- Công thức: `dp[i] = max(dp[i-1], nums[i] + dp[i-2])`.
- Điền bảng từ `i = 0` đến `i = n-1`.
- Thời gian: **O(n)**, không gian: **O(n)**.

### 4. DP tối ưu không gian (cách chọn trong code)

- Ở bước `i` chỉ cần `dp[i-1]` và `dp[i-2]` → chỉ cần hai biến `prev1`, `prev2`.
- Cập nhật: `curr = max(prev1, nums[i] + prev2)`, rồi `prev2, prev1 = prev1, curr`.
- Thời gian: **O(n)**, không gian: **O(1)**.

---

## Thuật toán được chọn: DP với O(1) không gian

### Ý tưởng

- Tại mỗi nhà `i` có hai lựa chọn:
  1. **Trộm nhà i** → nhận `nums[i]`, nhưng không được trộm nhà `i-1` → tổng tối đa khi kết thúc ở trạng thái “vừa trộm i” là `nums[i] + dp[i-2]`.
  2. **Bỏ qua nhà i** → tổng vẫn là tổng tối đa đến nhà `i-1`: `dp[i-1]`.
- Chọn lựa chọn tốt hơn: `dp[i] = max(dp[i-1], nums[i] + dp[i-2])`.

### Base case

- `dp[0] = nums[0]` (chỉ một nhà).
- `dp[1] = max(nums[0], nums[1])` (một hoặc hai nhà, không trộm cả hai).

### Ví dụ từng bước: `nums = [2, 7, 9, 3, 1]`

| Bước | i | prev2 (dp[i-2]) | prev1 (dp[i-1]) | Chọn: max(prev1, nums[i]+prev2) |
|------|---|-----------------|-----------------|----------------------------------|
| Khởi tạo | - | 2 | max(2,7)=7 | - |
| i=2 | 2 | 2 | 7 | max(7, 9+2)=**11** → prev1=11 |
| i=3 | 3 | 7 | 11 | max(11, 3+7)=**11** → prev1=11 |
| i=4 | 4 | 11 | 11 | max(11, 1+11)=**12** → prev1=12 |

Kết quả: **12** (trộm nhà 0, 2, 4).

---

## Giải thích code

### Khởi tạo và base case

```go
if n == 0 {
    return 0
}
if n == 1 {
    return nums[0]
}
prev2 := nums[0]
prev1 := max(nums[0], nums[1])
```

- Mảng rỗng → 0. Một nhà → trộm nhà đó.
- `prev2` đóng vai trò `dp[0]`, `prev1` đóng vai trò `dp[1]`.

### Vòng lặp DP

```go
for i := 2; i < n; i++ {
    curr := max(prev1, nums[i]+prev2)
    prev2, prev1 = prev1, curr
}
return prev1
```

- `curr` = tổng tối đa khi xét đến nhà `i`: hoặc không trộm `i` (`prev1`), hoặc trộm `i` (`nums[i] + prev2`).
- Sau mỗi bước, “trượt” cửa sổ: `prev2` = giá trị cũ của `prev1`, `prev1` = `curr`.
- Kết quả cuối cùng nằm ở `prev1` (tương đương `dp[n-1]`).

---

## Độ phức tạp

- **Thời gian:** O(n) — duyệt mảng một lần.
- **Không gian:** O(1) — chỉ dùng vài biến.

---

## Chạy test

Trong thư mục `house_robber`:

```bash
go test -v
```
