# Product of Array Except Self

## Mô tả bài toán

Cho một mảng số nguyên `nums`, hãy trả về một mảng `answer` sao cho `answer[i]` bằng tích của tất cả các phần tử trong `nums` ngoại trừ `nums[i]`.

Tích của bất kỳ tiền tố hoặc hậu tố nào của `nums` đều được đảm bảo nằm trong giới hạn số nguyên 32-bit.

**Yêu cầu:**
- Thuật toán chạy với độ phức tạp thời gian $O(n)$.
- Không sử dụng phép chia.

### Ví dụ 1:
- **Input:** `nums = [1,2,3,4]`
- **Output:** `[24,12,8,6]`

### Ví dụ 2:
- **Input:** `nums = [-1,1,0,-3,3]`
- **Output:** `[0,0,9,0,0]`

### Ràng buộc:
- `2 <= nums.length <= 10^5`
- `-30 <= nums[i] <= 30`
- Kết quả đảm bảo nằm trong giới hạn số nguyên 32-bit.

---

## Phân tích cách tiếp cận

### 1. Cách tiếp cận Brute Force (Ngây thơ)
- Với mỗi phần tử $i$, chạy một vòng lặp tính tích của tất cả các phần tử khác.
- **Độ phức tạp:** Thời gian $O(n^2)$, Không gian $O(1)$.
- **Hạn chế:** Không đáp ứng yêu cầu $O(n)$ khi mảng lớn.

### 2. Cách tiếp cận sử dụng phép chia
- Tính tổng tích của toàn bộ mảng. Sau đó với mỗi phần tử $i$, kết quả là `totalProduct / nums[i]`.
- **Hạn chế:**
    - Đề bài cấm sử dụng phép chia.
    - Gặp vấn đề khi mảng có chứa số 0.

### 3. Cách tiếp cận Tối ưu (Prefix & Suffix Products)
- Ý tưởng: Với mỗi index $i$, kết quả là tích của (tất cả các số bên trái $i$) nhân với (tất cả các số bên phải $i$).
- Chúng ta có thể tính tích các số bên trái và lưu vào một mảng, sau đó tính tích các số bên phải.
- **Độ phức tạp:** Thời gian $O(n)$, Không gian $O(n)$ nếu dùng 2 mảng phụ.

---

## Thuật toán tối ưu (O(n) Time, O(1) Space)

Chúng ta có thể tối ưu không gian bằng cách tận dụng mảng kết quả `res` để lưu trữ tích bên trái (prefix products), sau đó duyệt ngược lại để nhân trực tiếp tích bên phải (suffix products) vào mảng đó.

### Các bước thực hiện:

1. **Khởi tạo:** Tạo mảng `res` cùng kích thước với `nums`.
2. **Duyệt xuôi (Prefix):**
   - Đặt `res[0] = 1` (vì không có phần tử nào bên trái của index 0).
   - Với $i$ từ $1$ đến $n-1$: `res[i] = res[i-1] * nums[i-1]`.
   - Sau bước này, `res[i]` chứa tích của tất cả các số trước $i$.
3. **Duyệt ngược (Suffix):**
   - Sử dụng một biến `suffixProduct` khởi tạo bằng $1$.
   - Duyệt từ $i = n-1$ về $0$:
     - `res[i] = res[i] * suffixProduct` (Nhân tích bên trái với tích bên phải hiện tại).
     - `suffixProduct = suffixProduct * nums[i]` (Cập nhật tích bên phải cho phần tử tiếp theo bên trái).
4. **Kết quả:** Trả về mảng `res`.

### Ví dụ minh họa: `nums = [1, 2, 3, 4]`

1. **Prefix Pass:**
   - `res[0] = 1`
   - `res[1] = res[0] * nums[0] = 1 * 1 = 1`
   - `res[2] = res[1] * nums[1] = 1 * 2 = 2`
   - `res[3] = res[2] * nums[2] = 2 * 3 = 6`
   - Mảng tạm thời: `res = [1, 1, 2, 6]`

2. **Suffix Pass:**
   - Khởi tạo `suffixProduct = 1`
   - $i=3$: `res[3] = 6 * 1 = 6`, `suffixProduct = 1 * 4 = 4`
   - $i=2$: `res[2] = 2 * 4 = 8`, `suffixProduct = 4 * 3 = 12`
   - $i=1$: `res[1] = 1 * 12 = 12`, `suffixProduct = 12 * 2 = 24`
   - $i=0$: `res[0] = 1 * 24 = 24`, `suffixProduct = 24 * 1 = 24`
   - Kết quả cuối: `[24, 12, 8, 6]`

---

## Giải thích Code

```go
func ProductExceptSelf(nums []int) []int {
    n := len(nums)
    res := make([]int, n)

    // Bước 1: Tính Prefix Products
    res[0] = 1
    for i := 1; i < n; i++ {
        res[i] = res[i-1] * nums[i-1]
    }

    // Bước 2: Tính Suffix Products và nhân vào res
    suffixProduct := 1
    for i := n - 1; i >= 0; i-- {
        res[i] *= suffixProduct
        suffixProduct *= nums[i]
    }

    return res
}
```

- **Độ phức tạp thời gian:** $O(n)$ vì chúng ta chỉ duyệt qua mảng 2 lần độc lập.
- **Độ phức tạp không gian:** $O(1)$ (không tính không gian của mảng kết quả theo yêu cầu đề bài).

---

## Trả lời Follow-up

**Câu hỏi:** Bạn có thể giải bài toán với độ phức tạp không gian bổ sung $O(1)$ không? (Mảng kết quả không được tính là không gian bổ sung cho phân tích độ phức tạp không gian.)

**Trả lời:** Có, giải pháp trên đã thực hiện điều này bằng cách sử dụng chính mảng kết quả để lưu trữ các giá trị trung gian (prefix products) và sử dụng một biến đơn lẻ (`suffixProduct`) để lưu trữ tích dồn từ bên phải, thay vì sử dụng thêm một mảng hậu tố (suffix array) riêng biệt.
