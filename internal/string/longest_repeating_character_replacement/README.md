# Longest Repeating Character Replacement

## Mô tả bài toán

Cho một chuỗi `s` và một số nguyên `k`. Bạn có thể chọn bất kỳ ký tự nào trong chuỗi và thay đổi nó thành bất kỳ ký tự in hoa tiếng Anh nào khác. Bạn có thể thực hiện thao tác này tối đa `k` lần.

Trả về **độ dài của chuỗi con dài nhất** chứa cùng một ký tự mà bạn có thể đạt được sau khi thực hiện các thao tác trên.

**Ví dụ:**

### Example 1:
- Input: `s = "ABAB"`, `k = 2`
- Output: `4`
- Giải thích: Thay hai ký tự `'A'` thành `'B'` (hoặc ngược lại). Kết quả là `"BBBB"` hoặc `"AAAA"`, có độ dài 4.

### Example 2:
- Input: `s = "AABABBA"`, `k = 1`
- Output: `4`
- Giải thích: Thay ký tự `'A'` ở giữa thành `'B'` để tạo `"AABBBBA"`. Chuỗi con `"BBBB"` có độ dài 4.

**Ràng buộc:**
- `1 <= s.length <= 10^5`
- `s` chỉ chứa các ký tự in hoa tiếng Anh
- `0 <= k <= s.length`

## Phân tích thuật toán

### Cách tiếp cận 1: Brute Force — O(n³)

Kiểm tra tất cả các chuỗi con, với mỗi chuỗi con đếm ký tự xuất hiện nhiều nhất, nếu số ký tự cần thay ≤ k thì cập nhật kết quả.

**Độ phức tạp:**
- Thời gian: O(n³) — duyệt tất cả cặp (i, j) và đếm ký tự trong mỗi chuỗi con
- Không gian: O(1)

### Cách tiếp cận 2: Brute Force tối ưu — O(n² × 26)

Tối ưu bằng cách duy trì mảng đếm tần suất khi mở rộng chuỗi con.

**Độ phức tạp:**
- Thời gian: O(26 × n²) ≈ O(n²)
- Không gian: O(26) = O(1)

### Cách tiếp cận 3: Sliding Window — O(n) ⭐ Tối ưu

Sử dụng cửa sổ trượt với hai con trỏ `left` và `right`:

**Ý tưởng cốt lõi:**
- Trong một cửa sổ `[left, right]`, ký tự xuất hiện nhiều nhất có tần suất `maxFreq`
- Số ký tự cần thay đổi = `(right - left + 1) - maxFreq`
- Nếu số ký tự cần thay ≤ `k` → cửa sổ hợp lệ
- Nếu số ký tự cần thay > `k` → thu hẹp cửa sổ từ bên trái

**Trick quan trọng:** `maxFreq` không cần giảm khi thu hẹp cửa sổ!
- Ta chỉ quan tâm tìm cửa sổ **lớn hơn** cửa sổ tốt nhất đã tìm được
- Kết quả chỉ cải thiện khi `maxFreq` tăng
- Nên việc giữ `maxFreq` không giảm không ảnh hưởng đến tính đúng đắn

**Độ phức tạp:**
- Thời gian: O(n) — mỗi ký tự được thăm tối đa 2 lần
- Không gian: O(1) — mảng 26 phần tử

## Giải thích chi tiết thuật toán Sliding Window

### Nguyên lý hoạt động

Với cửa sổ `[left, right]`:
```
Chuỗi:    A  A  B  A  B  B  A
Index:     0  1  2  3  4  5  6
           ^              ^
          left           right
          
Window = "AABABB" (length = 6)
maxFreq = 3 (ký tự 'A' hoặc 'B')
Cần thay = 6 - 3 = 3
Nếu k >= 3: hợp lệ
Nếu k < 3: thu hẹp (left++)
```

### Ví dụ từng bước: s = "AABABBA", k = 1

```
Bước 1: right=0, char='A'
  count: A=1
  maxFreq = 1
  window = "A" (len=1), cần thay = 1-1 = 0 ≤ 1 ✓
  result = 1

Bước 2: right=1, char='A'
  count: A=2
  maxFreq = 2
  window = "AA" (len=2), cần thay = 2-2 = 0 ≤ 1 ✓
  result = 2

Bước 3: right=2, char='B'
  count: A=2, B=1
  maxFreq = 2
  window = "AAB" (len=3), cần thay = 3-2 = 1 ≤ 1 ✓
  result = 3

Bước 4: right=3, char='A'
  count: A=3, B=1
  maxFreq = 3
  window = "AABA" (len=4), cần thay = 4-3 = 1 ≤ 1 ✓
  result = 4

Bước 5: right=4, char='B'
  count: A=3, B=2
  maxFreq = 3
  window = "AABAB" (len=5), cần thay = 5-3 = 2 > 1 ✗
  → Thu hẹp: count['A']--, left=1
  count: A=2, B=2
  window = "ABAB" (len=4)
  result = 4

Bước 6: right=5, char='B'
  count: A=2, B=3
  maxFreq = 3
  window = "ABABB" (len=5), cần thay = 5-3 = 2 > 1 ✗
  → Thu hẹp: count['A']--, left=2
  count: A=1, B=3
  window = "BABB" (len=4)
  result = 4

Bước 7: right=6, char='A'
  count: A=2, B=3
  maxFreq = 3
  window = "BABBA" (len=5), cần thay = 5-3 = 2 > 1 ✗
  → Thu hẹp: count['B']--, left=3
  count: A=2, B=2
  window = "ABBA" (len=4)
  result = 4

Kết quả: 4
```

## Giải thích code

### Khởi tạo biến

```go
count := [26]int{}  // Mảng đếm tần suất 26 ký tự (A-Z)
left := 0           // Con trỏ trái của cửa sổ
maxFreq := 0        // Tần suất cao nhất của một ký tự trong cửa sổ
result := 0         // Kết quả: độ dài cửa sổ lớn nhất hợp lệ
```

- `count`: mảng cố định 26 phần tử, đếm số lần xuất hiện của mỗi ký tự trong cửa sổ hiện tại
- `s[right]-'A'`: chuyển ký tự thành chỉ số 0-25

### Vòng lặp chính — Mở rộng cửa sổ

```go
for right := 0; right < len(s); right++ {
    count[s[right]-'A']++

    if count[s[right]-'A'] > maxFreq {
        maxFreq = count[s[right]-'A']
    }
```

- Mỗi lần lặp, mở rộng cửa sổ sang phải bằng cách thêm `s[right]`
- Cập nhật tần suất và `maxFreq` nếu ký tự mới có tần suất cao hơn

### Kiểm tra và thu hẹp cửa sổ

```go
    windowLen := right - left + 1
    if windowLen-maxFreq > k {
        count[s[left]-'A']--
        left++
    }
```

- `windowLen - maxFreq`: số ký tự **không phải** ký tự chiếm đa số → cần thay đổi
- Nếu số cần thay > `k`: cửa sổ không hợp lệ → thu hẹp bằng cách bỏ ký tự bên trái
- Lưu ý: chỉ thu hẹp 1 bước (không dùng `while`) vì ta chỉ cần giữ cửa sổ không lớn hơn kích thước tối ưu hiện tại

### Cập nhật kết quả

```go
    if right-left+1 > result {
        result = right - left + 1
    }
}
return result
```

- Sau mỗi bước, kích thước cửa sổ hiện tại là ứng viên cho kết quả

## Tại sao maxFreq không cần giảm?

Đây là điểm **tinh tế nhất** của thuật toán:

1. **Kết quả chỉ cải thiện khi maxFreq tăng**: Vì `result = windowLen` chỉ lớn hơn khi `windowLen` tăng, mà `windowLen` chỉ tăng khi `maxFreq` tăng (do điều kiện `windowLen - maxFreq ≤ k`).

2. **maxFreq giảm không có lợi**: Nếu maxFreq giảm, cửa sổ hợp lệ cũng nhỏ hơn, không thể tốt hơn kết quả hiện tại.

3. **Cửa sổ chỉ co lại 1 bước**: Khi `windowLen - maxFreq > k`, ta chỉ thu hẹp 1 bước. Cửa sổ giữ nguyên kích thước (hoặc lớn hơn), đợi cho đến khi tìm được `maxFreq` mới lớn hơn.

Điều này biến thuật toán từ O(26n) (nếu phải quét lại maxFreq mỗi lần) thành O(n) thực sự.

## Độ phức tạp

### Thời gian: O(n)
- `right` duyệt từ 0 đến n-1: O(n)
- `left` chỉ tăng, không bao giờ giảm, tối đa n lần: O(n)
- Mỗi bước: O(1) (cập nhật count, so sánh maxFreq)
- Tổng: O(n)

### Không gian: O(1)
- Mảng `count[26]`: O(26) = O(1) (hằng số)
- Các biến `left`, `maxFreq`, `result`: O(1)

## Test Cases

| Test Case | Input | k | Output | Giải thích |
|-----------|-------|---|--------|------------|
| Example 1 | `"ABAB"` | 2 | 4 | Thay 2 ký tự để toàn bộ giống nhau |
| Example 2 | `"AABABBA"` | 1 | 4 | Thay 1 `'A'` thành `'B'` → `"BBBB"` |
| Single char | `"A"` | 0 | 1 | Chuỗi 1 ký tự |
| All same | `"AAAA"` | 2 | 4 | Không cần thay |
| k=0 | `"ABCDE"` | 0 | 1 | Không được thay, chuỗi con dài nhất = 1 |
| k=length | `"ABCDE"` | 5 | 5 | Thay hết → toàn bộ chuỗi |

## Chạy test

```bash
go test ./internal/string/longest_repeating_character_replacement/
```

Hoặc với verbose mode:

```bash
go test -v ./internal/string/longest_repeating_character_replacement/
```
