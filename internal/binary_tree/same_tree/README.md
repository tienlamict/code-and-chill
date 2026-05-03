# Same Tree - Kiểm tra hai cây nhị phân giống nhau

## Mô tả bài toán

Cho gốc của hai cây nhị phân `p` và `q`, hãy viết một hàm để kiểm tra xem chúng có giống nhau hay không.

Hai cây nhị phân được coi là giống nhau nếu chúng có cấu trúc giống hệt nhau và các nút tương ứng có cùng giá trị.

### Ví dụ 1:
**Input:** `p = [1,2,3], q = [1,2,3]`  
**Output:** `true`

### Ví dụ 2:
**Input:** `p = [1,2], q = [1,null,2]`  
**Output:** `false`

### Ví dụ 3:
**Input:** `p = [1,2,1], q = [1,1,2]`  
**Output:** `false`

## Ràng buộc

- Số lượng nút trong cả hai cây nằm trong khoảng `[0, 100]`.
- `-10^4 <= Node.val <= 10^4`

---

## Phân tích cách tiếp cận

### 1. Đệ quy (DFS) - Tiếp cận tối ưu

Ý tưởng là so sánh từng cặp nút tương ứng của hai cây cùng một lúc.

- **Cơ sở đệ quy:**
    - Nếu cả hai nút đang xét đều là `nil`, nghĩa là chúng ta đã đi hết nhánh và cả hai nhánh đều giống nhau cho đến điểm này -> Trả về `true`.
    - Nếu một trong hai nút là `nil` mà nút kia thì không, hoặc giá trị của hai nút khác nhau -> Hai cây không giống nhau -> Trả về `false`.
- **Bước đệ quy:**
    - Nếu hai nút hiện tại có cùng giá trị, ta tiếp tục kiểm tra các cây con bên trái và bên phải của chúng.
    - Hai cây chỉ giống nhau nếu cả cây con bên trái và cây con bên phải đều giống nhau.

#### Độ phức tạp:
- **Thời gian:** $O(N)$, trong đó $N$ là tổng số nút trong cây nhỏ hơn. Chúng ta phải duyệt qua mỗi nút ít nhất một lần.
- **Không gian:** $O(H)$, trong đó $H$ là chiều cao của cây. Đây là không gian sử dụng bởi stack đệ quy. Trong trường hợp xấu nhất (cây lệch), $H = N$.

---

## Giải thích thuật toán từng bước với ví dụ

Giả sử `p = [1, 2]`, `q = [1, nil, 2]`:

1. Gọi `isSameTree(p, q)`:
    - Cả hai không nil.
    - `p.Val (1) == q.Val (1)`. Tiếp tục đệ quy.
2. Kiểm tra nhánh trái: `isSameTree(p.Left, q.Left)`:
    - `p.Left` là nút có giá trị `2`.
    - `q.Left` là `nil`.
    - Vì một nút nil và một nút không nil, trả về `false`.
3. Kết quả cuối cùng là `true && false` -> `false`.

---

## Giải thích code

```go
func isSameTree(p *TreeNode, q *TreeNode) bool {
    // 1. Nếu cả hai nút đều null, chúng giống nhau
    if p == nil && q == nil {
        return true
    }

    // 2. Nếu một trong hai nút null hoặc giá trị của chúng khác nhau, chúng không giống nhau
    if p == nil || q == nil || p.Val != q.Val {
        return false
    }

    // 3. Đệ quy kiểm tra cây con bên trái và bên phải
    // Hai cây chỉ giống nhau nếu cả hai phía đều khớp
    return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
```

1. **Trường hợp cơ sở 1:** Khi duyệt đến lá của cả hai cây cùng lúc, điều đó có nghĩa là cấu trúc tại nhánh đó khớp nhau.
2. **Trường hợp sai lệch:** Nếu cấu trúc không khớp (một bên nil, một bên không) hoặc dữ liệu không khớp (giá trị khác nhau), ta dừng và trả về `false` ngay lập tức.
3. **Lan truyền kết quả:** Sử dụng toán tử `&&` để đảm bảo mọi phần của cây đều phải khớp nhau.
