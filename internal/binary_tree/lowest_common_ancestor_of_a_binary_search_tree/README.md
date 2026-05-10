# Lowest Common Ancestor of a Binary Search Tree (Tổ tiên chung thấp nhất của Cây Tìm kiếm Nhị phân)

## Mô tả bài toán

Cho một cây tìm kiếm nhị phân (BST), hãy tìm node tổ tiên chung thấp nhất (LCA) của hai node cho trước `p` và `q`.

Theo định nghĩa của LCA trên Wikipedia: “Tổ tiên chung thấp nhất được định nghĩa giữa hai node `p` và `q` là node thấp nhất trong cây `T` mà có cả `p` và `q` là node con cháu (trong đó chúng ta cho phép một node là con cháu của chính nó).”

### Ví dụ 1:
![Example 1](https://assets.leetcode.com/uploads/2018/12/14/binarysearchtree_improved.png)
**Input:** root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8  
**Output:** 6  
**Giải thích:** LCA của node 2 và 8 là 6.

### Ví dụ 2:
**Input:** root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4  
**Output:** 2  
**Giải thích:** LCA của node 2 và 4 là 2, vì một node có thể là con cháu của chính nó.

### Ví dụ 3:
**Input:** root = [2,1], p = 2, q = 1  
**Output:** 2  

### Ràng buộc:
- Số lượng node trong cây nằm trong khoảng `[2, 10^5]`.
- `-10^9 <= Node.val <= 10^9`
- Tất cả `Node.val` là duy nhất.
- `p != q`
- `p` và `q` chắc chắn tồn tại trong BST.

---

## Phân tích cách tiếp cận

### 1. Cách tiếp cận Duyệt cây thông thường (Dùng cho mọi Binary Tree)
Chúng ta có thể dùng thuật toán tìm LCA cho cây nhị phân bất kỳ. Duyệt đệ quy, nếu thấy `p` hoặc `q` thì trả về node đó. Nếu một node có kết quả từ cả hai phía trái và phải, thì node đó là LCA.
- **Độ phức tạp thời gian:** O(N)
- **Độ phức tạp không gian:** O(H) (chiều cao cây)

### 2. Cách tiếp cận Tối ưu (Dựa trên tính chất BST) - **Được chọn**
Vì đây là cây BST, chúng ta có một tính chất quan trọng:
- Node bên trái luôn nhỏ hơn node gốc.
- Node bên phải luôn lớn hơn node gốc.

Giả sử chúng ta đang ở node `root`:
- Nếu cả `p.Val` và `q.Val` đều nhỏ hơn `root.Val`: LCA phải nằm ở cây con bên trái.
- Nếu cả `p.Val` và `q.Val` đều lớn hơn `root.Val`: LCA phải nằm ở cây con bên phải.
- Nếu một node nhỏ hơn hoặc bằng `root.Val` và node kia lớn hơn hoặc bằng `root.Val`: Thì `root` chính là LCA (vì `p` và `q` bắt đầu phân nhánh tại đây).

#### Ưu điểm:
- Không cần duyệt hết cây.
- Có thể triển khai bằng vòng lặp để đạt độ phức tạp không gian O(1).

---

## Giải thích thuật toán bước từng bước

Giả sử `root = 6`, `p = 2`, `q = 4`:
1. **Bước 1:** So sánh `p=2`, `q=4` với `root=6`.
   - Cả 2 và 4 đều < 6.
   - Di chuyển sang trái: `root = root.Left` (node 2).
2. **Bước 2:** So sánh `p=2`, `q=4` với `root=2`.
   - `p=2` bằng `root`.
   - `q=4` lớn hơn `root`.
   - Vì không phải cả hai đều nhỏ hơn cũng không phải cả hai đều lớn hơn, ta dừng lại.
3. **Kết luận:** Trả về `root` (node 2).

---

## Giải thích code chi tiết

```go
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	for root != nil {
		if p.Val < root.Val && q.Val < root.Val {
			// Cả p và q đều nằm bên trái, tiếp tục tìm bên trái
			root = root.Left
		} else if p.Val > root.Val && q.Val > root.Val {
			// Cả p và q đều nằm bên phải, tiếp tục tìm bên phải
			root = root.Right
		} else {
			// p và q nằm ở hai phía hoặc một trong hai là root
			// Đây chính là điểm chia nhánh (LCA)
			return root
		}
	}
	return nil
}
```

1. **Vòng lặp `for`**: Duyệt từ gốc xuống dưới. Vì p và q luôn tồn tại, vòng lặp sẽ luôn tìm thấy kết quả.
2. **Điều kiện 1**: `p.Val < root.Val && q.Val < root.Val` -> Di chuyển xuống cây con trái.
3. **Điều kiện 2**: `p.Val > root.Val && q.Val > root.Val` -> Di chuyển xuống cây con phải.
4. **Trường hợp còn lại**: Khi p và q nằm về 2 phía của root (hoặc một node trùng với root), root đó chính là tổ tiên chung thấp nhất.

---

## Độ phức tạp

- **Thời gian:** `O(H)`, trong đó `H` là chiều cao của cây.
  - Trong trường hợp cây cân bằng, `H = log(N)`.
  - Trong trường hợp cây lệch (skewed tree), `H = N`.
- **Không gian:** `O(1)` vì chúng ta sử dụng vòng lặp và không tốn thêm bộ nhớ bổ trợ (khác với đệ quy tốn O(H) cho stack).
