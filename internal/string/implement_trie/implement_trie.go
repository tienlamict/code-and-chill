package string

// TrieNode đại diện cho một nút trong cây Trie.
// Mỗi nút chứa 26 con trỏ tới các nút con (tương ứng với 26 chữ cái thường),
// và một cờ đánh dấu kết thúc từ.
type TrieNode struct {
	children [26]*TrieNode
	isEnd    bool
}

// Trie là cấu trúc cây tiền tố.
type Trie struct {
	root *TrieNode
}

// Constructor khởi tạo một Trie mới với nút gốc rỗng.
func Constructor() Trie {
	return Trie{root: &TrieNode{}}
}

// Insert chèn một từ vào Trie.
// Duyệt từng ký tự, tạo nút mới nếu chưa tồn tại, đánh dấu isEnd ở nút cuối.
func (t *Trie) Insert(word string) {
	node := t.root
	for _, ch := range word {
		idx := ch - 'a'
		if node.children[idx] == nil {
			node.children[idx] = &TrieNode{}
		}
		node = node.children[idx]
	}
	node.isEnd = true
}

// Search kiểm tra xem từ word có tồn tại trong Trie hay không.
// Duyệt từng ký tự, trả về false nếu thiếu nút. Cuối cùng kiểm tra isEnd.
func (t *Trie) Search(word string) bool {
	node := t.traverse(word)
	return node != nil && node.isEnd
}

// StartsWith kiểm tra xem có từ nào trong Trie bắt đầu bằng prefix hay không.
func (t *Trie) StartsWith(prefix string) bool {
	return t.traverse(prefix) != nil
}

// traverse duyệt Trie theo chuỗi s, trả về nút cuối cùng hoặc nil nếu không tìm thấy.
func (t *Trie) traverse(s string) *TrieNode {
	node := t.root
	for _, ch := range s {
		idx := ch - 'a'
		if node.children[idx] == nil {
			return nil
		}
		node = node.children[idx]
	}
	return node
}
