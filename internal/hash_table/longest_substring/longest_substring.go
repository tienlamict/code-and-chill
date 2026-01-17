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
