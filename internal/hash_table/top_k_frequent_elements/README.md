# Top K Frequent Elements

## Mô tả bài toán

Cho một mảng số nguyên `nums` và một số nguyên `k`, trả về `k` phần tử xuất hiện nhiều nhất. Có thể trả về kết quả theo bất kỳ thứ tự nào.

## Ví dụ minh họa

**Ví dụ 1:**
```
Input:  nums = [1,1,1,2,2,3], k = 2
Output: [1,2]
```

**Ví dụ 2:**
```
Input:  nums = [1], k = 1
Output: [1]
```

**Ví dụ 3:**
```
Input:  nums = [1,2,1,2,1,2,3,1,3,2], k = 2
Output: [1,2]
```

## Ràng buộc

- `1 <= nums.length <= 10^5`
- `-10^4 <= nums[i] <= 10^4`
- `k` nằm trong khoảng `[1, số lượng phần tử duy nhất trong mảng]`
- Đảm bảo kết quả là duy nhất

## Phân tích các cách tiếp cận

### 1. Brute Force: Sắp xếp theo tần suất

Đếm tần suất từng phần tử, sau đó sắp xếp theo tần suất giảm dần và lấy k phần tử đầu.

```
Bước 1: Đếm freq: {1:3, 2:2, 3:1}
Bước 2: Sắp xếp: [(1,3), (2,2), (3,1)]
Bước 3: Lấy k=2 phần tử đầu → [1, 2]
```

- **Thời gian:** O(n log n) — do bước sắp xếp
- **Không gian:** O(n)
- **Nhược điểm:** Không đáp ứng yêu cầu follow-up (phải tốt hơn O(n log n))

---

### 2. Min-Heap kích thước k

Dùng min-heap để duy trì k phần tử có tần suất cao nhất.

- **Thời gian:** O(n log k)
- **Không gian:** O(n + k)
- **Ưu điểm:** Tốt khi k << n; tốt hơn O(n log n)

---

### 3. Bucket Sort — Tối ưu O(n) ✅

Quan sát then chốt: tần suất của một phần tử tối đa là `n` (toàn bộ mảng gồm phần tử đó). Vì vậy ta dùng mảng bucket độ dài `n+1`, bucket tại vị trí `i` chứa tất cả các phần tử có tần suất bằng `i`.

- **Thời gian:** O(n)
- **Không gian:** O(n)
- **Ưu điểm:** Đáp ứng yêu cầu follow-up, không cần sắp xếp

## Giải thích chi tiết thuật toán Bucket Sort

### Ý tưởng

Thay vì sắp xếp các phần tử theo tần suất (tốn O(n log n)), ta ánh xạ trực tiếp tần suất sang chỉ mục bucket — điều này có thể làm trong O(n).

### Các bước thực hiện

**Bước 1 — Đếm tần suất:**

```
nums = [1,1,1,2,2,3]

freq = {1: 3, 2: 2, 3: 1}
```

**Bước 2 — Tạo bucket:**

Tạo `buckets` có `n+1 = 7` ô (chỉ số 0..6):

```
Index:    0    1    2    3    4    5    6
Buckets: []  [3]  [2]  [1]  []  []  []
```

- Phần tử `3` có tần suất 1 → vào `buckets[1]`
- Phần tử `2` có tần suất 2 → vào `buckets[2]`
- Phần tử `1` có tần suất 3 → vào `buckets[3]`

**Bước 3 — Duyệt ngược, lấy k phần tử:**

Duyệt từ index cao nhất (6) xuống 0, thu thập phần tử cho đến khi đủ k=2:

```
i=6: [] → bỏ qua
i=5: [] → bỏ qua
i=4: [] → bỏ qua
i=3: [1] → result = [1]  (1 phần tử)
i=2: [2] → result = [1, 2]  (2 phần tử, đủ k)
```

Kết quả: `[1, 2]`

## Giải thích code

```go
// Bước 1: Đếm tần suất
freq := make(map[int]int)
for _, num := range nums {
    freq[num]++
}
```

Dùng hash map để đếm O(n).

```go
// Bước 2: Tạo bucket
buckets := make([][]int, len(nums)+1)
for num, count := range freq {
    buckets[count] = append(buckets[count], num)
}
```

Ánh xạ mỗi phần tử vào bucket tương ứng với tần suất của nó.

```go
// Bước 3: Duyệt ngược, lấy k phần tử
result := make([]int, 0, k)
for i := len(buckets) - 1; i >= 0 && len(result) < k; i-- {
    result = append(result, buckets[i]...)
}
return result[:k]
```

Duyệt từ tần suất cao nhất xuống, dừng khi đủ k phần tử. `result[:k]` đảm bảo trả đúng k phần tử ngay cả khi một bucket có nhiều hơn số phần tử còn thiếu.

## Phân tích độ phức tạp

| | Thời gian | Không gian |
|---|---|---|
| Đếm tần suất | O(n) | O(n) |
| Tạo bucket | O(n) | O(n) |
| Duyệt bucket | O(n) | O(k) |
| **Tổng** | **O(n)** | **O(n)** |

## Follow-up: Tốt hơn O(n log n)?

Thuật toán Bucket Sort đạt **O(n)** thời gian — tốt hơn đáng kể so với O(n log n). Điều này khả thi vì tần suất bị giới hạn trong [1, n], cho phép dùng mảng bucket thay vì sắp xếp tổng quát.
