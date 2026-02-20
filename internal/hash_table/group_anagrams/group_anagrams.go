package hash_table

import "sort"

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

// sortString sắp xếp các ký tự trong string theo thứ tự alphabet
func sortString(s string) string {
	// Chuyển string thành slice các rune để sắp xếp
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}
