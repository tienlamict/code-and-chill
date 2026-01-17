# Longest Substring Without Repeating Characters

## Mô tả bài toán

Cho một chuỗi `s`, tìm độ dài của chuỗi con dài nhất không có ký tự trùng lặp.

**Lưu ý:**
- Kết quả phải là một substring (chuỗi con liên tiếp), không phải subsequence (chuỗi con không liên tiếp)
- Chuỗi rỗng có độ dài 0

**Ví dụ:**

**Example 1:**
- Input: `s = "abcabcbb"`
- Output: `3`
- Giải thích: Câu trả lời là `"abc"`, có độ dài 3. Lưu ý rằng `"bca"` và `"cab"` cũng là câu trả lời đúng.

**Example 2:**
- Input: `s = "bbbbb"`
- Output: `1`
- Giải thích: Câu trả lời là `"b"`, có độ dài 1.

**Example 3:**
- Input: `s = "pwwkew"`
- Output: `3`
- Giải thích: Câu trả lời là `"wke"`, có độ dài 3. Lưu ý rằng câu trả lời phải là một substring, `"pwke"` là một subsequence và không phải là substring.

**Ràng buộc:**
- `0 <= s.length <= 5 * 10^4`
- `s` chỉ chứa các chữ cái tiếng Anh, chữ số, ký hiệu và khoảng trắng

## Phân tích thuật toán

### Cách tiếp cận 1: Brute Force (O(n³))

Kiểm tra tất cả các chuỗi con có thể có và tìm chuỗi dài nhất không có ký tự trùng lặp.

**Độ phức tạp:**
- Thời gian: O(n³) - có O(n²) chuỗi con, mỗi chuỗi cần O(n) để kiểm tra trùng lặp
- Không gian: O(min(n, m)) - với m là kích thước charset

### Cách tiếp cận 2: Sliding Window với Set (O(n))

Sử dụng cửa sổ trượt (sliding window) với một set để theo dõi các ký tự trong cửa sổ hiện tại.

**Ý tưởng:**
1. Duyệt qua chuỗi với hai con trỏ: `left` và `right`
2. Mở rộng cửa sổ bằng cách di chuyển `right`
3. Khi gặp ký tự trùng lặp, thu hẹp cửa sổ bằng cách di chuyển `left` cho đến khi loại bỏ ký tự trùng lặp
4. Theo dõi độ dài cửa sổ lớn nhất

**Độ phức tạp:**
- Thời gian: O(2n) = O(n) - mỗi ký tự được truy cập tối đa 2 lần (bởi left và right)
- Không gian: O(min(n, m))

### Cách tiếp cận 3: Sliding Window với Hash Map (O(n)) - Tối ưu

Sử dụng Hash Map để lưu trữ chỉ số cuối cùng mà mỗi ký tự xuất hiện, cho phép di chuyển `left` trực tiếp đến vị trí sau ký tự trùng lặp.

**Ý tưởng:**
1. Duyệt qua chuỗi với con trỏ `right`
2. Với mỗi ký tự, kiểm tra xem nó đã xuất hiện trong cửa sổ hiện tại chưa (thông qua map)
3. Nếu ký tự đã xuất hiện và nằm trong cửa sổ hiện tại (`lastIndex >= left`), di chuyển `left` đến vị trí ngay sau ký tự trùng lặp
4. Cập nhật map với chỉ số mới của ký tự
5. Tính độ dài cửa sổ hiện tại và cập nhật độ dài lớn nhất

**Độ phức tạp:**
- Thời gian: O(n) - chỉ duyệt qua chuỗi một lần
- Không gian: O(min(n, m)) - map lưu trữ tối đa min(n, m) ký tự

### Ví dụ minh họa

Với input: `s = "abcabcbb"`

```
Bước 1: right = 0, char = 'a'
  charMap = {}
  'a' chưa xuất hiện → left = 0
  charMap['a'] = 0
  currentLength = 1, maxLength = 1

Bước 2: right = 1, char = 'b'
  charMap = {'a': 0}
  'b' chưa xuất hiện → left = 0
  charMap['b'] = 1
  currentLength = 2, maxLength = 2

Bước 3: right = 2, char = 'c'
  charMap = {'a': 0, 'b': 1}
  'c' chưa xuất hiện → left = 0
  charMap['c'] = 2
  currentLength = 3, maxLength = 3

Bước 4: right = 3, char = 'a'
  charMap = {'a': 0, 'b': 1, 'c': 2}
  'a' đã xuất hiện tại index 0, và 0 >= left (0) → left = 0 + 1 = 1
  charMap['a'] = 3
  currentLength = 3 - 1 + 1 = 3, maxLength = 3

Bước 5: right = 4, char = 'b'
  charMap = {'a': 3, 'b': 1, 'c': 2}
  'b' đã xuất hiện tại index 1, và 1 >= left (1) → left = 1 + 1 = 2
  charMap['b'] = 4
  currentLength = 4 - 2 + 1 = 3, maxLength = 3

Bước 6: right = 5, char = 'c'
  charMap = {'a': 3, 'b': 4, 'c': 2}
  'c' đã xuất hiện tại index 2, và 2 >= left (2) → left = 2 + 1 = 3
  charMap['c'] = 5
  currentLength = 5 - 3 + 1 = 3, maxLength = 3

Bước 7: right = 6, char = 'b'
  charMap = {'a': 3, 'b': 4, 'c': 5}
  'b' đã xuất hiện tại index 4, và 4 >= left (3) → left = 4 + 1 = 5
  charMap['b'] = 6
  currentLength = 6 - 5 + 1 = 2, maxLength = 3

Bước 8: right = 7, char = 'b'
  charMap = {'a': 3, 'b': 6, 'c': 5}
  'b' đã xuất hiện tại index 6, và 6 >= left (5) → left = 6 + 1 = 7
  charMap['b'] = 7
  currentLength = 7 - 7 + 1 = 1, maxLength = 3

Kết quả: maxLength = 3
```

Với input: `s = "pwwkew"`

```
Bước 1-2: "pw" → maxLength = 2
Bước 3: gặp 'w' trùng → left = 3 (sau 'w' đầu tiên)
Bước 3-5: "wke" → maxLength = 3
Kết quả: maxLength = 3
```

## Giải thích code

### Cấu trúc hàm

```1:24:internal/hash_table/longest_substring/longest_substring.go
package hash_table

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	charMap := make(map[byte]int)
	maxLength := 0
	left := 0

	for right := 0; right < len(s); right++ {
		char := s[right]

		if lastIndex, exists := charMap[char]; exists && lastIndex >= left {
			left = lastIndex + 1
		}

		charMap[char] = right
		currentLength := right - left + 1

		if currentLength > maxLength {
			maxLength = currentLength
		}
	}

	return maxLength
}
```

### Chi tiết từng phần

#### 1. Xử lý chuỗi rỗng (dòng 3-5)
```go
if len(s) == 0 {
    return 0
}
```
- Kiểm tra trường hợp chuỗi rỗng và trả về 0 ngay lập tức

#### 2. Khởi tạo biến (dòng 7-9)
```go
charMap := make(map[byte]int)
maxLength := 0
left := 0
```
- `charMap`: Map lưu trữ chỉ số cuối cùng mà mỗi ký tự xuất hiện
  - Key: ký tự (byte)
  - Value: chỉ số cuối cùng mà ký tự xuất hiện
- `maxLength`: Độ dài lớn nhất của substring không trùng lặp đã tìm thấy
- `left`: Con trỏ trái của cửa sổ trượt

#### 3. Duyệt qua chuỗi (dòng 11-23)
```go
for right := 0; right < len(s); right++ {
    char := s[right]

    if lastIndex, exists := charMap[char]; exists && lastIndex >= left {
        left = lastIndex + 1
    }

    charMap[char] = right
    currentLength := right - left + 1

    if currentLength > maxLength {
        maxLength = currentLength
    }
}
```

**Giải thích từng bước:**

- **`right`**: Con trỏ phải của cửa sổ trượt, duyệt qua từng ký tự trong chuỗi

- **`char := s[right]`**: Lấy ký tự tại vị trí `right`

- **Kiểm tra ký tự trùng lặp (dòng 14-16)**:
  ```go
  if lastIndex, exists := charMap[char]; exists && lastIndex >= left {
      left = lastIndex + 1
  }
  ```
  - Kiểm tra xem ký tự đã xuất hiện chưa và có nằm trong cửa sổ hiện tại không
  - `exists`: Ký tự đã xuất hiện trong map
  - `lastIndex >= left`: Ký tự nằm trong cửa sổ hiện tại (từ `left` đến `right`)
  - Nếu cả hai điều kiện đúng: di chuyển `left` đến vị trí ngay sau ký tự trùng lặp
  - **Tại sao `lastIndex >= left`?** Đảm bảo rằng ký tự trùng lặp thực sự nằm trong cửa sổ hiện tại. Nếu `lastIndex < left`, có nghĩa là ký tự đó đã bị loại bỏ khỏi cửa sổ và không còn là trùng lặp nữa.

- **Cập nhật map (dòng 18)**:
  ```go
  charMap[char] = right
  ```
  - Lưu/cập nhật chỉ số mới nhất của ký tự

- **Tính độ dài và cập nhật maxLength (dòng 19-23)**:
  ```go
  currentLength := right - left + 1
  if currentLength > maxLength {
      maxLength = currentLength
  }
  ```
  - Tính độ dài cửa sổ hiện tại: `right - left + 1`
  - Cập nhật `maxLength` nếu cửa sổ hiện tại dài hơn

#### 4. Trả về kết quả (dòng 24)
```go
return maxLength
```
- Trả về độ dài lớn nhất của substring không trùng lặp

### Tại sao thuật toán này hoạt động?

1. **Sliding Window**: Sử dụng kỹ thuật cửa sổ trượt để duyệt qua tất cả các substring có thể mà không cần kiểm tra lại từ đầu
2. **Hash Map lookup**: Kiểm tra sự tồn tại và vị trí của ký tự trong O(1) thay vì O(n)
3. **Nhảy cóc**: Khi gặp ký tự trùng lặp, di chuyển `left` trực tiếp đến vị trí sau ký tự trùng lặp, tránh duyệt từng ký tự một
4. **Đảm bảo tính đúng**: Điều kiện `lastIndex >= left` đảm bảo rằng ta chỉ xử lý ký tự trùng lặp trong cửa sổ hiện tại

### Độ phức tạp

- **Thời gian**: O(n)
  - Duyệt qua chuỗi một lần: O(n)
  - Mỗi thao tác với map (insert, lookup): O(1)
  - Tổng: O(n)

- **Không gian**: O(min(n, m))
  - `m`: Kích thước của bảng ký tự (character set)
  - Trong trường hợp xấu nhất, map lưu trữ tối đa min(n, m) ký tự
  - Với ASCII: m = 128, Unicode: m có thể lớn hơn nhiều

## Test Cases

### Test Case 1
```go
Input: s = "abcabcbb"
Output: 3
```
Giải thích: Substring dài nhất không trùng lặp là `"abc"` (có thể là `"bca"` hoặc `"cab"`)

### Test Case 2
```go
Input: s = "bbbbb"
Output: 1
```
Giải thích: Tất cả các ký tự đều giống nhau, substring dài nhất chỉ có 1 ký tự

### Test Case 3
```go
Input: s = "pwwkew"
Output: 3
```
Giải thích: Substring dài nhất không trùng lặp là `"wke"`

### Test Case 4: Chuỗi rỗng
```go
Input: s = ""
Output: 0
```
Giải thích: Chuỗi rỗng có độ dài 0

### Test Case 5: Một ký tự
```go
Input: s = "a"
Output: 1
```
Giải thích: Substring dài nhất là chính chuỗi đó

### Test Case 6: Tất cả ký tự khác nhau
```go
Input: s = "abcdef"
Output: 6
```
Giải thích: Toàn bộ chuỗi là substring dài nhất không trùng lặp

### Test Case 7: Chuỗi có khoảng trắng
```go
Input: s = "a b c"
Output: 3
```
Giải thích: Mỗi phần tử (ký tự hoặc khoảng trắng) là duy nhất trong cửa sổ hiện tại

### Test Case 8: Chuỗi có chữ số và ký hiệu
```go
Input: s = "abc123!@#"
Output: 9
```
Giải thích: Tất cả các ký tự đều khác nhau

### Test Case 9: Trùng lặp ở cuối
```go
Input: s = "abcdabcd"
Output: 4
```
Giải thích: Substring dài nhất là `"abcd"`

### Test Case 10: Substring dài nhất ở đầu
```go
Input: s = "abcdefgabc"
Output: 7
```
Giải thích: Substring dài nhất là `"abcdefg"` ở đầu chuỗi

## Chạy test

Để chạy các test case:

```bash
go test ./internal/hash_table/longest_substring/
```

Hoặc chạy với verbose mode để xem chi tiết:

```bash
go test -v ./internal/hash_table/longest_substring/
```

## So sánh với các cách tiếp cận khác

| Cách tiếp cận | Độ phức tạp thời gian | Độ phức tạp không gian | Ghi chú |
|---------------|----------------------|----------------------|---------|
| Brute Force | O(n³) | O(min(n, m)) | Kiểm tra tất cả các chuỗi con |
| Sliding Window + Set | O(2n) = O(n) | O(min(n, m)) | Thu hẹp cửa sổ từng bước |
| Sliding Window + Hash Map | O(n) | O(min(n, m)) | **Tối ưu** - sử dụng trong code này |

**Kết luận:** Sliding Window với Hash Map là cách tiếp cận tối ưu nhất cho bài toán này vì:
- Độ phức tạp thời gian O(n) - tốt hơn O(n³) và O(2n)
- Cho phép "nhảy cóc" khi gặp ký tự trùng lặp, không cần duyệt từng ký tự
- Code đơn giản, dễ hiểu và hiệu quả
