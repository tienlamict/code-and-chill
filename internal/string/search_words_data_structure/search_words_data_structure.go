package string

// TrieNode đại diện cho một nút trong cây Trie
type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

// WordDictionary đại diện cho cấu trúc dữ liệu lưu trữ từ điển
type WordDictionary struct {
	root *TrieNode
}

// Constructor khởi tạo đối tượng WordDictionary
func Constructor() WordDictionary {
	return WordDictionary{
		root: &TrieNode{
			children: make(map[rune]*TrieNode),
		},
	}
}

// AddWord thêm một từ vào cấu trúc dữ liệu
func (this *WordDictionary) AddWord(word string) {
	curr := this.root
	for _, char := range word {
		if _, ok := curr.children[char]; !ok {
			curr.children[char] = &TrieNode{
				children: make(map[rune]*TrieNode),
			}
		}
		curr = curr.children[char]
	}
	curr.isEnd = true
}

// Search tìm kiếm xem từ có tồn tại trong từ điển hay không
// Hỗ trợ ký tự '.' có thể khớp với bất kỳ chữ cái nào
func (this *WordDictionary) Search(word string) bool {
	return this.searchRecursive(this.root, word)
}

// searchRecursive là hàm đệ quy để xử lý tìm kiếm với ký tự wildcard '.'
func (this *WordDictionary) searchRecursive(node *TrieNode, word string) bool {
	for i, char := range word {
		if char == '.' {
			// Nếu gặp '.', thử tất cả các nhánh con khả thi
			for _, child := range node.children {
				if this.searchRecursive(child, word[i+1:]) {
					return true
				}
			}
			return false
		} else {
			// Tìm kiếm ký tự thông thường
			if next, ok := node.children[char]; ok {
				node = next
			} else {
				return false
			}
		}
	}
	return node.isEnd
}
