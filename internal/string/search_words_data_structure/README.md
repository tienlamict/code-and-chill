# 211. Design Add and Search Words Data Structure

## Mô tả bài toán

Thiết kế một cấu trúc dữ liệu hỗ trợ việc thêm từ mới và tìm kiếm xem một từ có khớp với bất kỳ từ nào đã được thêm trước đó hay không.

Triển khai lớp `WordDictionary`:

- `WordDictionary()`: Khởi tạo đối tượng.
- `void addWord(word)`: Thêm một từ vào cấu trúc dữ liệu.
- `bool search(word)`: Trả về `true` nếu có bất kỳ chuỗi nào trong cấu trúc dữ liệu khớp với `word` hoặc `false` nếu không. `word` có thể chứa dấu chấm `'.'`, trong đó dấu chấm có thể khớp với bất kỳ chữ cái nào.

### Ví dụ

```text
Input
["WordDictionary","addWord","addWord","addWord","search","search","search","search"]
[[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]
Output
[null,null,null,null,false,true,true,true]

Explanation
WordDictionary wordDictionary = new WordDictionary();
wordDictionary.addWord("bad");
wordDictionary.addWord("dad");
wordDictionary.addWord("mad");
wordDictionary.search("pad"); // return False
wordDictionary.search("bad"); // return True
wordDictionary.search(".ad"); // return True
wordDictionary.search("b.."); // return True
```

### Ràng buộc

- `1 <= word.length <= 25`
- `word` trong `addWord` chỉ bao gồm các chữ cái tiếng Anh viết thường.
- `word` trong `search` bao gồm dấu `'.'` hoặc các chữ cái tiếng Anh viết thường.
- Có tối đa 2 dấu `'.'` trong một truy vấn tìm kiếm (Theo đề bài mới trên LeetCode, nhưng bài toán gốc có thể có nhiều hơn).
- Có tối đa $10^4$ lời gọi hàm `addWord` và `search`.

## Phân tích các cách tiếp cận

### 1. Brute Force (Mảng/Danh sách)
- **Lưu trữ**: Sử dụng một danh sách để lưu tất cả các từ.
- **Tìm kiếm**: Với mỗi truy vấn tìm kiếm, duyệt qua từng từ trong danh sách và so sánh xem nó có khớp với mẫu hay không (bao gồm cả việc xử lý dấu `'.'`).
- **Độ phức tạp**:
    - `addWord`: $O(1)$
    - `search`: $O(N \times L)$ với $N$ là số lượng từ và $L$ là độ dài từ.
- **Hạn chế**: Khi số lượng từ lên đến $10^4$, việc tìm kiếm sẽ rất chậm.

### 2. Trie (Cây tiền tố) - Cách tiếp cận tối ưu
- **Lưu trữ**: Sử dụng cấu trúc cây Trie. Mỗi nút của cây đại diện cho một ký tự.
- **Tìm kiếm**:
    - Với ký tự thông thường: Đi xuống nhánh tương ứng trong cây.
    - Với dấu `'.'`: Đây là bước quan trọng, chúng ta cần thử tất cả các nhánh con khả thi tại vị trí đó bằng cách sử dụng DFS (đệ quy).
- **Độ phức tạp**:
    - `addWord`: $O(L)$ với $L$ là độ dài từ.
    - `search`: Trong trường hợp xấu nhất không có dấu `'.'`, độ phức tạp là $O(L)$. Nếu có dấu `'.'`, độ phức tạp có thể tăng lên $O(26^L)$ trong trường hợp xấu nhất, nhưng thực tế sẽ nhanh hơn nhiều vì số lượng từ hữu hạn.

## Giải thích thuật toán Trie được chọn

Thuật toán Trie là lựa chọn phù hợp nhất cho bài toán này vì nó cho phép chúng ta chia sẻ các tiền tố chung và xử lý dấu wildcard `'.'` một cách có hệ thống thông qua việc duyệt cây.

### Các bước thực hiện:

1. **Khởi tạo**: Tạo một nút gốc (root) không chứa ký tự nào.
2. **addWord**:
    - Bắt đầu từ root.
    - Với mỗi ký tự trong từ, nếu chưa có nút con tương ứng, tạo mới.
    - Di chuyển xuống nút con.
    - Sau khi kết thúc từ, đánh dấu nút cuối cùng là `isEnd = true`.
3. **search (Đệ quy)**:
    - Nếu ký tự hiện tại là một chữ cái: Kiểm tra xem nút con tương ứng có tồn tại không. Nếu có, tiếp tục đệ quy xuống.
    - Nếu ký tự hiện tại là `'.'`: Duyệt qua tất cả các nút con hiện có của nút hiện tại. Với mỗi nút con, thực hiện đệ quy để tìm phần còn lại của từ. Nếu bất kỳ nhánh nào trả về `true`, kết quả cuối cùng là `true`.
    - Điều kiện dừng: Nếu duyệt hết từ, kiểm tra xem nút hiện tại có phải là điểm kết thúc của một từ không (`isEnd`).

## Giải thích Code

### Cấu trúc dữ liệu
Sử dụng một `map[rune]*TrieNode` để lưu các nút con, giúp linh hoạt hơn so với mảng cố định 26 phần tử (tiết kiệm bộ nhớ khi từ điển không chứa đủ 26 chữ cái tại mỗi cấp).

```go
type TrieNode struct {
    children map[rune]*TrieNode
    isEnd    bool
}
```

### Hàm Search đệ quy
Đây là phần quan trọng nhất để xử lý dấu `'.'`.

```go
func (this *WordDictionary) searchRecursive(node *TrieNode, word string) bool {
    for i, char := range word {
        if char == '.' {
            // Thử tất cả các nhánh con
            for _, child := range node.children {
                if this.searchRecursive(child, word[i+1:]) {
                    return true
                }
            }
            return false
        } else {
            // Tìm kiếm bình thường
            if next, ok := node.children[char]; ok {
                node = next
            } else {
                return false
            }
        }
    }
    return node.isEnd
}
```

## Phân tích độ phức tạp

- **Độ phức tạp thời gian**:
    - `addWord`: $O(L)$, trong đó $L$ là độ dài của từ được thêm. Chúng ta duyệt qua từng ký tự một lần.
    - `search`:
        - Trường hợp không có dấu `'.'`: $O(L)$.
        - Trường hợp có dấu `'.'`: Trong trường hợp xấu nhất (ví dụ: tìm kiếm "...."), chúng ta có thể phải duyệt qua nhiều nhánh. Tuy nhiên, với giới hạn 2 dấu `'.'` như đề bài gợi ý hoặc độ dài từ ngắn, thời gian phản hồi vẫn rất nhanh.
- **Độ phức tạp không gian**: $O(T \times L)$, trong đó $T$ là tổng số từ được thêm và $L$ là độ dài trung bình. Trong trường hợp tốt nhất, các từ chia sẻ nhiều tiền tố chung, giúp tiết kiệm không gian.
