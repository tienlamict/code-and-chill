# 212. Word Search II (Tìm kiếm từ II)

## Mô tả bài toán

Cho một bảng ký tự `board` kích thước `m x n` và một danh sách các chuỗi `words`. Hãy trả về tất cả các từ trong danh sách xuất hiện trên bảng.

Mỗi từ phải được xây dựng từ các ô nằm cạnh nhau theo chiều ngang hoặc chiều dọc. Một ô ký tự không được sử dụng quá một lần trong cùng một từ.

### Ví dụ 1:
![Example 1](https://assets.leetcode.com/uploads/2020/11/07/search1.jpg)
**Đầu vào:** board = [["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f","l","v"]], words = ["oath","pea","eat","rain"]
**Đầu ra:** ["eat","oath"]

### Ví dụ 2:
![Example 2](https://assets.leetcode.com/uploads/2020/11/07/search2.jpg)
**Đầu vào:** board = [["a","b"],["c","d"]], words = ["abcb"]
**Đầu ra:** []

### Ràng buộc:
- `m == board.length`
- `n == board[i].length`
- `1 <= m, n <= 12`
- `board[i][j]` là chữ cái tiếng Anh viết thường.
- `1 <= words.length <= 3 * 10^4`
- `1 <= words[i].length <= 10`
- `words[i]` gồm các chữ cái tiếng Anh viết thường.
- Tất cả các chuỗi trong `words` là duy nhất.

---

## Phân tích cách tiếp cận

### 1. Brute Force (Duyệt cạn)
- Với mỗi từ trong danh sách `words`, ta thực hiện tìm kiếm trên bảng (tương tự bài Word Search I).
- **Độ phức tạp:** $O(W \times M \times N \times 3^L)$, với $W$ là số lượng từ, $M \times N$ là kích thước bảng, và $L$ là độ dài trung bình của từ.
- **Nhược điểm:** Rất chậm khi số lượng từ lớn (lên tới 30,000 từ). Nhiều từ có tiền tố giống nhau sẽ bị duyệt lặp đi lặp lại.

### 2. Tối ưu với Trie và Backtracking (Được chọn)
- Thay vì duyệt từng từ, ta đưa tất cả các từ vào một cây **Trie**.
- Duyệt qua từng ô của bảng và bắt đầu tìm kiếm (Backtracking) đồng thời trên bảng và trên cây Trie.
- Nếu một đường đi trên bảng tương ứng với một đường đi trên Trie, ta tiếp tục. Nếu không, ta dừng lại ngay lập tức (cắt tỉa).

#### Ưu điểm:
- Xử lý hiệu quả các từ có chung tiền tố.
- Chỉ duyệt bảng một lần duy nhất cho tất cả các từ.
- **Tối ưu hóa (Pruning):** Khi một từ đã được tìm thấy, ta có thể xóa đánh dấu hoặc thậm chí xóa nút đó khỏi Trie nếu nó không còn dẫn đến từ nào khác.

---

## Chi tiết thuật toán

1. **Xây dựng Trie:** 
   - Duyệt qua danh sách `words` và thêm từng từ vào cây Trie. 
   - Tại nút kết thúc của mỗi từ, lưu chính từ đó vào trường `word` để dễ dàng lấy kết quả.

2. **Duyệt bảng:**
   - Với mỗi ô `(r, c)` trên bảng, nếu ký tự đó tồn tại ở cấp đầu tiên của Trie, bắt đầu hàm `backtrack`.

3. **Hàm Backtrack:**
   - Đánh dấu ô hiện tại là đã dùng (ví dụ: thay bằng ký tự `#`).
   - Kiểm tra xem tại nút Trie hiện tại có lưu một từ nào không. Nếu có, thêm vào kết quả và xóa trường `word` của nút đó để tránh trùng lặp kết quả.
   - Di chuyển sang 4 ô lân cận (lên, xuống, trái, phải).
   - Nếu ô lân cận hợp lệ và ký tự đó có trong các nút con của Trie hiện tại, tiếp tục đệ quy.
   - Sau khi đệ quy xong, khôi phục lại ký tự ban đầu của ô (Backtracking).

4. **Tối ưu hóa nâng cao (Pruning):**
   - Sau khi duyệt qua tất cả các con của một nút Trie, nếu nút đó không còn nút con nào khác (là nút lá), ta có thể xóa nó khỏi nút cha để giảm bớt không gian tìm kiếm cho các lượt sau.

---

## Giải thích code

### Cấu trúc Trie
```go
type TrieNode struct {
    children map[byte]*TrieNode
    word     string // Lưu từ tại đây nếu nút này là kết thúc của một từ
}
```

### Hàm tìm kiếm chính
Sử dụng backtracking để duyệt bảng:
```go
func backtrack(board [][]byte, r, c int, parent *TrieNode, result *[]string) {
    char := board[r][c]
    currNode := parent.children[char]

    if currNode.word != "" {
        *result = append(*result, currNode.word)
        currNode.word = "" // Tránh trùng lặp
    }

    board[r][c] = '#' // Đánh dấu đã dùng

    // ... duyệt 4 hướng lân cận ...

    board[r][c] = char // Backtrack

    // Pruning: Cắt tỉa cây Trie
    if len(currNode.children) == 0 {
        delete(parent.children, char)
    }
}
```

---

## Phân tích độ phức tạp

- **Thời gian:** $O(M \times N \times 3^L)$, trong đó $M \times N$ là số ô của bảng, $L$ là độ dài tối đa của từ.
  - Việc xây dựng Trie mất $O(\Sigma \text{ độ dài các từ})$.
  - Trong quá trình backtracking, tại mỗi ô ta có 3 hướng di chuyển (trừ hướng vừa tới) và đi sâu tối đa $L$ bước.
  - Nhờ việc cắt tỉa (pruning), thời gian thực tế sẽ nhanh hơn nhiều.
- **Không gian:** $O(\Sigma \text{ độ dài các từ})$ để lưu trữ cây Trie. Trong trường hợp xấu nhất, các từ không chung tiền tố.
