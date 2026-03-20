# Word Search (LeetCode 79)

## Mô tả bài toán

Cho lưới ký tự `board` kích thước `m x n` và chuỗi `word`. Trả về `true` nếu có thể xây dựng `word` bằng cách đi qua các ô **liền kề theo chiều ngang hoặc dọc** (không đi chéo). **Mỗi ô chỉ được dùng tối đa một lần** trong một đường đi.

## Ví dụ

- `board` như đề, `word = "ABCCED"` → `true`
- `word = "SEE"` → `true`
- `word = "ABCB"` → `false`

## Ràng buộc

- `1 <= m, n <= 6`
- `1 <= len(word) <= 15`
- `board` và `word` chỉ gồm chữ cái tiếng Anh (hoa/thường)

## Follow-up: Search pruning

Có thể cắt tỉa sớm để nhanh hơn khi board lớn hơn, ví dụ:

1. **Đếm ký tự**: Nếu `word` cần nhiều ký tự `c` hơn số lần `c` xuất hiện trên `board` → `false` ngay (đã dùng trong code).
2. **Điều chỉnh thứ tự duyệt**: Bắt đầu từ ô trùng ký tự đầu `word[0]` (đã ngầm qua vòng lặp).
3. **Trie / tối ưu khác**: Với nhiều từ hoặc board rất lớn có thể kết hợp cấu trúc khác; bài này ràng buộc nhỏ nên DFS + pruning đủ.

---

## Thuật toán chính

### Backtracking (DFS)

Với mỗi ô `(i,j)` làm điểm bắt đầu, gọi `dfs(i, j, k)` nghĩa là đang cần khớp `word[k]`.

1. Nếu `k == len(word)` → đã ghép đủ → `true`.
2. Nếu ra ngoài lưới hoặc `board[i][j] != word[k]` → `false`.
3. Lưu ký tự gốc, đặt `board[i][j] = '#'` (hoặc ký tự đánh dấu) để không quay lại ô này.
4. Thử 4 hướng với `k+1`.
5. Khôi phục `board[i][j]` về ký tự gốc.

### Độ phức tạp

- **Thời gian**: worst-case khoảng `O(m * n * 3^L)` với `L = len(word)` (mỗi bước tối đa 3 nhánh vì không đi ngược lại ô trước).
- **Không gian**: `O(L)` cho stack đệ quy (độ sâu tối đa `L`).

---

## Code

Xem `word_search.go`: hàm `exist` và `dfsExist`, pruning trong `hasEnoughLetters`.

## Chạy test

```bash
go test ./internal/string/word_search/
```
