# Course Schedule (LeetCode 207)

## Mô tả bài toán

Có `numCourses` khóa học đánh số `0 .. numCourses-1`. Mảng `prerequisites[i] = [ai, bi]` nghĩa là **phải học `bi` trước** mới được học `ai`.

Trả về `true` nếu có thể hoàn thành **tất cả** khóa học, ngược lại `false`.

### Ví dụ

- `numCourses = 2`, `prerequisites = [[1,0]]` → `true` (học 0 rồi 1).
- `numCourses = 2`, `prerequisites = [[1,0],[0,1]]` → `false` (chu trình phụ thuộc).

### Ràng buộc

- `1 <= numCourses <= 2000`
- `0 <= prerequisites.length <= 5000`
- Mỗi phần tử độ dài 2, `0 <= ai, bi < numCourses`, các cặp **duy nhất**.

---

## Mô hình đồ thị

- Mỗi khóa học là một **đỉnh**.
- `[ai, bi]` → cạnh có hướng **`bi -> ai`** (phải qua `bi` trước `ai`).

Có thể xếp lịch học hết **khi và chỉ khi** đồ thị **không có chu trình** (DAG) ⇔ tồn tại **thứ tự topo** (topological order).

---

## Cách tiếp cận

### 1. Kahn — BFS + indegree (đã dùng trong code) ⭐

1. Tính **bậc vào** mỗi đỉnh; xây **danh sách kề** `adj`.
2. Cho tất cả đỉnh có `indegree == 0` vào hàng đợi.
3. Lần lượt lấy đỉnh `u` khỏi queue, tăng biến đếm; với mỗi `v` kề `u`, giảm `indegree[v]`; nếu về 0 thì đưa `v` vào queue.
4. Nếu số đỉnh đã xử lý bằng `numCourses` → không chu trình → `true`.

**Độ phức tạp:** O(V + E) thời gian, O(V + E) không gian.

### 2. DFS — tô màu / trạng thái

Duyệt DFS, đánh dấu đỉnh: chưa thăm / đang trong stack / đã xong. Nếu gặp cạnh tới đỉnh **đang trong stack** → có chu trình → `false`.

Cũng O(V + E).

---

## Code chính

Xem `course_schedule.go`: hàm `canFinish`.

---

## Chạy test

```bash
go test ./internal/graph/course_schedule/
```
