package string

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
