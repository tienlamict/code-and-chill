package hash_table

// minWindow trả về chuỗi con ngắn nhất của s sao cho chứa đủ mọi ký tự của t (tính cả số lần xuất hiện).
// Nếu không tồn tại thì trả về "".
//
// Sliding Window (O(m+n)):
// - need[c]: số lần ký tự c cần có (t).
// - window[c]: số lần ký tự c đang có trong cửa sổ s[left:right].
// - required: số lượng ký tự khác nhau trong t cần thỏa.
// - formed: số lượng ký tự khác nhau hiện đang thỏa (window[c] == need[c]).
//
// Mở rộng right để "đủ điều kiện" (formed == required), sau đó co left để tối thiểu hóa.
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

