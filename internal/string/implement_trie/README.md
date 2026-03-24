# 208. Implement Trie (Prefix Tree)

## Mô tả bài toán

Trie (đọc là "try") hay còn gọi là **cây tiền tố** (prefix tree) là một cấu trúc dữ liệu dạng cây dùng để lưu trữ và tra cứu các chuỗi ký tự một cách hiệu quả. Ứng dụng phổ biến: tính năng gợi ý từ (autocomplete), kiểm tra chính tả (spellchecker).

Implement class `Trie` với các phương thức:

- `Trie()` — khởi tạo đối tượng Trie.
- `void insert(word)` — chèn chuỗi `word` vào Trie.
- `boolean search(word)` — trả về `true` nếu `word` đã được chèn vào Trie trước đó.
- `boolean startsWith(prefix)` — trả về `true` nếu có bất kỳ từ nào đã chèn có tiền tố là `prefix`.

## Ví dụ minh họa

```
Input:
["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
[[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]

Output:
[null, null, true, false, true, null, true]

Giải thích:
Trie trie = new Trie();
trie.insert("apple");
trie.search("apple");    // true  — "apple" đã được chèn
trie.search("app");      // false — "app" chưa được chèn (chỉ là tiền tố)
trie.startsWith("app");  // true  — "apple" có tiền tố "app"
trie.insert("app");
trie.search("app");      // true  — "app" đã được chèn
```

## Ràng buộc

- `1 <= word.length, prefix.length <= 2000`
- `word` và `prefix` chỉ gồm chữ cái thường.
- Tổng số lần gọi hàm không quá `3 * 10^4`.

## Phân tích các cách tiếp cận

### 1. Dùng Hash Map (đơn giản nhất)

Lưu tất cả các từ vào một `map[string]bool`. Với `startsWith`, duyệt toàn bộ map để kiểm tra.

- **Insert**: O(1)
- **Search**: O(1)
- **StartsWith**: O(n * m) — n là số từ, m là độ dài prefix
- **Nhược điểm**: không hiệu quả cho `startsWith`, không tận dụng cấu trúc chung của các tiền tố.

### 2. Cây Trie (tối ưu — được chọn)

Xây dựng cây với mỗi nút đại diện cho một ký tự. Các từ có chung tiền tố sẽ dùng chung nhánh cây.

- **Insert**: O(m) — m là độ dài từ
- **Search**: O(m)
- **StartsWith**: O(m) — m là độ dài prefix
- **Bộ nhớ**: O(n * m * 26) trong trường hợp xấu nhất, nhưng các từ cùng tiền tố chia sẻ nút nên thực tế tiết kiệm hơn nhiều.

## Giải thích thuật toán Trie

### Cấu trúc nút

```
TrieNode {
    children: [26]*TrieNode  // 26 con trỏ cho 'a'..'z'
    isEnd:    bool           // đánh dấu đây là kết thúc của một từ
}
```

Mỗi cạnh trong cây ngầm đại diện cho một ký tự (vị trí index = ký tự - 'a').

### Insert "apple"

```
root
 └─ [a] → node_a
           └─ [p] → node_ap
                    └─ [p] → node_app
                             └─ [l] → node_appl
                                      └─ [e] → node_apple (isEnd=true)
```

### Insert "app" (sau khi đã có "apple")

Các nút `a`, `p`, `p` đã tồn tại → chỉ đánh dấu `node_app.isEnd = true`:

```
root
 └─ [a] → node_a
           └─ [p] → node_ap
                    └─ [p] → node_app (isEnd=true)  ← đánh dấu thêm
                             └─ [l] → node_appl
                                      └─ [e] → node_apple (isEnd=true)
```

### Search "app"

Duyệt `a → p → p`, đến `node_app`, kiểm tra `isEnd = true` → trả về `true`.

### StartsWith "app"

Duyệt `a → p → p`, đến `node_app` (không phải nil) → trả về `true`.

## Giải thích code

### Hàm `traverse` — dùng chung cho Search và StartsWith

```go
func (t *Trie) traverse(s string) *TrieNode {
    node := t.root
    for _, ch := range s {
        idx := ch - 'a'
        if node.children[idx] == nil {
            return nil  // ký tự không tồn tại → dừng sớm
        }
        node = node.children[idx]
    }
    return node  // trả về nút cuối của chuỗi s
}
```

Tách logic duyệt cây ra hàm riêng để tránh lặp code.

### Search và StartsWith

```go
func (t *Trie) Search(word string) bool {
    node := t.traverse(word)
    return node != nil && node.isEnd  // phải tồn tại VÀ là kết thúc từ
}

func (t *Trie) StartsWith(prefix string) bool {
    return t.traverse(prefix) != nil  // chỉ cần tồn tại đường đi
}
```

Sự khác biệt duy nhất: `search` kiểm tra thêm `isEnd`, còn `startsWith` chỉ cần tìm được nút cuối.

## Phân tích độ phức tạp

| Thao tác     | Thời gian | Không gian     |
|--------------|-----------|----------------|
| `Insert`     | O(m)      | O(m) mỗi từ mới |
| `Search`     | O(m)      | O(1)           |
| `StartsWith` | O(m)      | O(1)           |

Trong đó `m` là độ dài chuỗi đầu vào. Không gian tổng: O(n * m * 26) — n từ, mỗi từ tối đa m nút, mỗi nút 26 con trỏ.
