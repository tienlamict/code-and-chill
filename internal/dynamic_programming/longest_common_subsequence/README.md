# Longest Common Subsequence

## Mô tả bài toán

Cho hai chuỗi `text1` và `text2`, trả về độ dài của **dãy con chung dài nhất (LCS)**. Nếu không có dãy con chung, trả về `0`.

**Dãy con (subsequence)** của một chuỗi là chuỗi mới được tạo ra bằng cách xóa một số ký tự (có thể không xóa) mà không thay đổi thứ tự tương đối của các ký tự còn lại.

Ví dụ: `"ace"` là dãy con của `"abcde"`.

**Dãy con chung** là dãy con xuất hiện trong cả hai chuỗi.

---

## Ví dụ minh họa

**Ví dụ 1:**
```
Input:  text1 = "abcde", text2 = "ace"
Output: 3
Giải thích: Dãy con chung dài nhất là "ace", độ dài 3.
```

**Ví dụ 2:**
```
Input:  text1 = "abc", text2 = "abc"
Output: 3
Giải thích: Dãy con chung dài nhất là "abc", độ dài 3.
```

**Ví dụ 3:**
```
Input:  text1 = "abc", text2 = "def"
Output: 0
Giải thích: Không có dãy con chung, trả về 0.
```

---

## Ràng buộc

- `1 <= text1.length, text2.length <= 1000`
- `text1` và `text2` chỉ gồm các chữ cái viết thường.

---

## Phân tích các cách tiếp cận

### 1. Brute Force (đệ quy thuần túy)

Duyệt qua tất cả các dãy con của `text1` và kiểm tra xem nó có là dãy con của `text2` không.

- **Độ phức tạp thời gian:** O(2^m * n) — với m = len(text1), n = len(text2)
- **Độ phức tạp không gian:** O(m) — stack đệ quy
- **Nhược điểm:** Quá chậm, không khả thi với chuỗi dài.

---

### 2. Đệ quy + Memoization (Top-down DP)

Định nghĩa hàm `lcs(i, j)` = LCS của `text1[i:]` và `text2[j:]`.

```
lcs(i, j) =
  lcs(i+1, j+1) + 1          nếu text1[i] == text2[j]
  max(lcs(i+1, j), lcs(i, j+1))  ngược lại
```

Lưu kết quả vào bảng memo để tránh tính lại.

- **Độ phức tạp thời gian:** O(m * n)
- **Độ phức tạp không gian:** O(m * n) — bảng memo + stack đệ quy

---

### 3. Dynamic Programming Bottom-up 2D ✅ (Được chọn)

Xây dựng bảng `dp[i][j]` = LCS của `text1[0..i-1]` và `text2[0..j-1]` theo thứ tự từ dưới lên.

- **Độ phức tạp thời gian:** O(m * n)
- **Độ phức tạp không gian:** O(m * n)
- **Ưu điểm:** Không có overhead của đệ quy, dễ hiểu và triển khai.

---

### 4. DP tối ưu không gian (1D)

Chỉ cần lưu 2 hàng (`prev` và `curr`) thay vì toàn bộ bảng 2D.

- **Độ phức tạp thời gian:** O(m * n)
- **Độ phức tạp không gian:** O(min(m, n))

---

## Giải thích chi tiết thuật toán DP Bottom-up

### Định nghĩa trạng thái

```
dp[i][j] = độ dài LCS của text1[0..i-1] và text2[0..j-1]
```

### Trạng thái ban đầu

```
dp[0][j] = 0   (text1 rỗng, không có LCS)
dp[i][0] = 0   (text2 rỗng, không có LCS)
```

### Công thức chuyển trạng thái

```
Nếu text1[i-1] == text2[j-1]:
    dp[i][j] = dp[i-1][j-1] + 1    ← ký tự khớp, mở rộng LCS

Ngược lại:
    dp[i][j] = max(dp[i-1][j], dp[i][j-1])   ← bỏ ký tự cuối của text1 hoặc text2
```

### Ví dụ từng bước: text1 = "abcde", text2 = "ace"

Bảng DP (hàng = text1, cột = text2):

```
     ""  a   c   e
""  [ 0   0   0   0 ]
a   [ 0   1   1   1 ]
b   [ 0   1   1   1 ]
c   [ 0   1   2   2 ]
d   [ 0   1   2   2 ]
e   [ 0   1   2   3 ]
```

- `dp[1][1]`: 'a' == 'a' → dp[0][0] + 1 = 1
- `dp[3][2]`: 'c' == 'c' → dp[2][1] + 1 = 2
- `dp[5][3]`: 'e' == 'e' → dp[4][2] + 1 = 3

Kết quả: `dp[5][3] = 3`

---

## Giải thích code

```go
dp := make([][]int, m+1)
for i := range dp {
    dp[i] = make([]int, n+1)
}
```
Khởi tạo bảng dp kích thước (m+1) x (n+1), mặc định là 0 (base case).

```go
if text1[i-1] == text2[j-1] {
    dp[i][j] = dp[i-1][j-1] + 1
}
```
Nếu ký tự tại vị trí `i` của text1 khớp với ký tự tại vị trí `j` của text2, LCS tăng thêm 1 so với LCS của phần trước đó.

```go
} else {
    if dp[i-1][j] > dp[i][j-1] {
        dp[i][j] = dp[i-1][j]
    } else {
        dp[i][j] = dp[i][j-1]
    }
}
```
Nếu không khớp, lấy giá trị lớn hơn giữa: bỏ ký tự cuối text1 (`dp[i-1][j]`) hoặc bỏ ký tự cuối text2 (`dp[i][j-1]`).

---

## Phân tích độ phức tạp

| | Thời gian | Không gian |
|---|---|---|
| Brute Force | O(2^m * n) | O(m) |
| Top-down DP | O(m * n) | O(m * n) |
| **Bottom-up DP** | **O(m * n)** | **O(m * n)** |
| DP tối ưu 1D | O(m * n) | O(min(m, n)) |

Với m, n ≤ 1000: bảng dp có tối đa 10^6 ô — hoàn toàn khả thi.
