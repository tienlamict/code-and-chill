# Valid Anagram

## Mô tả bài toán

Cho hai chuỗi `s` và `t`, trả về `true` nếu `t` là một anagram của `s`, và `false` nếu ngược lại.

Một **Anagram** là một từ hoặc cụm từ được hình thành bằng cách sắp xếp lại các chữ cái của một từ hoặc cụm từ khác, thông thường sử dụng tất cả các chữ cái ban đầu đúng một lần.

### Ví dụ 1:
- **Input:** `s = "anagram"`, `t = "nagaram"`
- **Output:** `true`

### Ví dụ 2:
- **Input:** `s = "rat"`, `t = "car"`
- **Output:** `false`

### Ràng buộc:
- `1 <= s.length, t.length <= 5 * 10^4`
- `s` và `t` chỉ bao gồm các chữ cái tiếng Anh thường.

---

## Phân tích cách tiếp cận

### 1. Sắp xếp (Sorting)
- **Ý tưởng:** Nếu hai chuỗi là anagram của nhau, sau khi sắp xếp, chúng phải hoàn toàn giống nhau.
- **Độ phức tạp thời gian:** $O(N \log N)$, với $N$ là độ dài của chuỗi.
- **Độ phức tạp không gian:** $O(1)$ hoặc $O(N)$ tùy thuộc vào thuật toán sắp xếp và ngôn ngữ lập trình.

### 2. Bảng băm hoặc Mảng đếm (Hash Table / Frequency Counter) - **Tối ưu**
- **Ý tưởng:** Đếm tần suất xuất hiện của từng ký tự trong cả hai chuỗi. Nếu tần suất của mọi ký tự đều khớp nhau, thì đó là anagram.
- **Độ phức tạp thời gian:** $O(N)$, vì chúng ta chỉ cần duyệt qua chuỗi một vài lần.
- **Độ phức tạp không gian:** $O(1)$ vì số lượng ký tự tiếng Anh thường là cố định (26).

---

## Thuật toán được chọn: Mảng đếm (Frequency Counter)

### Giải thích chi tiết:
1. **Kiểm tra độ dài:** Nếu `len(s) != len(t)`, chắc chắn không phải anagram, trả về `false`.
2. **Khởi tạo mảng đếm:** Sử dụng một mảng `count` kích thước 26 (tương ứng với 26 chữ cái 'a'-'z').
3. **Duyệt và đếm:**
   - Với mỗi ký tự trong `s`, tăng giá trị tương ứng trong mảng `count`.
   - Với mỗi ký tự trong `t`, giảm giá trị tương ứng trong mảng `count`.
   - Có thể thực hiện đồng thời trong một vòng lặp vì độ dài hai chuỗi bằng nhau.
4. **Kiểm tra kết quả:** Sau khi duyệt xong, nếu tất cả các phần tử trong mảng `count` đều bằng 0, nghĩa là số lượng từng ký tự trong `s` và `t` bằng nhau -> Trả về `true`. Ngược lại trả về `false`.

### Ví dụ từng bước (`s = "rat"`, `t = "car"`):
1. `len("rat") == len("car")` (3 == 3) -> Tiếp tục.
2. `count` khởi tạo: `[0, 0, ..., 0]`.
3. Vòng lặp:
   - `i = 0`: `s[0] = 'r'`, `t[0] = 'c'`. `count['r']++`, `count['c']--`.
   - `i = 1`: `s[1] = 'a'`, `t[1] = 'a'`. `count['a']++`, `count['a']--` (triệt tiêu).
   - `i = 2`: `s[2] = 't'`, `t[2] = 'r'`. `count['t']++`, `count['r']--`.
4. Kết quả `count`: `count['c'] = -1`, `count['t'] = 1`, các cái khác bằng 0.
5. Kiểm tra thấy có phần tử khác 0 -> Trả về `false`.

---

## Giải thích Code

```go
func isAnagram(s string, t string) bool {
    // 1. Kiểm tra độ dài
    if len(s) != len(t) {
        return false
    }

    // 2. Mảng đếm 26 ký tự (a-z)
    count := [26]int{}

    // 3. Duyệt và cập nhật tần suất
    for i := 0; i < len(s); i++ {
        count[s[i]-'a']++
        count[t[i]-'a']--
    }

    // 4. Kiểm tra xem tất cả có bằng 0 không
    for _, c := range count {
        if c != 0 {
            return false
        }
    }

    return true
}
```

---

## Phân tích độ phức tạp
- **Thời gian:** $O(N)$, trong đó $N$ là độ dài của chuỗi. Chúng ta duyệt qua chuỗi một lần và duyệt qua mảng đếm cố định (26 phần tử).
- **Không gian:** $O(1)$, vì mảng đếm luôn có kích thước 26, không phụ thuộc vào độ dài đầu vào.

---

## Follow-up: Unicode characters
**Câu hỏi:** Nếu đầu vào chứa ký tự Unicode thì sao? Bạn sẽ thích nghi giải pháp như thế nào?

**Trả lời:**
Nếu chuỗi chứa ký tự Unicode (ví dụ: tiếng Việt có dấu, emoji, v.v.), mảng cố định 26 phần tử sẽ không còn đủ.
- **Giải pháp:** Sử dụng một **Hash Map** (`map[rune]int` trong Go) thay vì mảng.
- **Cách hoạt động:** Thay vì dùng `s[i]-'a'` làm index, ta dùng chính ký tự (`rune`) đó làm key trong map.
- **Độ phức tạp:** Thời gian vẫn là $O(N)$, nhưng không gian sẽ là $O(K)$ với $K$ là số lượng ký tự duy nhất trong chuỗi.
