# Word Break

## Mô tả bài toán

Cho một chuỗi `s` và một từ điển các chuỗi `wordDict`, trả về `true` nếu `s` có thể được phân tách thành một chuỗi các từ được phân cách bởi khoảng trắng, trong đó mỗi từ đều có trong `wordDict`.

**Lưu ý:**
- Cùng một từ trong từ điển có thể được sử dụng nhiều lần trong quá trình phân tách
- Không cần sử dụng tất cả các từ trong từ điển

**Ví dụ:**

### Example 1:
- Input: `s = "leetcode"`, `wordDict = ["leet","code"]`
- Output: `true`
- Giải thích: `"leetcode"` có thể được phân tách thành `"leet code"`

### Example 2:
- Input: `s = "applepenapple"`, `wordDict = ["apple","pen"]`
- Output: `true`
- Giải thích: `"applepenapple"` có thể được phân tách thành `"apple pen apple"`. Lưu ý rằng từ "apple" được sử dụng lại nhiều lần.

### Example 3:
- Input: `s = "catsandog"`, `wordDict = ["cats","dog","sand","and","cat"]`
- Output: `false`
- Giải thích: Không có cách nào để phân tách `"catsandog"` thành các từ trong từ điển

**Ràng buộc:**
- `1 <= s.length <= 300`
- `1 <= wordDict.length <= 1000`
- `1 <= wordDict[i].length <= 20`
- `s` và `wordDict[i]` chỉ chứa chữ cái tiếng Anh viết thường
- Tất cả các chuỗi trong `wordDict` là duy nhất

## Phân tích thuật toán

### Cách tiếp cận: Dynamic Programming (DP)

Đây là một bài toán kinh điển sử dụng Dynamic Programming. Ý tưởng chính là:

**Vấn đề con:** Với mỗi vị trí `i` trong chuỗi, kiểm tra xem chuỗi con `s[0:i]` có thể được phân tách thành các từ trong `wordDict` hay không.

**Quan hệ đệ quy:**
- `dp[i] = true` nếu tồn tại một vị trí `j < i` sao cho:
  - `dp[j] = true` (chuỗi `s[0:j]` có thể phân tách được)
  - `s[j:i]` là một từ trong `wordDict`

**Trường hợp cơ sở:**
- `dp[0] = true` (chuỗi rỗng luôn có thể phân tách được)

### Các bước thực hiện

1. **Tạo Hash Set từ wordDict**: Chuyển `wordDict` thành một map để kiểm tra sự tồn tại của từ trong O(1) thay vì O(n)

2. **Khởi tạo mảng DP**: 
   - `dp[i]` đại diện cho việc chuỗi `s[0:i]` có thể phân tách được hay không
   - `dp[0] = true` (chuỗi rỗng)

3. **Duyệt qua từng vị trí**: Với mỗi vị trí `i` từ 1 đến n:
   - Kiểm tra tất cả các vị trí `j` từ 0 đến i-1
   - Nếu `dp[j] = true` và `s[j:i]` là một từ trong `wordDict`, thì `dp[i] = true`

4. **Kết quả**: `dp[n]` cho biết toàn bộ chuỗi `s` có thể phân tách được hay không

### Ví dụ minh họa

Với input: `s = "leetcode"`, `wordDict = ["leet","code"]`

```
Khởi tạo:
  dp[0] = true (chuỗi rỗng)
  wordSet = {"leet": true, "code": true}

i = 1: s[0:1] = "l"
  j = 0: dp[0] = true, s[0:1] = "l" không có trong wordSet
  dp[1] = false

i = 2: s[0:2] = "le"
  j = 0: dp[0] = true, s[0:2] = "le" không có trong wordSet
  j = 1: dp[1] = false
  dp[2] = false

i = 3: s[0:3] = "lee"
  j = 0: dp[0] = true, s[0:3] = "lee" không có trong wordSet
  j = 1, 2: dp[1], dp[2] = false
  dp[3] = false

i = 4: s[0:4] = "leet"
  j = 0: dp[0] = true, s[0:4] = "leet" có trong wordSet ✓
  dp[4] = true

i = 5: s[0:5] = "leetc"
  j = 0: dp[0] = true, s[0:5] = "leetc" không có trong wordSet
  j = 1, 2, 3: dp[1], dp[2], dp[3] = false
  j = 4: dp[4] = true, s[4:5] = "c" không có trong wordSet
  dp[5] = false

i = 6: s[0:6] = "leetco"
  j = 0: dp[0] = true, s[0:6] = "leetco" không có trong wordSet
  j = 1, 2, 3: dp[1], dp[2], dp[3] = false
  j = 4: dp[4] = true, s[4:6] = "co" không có trong wordSet
  j = 5: dp[5] = false
  dp[6] = false

i = 7: s[0:7] = "leetcod"
  j = 0: dp[0] = true, s[0:7] = "leetcod" không có trong wordSet
  j = 1, 2, 3: dp[1], dp[2], dp[3] = false
  j = 4: dp[4] = true, s[4:7] = "cod" không có trong wordSet
  j = 5, 6: dp[5], dp[6] = false
  dp[7] = false

i = 8: s[0:8] = "leetcode"
  j = 0: dp[0] = true, s[0:8] = "leetcode" không có trong wordSet
  j = 1, 2, 3: dp[1], dp[2], dp[3] = false
  j = 4: dp[4] = true, s[4:8] = "code" có trong wordSet ✓
  dp[8] = true

Kết quả: dp[8] = true
```

Với input: `s = "catsandog"`, `wordDict = ["cats","dog","sand","and","cat"]`

```
Khởi tạo:
  dp[0] = true
  wordSet = {"cats": true, "dog": true, "sand": true, "and": true, "cat": true}

i = 1, 2: dp[1] = false, dp[2] = false

i = 3: s[0:3] = "cat"
  j = 0: dp[0] = true, s[0:3] = "cat" có trong wordSet ✓
  dp[3] = true

i = 4: s[0:4] = "cats"
  j = 0: dp[0] = true, s[0:4] = "cats" có trong wordSet ✓
  dp[4] = true

i = 5, 6: dp[5] = false, dp[6] = false

i = 7: s[0:7] = "catsand"
  j = 0: dp[0] = true, s[0:7] = "catsand" không có trong wordSet
  j = 3: dp[3] = true, s[3:7] = "sand" có trong wordSet ✓
  dp[7] = true

i = 8: s[0:8] = "catsando"
  j = 0: dp[0] = true, s[0:8] = "catsando" không có trong wordSet
  j = 3: dp[3] = true, s[3:8] = "sando" không có trong wordSet
  j = 4: dp[4] = true, s[4:8] = "sand" có trong wordSet, nhưng cần kiểm tra tiếp
  j = 4: dp[4] = true, s[4:8] = "sand" có trong wordSet ✓
  dp[8] = true (từ "cats" + "sand")

i = 9: s[0:9] = "catsandog"
  j = 0: dp[0] = true, s[0:9] = "catsandog" không có trong wordSet
  j = 3: dp[3] = true, s[3:9] = "sandog" không có trong wordSet
  j = 4: dp[4] = true, s[4:9] = "sandog" không có trong wordSet
  j = 7: dp[7] = true, s[7:9] = "og" không có trong wordSet
  j = 8: dp[8] = true, s[8:9] = "g" không có trong wordSet
  dp[9] = false

Kết quả: dp[9] = false
```

## Giải thích code

### Cấu trúc hàm

```1:35:internal/array/word_break/word_break.go
package array

// wordBreak kiểm tra xem chuỗi s có thể được phân tách thành các từ trong wordDict hay không.
//
// Sử dụng Dynamic Programming với memoization:
// - dp[i] = true nếu s[0:i] có thể được phân tách thành các từ trong wordDict
// - Với mỗi vị trí i, kiểm tra tất cả các từ trong wordDict:
//   Nếu s[j:i] là một từ trong wordDict và dp[j] = true, thì dp[i] = true
func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	
	// Tạo map để kiểm tra từ nhanh hơn (O(1) thay vì O(n))
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}
	
	// dp[i] = true nếu s[0:i] có thể được phân tách thành các từ trong wordDict
	dp := make([]bool, n+1)
	dp[0] = true // Chuỗi rỗng luôn có thể phân tách được
	
	// Duyệt qua từng vị trí trong chuỗi
	for i := 1; i <= n; i++ {
		// Kiểm tra tất cả các vị trí j trước đó
		for j := 0; j < i; j++ {
			// Nếu s[0:j] có thể phân tách được (dp[j] = true)
			// và s[j:i] là một từ trong wordDict
			// thì s[0:i] cũng có thể phân tách được
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
				break // Đã tìm thấy cách phân tách, không cần kiểm tra tiếp
			}
		}
	}
	
	return dp[n]
}
```

### Chi tiết từng phần

#### 1. Tạo Hash Set từ wordDict (dòng 11-14)
```go
wordSet := make(map[string]bool)
for _, word := range wordDict {
    wordSet[word] = true
}
```
- Chuyển `wordDict` từ slice thành map để kiểm tra sự tồn tại của từ trong O(1) thay vì O(n)
- Giảm độ phức tạp thời gian từ O(n²) xuống O(n²) nhưng với hằng số nhỏ hơn nhiều

#### 2. Khởi tạo mảng DP (dòng 16-18)
```go
dp := make([]bool, n+1)
dp[0] = true
```
- `dp[i]` đại diện cho việc chuỗi con `s[0:i]` có thể phân tách được hay không
- `dp[0] = true` là trường hợp cơ sở: chuỗi rỗng luôn có thể phân tách được

#### 3. Vòng lặp chính - Duyệt qua từng vị trí (dòng 20-32)
```go
for i := 1; i <= n; i++ {
    for j := 0; j < i; j++ {
        if dp[j] && wordSet[s[j:i]] {
            dp[i] = true
            break
        }
    }
}
```

**Giải thích:**
- Với mỗi vị trí `i`, kiểm tra tất cả các vị trí `j` trước đó (0 đến i-1)
- Nếu:
  - `dp[j] = true`: Chuỗi `s[0:j]` đã có thể phân tách được
  - `s[j:i]` là một từ trong `wordDict`: Phần còn lại từ `j` đến `i` là một từ hợp lệ
- Thì `dp[i] = true`: Chuỗi `s[0:i]` có thể phân tách được
- `break`: Khi đã tìm thấy một cách phân tách, không cần kiểm tra các vị trí `j` khác

#### 4. Trả về kết quả (dòng 34)
```go
return dp[n]
```
- `dp[n]` cho biết toàn bộ chuỗi `s` (từ vị trí 0 đến n) có thể phân tách được hay không

### Độ phức tạp

- **Thời gian**: O(n² × m)
  - `n`: độ dài chuỗi `s`
  - `m`: độ dài trung bình của các từ trong `wordDict`
  - Vòng lặp ngoài: O(n)
  - Vòng lặp trong: O(n)
  - Kiểm tra substring `s[j:i]` và lookup trong map: O(m)
  - Tổng: O(n² × m)

- **Không gian**: O(n + k)
  - `n`: mảng `dp` có kích thước n+1
  - `k`: hash set `wordSet` lưu trữ tất cả các từ trong `wordDict`
  - Trong trường hợp xấu nhất: O(n + k) với k là tổng độ dài tất cả các từ

### Tối ưu hóa

Có thể tối ưu thêm bằng cách:
1. **Giới hạn độ dài từ**: Chỉ kiểm tra các substring có độ dài bằng với các từ trong `wordDict`
2. **Early termination**: Nếu không có từ nào trong `wordDict` có thể bắt đầu từ vị trí hiện tại, có thể bỏ qua một số kiểm tra

Tuy nhiên, giải pháp hiện tại đã đủ hiệu quả và dễ hiểu cho hầu hết các trường hợp.

## Test Cases

### Test Case 1: Example 1
```go
Input: s = "leetcode", wordDict = ["leet","code"]
Output: true
```
Phân tách: `"leet" + "code"`

### Test Case 2: Example 2
```go
Input: s = "applepenapple", wordDict = ["apple","pen"]
Output: true
```
Phân tách: `"apple" + "pen" + "apple"` (từ "apple" được sử dụng lại)

### Test Case 3: Example 3
```go
Input: s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]
Output: false
```
Không thể phân tách: Có thể tạo "cats" + "and" nhưng "og" không phải là từ hợp lệ, hoặc "cat" + "sand" nhưng "og" vẫn không hợp lệ.

### Test Case 4: Single word match
```go
Input: s = "leetcode", wordDict = ["leetcode"]
Output: true
```
Toàn bộ chuỗi là một từ trong từ điển.

### Test Case 5: Single word no match
```go
Input: s = "leetcode", wordDict = ["leet"]
Output: false
```
Chuỗi không thể phân tách hoàn toàn.

### Test Case 6: Empty string
```go
Input: s = "", wordDict = ["leet","code"]
Output: true
```
Chuỗi rỗng luôn có thể phân tách được (trường hợp cơ sở).

### Test Case 7: Repeated words
```go
Input: s = "aaaaaaa", wordDict = ["aaaa","aaa"]
Output: true
```
Có nhiều cách phân tách khác nhau với các từ lặp lại.

## Chạy test

Để chạy các test case:

```bash
go test ./internal/array/word_break/
```

Hoặc chạy với verbose mode để xem chi tiết:

```bash
go test -v ./internal/array/word_break/
```

## Mở rộng

### Word Break II

Một biến thể của bài toán này là **Word Break II** (LeetCode 140), yêu cầu trả về tất cả các cách phân tách có thể thay vì chỉ kiểm tra xem có thể phân tách được hay không. Giải pháp sẽ sử dụng DP kết hợp với backtracking để tìm tất cả các cách phân tách.
