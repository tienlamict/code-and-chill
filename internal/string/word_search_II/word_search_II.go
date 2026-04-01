package word_search_II

// TrieNode đại diện cho một nút trong cây Trie
type TrieNode struct {
	children map[byte]*TrieNode
	word     string // Lưu trữ từ tại nút kết thúc để dễ dàng lấy kết quả
}

// buildTrie xây dựng cây Trie từ danh sách các từ
func buildTrie(words []string) *TrieNode {
	root := &TrieNode{children: make(map[byte]*TrieNode)}
	for _, w := range words {
		node := root
		for i := 0; i < len(w); i++ {
			char := w[i]
			if _, ok := node.children[char]; !ok {
				node.children[char] = &TrieNode{children: make(map[byte]*TrieNode)}
			}
			node = node.children[char]
		}
		node.word = w
	}
	return root
}

func findWords(board [][]byte, words []string) []string {
	root := buildTrie(words)
	result := []string{}
	rows := len(board)
	cols := len(board[0])

	// Duyệt qua từng ô trên bảng để bắt đầu tìm kiếm
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if _, ok := root.children[board[r][c]]; ok {
				backtrack(board, r, c, root, &result)
			}
		}
	}

	return result
}

func backtrack(board [][]byte, r, c int, parent *TrieNode, result *[]string) {
	char := board[r][c]
	currNode := parent.children[char]

	// Nếu tìm thấy một từ, thêm vào kết quả và xóa word để tránh trùng lặp
	if currNode.word != "" {
		*result = append(*result, currNode.word)
		currNode.word = ""
	}

	// Đánh dấu ô hiện tại là đã dùng
	board[r][c] = '#'

	// Các hướng di chuyển: lên, xuống, trái, phải
	dr := []int{-1, 1, 0, 0}
	dc := []int{0, 0, -1, 1}

	for i := 0; i < 4; i++ {
		nr, nc := r+dr[i], c+dc[i]
		if nr >= 0 && nr < len(board) && nc >= 0 && nc < len(board[0]) {
			nextChar := board[nr][nc]
			if _, ok := currNode.children[nextChar]; ok {
				backtrack(board, nr, nc, currNode, result)
			}
		}
	}

	// Khôi phục lại ô hiện tại (Backtracking)
	board[r][c] = char

	// Tối ưu hóa: Cắt tỉa cây Trie (Pruning)
	// Nếu một nút lá không còn dẫn đến từ nào khác, có thể xóa nó
	if len(currNode.children) == 0 {
		delete(parent.children, char)
	}
}
