# Clone Graph

## Mô tả bài toán

Cho một tham chiếu đến một node trong một đồ thị vô hướng liên thông. Trả về một bản sao sâu (deep copy) của đồ thị.

Mỗi node trong đồ thị chứa một giá trị (int) và một danh sách (List[Node]) các node láng giềng.

```go
type Node struct {
    Val       int
    Neighbors []*Node
}
```

**Định dạng test case:**

Để đơn giản, giá trị của mỗi node giống với chỉ số của node đó (đánh số từ 1). Ví dụ, node đầu tiên có `val == 1`, node thứ hai có `val == 2`, và cứ thế. Đồ thị được biểu diễn trong test case bằng cách sử dụng danh sách kề (adjacency list).

Danh sách kề là một tập hợp các danh sách không có thứ tự dùng để biểu diễn một đồ thị hữu hạn. Mỗi danh sách mô tả tập hợp các node láng giềng của một node trong đồ thị.

Node đã cho luôn là node đầu tiên với `val = 1`. Bạn phải trả về bản sao của node đã cho như một tham chiếu đến đồ thị đã được clone.

**Ràng buộc:**
- Số lượng node trong đồ thị nằm trong khoảng [0, 100]
- `1 <= Node.val <= 100`
- `Node.val` là duy nhất cho mỗi node
- Không có cạnh lặp lại và không có vòng lặp tự thân trong đồ thị
- Đồ thị là liên thông và tất cả các node có thể được truy cập từ node đã cho

### Ví dụ 1:

```
Input: adjList = [[2,4],[1,3],[2,4],[1,3]]
Output: [[2,4],[1,3],[2,4],[1,3]]
```

**Giải thích:** Có 4 node trong đồ thị.
- Node 1 (val = 1) có láng giềng là node 2 (val = 2) và node 4 (val = 4)
- Node 2 (val = 2) có láng giềng là node 1 (val = 1) và node 3 (val = 3)
- Node 3 (val = 3) có láng giềng là node 2 (val = 2) và node 4 (val = 4)
- Node 4 (val = 4) có láng giềng là node 1 (val = 1) và node 3 (val = 3)

### Ví dụ 2:

```
Input: adjList = [[]]
Output: [[]]
```

**Giải thích:** Lưu ý rằng input chứa một danh sách rỗng. Đồ thị chỉ gồm một node với val = 1 và nó không có láng giềng nào.

### Ví dụ 3:

```
Input: adjList = []
Output: []
```

**Giải thích:** Đây là một đồ thị rỗng, không có node nào.

## Phân tích thuật toán

### Ý tưởng

Bài toán yêu cầu tạo một bản sao sâu (deep copy) của đồ thị, nghĩa là:
1. Mỗi node phải được tạo mới (không phải tham chiếu đến node cũ)
2. Cấu trúc của đồ thị phải được giữ nguyên
3. Các mối quan hệ láng giềng phải được sao chép chính xác

Thách thức chính là xử lý **chu trình** (cycles) trong đồ thị. Nếu không cẩn thận, thuật toán có thể bị lặp vô hạn hoặc tạo ra nhiều bản sao của cùng một node.

**Giải pháp:** Sử dụng **Depth-First Search (DFS)** kết hợp với **Hash Map** để:
- Theo dõi các node đã được clone
- Tránh clone lại các node đã xử lý
- Xử lý các chu trình một cách đúng đắn

### Các bước thực hiện

1. **Kiểm tra edge case**: Nếu node đầu vào là `nil`, trả về `nil`
2. **Tạo Hash Map**: Dùng map để lưu trữ mapping từ node gốc → node đã clone
3. **DFS Recursive**:
   - Nếu node đã được clone (có trong map) → trả về node clone đó
   - Tạo node clone mới với cùng giá trị
   - Thêm node clone vào map
   - Đệ quy clone tất cả các node láng giềng
   - Thêm các node láng giềng đã clone vào danh sách láng giềng của node clone

### Ví dụ minh họa

Với input: `[[2,4],[1,3],[2,4],[1,3]]`

```
Bước 1: Bắt đầu từ node 1 (val=1)
  - Tạo node clone mới: clone1 (val=1)
  - visited[original1] = clone1
  
Bước 2: Clone láng giềng của node 1: [2, 4]
  - Clone node 2:
    - Tạo node clone mới: clone2 (val=2)
    - visited[original2] = clone2
    - Clone láng giềng của node 2: [1, 3]
      - Clone node 1: Đã có trong visited → dùng clone1
      - Clone node 3:
        - Tạo node clone mới: clone3 (val=3)
        - visited[original3] = clone3
        - Clone láng giềng của node 3: [2, 4]
          - Clone node 2: Đã có trong visited → dùng clone2
          - Clone node 4:
            - Tạo node clone mới: clone4 (val=4)
            - visited[original4] = clone4
            - Clone láng giềng của node 4: [1, 3]
              - Clone node 1: Đã có trong visited → dùng clone1
              - Clone node 3: Đã có trong visited → dùng clone3
            - clone4.Neighbors = [clone1, clone3]
        - clone3.Neighbors = [clone2, clone4]
      - clone2.Neighbors = [clone1, clone3]
  - Clone node 4: Đã có trong visited → dùng clone4
  - clone1.Neighbors = [clone2, clone4]

Kết quả: Đồ thị clone hoàn chỉnh với cấu trúc giống hệt đồ thị gốc
```

## Giải thích code

### Cấu trúc Node

```go
type Node struct {
    Val       int
    Neighbors []*Node
}
```

- `Val`: Giá trị của node (1-indexed trong test cases)
- `Neighbors`: Danh sách các node láng giềng

### Hàm cloneGraph

```17:36:internal/hash_table/clone_graph/clone_graph.go
// cloneGraph returns a deep copy of the graph using DFS
func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	// Map to store original node -> cloned node mapping
	visited := make(map[*Node]*Node)

	var dfs func(*Node) *Node
	dfs = func(original *Node) *Node {
		// If node already cloned, return the clone
		if cloned, exists := visited[original]; exists {
			return cloned
		}

		// Create a new node with the same value
		cloned := &Node{
			Val:       original.Val,
			Neighbors: make([]*Node, 0),
		}

		// Mark this node as visited/cloned
		visited[original] = cloned

		// Recursively clone all neighbors
		for _, neighbor := range original.Neighbors {
			cloned.Neighbors = append(cloned.Neighbors, dfs(neighbor))
		}

		return cloned
	}

	return dfs(node)
}
```

### Chi tiết từng phần

#### 1. Xử lý edge case (dòng 18-21)
```go
if node == nil {
    return nil
}
```
- Xử lý trường hợp đồ thị rỗng

#### 2. Tạo Hash Map (dòng 23)
```go
visited := make(map[*Node]*Node)
```
- Map lưu trữ mapping từ node gốc → node clone
- Cho phép kiểm tra và truy xuất node đã clone trong O(1)
- **Quan trọng**: Ngăn chặn việc clone lại node đã xử lý, xử lý chu trình

#### 3. Hàm DFS đệ quy (dòng 25-35)
```go
var dfs func(*Node) *Node
dfs = func(original *Node) *Node {
    // Kiểm tra node đã được clone chưa
    if cloned, exists := visited[original]; exists {
        return cloned
    }
    
    // Tạo node clone mới
    cloned := &Node{
        Val:       original.Val,
        Neighbors: make([]*Node, 0),
    }
    
    // Đánh dấu đã clone
    visited[original] = cloned
    
    // Clone đệ quy tất cả láng giềng
    for _, neighbor := range original.Neighbors {
        cloned.Neighbors = append(cloned.Neighbors, dfs(neighbor))
    }
    
    return cloned
}
```

**Giải thích từng bước:**

1. **Kiểm tra node đã clone (dòng 26-29)**:
   - Nếu node đã có trong map → trả về node clone đã tạo trước đó
   - **Quan trọng**: Xử lý chu trình - khi gặp lại node đã xử lý, dùng bản sao đã tạo

2. **Tạo node clone mới (dòng 31-34)**:
   - Tạo node mới với cùng giá trị
   - Khởi tạo danh sách láng giềng rỗng

3. **Đánh dấu đã clone (dòng 36)**:
   - Lưu node clone vào map **trước** khi clone láng giềng
   - **Quan trọng**: Ngăn chặn infinite loop khi có chu trình

4. **Clone đệ quy láng giềng (dòng 38-40)**:
   - Duyệt qua từng node láng giềng
   - Gọi đệ quy `dfs(neighbor)` để clone node láng giềng
   - Thêm node láng giềng đã clone vào danh sách láng giềng

5. **Trả về node clone (dòng 42)**:
   - Trả về node clone đã hoàn chỉnh với tất cả láng giềng

#### 4. Gọi DFS từ node gốc (dòng 44)
```go
return dfs(node)
```
- Bắt đầu quá trình clone từ node gốc

### Tại sao phải lưu vào map TRƯỚC khi clone láng giềng?

Đây là điểm quan trọng để xử lý chu trình:

**Ví dụ:** Node A láng giềng với Node B, và Node B láng giềng với Node A (chu trình 2 node)

```
Nếu lưu vào map SAU khi clone láng giềng:
1. Clone A → chưa có trong map
2. Clone láng giềng của A (B) → chưa có trong map
3. Clone láng giềng của B (A) → A chưa có trong map → clone lại A → VÔ HẠN!

Nếu lưu vào map TRƯỚC khi clone láng giềng:
1. Clone A → lưu vào map ngay
2. Clone láng giềng của A (B) → B chưa có → clone B → lưu vào map ngay
3. Clone láng giềng của B (A) → A đã có trong map → dùng bản sao đã tạo ✓
```

### Độ phức tạp

- **Thời gian**: O(V + E)
  - V: số lượng node (vertices)
  - E: số lượng cạnh (edges)
  - Mỗi node được truy cập đúng một lần
  - Mỗi cạnh được xử lý đúng một lần
  
- **Không gian**: O(V)
  - Hash map lưu trữ tối đa V node
  - Call stack DFS có độ sâu tối đa V (trong trường hợp đồ thị đường thẳng)

## Test Cases

### Test Case 1 - Example 1
```go
Input: [[2,4],[1,3],[2,4],[1,3]]
Output: [[2,4],[1,3],[2,4],[1,3]]
```
Đồ thị hình vuông với 4 node, có chu trình

### Test Case 2 - Example 2
```go
Input: [[]]
Output: [[]]
```
Chỉ có 1 node, không có láng giềng

### Test Case 3 - Example 3
```go
Input: []
Output: []
```
Đồ thị rỗng

### Test Case 4 - Hai node kết nối
```go
Input: [[2],[1]]
Output: [[2],[1]]
```
Hai node kết nối với nhau tạo thành chu trình

## Chạy test

Để chạy các test case:

```bash
go test ./internal/hash_table/clone_graph/
```

Hoặc chạy với verbose mode để xem chi tiết:

```bash
go test -v ./internal/hash_table/clone_graph/
```

## Lưu ý

1. **Deep Copy vs Shallow Copy**:
   - Deep copy: Tạo đối tượng mới hoàn toàn, không chia sẻ tham chiếu
   - Shallow copy: Chỉ sao chép con trỏ, vẫn trỏ đến đối tượng gốc
   - Bài toán yêu cầu deep copy, nên mỗi node phải là đối tượng mới

2. **Xử lý chu trình**:
   - Hash map là giải pháp quan trọng để xử lý chu trình
   - Phải lưu node vào map trước khi clone láng giềng

3. **Cấu trúc đồ thị**:
   - Đồ thị vô hướng: nếu A láng giềng với B, thì B cũng láng giềng với A
   - Đồ thị liên thông: tất cả node có thể truy cập từ node gốc
   - Không có cạnh lặp lại và không có vòng lặp tự thân
