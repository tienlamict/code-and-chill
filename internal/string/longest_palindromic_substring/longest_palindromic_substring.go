package string

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	start := 0
	maxLength := 1

	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)

		length := len1
		if len2 > len1 {
			length = len2
		}

		if length > maxLength {
			maxLength = length
			start = i - (length-1)/2
		}
	}

	return s[start : start+maxLength]
}

func expandAroundCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}
