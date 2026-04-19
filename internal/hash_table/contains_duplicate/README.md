# 217. Contains Duplicate

## Mô tả bài toán

Cho một mảng số nguyên `nums`, trả về `true` nếu bất kỳ giá trị nào xuất hiện **ít nhất hai lần**, và trả về `false` nếu mọi phần tử đều phân biệt.

## Ví dụ minh họa

**Ví dụ 1:**
```
Input:  nums = [1, 2, 3, 1]
Output: true
Giải thích: Phần tử 1 xuất hiện ở chỉ số 0 và 3.
```

**Ví dụ 2:**
```
Input:  nums = [1, 2, 3, 4]
Output: false
Giải thích: Tất cả phần tử đều phân biệt.
```

**Ví dụ 3:**
```
Input:  nums = [1, 1, 1, 3, 3, 4, 3, 2, 4, 2]
Output: true
```

## Ràng buộc

- `1 <= nums.length <= 10^5`
- `-10^9 <= nums[i] <= 10^9`

---

## Phân tích các cách tiếp cận

### 1. Brute Force — O(n²) thời gian, O(1) không gian

So sánh từng cặp phần tử với nhau.

```go
for i := 0; i < len(nums); i++ {
    for j := i + 1; j < len(nums); j++ {
        if nums[i] == nums[j] { return true }
    }
}
```

- **Ưu điểm:** Không dùng bộ nhớ thêm.
- **Nhược điểm:** Quá chậm với mảng lớn (10^5 phần tử → 10^10 phép so sánh).

### 2. Sắp xếp — O(n log n) thời gian, O(1) không gian

Sắp xếp mảng, rồi kiểm tra các phần tử liền kề.

```go
sort.Ints(nums)
for i := 1; i < len(nums); i++ {
    if nums[i] == nums[i-1] { return true }
}
```

- **Ưu điểm:** Tiết kiệm bộ nhớ.
- **Nhược điểm:** Làm thay đổi mảng gốc; chậm hơn hash set.

### 3. Hash Set — O(n) thời gian, O(n) không gian ✅

Dùng hash map/set để tra cứu O(1), duyệt mảng một lần duy nhất.

- **Ưu điểm:** Nhanh nhất, dừng sớm ngay khi tìm thấy trùng.
- **Nhược điểm:** Tốn thêm O(n) bộ nhớ.

---

## Giải thích thuật toán được chọn (Hash Set)

### Ý tưởng

Duyệt qua mảng, với mỗi phần tử:
1. Kiểm tra xem nó có trong `seen` chưa.
2. Nếu có → trả về `true` ngay lập tức.
3. Nếu chưa → thêm vào `seen` và tiếp tục.

### Minh họa từng bước với `[1, 2, 3, 1]`

| Bước | Phần tử | seen trước | Có trùng? | seen sau       |
|------|---------|------------|-----------|----------------|
| 1    | 1       | {}         | Không     | {1}            |
| 2    | 2       | {1}        | Không     | {1, 2}         |
| 3    | 3       | {1, 2}     | Không     | {1, 2, 3}      |
| 4    | 1       | {1, 2, 3}  | **Có** ✓  | — (dừng lại)   |

→ Trả về `true`.

---

## Giải thích code

```go
func containsDuplicate(nums []int) bool {
    seen := make(map[int]struct{})         // dùng struct{} để tiết kiệm bộ nhớ
    for _, n := range nums {
        if _, exists := seen[n]; exists {  // tra cứu O(1)
            return true                    // dừng sớm khi tìm thấy
        }
        seen[n] = struct{}{}               // đánh dấu đã gặp
    }
    return false
}
```

- `map[int]struct{}` thay vì `map[int]bool` vì `struct{}` không chiếm bộ nhớ (zero-size).
- Trả về `true` ngay khi tìm thấy trùng, không cần duyệt hết mảng.

---

## Phân tích độ phức tạp

| | Độ phức tạp |
|---|---|
| **Thời gian** | O(n) — duyệt mảng một lần |
| **Không gian** | O(n) — hash set chứa tối đa n phần tử |
