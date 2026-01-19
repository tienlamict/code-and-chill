# Longest Palindromic Substring

## Mô tả bài toán

Cho một chuỗi `s`, trả về chuỗi con palindromic dài nhất trong `s`.

**Lưu ý:**
- Palindrome là chuỗi đọc xuôi hay đọc ngược đều giống nhau (ví dụ: "aba", "abba", "racecar")
- Kết quả phải là một substring (chuỗi con liên tiếp)

**Ví dụ:**

**Example 1:**
- Input: `s = "babad"`
- Output: `"bab"`
- Giải thích: `"aba"` cũng là một câu trả lời hợp lệ.

**Example 2:**
- Input: `s = "cbbd"`
- Output: `"bb"`

**Ràng buộc:**
- `1 <= s.length <= 1000`
- `s` chỉ chứa chữ số và chữ cái tiếng Anh

## Phân tích thuật toán

### Cách tiếp cận 1: Brute Force (O(n³))

Kiểm tra tất cả các substring có thể có và xác định substring nào là palindrome dài nhất.

**Độ phức tạp:**
- Thời gian: O(n³) - có O(n²) substring, mỗi substring cần O(n) để kiểm tra palindrome
- Không gian: O(1)

### Cách tiếp cận 2: Dynamic Programming (O(n²))

Sử dụng bảng DP để lưu trữ kết quả của các substring nhỏ hơn.

**Ý tưởng:**
- `dp[i][j]` = true nếu substring từ `i` đến `j` là palindrome
- `dp[i][j] = (s[i] == s[j]) && dp[i+1][j-1]`

**Độ phức tạp:**
- Thời gian: O(n²)
- Không gian: O(n²)

### Cách tiếp cận 3: Expand Around Centers (O(n²)) - Tối ưu về không gian

Mỗi palindrome đều có một center. Ta có thể mở rộng từ center ra hai bên để tìm palindrome dài nhất.

**Ý tưởng:**
1. Mỗi vị trí trong chuỗi có thể là center của một palindrome
2. Có hai loại center:
   - **Odd length**: Center là một ký tự (ví dụ: "aba", center ở 'b')
   - **Even length**: Center là giữa hai ký tự (ví dụ: "abba", center giữa hai 'b')
3. Với mỗi center, mở rộng ra hai bên cho đến khi không còn matching
4. Theo dõi palindrome dài nhất

**Độ phức tạp:**
- Thời gian: O(n²) - có n vị trí có thể làm center, mỗi vị trí tối đa mở rộng O(n)
- Không gian: O(1) - chỉ sử dụng biến để lưu trữ

### Ví dụ minh họa

Với input: `s = "babad"`

```
Center tại index 0 ('b'):
  Expand: b (length 1)

Center tại index 1 ('a'):
  Expand: b-a-b (length 3) ✓ Tìm thấy "bab"

Center giữa index 0-1:
  Expand: không match

Center tại index 2 ('b'):
  Expand: a-b-a (length 3) ✓ "aba" cũng hợp lệ

Center giữa index 1-2:
  Expand: không match

Center tại index 3 ('a'):
  Expand: a (length 1)

Center tại index 4 ('d'):
  Expand: d (length 1)

Kết quả: "bab" hoặc "aba" (độ dài 3)
```

Với input: `s = "cbbd"`

```
Center tại index 0 ('c'):
  Expand: c (length 1)

Center tại index 1 ('b'):
  Expand: c-b (không match)
  
Center giữa index 1-2:
  Expand: b-b (length 2) ✓ Tìm thấy "bb"

Center tại index 2 ('b'):
  Expand: b-d (không match)

Center tại index 3 ('d'):
  Expand: d (length 1)

Kết quả: "bb" (độ dài 2)
```

## Giải thích code

### Cấu trúc hàm

```1:29:internal/hash_table/longest_palindromic_substring/longest_palindromic_substring.go
package hash_table

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	start := 0
	maxLength := 1

	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)

		length := len1
		if len2 > len1 {
			length = len2
		}

		if length > maxLength {
			maxLength = length
			start = i - (length-1)/2
		}
	}

	return s[start : start+maxLength]
}

func expandAroundCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}
```

### Chi tiết từng phần

#### 1. Xử lý chuỗi rỗng (dòng 3-5)
```go
if len(s) == 0 {
    return ""
}
```
- Kiểm tra trường hợp chuỗi rỗng và trả về chuỗi rỗng ngay lập tức

#### 2. Khởi tạo biến (dòng 7-8)
```go
start := 0
maxLength := 1
```
- `start`: Vị trí bắt đầu của palindrome dài nhất
- `maxLength`: Độ dài của palindrome dài nhất (mặc định là 1 vì mỗi ký tự đều là palindrome)

#### 3. Duyệt qua các center tiềm năng (dòng 10-22)
```go
for i := 0; i < len(s); i++ {
    len1 := expandAroundCenter(s, i, i)
    len2 := expandAroundCenter(s, i, i+1)

    length := len1
    if len2 > len1 {
        length = len2
    }

    if length > maxLength {
        maxLength = length
        start = i - (length-1)/2
    }
}
```

**Giải thích từng bước:**

- **`len1 := expandAroundCenter(s, i, i)`**: Tìm palindrome có độ dài lẻ với center tại vị trí `i`
  - Truyền cùng một giá trị cho `left` và `right` nghĩa là center là một ký tự
  
- **`len2 := expandAroundCenter(s, i, i+1)`**: Tìm palindrome có độ dài chẵn với center giữa `i` và `i+1`
  - Truyền `i` và `i+1` nghĩa là center nằm giữa hai ký tự

- **Chọn độ dài lớn hơn**: So sánh `len1` và `len2`, chọn giá trị lớn hơn

- **Cập nhật kết quả**: Nếu tìm thấy palindrome dài hơn:
  - Cập nhật `maxLength`
  - Tính `start` position: `i - (length-1)/2`
    - Công thức này hoạt động cho cả odd và even length:
      - Odd (length=3): `start = i - (3-1)/2 = i - 1`
      - Even (length=2): `start = i - (2-1)/2 = i - 0 = i`

#### 4. Hàm expandAroundCenter (dòng 25-29)
```go
func expandAroundCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}
```

**Giải thích:**
- **Tham số**: `left` và `right` là hai con trỏ bắt đầu tại center (hoặc gần center)
- **Vòng lặp mở rộng**: Tiếp tục mở rộng ra hai bên trong khi:
  - `left >= 0`: Chưa vượt quá đầu chuỗi
  - `right < len(s)`: Chưa vượt quá cuối chuỗi
  - `s[left] == s[right]`: Hai ký tự ở hai bên vẫn khớp nhau
- **Trả về độ dài**: `right - left - 1`
  - Sau khi vòng lặp kết thúc, `left` và `right` đã vượt quá palindrome
  - Độ dài thực tế = `(right - 1) - (left + 1) + 1 = right - left - 1`

**Ví dụ minh họa expandAroundCenter:**

Với `s = "babad"`, `i = 1` (ký tự 'a'):
```
len1 = expandAroundCenter(s, 1, 1):
  left=1, right=1: s[1]=='a', s[1]=='a' → match
  left=0, right=2: s[0]=='b', s[2]=='b' → match
  left=-1, right=3: out of bounds → stop
  Return: 3 - (-1) - 1 = 3 ✓ "bab"

len2 = expandAroundCenter(s, 1, 2):
  left=1, right=2: s[1]=='a', s[2]=='b' → không match
  Return: 2 - 1 - 1 = 0
```

#### 5. Trả về kết quả (dòng 24)
```go
return s[start : start+maxLength]
```
- Trả về substring từ vị trí `start` với độ dài `maxLength`

### Tại sao thuật toán này hoạt động?

1. **Tính chất palindrome**: Mỗi palindrome đều có một center (hoặc giữa hai ký tự cho even length)
2. **Bao phủ đầy đủ**: Duyệt qua tất cả các vị trí có thể làm center, đảm bảo không bỏ sót palindrome nào
3. **Mở rộng hiệu quả**: Mở rộng từ center ra hai bên là cách tự nhiên và hiệu quả để tìm palindrome
4. **Tối ưu không gian**: Chỉ sử dụng O(1) không gian phụ, không cần mảng DP

### Độ phức tạp

- **Thời gian**: O(n²)
  - Vòng lặp ngoài: O(n) - duyệt qua n vị trí
  - Vòng lặp trong (expandAroundCenter): O(n) trong trường hợp xấu nhất (toàn bộ chuỗi là palindrome)
  - Tổng: O(n × n) = O(n²)

- **Không gian**: O(1)
  - Chỉ sử dụng một số biến để lưu trữ: `start`, `maxLength`, `i`, `len1`, `len2`, `length`
  - Hàm `expandAroundCenter` chỉ sử dụng tham số và biến cục bộ

## Test Cases

### Test Case 1
```go
Input: s = "babad"
Output: "bab"
```
Giải thích: `"bab"` hoặc `"aba"` đều là palindrome dài nhất (độ dài 3)

### Test Case 2
```go
Input: s = "cbbd"
Output: "bb"
```
Giải thích: `"bb"` là palindrome dài nhất (độ dài 2)

### Test Case 3: Một ký tự
```go
Input: s = "a"
Output: "a"
```
Giải thích: Chính ký tự đó là palindrome dài nhất

### Test Case 4: Tất cả ký tự giống nhau
```go
Input: s = "aaa"
Output: "aaa"
```
Giải thích: Toàn bộ chuỗi là palindrome

### Test Case 5: Palindrome ở đầu
```go
Input: s = "abacdfgdcaba"
Output: "aba"
```
Giải thích: Palindrome `"aba"` ở đầu chuỗi

### Test Case 6: Toàn bộ chuỗi là palindrome
```go
Input: s = "racecar"
Output: "racecar"
```
Giải thích: Toàn bộ chuỗi là palindrome

### Test Case 7: Palindrome độ dài chẵn
```go
Input: s = "abba"
Output: "abba"
```
Giải thích: Palindrome có độ dài chẵn với center giữa hai ký tự 'b'

### Test Case 8: Không có palindrome dài hơn 1
```go
Input: s = "abc"
Output: "a"
```
Giải thích: Mỗi ký tự đều là palindrome, ký tự đầu tiên được trả về

### Test Case 9: Chuỗi có chữ số
```go
Input: s = "a1b2c3c2b1a"
Output: "a1b2c3c2b1a"
```
Giải thích: Toàn bộ chuỗi là palindrome

### Test Case 10: Chuỗi dài
```go
Input: s = "forgeeksskeegfor"
Output: "geeksskeeg"
```
Giải thích: Palindrome dài nhất ở giữa chuỗi

## Chạy test

Để chạy các test case:

```bash
go test ./internal/hash_table/longest_palindromic_substring/
```

Hoặc chạy với verbose mode để xem chi tiết:

```bash
go test -v ./internal/hash_table/longest_palindromic_substring/
```

## So sánh với các cách tiếp cận khác

| Cách tiếp cận | Độ phức tạp thời gian | Độ phức tạp không gian | Ghi chú |
|---------------|----------------------|----------------------|---------|
| Brute Force | O(n³) | O(1) | Kiểm tra tất cả các substring |
| Dynamic Programming | O(n²) | O(n²) | Sử dụng bảng DP |
| Expand Around Centers | O(n²) | O(1) | **Tối ưu về không gian** - sử dụng trong code này |
| Manacher's Algorithm | O(n) | O(n) | Phức tạp hơn, khó implement |

**Kết luận:** Expand Around Centers là cách tiếp cận cân bằng tốt vì:
- Độ phức tạp thời gian O(n²) - tốt hơn O(n³)
- Độ phức tạp không gian O(1) - tốt hơn O(n²) của DP
- Code đơn giản, dễ hiểu và maintain
- Với constraint `n <= 1000`, O(n²) là đủ nhanh

## Lưu ý về Manacher's Algorithm

Manacher's Algorithm có thể giải bài toán này trong O(n) thời gian, nhưng:
- Thuật toán phức tạp hơn, khó hiểu và implement
- Với constraint `n <= 1000`, O(n²) đã đủ nhanh (tối đa 1,000,000 operations)
- Expand Around Centers dễ hiểu và maintain hơn

Vì vậy, Expand Around Centers là lựa chọn phù hợp cho bài toán này.
