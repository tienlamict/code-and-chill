# Construct Binary Tree from Preorder and Inorder Traversal

## Mô tả bài toán

Cho hai mảng số nguyên `preorder` và `inorder`, trong đó:
- `preorder` là kết quả duyệt **preorder** (gốc → trái → phải) của một cây nhị phân.
- `inorder` là kết quả duyệt **inorder** (trái → gốc → phải) của cùng cây đó.

Hãy xây dựng và trả về cây nhị phân tương ứng.

## Ví dụ minh họa

**Ví dụ 1:**
```
Input:  preorder = [3,9,20,15,7]
        inorder  = [9,3,15,20,7]
Output: [3,9,20,null,null,15,7]

Cây:
        3
       / \
      9  20
        /  \
       15   7
```

**Ví dụ 2:**
```
Input:  preorder = [-1]
        inorder  = [-1]
Output: [-1]
```

## Ràng buộc

- `1 <= preorder.length <= 3000`
- `inorder.length == preorder.length`
- `-3000 <= preorder[i], inorder[i] <= 3000`
- Các giá trị trong `preorder` và `inorder` đều **duy nhất**.
- Mỗi giá trị trong `inorder` cũng xuất hiện trong `preorder`.
- `preorder` đảm bảo là preorder traversal hợp lệ của cây.
- `inorder` đảm bảo là inorder traversal hợp lệ của cây.

## Phân tích các cách tiếp cận

### Cách 1: Brute Force (Tìm kiếm tuyến tính)

- Phần tử đầu tiên của `preorder` là root.
- Duyệt tuyến tính `inorder` để tìm vị trí root.
- Đệ quy xây dựng cây con trái và phải.

**Độ phức tạp:**
- Thời gian: O(n²) — mỗi lần đệ quy duyệt tuyến tính mảng inorder.
- Không gian: O(n) — đệ quy stack.

### Cách 2: Divide & Conquer + HashMap (Tối ưu) ✅

- Dùng HashMap lưu trước `value → index` trong `inorder` để tra cứu O(1).
- Dùng chỉ số (index) thay vì tạo mảng con mới, tránh tốn thêm bộ nhớ.
- Đệ quy phân chia: mỗi lần xác định root, tính kích thước cây con trái để xác định vùng `preorder` và `inorder` tương ứng.

**Độ phức tạp:**
- Thời gian: O(n) — mỗi nút được xử lý đúng một lần.
- Không gian: O(n) — HashMap + đệ quy stack (O(h) với h là chiều cao cây, tệ nhất O(n)).

## Giải thích thuật toán được chọn

### Nguyên lý cốt lõi

1. **Preorder**: phần tử **đầu tiên** luôn là **root** của cây (hoặc cây con hiện tại).
2. **Inorder**: root chia mảng thành hai phần — bên **trái** là cây con trái, bên **phải** là cây con phải.

### Các bước thực hiện (Ví dụ 1)

```
preorder = [3, 9, 20, 15, 7]
inorder  = [9, 3, 15, 20, 7]
```

**Bước 1:** `preorder[0] = 3` → root = 3.
Tìm 3 trong inorder tại index 1 (`mid = 1`).
- Inorder trái: `[9]` (index 0..0) → cây con trái có **1 nút**.
- Inorder phải: `[15, 20, 7]` (index 2..4) → cây con phải có **3 nút**.

```
preorder trái : [9]        (preStart+1 = 1, size=1)
preorder phải : [20, 15, 7](preStart+1+1 = 2)
```

**Bước 2:** Đệ quy cây con trái: `preorder=[9]`, `inorder=[9]`
- Root = 9, không có con → node lá.

**Bước 3:** Đệ quy cây con phải: `preorder=[20, 15, 7]`, `inorder=[15, 20, 7]`
- Root = 20, tìm trong inorder tại index 3 (`mid = 3`).
- Inorder trái: `[15]`, phải: `[7]`.
- Đệ quy → node 15 (lá), node 7 (lá).

**Kết quả:**
```
        3
       / \
      9  20
        /  \
       15   7
```

## Giải thích code

```go
// Tạo map tra cứu O(1): value → index trong inorder
inorderIdx := make(map[int]int, len(inorder))
for i, v := range inorder {
    inorderIdx[v] = i
}
```

Thay vì duyệt tuyến tính O(n) mỗi lần, map cho phép tìm vị trí root trong O(1).

```go
var build func(preStart, inStart, inEnd int) *TreeNode
build = func(preStart, inStart, inEnd int) *TreeNode {
    if inStart > inEnd {
        return nil
    }

    rootVal := preorder[preStart]   // phần tử đầu preorder = root
    root := &TreeNode{Val: rootVal}

    mid := inorderIdx[rootVal]      // vị trí root trong inorder
    leftSize := mid - inStart       // số nút cây con trái

    root.Left = build(preStart+1, inStart, mid-1)
    root.Right = build(preStart+1+leftSize, mid+1, inEnd)

    return root
}
```

- `preStart+1`: bỏ qua root hiện tại, phần còn lại bắt đầu bằng root của cây con trái.
- `preStart+1+leftSize`: sau `leftSize` nút của cây con trái là root cây con phải.
- `inStart > inEnd`: điều kiện dừng khi không còn nút nào trong vùng inorder.

## Phân tích độ phức tạp

| | Thời gian | Không gian |
|---|---|---|
| Brute Force | O(n²) | O(n) |
| HashMap + Divide & Conquer | **O(n)** | **O(n)** |

- **Thời gian O(n):** Mỗi trong n nút được tạo đúng một lần, tra cứu inorder O(1).
- **Không gian O(n):** HashMap chiếm O(n); call stack đệ quy O(h) — O(log n) cây cân bằng, O(n) cây skewed.
