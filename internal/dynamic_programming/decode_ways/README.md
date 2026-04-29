# 91. Decode Ways

## Mô tả bài toán

Bạn nhận được một tin nhắn bí mật được mã hóa dưới dạng chuỗi số. Tin nhắn được giải mã theo bảng ánh xạ sau:

```
"1" -> 'A'
"2" -> 'B'
...
"25" -> 'Y'
"26" -> 'Z'
```

Do một số mã nằm trong mã khác (ví dụ: "2" và "5" trong "25"), một chuỗi có thể có nhiều cách giải mã khác nhau.

Cho chuỗi `s` chỉ chứa các chữ số, hãy trả về **số cách giải mã** chuỗi đó. Nếu không thể giải mã theo bất kỳ cách nào hợp lệ, trả về `0`.

---

## Ví dụ minh họa

**Ví dụ 1:**
```
Input:  s = "12"
Output: 2
Giải thích: "12" -> "AB" (1, 2) hoặc "L" (12)
```

**Ví dụ 2:**
```
Input:  s = "226"
Output: 3
Giải thích: "BZ" (2, 26), "VF" (22, 6), "BBF" (2, 2, 6)
```

**Ví dụ 3:**
```
Input:  s = "06"
Output: 0
Giải thích: "06" không hợp lệ vì có leading zero. Chỉ "6" mới ánh xạ thành 'F'.
```

---

## Ràng buộc

- `1 <= s.length <= 100`
- `s` chỉ chứa các chữ số và có thể có leading zero

---

## Phân tích các cách tiếp cận

### 1. Brute Force (Đệ quy)

**Ý tưởng:** Tại mỗi vị trí, thử lấy 1 hoặc 2 chữ số làm một ký tự, rồi đệ quy cho phần còn lại.

```
decode("226"):
  lấy "2" -> decode("26") -> lấy "2" -> decode("6") -> 1 cách
                           -> lấy "26" -> decode("") -> 1 cách
  lấy "22" -> decode("6") -> 1 cách
Tổng: 3 cách
```

- **Thời gian:** O(2^n) — bùng nổ tổ hợp
- **Không gian:** O(n) — stack đệ quy

### 2. Đệ quy có Memoization (Top-down DP)

**Ý tưởng:** Lưu kết quả của các bài toán con đã tính để tránh tính lại.

- **Thời gian:** O(n)
- **Không gian:** O(n)

### 3. Dynamic Programming (Bottom-up) — Cách được chọn

**Ý tưởng:** Xây dựng từ dưới lên, `dp[i]` lưu số cách giải mã chuỗi `s[0..i-1]`.

- **Thời gian:** O(n)
- **Không gian:** O(n), tối ưu thêm thành O(1)

---

## Giải thích thuật toán chi tiết

### Định nghĩa

Đặt `dp[i]` = số cách giải mã chuỗi `s[0..i-1]` (chuỗi có độ dài `i`).

### Trạng thái ban đầu

```
dp[0] = 1  // chuỗi rỗng: 1 cách giải mã (base case)
dp[1] = 1  // ký tự đầu tiên: 1 cách, nếu s[0] != '0'
           // dp[1] = 0        nếu s[0] == '0'
```

### Công thức chuyển trạng thái

Tại mỗi vị trí `i` (từ 2 đến n):

1. **Lấy 1 chữ số** `s[i-1]`:
   - Hợp lệ khi `s[i-1] != '0'` (các mã từ 1-9)
   - Nếu hợp lệ: `dp[i] += dp[i-1]`

2. **Lấy 2 chữ số** `s[i-2..i-1]`:
   - Hợp lệ khi `10 <= twoDigit <= 26`
   - Nếu hợp lệ: `dp[i] += dp[i-2]`

### Minh họa với "11106"

```
s    =  1  1  1  0  6
index=  0  1  2  3  4

dp[0] = 1 (base)
dp[1] = 1 (s[0]='1' != '0')

i=2: s[1]='1' (oneDigit=1 >= 1) -> dp[2] += dp[1] = 1
     s[0..1]="11" (twoDigit=11, 10<=11<=26) -> dp[2] += dp[0] = 1
     dp[2] = 2

i=3: s[2]='1' (oneDigit=1 >= 1) -> dp[3] += dp[2] = 2
     s[1..2]="11" (twoDigit=11, valid) -> dp[3] += dp[1] = 1
     dp[3] = 3

i=4: s[3]='0' (oneDigit=0, KHÔNG hợp lệ)
     s[2..3]="10" (twoDigit=10, valid) -> dp[4] += dp[2] = 2
     dp[4] = 2

i=5: s[4]='6' (oneDigit=6 >= 1) -> dp[5] += dp[4] = 2
     s[3..4]="06" (twoDigit=6, KHÔNG hợp lệ vì < 10)
     dp[5] = 2

Kết quả: dp[5] = 2
```

Hai cách: "AAJF" (1,1,10,6) và "KJF" (11,10,6).

### Tối ưu không gian O(1)

Nhận thấy `dp[i]` chỉ phụ thuộc vào `dp[i-1]` và `dp[i-2]`, ta chỉ cần 2 biến:

```
prev2 = dp[i-2]
prev1 = dp[i-1]
```

---

## Giải thích code

```go
func numDecodings(s string) int {
    n := len(s)
    if n == 0 || s[0] == '0' {
        return 0  // Chuỗi rỗng hoặc bắt đầu bằng '0': không giải mã được
    }

    prev2 := 1  // dp[0] = 1
    prev1 := 1  // dp[1] = 1 (s[0] != '0' đã kiểm tra)

    for i := 2; i <= n; i++ {
        curr := 0

        // Thử lấy 1 chữ số: s[i-1]
        oneDigit := s[i-1] - '0'
        if oneDigit >= 1 {
            curr += prev1
        }

        // Thử lấy 2 chữ số: s[i-2..i-1]
        twoDigit := int(s[i-2]-'0')*10 + int(s[i-1]-'0')
        if twoDigit >= 10 && twoDigit <= 26 {
            curr += prev2
        }

        prev2 = prev1
        prev1 = curr
    }

    return prev1
}
```

**Điểm quan trọng:**
- `s[i-1] - '0'`: chuyển byte sang số nguyên
- `oneDigit >= 1`: loại bỏ trường hợp `'0'` đứng riêng (không có mã 0)
- `twoDigit >= 10`: loại bỏ `"01"`, `"05"`, ... (leading zero trong 2 chữ số)
- `twoDigit <= 26`: chỉ có mã từ 1 đến 26

---

## Độ phức tạp

| | Độ phức tạp |
|---|---|
| **Thời gian** | O(n) — duyệt một lần qua chuỗi |
| **Không gian** | O(1) — chỉ dùng 2 biến |

---

## Các trường hợp đặc biệt cần lưu ý

| Chuỗi | Kết quả | Lý do |
|---|---|---|
| `"0"` | 0 | Không có mã nào là "0" |
| `"06"` | 0 | Leading zero, "06" không hợp lệ |
| `"10"` | 1 | Chỉ có 1 cách: "10" -> 'J' |
| `"30"` | 0 | "30" > 26 không hợp lệ, "3" + "0" không hợp lệ |
| `"100"` | 0 | "10" + "0": "0" không giải mã được |
