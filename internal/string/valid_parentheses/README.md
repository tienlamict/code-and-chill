# Valid Parentheses

## Mô tả bài toán

Cho một chuỗi `s` chỉ chứa các ký tự `'('`, `')'`, `'{'`, `'}'`, `'['` và `']'`, xác định xem chuỗi input có hợp lệ hay không.

Một chuỗi input được coi là hợp lệ nếu:

1. Dấu ngoặc mở phải được đóng bởi cùng loại dấu ngoặc
2. Dấu ngoặc mở phải được đóng theo đúng thứ tự
3. Mỗi dấu ngoặc đóng phải có một dấu ngoặc mở tương ứng cùng loại

**Ví dụ:**

### Example 1:
- Input: `s = "()"`
- Output: `true`

### Example 2:
- Input: `s = "()[]{}"`
- Output: `true`

### Example 3:
- Input: `s = "(]"`
- Output: `false`

### Example 4:
- Input: `s = "([])"`
- Output: `true`

### Example 5:
- Input: `s = "([)]"`
- Output: `false`

**Ràng buộc:**
- `1 <= s.length <= 10^4`
- `s` chỉ chứa các ký tự `'()[]{}'`

## Phân tích thuật toán

### Cách tiếp cận: Stack

Đây là bài toán kinh điển sử dụng cấu trúc dữ liệu **Stack** (ngăn xếp).

**Ý tưởng:**
- Sử dụng nguyên lý LIFO (Last In First Out) của Stack
- Dấu ngoặc mở gần nhất phải được đóng trước (giống như stack)
- Khi gặp dấu ngoặc mở, push vào stack
- Khi gặp dấu ngoặc đóng, kiểm tra xem có match với dấu ngoặc mở trên đỉnh stack không

**Các bước thực hiện:**

1. **Tạo Stack**: Sử dụng slice hoặc stack để lưu các dấu ngoặc mở
2. **Duyệt qua chuỗi**: Với mỗi ký tự trong chuỗi:
   - Nếu là dấu ngoặc mở `'('`, `'{'`, `'['`: Push vào stack
   - Nếu là dấu ngoặc đóng `')'`, `'}'`, `']'`:
     - Kiểm tra xem stack có rỗng không (nếu rỗng → không hợp lệ)
     - Lấy dấu ngoặc mở trên đỉnh stack
     - Kiểm tra xem có match không (nếu không match → không hợp lệ)
     - Pop khỏi stack nếu match
3. **Kiểm tra kết quả**: Sau khi duyệt xong, stack phải rỗng (nếu không rỗng → không hợp lệ)

### Ví dụ minh họa

Với input: `s = "()"`

```
Bước 0: stack = []
Bước 1: char = '(' → push vào stack
        stack = ['(']
Bước 2: char = ')' → kiểm tra stack
        stack không rỗng
        top = '(' → match với ')'
        pop khỏi stack
        stack = []
Kết quả: stack rỗng → true
```

Với input: `s = "()[]{}"`

```
Bước 0: stack = []
Bước 1: char = '(' → push
        stack = ['(']
Bước 2: char = ')' → match và pop
        stack = []
Bước 3: char = '[' → push
        stack = ['[']
Bước 4: char = ']' → match và pop
        stack = []
Bước 5: char = '{' → push
        stack = ['{']
Bước 6: char = '}' → match và pop
        stack = []
Kết quả: stack rỗng → true
```

Với input: `s = "(]"`

```
Bước 0: stack = []
Bước 1: char = '(' → push
        stack = ['(']
Bước 2: char = ']' → kiểm tra stack
        stack không rỗng
        top = '(' → không match với ']'
        return false
Kết quả: false
```

Với input: `s = "([)]"`

```
Bước 0: stack = []
Bước 1: char = '(' → push
        stack = ['(']
Bước 2: char = '[' → push
        stack = ['(', '[']
Bước 3: char = ')' → kiểm tra stack
        stack không rỗng
        top = '[' → không match với ')'
        return false
Kết quả: false
```

Với input: `s = "([])"`

```
Bước 0: stack = []
Bước 1: char = '(' → push
        stack = ['(']
Bước 2: char = '[' → push
        stack = ['(', '[']
Bước 3: char = ']' → kiểm tra stack
        top = '[' → match với ']'
        pop → stack = ['(']
Bước 4: char = ')' → kiểm tra stack
        top = '(' → match với ')'
        pop → stack = []
Kết quả: stack rỗng → true
```

## Giải thích code

### Cấu trúc hàm

```1:48:internal/string/valid_parentheses/valid_parentheses.go
package string

// isValid kiểm tra xem chuỗi chứa các dấu ngoặc có hợp lệ hay không.
//
// Sử dụng Stack để kiểm tra tính hợp lệ:
// - Khi gặp dấu ngoặc mở '(', '{', '[', push vào stack
// - Khi gặp dấu ngoặc đóng ')', '}', ']', kiểm tra xem có match với dấu ngoặc mở trên đỉnh stack không
// - Nếu match, pop khỏi stack; nếu không match hoặc stack rỗng, return false
// - Cuối cùng, stack phải rỗng để chuỗi hợp lệ
//
// Độ phức tạp: O(n) thời gian, O(n) không gian
func isValid(s string) bool {
	// Stack để lưu các dấu ngoặc mở
	stack := []rune{}

	// Map để ánh xạ dấu ngoặc đóng với dấu ngoặc mở tương ứng
	closingToOpening := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	// Duyệt qua từng ký tự trong chuỗi
	for _, char := range s {
		// Nếu là dấu ngoặc đóng
		if opening, isClosing := closingToOpening[char]; isClosing {
			// Kiểm tra xem stack có rỗng không
			if len(stack) == 0 {
				return false
			}

			// Lấy dấu ngoặc mở trên đỉnh stack
			top := stack[len(stack)-1]

			// Nếu không match, chuỗi không hợp lệ
			if top != opening {
				return false
			}

			// Pop khỏi stack (match thành công)
			stack = stack[:len(stack)-1]
		} else {
			// Nếu là dấu ngoặc mở, push vào stack
			stack = append(stack, char)
		}
	}

	// Stack phải rỗng để chuỗi hợp lệ
	return len(stack) == 0
}
```

### Chi tiết từng phần

#### 1. Khởi tạo Stack (dòng 14)
```go
stack := []rune{}
```
- Sử dụng slice của `rune` để lưu các dấu ngoặc mở
- `rune` là kiểu dữ liệu trong Go để xử lý Unicode characters

#### 2. Tạo Map ánh xạ (dòng 16-20)
```go
closingToOpening := map[rune]rune{
    ')': '(',
    '}': '{',
    ']': '[',
}
```
- Map để ánh xạ dấu ngoặc đóng với dấu ngoặc mở tương ứng
- Giúp kiểm tra match nhanh chóng trong O(1)

#### 3. Vòng lặp duyệt chuỗi (dòng 22-42)
```go
for _, char := range s {
    if opening, isClosing := closingToOpening[char]; isClosing {
        // Xử lý dấu ngoặc đóng
    } else {
        // Xử lý dấu ngoặc mở
    }
}
```

**Xử lý dấu ngoặc đóng (dòng 24-37):**
```go
if opening, isClosing := closingToOpening[char]; isClosing {
    if len(stack) == 0 {
        return false
    }
    top := stack[len(stack)-1]
    if top != opening {
        return false
    }
    stack = stack[:len(stack)-1]
}
```

**Giải thích:**
- `isClosing`: Kiểm tra xem ký tự có phải là dấu ngoặc đóng không
- `len(stack) == 0`: Nếu stack rỗng, không có dấu ngoặc mở để match → return false
- `top := stack[len(stack)-1]`: Lấy dấu ngoặc mở trên đỉnh stack
- `top != opening`: Nếu không match → return false
- `stack = stack[:len(stack)-1]`: Pop khỏi stack (match thành công)

**Xử lý dấu ngoặc mở (dòng 38-40):**
```go
else {
    stack = append(stack, char)
}
```
- Nếu là dấu ngoặc mở, push vào stack

#### 4. Kiểm tra kết quả (dòng 45)
```go
return len(stack) == 0
```
- Sau khi duyệt xong, stack phải rỗng
- Nếu stack không rỗng, có dấu ngoặc mở chưa được đóng → return false

### Độ phức tạp

- **Thời gian**: O(n)
  - Duyệt qua chuỗi một lần: O(n)
  - Mỗi thao tác push/pop trên stack: O(1)
  - Tổng: O(n)

- **Không gian**: O(n)
  - Trong trường hợp xấu nhất, tất cả các ký tự đều là dấu ngoặc mở
  - Stack có thể chứa tối đa n/2 phần tử (nếu chuỗi hợp lệ)
  - Trong trường hợp xấu nhất: O(n)

### Tại sao thuật toán này đúng?

1. **Nguyên lý LIFO**: Dấu ngoặc mở gần nhất phải được đóng trước, phù hợp với tính chất của Stack.

2. **Kiểm tra match**: Sử dụng map để kiểm tra match nhanh chóng và chính xác.

3. **Xử lý edge cases**:
   - Stack rỗng khi gặp dấu ngoặc đóng → không hợp lệ
   - Stack không rỗng sau khi duyệt xong → không hợp lệ
   - Dấu ngoặc không match → không hợp lệ

4. **Tính đầy đủ**: Duyệt qua tất cả các ký tự và kiểm tra mọi trường hợp.

## Test Cases

### Test Case 1: Example 1
```go
Input: s = "()"
Output: true
```
Dấu ngoặc đơn đơn giản.

### Test Case 2: Example 2
```go
Input: s = "()[]{}"
Output: true
```
Nhiều loại dấu ngoặc khác nhau.

### Test Case 3: Example 3
```go
Input: s = "(]"
Output: false
```
Dấu ngoặc không match.

### Test Case 4: Example 4
```go
Input: s = "([])"
Output: true
```
Dấu ngoặc lồng nhau hợp lệ.

### Test Case 5: Example 5
```go
Input: s = "([)]"
Output: false
```
Dấu ngoặc lồng nhau nhưng sai thứ tự.

### Test Case 6: Empty string
```go
Input: s = ""
Output: true
```
Chuỗi rỗng được coi là hợp lệ.

### Test Case 7: Single opening
```go
Input: s = "("
Output: false
```
Dấu ngoặc mở không được đóng.

### Test Case 8: Nested valid
```go
Input: s = "({[]})"
Output: true
```
Nhiều lớp lồng nhau hợp lệ.

### Test Case 9: Unmatched opening
```go
Input: s = "((("
Output: false
```
Nhiều dấu ngoặc mở không được đóng.

### Test Case 10: Wrong order
```go
Input: s = "([)]"
Output: false
```
Dấu ngoặc đóng sai thứ tự.

## Chạy test

Để chạy các test case:

```bash
go test ./internal/string/valid_parentheses/
```

Hoặc chạy với verbose mode để xem chi tiết:

```bash
go test -v ./internal/string/valid_parentheses/
```

## Mở rộng

### Generate Parentheses (LeetCode 22)

Một biến thể của bài toán này là tạo tất cả các chuỗi dấu ngoặc hợp lệ với n cặp dấu ngoặc. Có thể giải bằng backtracking.

### Longest Valid Parentheses (LeetCode 32)

Tìm độ dài của chuỗi dấu ngoặc hợp lệ dài nhất trong một chuỗi. Yêu cầu kỹ thuật phức tạp hơn với dynamic programming hoặc stack.

### Remove Invalid Parentheses (LeetCode 301)

Xóa số lượng dấu ngoặc tối thiểu để chuỗi trở thành hợp lệ. Yêu cầu backtracking và xử lý duplicate.

### Ứng dụng thực tế

Bài toán này có ứng dụng trong:
- **Compiler và Parser**: Kiểm tra cú pháp trong code
- **Expression Evaluation**: Đánh giá biểu thức toán học
- **XML/JSON Parsing**: Kiểm tra tính hợp lệ của cấu trúc
- **Code Editor**: Highlight matching brackets
- **Syntax Validation**: Kiểm tra cú pháp trong các ngôn ngữ lập trình
