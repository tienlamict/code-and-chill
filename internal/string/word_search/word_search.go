package string

// exist kiểm tra xem có thể ghép được chuỗi word từ các ô liền kề (ngang/dọc) trên board hay không.
// Mỗi ô chỉ được dùng tối đa một lần trong một đường đi.
//
// Thuật toán: Backtracking (DFS) từ mọi ô làm điểm bắt đầu.
// - Tại (i,j), nếu board[i][j] == word[k] thì đánh dấu ô đã thăm (tạm thời đổi byte),
//   thử 4 hướng với chỉ số k+1, sau đó khôi phục lại ký tự gốc.
//
// Pruning (Follow-up): trước khi DFS, đếm tần suất ký tự trên board và trong word.
// Nếu word cần nhiều ký tự c hơn số lần xuất hiện trên board thì trả về false ngay.
//
// Độ phức tạp: worst-case O(m * n * 3^L) với L = len(word) (mỗi bước tối đa 3 nhánh,
// vì không quay lại ô vừa đi), với ràng buộc đề bài (m,n <= 6, L <= 15) là chấp nhận được.
func exist(board [][]byte, word string) bool {
	if len(word) == 0 {
		return true
	}
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}

	if !hasEnoughLetters(board, word) {
		return false
	}

	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfsExist(board, i, j, 0, word) {
				return true
			}
		}
	}
	return false
}

// hasEnoughLetters pruning: nếu word yêu cầu nhiều ký tự hơn board có thì không thể tồn tại đường đi.
func hasEnoughLetters(board [][]byte, word string) bool {
	var cnt [256]int
	for i := range board {
		for j := range board[i] {
			cnt[board[i][j]]++
		}
	}
	var need [256]int
	for i := 0; i < len(word); i++ {
		need[word[i]]++
	}
	for c := 0; c < 256; c++ {
		if need[c] > cnt[c] {
			return false
		}
	}
	return true
}

func dfsExist(board [][]byte, i, j, k int, word string) bool {
	if k == len(word) {
		return true
	}
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
		return false
	}
	if board[i][j] != word[k] {
		return false
	}

	orig := board[i][j]
	board[i][j] = '#' // đánh dấu đã thăm (không dùng lại ô này)

	found := dfsExist(board, i+1, j, k+1, word) ||
		dfsExist(board, i-1, j, k+1, word) ||
		dfsExist(board, i, j+1, k+1, word) ||
		dfsExist(board, i, j-1, k+1, word)

	board[i][j] = orig
	return found
}
