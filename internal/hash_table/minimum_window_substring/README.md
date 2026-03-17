# Minimum Window Substring (LeetCode 76)

## Mô tả bài toán

Cho hai chuỗi `s` và `t`. Hãy tìm **chuỗi con liên tiếp ngắn nhất** của `s` sao cho chứa **đủ mọi ký tự** trong `t` (tính cả ký tự trùng lặp).

Nếu không tồn tại chuỗi con thỏa điều kiện, trả về chuỗi rỗng `""`.

Lưu ý: test đảm bảo **đáp án là duy nhất**.

---

## Ví dụ

### Example 1
- Input: `s = "ADOBECODEBANC"`, `t = "ABC"`
- Output: `"BANC"`

### Example 2
- Input: `s = "a"`, `t = "a"`
- Output: `"a"`

### Example 3
- Input: `s = "a"`, `t = "aa"`
- Output: `""`

---

## Ràng buộc

- `1 <= len(s), len(t) <= 10^5`
- `s` và `t` chỉ gồm chữ cái tiếng Anh hoa và thường.

---

## Phân tích & hướng giải

### Cách 1: Brute force (không khả thi)

Xét mọi substring của `s` và kiểm tra có chứa đủ `t` không.

- Thời gian: \(O(m^2 \cdot \Sigma)\) (rất chậm với `m` đến \(10^5\))
- Không dùng được.

### Cách 2: Sliding Window + Counting (O(m+n)) ⭐

**Ý tưởng**:

- Dùng hai mảng đếm:
  - `need[c]`: số lần ký tự `c` cần có (từ `t`)
  - `window[c]`: số lần ký tự `c` đang có trong cửa sổ hiện tại của `s`
- `required`: số lượng ký tự *khác nhau* trong `t` cần thỏa
- `formed`: số lượng ký tự khác nhau hiện đang thỏa điều kiện `window[c] == need[c]`

**Cách chạy**:

1. Mở rộng con trỏ `right` để cửa sổ chứa đủ ký tự cần thiết (`formed == required`).
2. Khi đã đủ, co `left` để tối thiểu hóa độ dài cửa sổ.
3. Lưu lại cửa sổ tốt nhất (ngắn nhất) trong quá trình co.

**Vì sao đúng?**

- Khi `formed == required`, cửa sổ hiện tại chắc chắn hợp lệ.
- Co `left` cho đến khi cửa sổ mất hợp lệ đảm bảo ta tìm được cửa sổ ngắn nhất ứng với mỗi `right`.
- Duyệt `right` từ trái sang phải đảm bảo không bỏ sót đáp án tối ưu.

---

## Độ phức tạp

- **Thời gian**: \(O(m+n)\)
  - Mỗi con trỏ `left` và `right` chỉ tăng tối đa `m` lần.
  - Khởi tạo `need` mất `n`.
- **Không gian**: \(O(1)\)
  - Vì bảng chữ cái giới hạn (ASCII cho chữ hoa/thường), dùng mảng cố định kích thước 128.

---

## Code chính (trích đoạn)

```12:82:internal/hash_table/minimum_window_substring/minimum_window_substring.go
func minWindow(s string, t string) string {
	if len(t) == 0 || len(s) == 0 {
		return ""
	}
	if len(t) > len(s) {
		return ""
	}

	var need [128]int
	required := 0
	for i := 0; i < len(t); i++ {
		c := t[i]
		if need[c] == 0 {
			required++
		}
		need[c]++
	}

	var window [128]int
	formed := 0

	bestLen := int(^uint(0) >> 1) // MaxInt
	bestL := 0

	left := 0
	for right := 0; right < len(s); right++ {
		cr := s[right]
		window[cr]++

		if need[cr] > 0 && window[cr] == need[cr] {
			formed++
		}

		for left <= right && formed == required {
			if currLen := right - left + 1; currLen < bestLen {
				bestLen = currLen
				bestL = left
			}

			cl := s[left]
			window[cl]--
			if need[cl] > 0 && window[cl] < need[cl] {
				formed--
			}
			left++
		}
	}

	if bestLen == int(^uint(0)>>1) {
		return ""
	}
	return s[bestL : bestL+bestLen]
}
```

---

## Test

Chạy test:

```bash
go test ./internal/hash_table/minimum_window_substring/
```

