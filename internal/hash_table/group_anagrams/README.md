# Group Anagrams

## Mô tả bài toán

Cho một mảng các chuỗi `strs`, nhóm các anagrams lại với nhau. Bạn có thể trả về câu trả lời theo bất kỳ thứ tự nào.

**Anagram** là một từ hoặc cụm từ được tạo thành bằng cách sắp xếp lại các ký tự của một từ hoặc cụm từ khác, thường sử dụng tất cả các ký tự gốc đúng một lần.

**Ví dụ:**

### Example 1:
- Input: `strs = ["eat","tea","tan","ate","nat","bat"]`
- Output: `[["bat"],["nat","tan"],["ate","eat","tea"]]`
- Giải thích:
  - Không có string nào trong `strs` có thể được sắp xếp lại để tạo thành "bat".
  - Các string "nat" và "tan" là anagrams vì chúng có thể được sắp xếp lại để tạo thành nhau.
  - Các string "ate", "eat", và "tea" là anagrams vì chúng có thể được sắp xếp lại để tạo thành nhau.

```
Anagrams groups:
- Group 1: ["bat"] (không có anagram khác)
- Group 2: ["nat", "tan"] (cùng ký tự: a, n, t)
- Group 3: ["ate", "eat", "tea"] (cùng ký tự: a, e, t)
```

### Example 2:
- Input: `strs = [""]`
- Output: `[[""]]`
- Giải thích: Chỉ có một chuỗi rỗng.

### Example 3:
- Input: `strs = ["a"]`
- Output: `[["a"]]`
- Giải thích: Chỉ có một chuỗi.

**Ràng buộc:**
- `1 <= strs.length <= 10^4`
- `0 <= strs[i].length <= 100`
- `strs[i]` chỉ chứa chữ cái tiếng Anh viết thường

## Phân tích thuật toán

### Cách tiếp cận 1: Brute Force (O(n² * k) time, O(n * k) space)

Với mỗi string, so sánh với tất cả các string khác để tìm anagrams:

```go
func groupAnagramsBruteForce(strs []string) [][]string {
    result := [][]string{}
    used := make([]bool, len(strs))
    
    for i := 0; i < len(strs); i++ {
        if used[i] {
            continue
        }
        
        group := []string{strs[i]}
        used[i] = true
        
        for j := i + 1; j < len(strs); j++ {
            if !used[j] && areAnagrams(strs[i], strs[j]) {
                group = append(group, strs[j])
                used[j] = true
            }
        }
        
        result = append(result, group)
    }
    
    return result
}

func areAnagrams(s1, s2 string) bool {
    if len(s1) != len(s2) {
        return false
    }
    
    count1 := make(map[rune]int)
    count2 := make(map[rune]int)
    
    for _, char := range s1 {
        count1[char]++
    }
    for _, char := range s2 {
        count2[char]++
    }
    
    return reflect.DeepEqual(count1, count2)
}
```

**Độ phức tạp:**
- Thời gian: O(n² * k) - với n là số lượng strings và k là độ dài trung bình
- Không gian: O(n * k)

**Nhược điểm:** Quá chậm với n lên tới 10^4

### Cách tiếp cận 2: Hash Map với Sorted String (O(n * k * log(k)) time, O(n * k) space) ⭐ Tối ưu

**Ý tưởng chính:**

Các anagrams có cùng các ký tự, chỉ khác thứ tự. Nếu sắp xếp các ký tự trong mỗi string, các anagrams sẽ có cùng kết quả sau khi sắp xếp.

**Thuật toán:**

1. **Tạo hash map**: Key là string đã được sắp xếp, Value là danh sách các anagrams
2. **Duyệt qua từng string**:
   - Sắp xếp các ký tự trong string để tạo key
   - Thêm string vào nhóm tương ứng trong map
3. **Chuyển đổi map thành kết quả**: Lấy tất cả các nhóm từ map

**Độ phức tạp:**
- Thời gian: O(n * k * log(k))
  - Duyệt qua n strings: O(n)
  - Sắp xếp mỗi string (độ dài k): O(k * log(k))
  - Tổng: O(n * k * log(k))
- Không gian: O(n * k)
  - Map lưu trữ tất cả strings

**Ưu điểm:**
- Đơn giản, dễ hiểu
- Hiệu quả với n lớn
- Code ngắn gọn

### Cách tiếp cận 3: Hash Map với Character Count (O(n * k) time, O(n * k) space)

Thay vì sắp xếp, đếm số lượng mỗi ký tự và dùng làm key:

```go
func groupAnagramsCount(strs []string) [][]string {
    anagramMap := make(map[string][]string)
    
    for _, str := range strs {
        // Đếm số lượng mỗi ký tự
        count := make([]int, 26)
        for _, char := range str {
            count[char-'a']++
        }
        
        // Tạo key từ count array
        key := ""
        for i := 0; i < 26; i++ {
            key += string(rune('a'+i)) + strconv.Itoa(count[i])
        }
        
        anagramMap[key] = append(anagramMap[key], str)
    }
    
    result := make([][]string, 0, len(anagramMap))
    for _, group := range anagramMap {
        result = append(result, group)
    }
    
    return result
}
```

**Độ phức tạp:**
- Thời gian: O(n * k) - không cần sort
- Không gian: O(n * k)

**Ưu điểm:** Nhanh hơn một chút (không cần sort)
**Nhược điểm:** Key phức tạp hơn, code dài hơn

### So sánh các cách tiếp cận

| Cách tiếp cận | Thời gian | Không gian | Ưu điểm | Nhược điểm |
|--------------|-----------|------------|---------|------------|
| Brute Force | O(n² * k) | O(n * k) | Đơn giản | Quá chậm |
| Sorted String | O(n * k * log(k)) | O(n * k) | Đơn giản, dễ hiểu | Cần sort |
| Character Count | O(n * k) | O(n * k) | Nhanh nhất | Code phức tạp hơn |

## Giải thích chi tiết thuật toán Sorted String

### Ví dụ minh họa

Với input: `strs = ["eat","tea","tan","ate","nat","bat"]`

```
Bước 0: Khởi tạo
anagramMap = {}

Bước 1: Xử lý "eat"
sortedStr = sortString("eat") = "aet"
anagramMap["aet"] = ["eat"]
anagramMap = {"aet": ["eat"]}

Bước 2: Xử lý "tea"
sortedStr = sortString("tea") = "aet"
anagramMap["aet"] = ["eat", "tea"]
anagramMap = {"aet": ["eat", "tea"]}

Bước 3: Xử lý "tan"
sortedStr = sortString("tan") = "ant"
anagramMap["ant"] = ["tan"]
anagramMap = {"aet": ["eat", "tea"], "ant": ["tan"]}

Bước 4: Xử lý "ate"
sortedStr = sortString("ate") = "aet"
anagramMap["aet"] = ["eat", "tea", "ate"]
anagramMap = {"aet": ["eat", "tea", "ate"], "ant": ["tan"]}

Bước 5: Xử lý "nat"
sortedStr = sortString("nat") = "ant"
anagramMap["ant"] = ["tan", "nat"]
anagramMap = {"aet": ["eat", "tea", "ate"], "ant": ["tan", "nat"]}

Bước 6: Xử lý "bat"
sortedStr = sortString("bat") = "abt"
anagramMap["abt"] = ["bat"]
anagramMap = {"aet": ["eat", "tea", "ate"], "ant": ["tan", "nat"], "abt": ["bat"]}

Bước 7: Chuyển đổi map thành kết quả
result = [["eat", "tea", "ate"], ["tan", "nat"], ["bat"]]

Kết quả: [["bat"],["nat","tan"],["ate","eat","tea"]]
(Lưu ý: thứ tự có thể khác nhau)
```

### Tại sao sắp xếp các ký tự tạo ra key duy nhất?

**Chứng minh:**

1. **Các anagrams có cùng key**: Nếu hai string là anagrams, chúng có cùng các ký tự (chỉ khác thứ tự). Sau khi sắp xếp, chúng sẽ có cùng kết quả.

2. **Các non-anagrams có key khác nhau**: Nếu hai string không phải anagrams, chúng có ít nhất một ký tự khác nhau về số lượng hoặc loại. Sau khi sắp xếp, chúng sẽ có kết quả khác nhau.

**Ví dụ:**
- "eat" và "tea": Cùng có 1 'a', 1 'e', 1 't' → sau sort đều thành "aet" ✓
- "eat" và "bat": "eat" có 1 'e', "bat" có 1 'b' → sau sort thành "aet" và "abt" (khác nhau) ✓

## Giải thích code

### Hàm groupAnagrams

```9:44:internal/hash_table/group_anagrams/group_anagrams.go
// groupAnagrams nhóm các anagrams lại với nhau.
//
// Sử dụng Hash Map với sorted string làm key:
// - Anagrams là các từ có cùng ký tự nhưng khác thứ tự
// - Sắp xếp các ký tự trong mỗi string để tạo key duy nhất
// - Các anagrams sẽ có cùng key sau khi sắp xếp
// - Sử dụng map để nhóm các string có cùng key
//
// Ví dụ:
// - "eat", "tea", "ate" → sau khi sort đều thành "aet"
// - "bat" → sau khi sort thành "abt" (khác với "aet")
//
// Độ phức tạp: O(n * k * log(k)) thời gian, O(n * k) không gian
// với n là số lượng strings và k là độ dài trung bình của mỗi string
func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return [][]string{}
	}

	// Map để lưu trữ: key (sorted string) -> danh sách các anagrams
	anagramMap := make(map[string][]string)

	// Duyệt qua từng string trong mảng
	for _, str := range strs {
		// Sắp xếp các ký tự trong string để tạo key
		// Ví dụ: "eat" → "aet", "tea" → "aet"
		sortedStr := sortString(str)

		// Thêm string vào nhóm tương ứng
		anagramMap[sortedStr] = append(anagramMap[sortedStr], str)
	}

	// Chuyển đổi map thành slice các nhóm
	result := make([][]string, 0, len(anagramMap))
	for _, group := range anagramMap {
		result = append(result, group)
	}

	return result
}
```

**Chi tiết từng phần:**

1. **Xử lý mảng rỗng (dòng 18-20):**
   ```go
   if len(strs) == 0 {
       return [][]string{}
   }
   ```
   - Trả về mảng rỗng nếu input rỗng

2. **Khởi tạo map (dòng 22-23):**
   ```go
   anagramMap := make(map[string][]string)
   ```
   - Map lưu trữ: key là string đã sắp xếp, value là danh sách các anagrams

3. **Duyệt và nhóm (dòng 25-32):**
   ```go
   for _, str := range strs {
       sortedStr := sortString(str)
       anagramMap[sortedStr] = append(anagramMap[sortedStr], str)
   }
   ```
   - Với mỗi string, sắp xếp các ký tự để tạo key
   - Thêm string vào nhóm tương ứng trong map

4. **Chuyển đổi thành kết quả (dòng 34-38):**
   ```go
   result := make([][]string, 0, len(anagramMap))
   for _, group := range anagramMap {
       result = append(result, group)
   }
   ```
   - Chuyển đổi map thành slice các nhóm

### Hàm sortString

```46:54:internal/hash_table/group_anagrams/group_anagrams.go
// sortString sắp xếp các ký tự trong string theo thứ tự alphabet
func sortString(s string) string {
	// Chuyển string thành slice các rune để sắp xếp
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}
```

**Chi tiết:**
- Chuyển string thành slice các rune (để xử lý Unicode đúng cách)
- Sắp xếp các rune theo thứ tự alphabet
- Chuyển đổi lại thành string

## Độ phức tạp

### Thời gian: O(n * k * log(k))

- **Duyệt qua n strings**: O(n)
- **Sắp xếp mỗi string (độ dài k)**: O(k * log(k))
- **Tổng**: O(n * k * log(k))

**Phân tích chi tiết:**
- Vòng lặp chính: O(n)
- Trong mỗi lần lặp:
  - `sortString`: O(k * log(k)) với k là độ dài string
  - Thêm vào map: O(1) (average case)
- Chuyển đổi map thành slice: O(n) (số lượng nhóm ≤ n)

### Không gian: O(n * k)

- **Map `anagramMap`**: O(n * k)
  - Lưu trữ tất cả n strings, mỗi string có độ dài trung bình k
- **Slice `result`**: O(n * k)
  - Chứa tất cả các nhóm (không tạo thêm string mới, chỉ tham chiếu)
- **Tổng**: O(n * k)

**Lưu ý:** Trong thực tế, không gian có thể nhỏ hơn một chút vì các string được chia sẻ tham chiếu, nhưng worst case vẫn là O(n * k).

## Test Cases

### Test Case 1: Example 1
```go
Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
```
Có 3 nhóm anagrams khác nhau.

### Test Case 2: Example 2
```go
Input: strs = [""]
Output: [[""]]
```
Chỉ có một chuỗi rỗng.

### Test Case 3: Example 3
```go
Input: strs = ["a"]
Output: [["a"]]
```
Chỉ có một chuỗi.

### Test Case 4: Tất cả là anagrams
```go
Input: strs = ["eat", "tea", "ate"]
Output: [["ate","eat","tea"]]
```
Tất cả các string đều là anagrams của nhau.

### Test Case 5: Không có anagrams
```go
Input: strs = ["abc", "def", "ghi"]
Output: [["abc"],["def"],["ghi"]]
```
Mỗi string là một nhóm riêng.

### Test Case 6: Nhiều nhóm
```go
Input: strs = ["abc", "bca", "cab", "xyz", "yzx", "a"]
Output: [["a"],["abc","bca","cab"],["xyz","yzx"]]
```
Có 3 nhóm với số lượng phần tử khác nhau.

### Test Case 7: Chuỗi một ký tự
```go
Input: strs = ["a", "b", "a", "b", "c"]
Output: [["a","a"],["b","b"],["c"]]
```
Các chuỗi một ký tự được nhóm theo giá trị.

### Test Case 8: Chuỗi trùng lặp
```go
Input: strs = ["eat", "eat", "tea"]
Output: [["eat","eat","tea"]]
```
Các chuỗi trùng lặp được nhóm cùng với anagrams.

### Test Case 9: Chuỗi dài
```go
Input: strs = ["listen", "silent", "enlist"]
Output: [["enlist","listen","silent"]]
```
Anagrams với chuỗi dài hơn.

### Test Case 10: Mảng rỗng
```go
Input: strs = []
Output: []
```
Trả về mảng rỗng.

## Chạy test

Để chạy các test case:

```bash
go test ./internal/hash_table/group_anagrams/
```

Hoặc chạy với verbose mode để xem chi tiết:

```bash
go test -v ./internal/hash_table/group_anagrams/
```

## Mở rộng

### Biến thể: Group Anagrams với Case-Insensitive

Nếu cần xử lý cả chữ hoa và chữ thường:

```go
func sortStringCaseInsensitive(s string) string {
    runes := []rune(strings.ToLower(s))
    sort.Slice(runes, func(i, j int) bool {
        return runes[i] < runes[j]
    })
    return string(runes)
}
```

### Biến thể: Group Anagrams với Unicode

Xử lý các ký tự Unicode phức tạp:

```go
func sortStringUnicode(s string) string {
    runes := []rune(s)
    sort.Slice(runes, func(i, j int) bool {
        return runes[i] < runes[j]
    })
    return string(runes)
}
```

Code hiện tại đã xử lý Unicode đúng cách bằng cách sử dụng `[]rune`.

### Biến thể: Tìm số lượng nhóm anagrams lớn nhất

```go
func maxAnagramGroupSize(strs []string) int {
    groups := groupAnagrams(strs)
    maxSize := 0
    for _, group := range groups {
        if len(group) > maxSize {
            maxSize = len(group)
        }
    }
    return maxSize
}
```

### Ứng dụng thực tế

Thuật toán này được sử dụng trong:

- **Từ điển**: Nhóm các từ có cùng ký tự
- **Trò chơi chữ**: Tìm các từ có thể được tạo từ cùng một bộ ký tự
- **Phân tích văn bản**: Phát hiện các biến thể của cùng một từ
- **Database indexing**: Tạo index cho các từ tương tự
- **Spell checker**: Tìm các từ có thể là lỗi chính tả

### Tips và Tricks

1. **Sử dụng sorted string làm key**: Đơn giản và hiệu quả
2. **Xử lý Unicode đúng cách**: Sử dụng `[]rune` thay vì `[]byte`
3. **Tối ưu không gian**: Có thể sử dụng character count nếu cần tốc độ cao hơn
4. **Edge cases**: Mảng rỗng, chuỗi rỗng, một phần tử
5. **So sánh kết quả**: Cần sắp xếp để so sánh vì thứ tự không quan trọng

### Tối ưu hóa

1. **Character Count thay vì Sort**: 
   - Nếu k nhỏ, sort vẫn nhanh
   - Nếu k lớn, character count có thể nhanh hơn
   - Trade-off: Code phức tạp hơn

2. **Pre-allocate slice**: 
   ```go
   result := make([][]string, 0, len(anagramMap))
   ```
   - Giảm số lần reallocate

3. **String Builder cho key**: 
   - Nếu dùng character count, có thể dùng string builder để tạo key hiệu quả hơn

### So sánh với các cách tiếp cận khác

| Đặc điểm | Sorted String | Character Count |
|----------|---------------|-----------------|
| Độ phức tạp thời gian | O(n*k*log(k)) | O(n*k) |
| Độ phức tạp không gian | O(n*k) | O(n*k) |
| Độ phức tạp code | Đơn giản | Phức tạp hơn |
| Hiệu suất thực tế | Tốt với k nhỏ | Tốt với k lớn |
| Dễ hiểu | Rất dễ | Trung bình |

### Lưu ý về Unicode

Go xử lý Unicode tốt với `[]rune`:
- `string` trong Go là UTF-8 encoded
- `[]rune` là slice của Unicode code points
- Sắp xếp `[]rune` sẽ sắp xếp theo code point, không phải theo byte

Ví dụ:
- "café" có 4 runes: 'c', 'a', 'f', 'é'
- Sau sort: "acéf" (đúng)
- Nếu dùng `[]byte`: có thể sai với các ký tự multi-byte
