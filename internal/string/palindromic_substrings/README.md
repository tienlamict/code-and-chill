# Palindromic Substrings

## Mô tả bài toán

Cho một chuỗi `s`, trả về số lượng substring palindromic trong chuỗi đó.

**Lưu ý:**
- Palindrome là chuỗi đọc xuôi hay đọc ngược đều giống nhau (ví dụ: "a", "aa", "aba", "abba")
- Substring là một chuỗi con liên tiếp trong chuỗi gốc
- Mỗi ký tự đơn lẻ cũng được tính là một palindrome

**Ví dụ:**

### Example 1:
- Input: `s = "abc"`
- Output: `3`
- Giải thích: Có 3 substring palindromic: `"a"`, `"b"`, `"c"`

### Example 2:
- Input: `s = "aaa"`
- Output: `6`
- Giải thích: Có 6 substring palindromic: `"a"`, `"a"`, `"a"`, `"aa"`, `"aa"`, `"aaa"`

**Ràng buộc:**
- `1 <= s.length <= 1000`
- `s` chỉ chứa chữ cái tiếng Anh viết thường

## Phân tích thuật toán

### Cách tiếp cận 1: Brute Force (O(n³))

Kiểm tra tất cả các substring có thể có và đếm số lượng substring là palindrome.

**Độ phức tạp:**
- Thời gian: O(n³) - có O(n²) substring, mỗi substring cần O(n) để kiểm tra palindrome
- Không gian: O(1)

### Cách tiếp cận 2: Expand Around Centers (O(n²)) - Tối ưu

Mỗi palindrome đều có một center. Ta có thể mở rộng từ center ra hai bên và đếm tất cả các palindrome có center tại đó.

**Ý tưởng:**
1. Mỗi vị trí trong chuỗi có thể là center của một hoặc nhiều palindrome
2. Có hai loại center:
   - **Odd length**: Center là một ký tự (ví dụ: "aba", center ở 'b')
   - **Even length**: Center là giữa hai ký tự (ví dụ: "abba", center giữa hai 'b')
3. Với mỗi center, mở rộng ra hai bên và đếm mỗi lần mở rộng thành công
4. Mỗi lần mở rộng thành công = tìm thấy một palindrome mới

**Độ phức tạp:**
- Thời gian: O(n²) - có n vị trí có thể làm center, mỗi vị trí tối đa mở rộng O(n)
- Không gian: O(1) - chỉ sử dụng biến để đếm

### Ví dụ minh họa

Với input: `s = "abc"`

```
Center tại index 0 ('a'):
  Expand: a → count = 1
  Expand tiếp: không match → dừng
  Tổng: 1 palindrome ("a")

Center giữa index 0-1:
  Expand: không match → dừng
  Tổng: 0 palindrome

Center tại index 1 ('b'):
  Expand: b → count = 1
  Expand tiếp: không match → dừng
  Tổng: 1 palindrome ("b")

Center giữa index 1-2:
  Expand: không match → dừng
  Tổng: 0 palindrome

Center tại index 2 ('c'):
  Expand: c → count = 1
  Expand tiếp: không match → dừng
  Tổng: 1 palindrome ("c")

Tổng cộng: 1 + 0 + 1 + 0 + 1 = 3 palindromes
```

Với input: `s = "aaa"`

```
Center tại index 0 ('a'):
  Expand: a → count = 1 ("a")
  Expand: a-a → count = 2 ("aa")
  Expand: a-a-a → count = 3 ("aaa")
  Tổng: 3 palindromes

Center giữa index 0-1:
  Expand: a-a → count = 1 ("aa")
  Expand: a-a-a → count = 2 ("aaa")
  Tổng: 2 palindromes

Center tại index 1 ('a'):
  Expand: a → count = 1 ("a")
  Expand: a-a → count = 2 ("aa")
  Tổng: 2 palindromes

Center giữa index 1-2:
  Expand: a-a → count = 1 ("aa")
  Tổng: 1 palindrome

Center tại index 2 ('a'):
  Expand: a → count = 1 ("a")
  Tổng: 1 palindrome

Tổng cộng: 3 + 2 + 2 + 1 + 1 = 9 palindromes? 
→ KHÔNG! Mỗi palindrome chỉ được đếm một lần.

Thực tế:
- "a" tại index 0: đếm 1 lần (từ center 0)
- "a" tại index 1: đếm 1 lần (từ center 1)  
- "a" tại index 2: đếm 1 lần (từ center 2)
- "aa" từ index 0-1: đếm 1 lần (từ center 0-1)
- "aa" từ index 1-2: đếm 1 lần (từ center 1-2)
- "aaa" từ index 0-2: đếm 1 lần (từ center 0)

Tổng: 6 palindromes ✓
```

**Lưu ý quan trọng:** Mỗi palindrome được đếm đúng một lần khi nó được tìm thấy từ center của nó. Không có palindrome nào bị đếm trùng.

## Giải thích code

### Cấu trúc hàm

```5:25:internal/string/palindromic_substrings/palindromic_substrings.go
// countSubstrings returns the number of palindromic substrings in s
func countSubstrings(s string) int {
	count := 0

	// Expand around each possible center
	for i := 0; i < len(s); i++ {
		// Count palindromes with odd length (center at i)
		count += expandAndCount(s, i, i)

		// Count palindromes with even length (center between i and i+1)
		count += expandAndCount(s, i, i+1)
	}

	return count
}

// expandAndCount expands around center and counts palindromic substrings
func expandAndCount(s string, left, right int) int {
	count := 0
	for left >= 0 && right < len(s) && s[left] == s[right] {
		count++
		left--
		right++
	}
	return count
}
```

### Chi tiết từng phần

#### 1. Khởi tạo biến đếm (dòng 6)
```go
count := 0
```
- Biến lưu tổng số lượng substring palindromic

#### 2. Duyệt qua các center tiềm năng (dòng 8-15)
```go
for i := 0; i < len(s); i++ {
    // Count palindromes with odd length (center at i)
    count += expandAndCount(s, i, i)

    // Count palindromes with even length (center between i and i+1)
    count += expandAndCount(s, i, i+1)
}
```

**Giải thích từng bước:**

- **`expandAndCount(s, i, i)`**: Đếm palindromes có độ dài lẻ với center tại vị trí `i`
  - Truyền cùng một giá trị cho `left` và `right` nghĩa là center là một ký tự
  - Ví dụ: với `i = 1` trong `"aba"`, sẽ tìm các palindrome: `"b"`, `"aba"`
  
- **`expandAndCount(s, i, i+1)`**: Đếm palindromes có độ dài chẵn với center giữa `i` và `i+1`
  - Truyền `i` và `i+1` nghĩa là center nằm giữa hai ký tự
  - Ví dụ: với `i = 1` trong `"abba"`, sẽ tìm các palindrome: `"bb"`, `"abba"`

- **Cộng dồn**: Mỗi lần gọi `expandAndCount` trả về số lượng palindrome tìm được từ center đó, cộng vào tổng

#### 3. Hàm expandAndCount (dòng 18-25)
```go
func expandAndCount(s string, left, right int) int {
	count := 0
	for left >= 0 && right < len(s) && s[left] == s[right] {
		count++
		left--
		right++
	}
	return count
}
```

**Giải thích:**
- **Tham số**: `left` và `right` là hai con trỏ bắt đầu tại center (hoặc gần center)
- **Khởi tạo count**: Bắt đầu từ 0
- **Vòng lặp mở rộng**: Tiếp tục mở rộng ra hai bên trong khi:
  - `left >= 0`: Chưa vượt quá đầu chuỗi
  - `right < len(s)`: Chưa vượt quá cuối chuỗi
  - `s[left] == s[right]`: Hai ký tự ở hai bên vẫn khớp nhau
- **Đếm palindrome**: Mỗi lần mở rộng thành công, `count++` vì ta tìm thấy một palindrome mới
- **Mở rộng**: Di chuyển `left--` và `right++` để kiểm tra palindrome lớn hơn
- **Trả về**: Số lượng palindrome tìm được từ center này

**Ví dụ minh họa expandAndCount:**

Với `s = "aaa"`, `i = 0` (ký tự 'a' đầu tiên):

```
expandAndCount(s, 0, 0): // Odd length
  left=0, right=0: s[0]=='a', s[0]=='a' → match → count = 1 ("a")
  left=-1, right=1: out of bounds → stop
  Return: 1

expandAndCount(s, 0, 1): // Even length
  left=0, right=1: s[0]=='a', s[1]=='a' → match → count = 1 ("aa")
  left=-1, right=2: out of bounds → stop
  Return: 1

Tổng từ center 0: 1 + 1 = 2 palindromes
```

Với `s = "aaa"`, `i = 1` (ký tự 'a' giữa):

```
expandAndCount(s, 1, 1): // Odd length
  left=1, right=1: s[1]=='a', s[1]=='a' → match → count = 1 ("a" tại index 1)
  left=0, right=2: s[0]=='a', s[2]=='a' → match → count = 2 ("aaa")
  left=-1, right=3: out of bounds → stop
  Return: 2

expandAndCount(s, 1, 2): // Even length
  left=1, right=2: s[1]=='a', s[2]=='a' → match → count = 1 ("aa" từ index 1-2)
  left=0, right=3: out of bounds → stop
  Return: 1

Tổng từ center 1: 2 + 1 = 3 palindromes
```

#### 4. Trả về kết quả (dòng 17)
```go
return count
```
- Trả về tổng số lượng substring palindromic

### Tại sao thuật toán này hoạt động?

1. **Tính chất palindrome**: Mỗi palindrome đều có một center (hoặc giữa hai ký tự cho even length)
2. **Bao phủ đầy đủ**: Duyệt qua tất cả các vị trí có thể làm center, đảm bảo không bỏ sót palindrome nào
3. **Đếm chính xác**: Mỗi palindrome được đếm đúng một lần khi nó được tìm thấy từ center của nó
4. **Mở rộng hiệu quả**: Mở rộng từ center ra hai bên và đếm mỗi lần mở rộng thành công

### Tại sao mỗi palindrome chỉ được đếm một lần?

Mỗi palindrome có một center duy nhất (hoặc một cặp center duy nhất cho even length):
- Palindrome độ dài lẻ: Center là ký tự ở giữa (duy nhất)
- Palindrome độ dài chẵn: Center là giữa hai ký tự ở giữa (duy nhất)

Ví dụ với `"aaa"`:
- `"a"` tại index 0: Center tại index 0 → đếm 1 lần
- `"a"` tại index 1: Center tại index 1 → đếm 1 lần
- `"a"` tại index 2: Center tại index 2 → đếm 1 lần
- `"aa"` từ index 0-1: Center giữa index 0-1 → đếm 1 lần
- `"aa"` từ index 1-2: Center giữa index 1-2 → đếm 1 lần
- `"aaa"`: Center tại index 1 → đếm 1 lần

Tổng: 6 palindromes, mỗi palindrome được đếm đúng 1 lần ✓

### Độ phức tạp

- **Thời gian**: O(n²)
  - Vòng lặp ngoài: O(n) - duyệt qua n vị trí
  - Vòng lặp trong (expandAndCount): O(n) trong trường hợp xấu nhất (toàn bộ chuỗi là palindrome)
  - Tổng: O(n × n) = O(n²)

- **Không gian**: O(1)
  - Chỉ sử dụng một số biến để lưu trữ: `count`, `i`, `left`, `right`
  - Hàm `expandAndCount` chỉ sử dụng tham số và biến cục bộ

## Test Cases

### Test Case 1 - Example 1
```go
Input: s = "abc"
Output: 3
```
Giải thích: 3 palindromes: `"a"`, `"b"`, `"c"`

### Test Case 2 - Example 2
```go
Input: s = "aaa"
Output: 6
```
Giải thích: 6 palindromes: `"a"` (3 lần), `"aa"` (2 lần), `"aaa"` (1 lần)

### Test Case 3 - Một ký tự
```go
Input: s = "a"
Output: 1
```
Giải thích: Chính ký tự đó là palindrome duy nhất

### Test Case 4 - Hai ký tự giống nhau
```go
Input: s = "aa"
Output: 3
```
Giải thích: 3 palindromes: `"a"` (2 lần), `"aa"` (1 lần)

### Test Case 5 - Hai ký tự khác nhau
```go
Input: s = "ab"
Output: 2
```
Giải thích: 2 palindromes: `"a"`, `"b"`

### Test Case 6 - Chuỗi palindrome
```go
Input: s = "racecar"
Output: 10
```
Giải thích: Bao gồm các palindrome như `"r"`, `"a"`, `"c"`, `"e"`, `"racecar"`, v.v.

### Test Case 7 - Tất cả ký tự giống nhau (4 ký tự)
```go
Input: s = "aaaa"
Output: 10
```
Giải thích: 
- `"a"`: 4 lần
- `"aa"`: 3 lần
- `"aaa"`: 2 lần
- `"aaaa"`: 1 lần
- Tổng: 4 + 3 + 2 + 1 = 10

### Test Case 8 - Palindrome độ dài chẵn
```go
Input: s = "abba"
Output: 6
```
Giải thích: Bao gồm `"a"`, `"b"`, `"b"`, `"a"`, `"bb"`, `"abba"`

### Test Case 9 - Không có palindrome dài hơn 1
```go
Input: s = "abcdef"
Output: 6
```
Giải thích: Mỗi ký tự là một palindrome

### Test Case 10 - Chuỗi dài hơn
```go
Input: s = "abcba"
Output: 7
```
Giải thích: Bao gồm các palindrome như `"a"`, `"b"`, `"c"`, `"b"`, `"a"`, `"bcb"`, `"abcba"`

## Chạy test

Để chạy các test case:

```bash
go test ./internal/string/palindromic_substrings/
```

Hoặc chạy với verbose mode để xem chi tiết:

```bash
go test -v ./internal/string/palindromic_substrings/
```

## So sánh với các cách tiếp cận khác

| Cách tiếp cận | Độ phức tạp thời gian | Độ phức tạp không gian | Ghi chú |
|---------------|----------------------|----------------------|---------|
| Brute Force | O(n³) | O(1) | Kiểm tra tất cả các substring |
| Dynamic Programming | O(n²) | O(n²) | Sử dụng bảng DP |
| Expand Around Centers | O(n²) | O(1) | **Tối ưu về không gian** - sử dụng trong code này |

**Kết luận:** Expand Around Centers là cách tiếp cận tốt vì:
- Độ phức tạp thời gian O(n²) - tốt hơn O(n³) của Brute Force
- Độ phức tạp không gian O(1) - tốt hơn O(n²) của DP
- Code đơn giản, dễ hiểu và maintain
- Với constraint `n <= 1000`, O(n²) là đủ nhanh

## So sánh với bài Longest Palindromic Substring

Bài toán này khác với "Longest Palindromic Substring" ở chỗ:

| Khía cạnh | Longest Palindromic Substring | Palindromic Substrings |
|-----------|------------------------------|------------------------|
| Mục tiêu | Tìm substring palindrome **dài nhất** | **Đếm** tất cả substring palindrome |
| Kết quả | Một chuỗi (substring) | Một số nguyên (số lượng) |
| Thuật toán | Expand và lưu substring dài nhất | Expand và đếm mỗi lần thành công |
| Độ phức tạp | O(n²) thời gian, O(1) không gian | O(n²) thời gian, O(1) không gian |

Cả hai đều sử dụng kỹ thuật "Expand Around Centers" nhưng với mục đích khác nhau.

## Lưu ý

1. **Đếm chính xác**: Mỗi palindrome được đếm đúng một lần khi nó được tìm thấy từ center của nó
2. **Hai loại center**: Phải xử lý cả odd length (center tại ký tự) và even length (center giữa hai ký tự)
3. **Edge cases**: 
   - Chuỗi rỗng: không có palindrome (theo constraint, `s.length >= 1`)
   - Một ký tự: luôn có 1 palindrome
   - Tất cả ký tự giống nhau: số lượng palindrome = n + (n-1) + ... + 1 = n(n+1)/2
