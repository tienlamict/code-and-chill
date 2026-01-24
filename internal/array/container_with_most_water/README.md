# Container With Most Water

## Đề bài
Bạn được cho mảng số nguyên `height` độ dài `n`. Có `n` đường thẳng đứng, trong đó đường thẳng thứ `i` có hai đầu mút là `(i, 0)` và `(i, height[i])`.

Hãy chọn **hai** đường thẳng sao cho cùng với trục hoành tạo thành một “container” chứa được **nhiều nước nhất**.

Trả về **diện tích lớn nhất** (lượng nước tối đa) mà container có thể chứa.

**Lưu ý:** Không được nghiêng container.

### Ví dụ

**Ví dụ 1**
- Input: `height = [1,8,6,2,5,4,8,3,7]`
- Output: `49`

**Ví dụ 2**
- Input: `height = [1,1]`
- Output: `1`

### Ràng buộc
- `2 <= n <= 10^5`
- `0 <= height[i] <= 10^4`

---

## Ý tưởng
Chọn hai vị trí `i < j`:
- **Chiều rộng**: `j - i`
- **Chiều cao hữu dụng** (bị giới hạn bởi cạnh thấp hơn): `min(height[i], height[j])`
- **Diện tích**:

\[
area(i, j) = (j - i) \times \min(height[i], height[j])
\]

Brute force \(O(n^2)\) sẽ TLE vì `n` lên tới `10^5`.

---

## Cách giải: Two Pointers (Hai con trỏ)
Dùng hai con trỏ:
- `left = 0` (đầu mảng)
- `right = n - 1` (cuối mảng)

Ở mỗi bước:
1. Tính `area = (right-left) * min(height[left], height[right])`, cập nhật đáp án lớn nhất.
2. **Di chuyển con trỏ tại cạnh thấp hơn**:
   - Nếu `height[left] < height[right]` thì `left++`
   - Ngược lại `right--`

### Vì sao phải di chuyển cạnh thấp hơn?
Giả sử `height[left] <= height[right]`.
- Chiều cao hiện tại bị giới hạn bởi `height[left]`.
- Nếu bạn giảm `right` (giảm chiều rộng) nhưng vẫn giữ `left` (chiều cao giới hạn không tăng), thì diện tích **không thể lớn hơn** một cách chắc chắn.
- Chỉ khi tăng được “cạnh thấp hơn” (tức tìm `height[left]` lớn hơn) thì mới có cơ hội bù lại việc chiều rộng giảm.

Do đó ta luôn bỏ cạnh thấp hơn để tìm cơ hội tốt hơn.

---

## Độ phức tạp
- **Thời gian**: \(O(n)\) (mỗi con trỏ chỉ đi tối đa `n` bước)
- **Bộ nhớ**: \(O(1)\)

---

## Code
- Solution: `container_with_most_water.go`
- Tests: `container_with_most_water_test.go`


