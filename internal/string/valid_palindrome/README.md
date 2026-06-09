# Valid Palindrome (Chuỗi Đối Xứng Hợp Lệ)

## Mô tả bài toán

Một chuỗi được gọi là chuỗi đối xứng (palindrome) nếu sau khi chuyển tất cả các chữ cái viết hoa thành chữ thường và loại bỏ tất cả các ký tự không phải chữ cái và số (non-alphanumeric), nó đọc từ trái sang phải hay từ phải sang trái đều giống nhau.

Các ký tự chữ cái và số bao gồm các chữ cái (a-z, A-Z) và các chữ số (0-9).

Cho một chuỗi `s`, trả về `true` nếu nó là chuỗi đối xứng, ngược lại trả về `false`.

### Ví dụ 1:
**Input:** `s = "A man, a plan, a canal: Panama"`
**Output:** `true`
**Giải thích:** "amanaplanacanalpanama" là một chuỗi đối xứng.

### Ví dụ 2:
**Input:** `s = "race a car"`
**Output:** `false`
**Giải thích:** "raceacar" không phải là chuỗi đối xứng.

### Ví dụ 3:
**Input:** `s = " "`
**Output:** `true`
**Giải thích:** Sau khi loại bỏ các ký tự không phải chữ cái và số, `s` trở thành một chuỗi rỗng "". Vì chuỗi rỗng đọc xuôi hay ngược đều giống nhau, nên nó là một chuỗi đối xứng.

### Ràng buộc:
* `1 <= s.length <= 2 * 10^5`
* `s` chỉ chứa các ký tự ASCII có thể in được.

---

## Phân tích các cách tiếp cận

### 1. Cách tiếp cận Brute Force (Sử dụng mảng phụ)
- **Ý tưởng:** Duyệt qua chuỗi `s`, lọc các ký tự alphanumeric, chuyển thành chữ thường và lưu vào một mảng hoặc chuỗi mới. Sau đó so sánh chuỗi mới này với chuỗi đảo ngược của nó.
- **Độ phức tạp:**
    - Thời gian: `O(n)` để lọc và `O(n)` để đảo ngược/so sánh => Tổng `O(n)`.
    - Không gian: `O(n)` để lưu trữ chuỗi đã lọc.

### 2. Cách tiếp cận Tối ưu (Two Pointers - Hai con trỏ)
- **Ý tưởng:** Sử dụng hai con trỏ `left` (bắt đầu từ đầu chuỗi) và `right` (bắt đầu từ cuối chuỗi).
    - Di chuyển `left` sang phải cho đến khi gặp ký tự alphanumeric.
    - Di chuyển `right` sang trái cho đến khi gặp ký tự alphanumeric.
    - So sánh hai ký tự tại `left` và `right` sau khi chuyển về chữ thường.
    - Nếu không giống nhau, trả về `false`.
    - Lặp lại cho đến khi `left >= right`.
- **Độ phức tạp:**
    - Thời gian: `O(n)` vì mỗi ký tự được duyệt tối đa một lần.
    - Không gian: `O(1)` vì không sử dụng thêm bộ nhớ đáng kể.

---

## Thuật toán được chọn: Two Pointers

Chúng ta chọn cách tiếp cận Two Pointers vì nó tối ưu về mặt không gian (`O(1)`).

### Ví dụ từng bước:
`s = "A man, a plan, a canal: Panama"`

1. `left = 0 ('A')`, `right = 29 ('a')`. Cả hai đều là alphanumeric. So sánh 'a' và 'a' => Khớp.
2. `left = 1 (' ')`, `right = 28 ('m')`. `left` không phải alphanumeric, tăng `left`.
3. `left = 2 ('m')`, `right = 28 ('m')`. So sánh 'm' và 'm' => Khớp.
... và tiếp tục như vậy cho đến hết.

---

## Giải thích Code

```go
func isPalindrome(s string) bool {
    left, right := 0, len(s)-1

    for left < right {
        // Bỏ qua các ký tự không phải alphanumeric từ bên trái
        for left < right && !isAlphanumeric(s[left]) {
            left++
        }
        // Bỏ qua các ký tự không phải alphanumeric từ bên phải
        for left < right && !isAlphanumeric(s[right]) {
            right--
        }

        // So sánh hai ký tự sau khi chuyển về chữ thường
        if toLower(s[left]) != toLower(s[right]) {
            return false
        }

        left++
        right--
    }

    return true
}
```

- **Hàm hỗ trợ `isAlphanumeric`**: Kiểm tra ký tự có nằm trong khoảng 'a'-'z', 'A'-'Z', hoặc '0'-'9'.
- **Hàm hỗ trợ `toLower`**: Nếu ký tự là chữ hoa ('A'-'Z'), cộng thêm 32 (khoảng cách giữa 'A' và 'a') để chuyển thành chữ thường.

---

## Phân tích độ phức tạp

- **Độ phức tạp thời gian:** `O(n)`, trong đó `n` là độ dài của chuỗi. Chúng ta chỉ duyệt qua chuỗi một lần duy nhất.
- **Độ phức tạp không gian:** `O(1)`, chúng ta chỉ sử dụng một vài biến con trỏ và không dùng thêm cấu trúc dữ liệu nào phụ thuộc vào `n`.
