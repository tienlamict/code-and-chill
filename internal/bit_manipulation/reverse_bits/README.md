# Reverse Bits

## Mô tả bài toán

Đảo ngược các bit của một số nguyên 32-bit đã cho.

**Ví dụ:**

### Example 1:
- Input: `n = 43261596`
- Output: `964176192`
- Giải thích:
  ```
  Integer          Binary
  43261596         00000010100101000001111010011100
  964176192        00111001011110000010100101000000
  ```

### Example 2:
- Input: `n = 2147483644`
- Output: `1073741822`
- Giải thích:
  ```
  Integer          Binary
  2147483644       01111111111111111111111111111100
  1073741822       00111111111111111111111111111110
  ```

## Ràng buộc

- `0 <= n <= 2^31 - 2`
- `n` là số chẵn

## Phân tích các cách tiếp cận

### 1. Brute Force - Chuyển đổi sang chuỗi nhị phân

**Ý tưởng**: Chuyển số thành chuỗi nhị phân, đảo ngược chuỗi, rồi chuyển lại thành số.

**Cách hoạt động**:
- Chuyển số thành chuỗi nhị phân 32-bit (có padding với số 0 ở đầu)
- Đảo ngược chuỗi
- Chuyển chuỗi đã đảo ngược thành số nguyên

**Độ phức tạp**:
- Thời gian: O(32) = O(1) - duyệt qua 32 ký tự
- Không gian: O(32) = O(1) - lưu chuỗi nhị phân

**Nhược điểm**: Cần xử lý chuỗi, không tối ưu về mặt hiệu suất.

### 2. Bit Manipulation (Thuật toán được chọn)

**Ý tưởng**: Sử dụng các phép toán bit để đảo ngược từng bit một cách trực tiếp.

**Cách hoạt động**:
- Duyệt qua tất cả 32 bit của số đầu vào
- Với mỗi bit, lấy bit cuối cùng (LSB - Least Significant Bit) của `n`
- Thêm bit đó vào kết quả (đã được dịch chuyển phù hợp)
- Dịch chuyển `n` sang phải 1 bit và kết quả sang trái 1 bit

**Độ phức tạp**:
- Thời gian: O(1) - luôn duyệt đúng 32 bit
- Không gian: O(1) - chỉ sử dụng biến kết quả

**Ưu điểm**: 
- Hiệu quả cao, không cần chuyển đổi sang chuỗi
- Code ngắn gọn và dễ hiểu
- Tận dụng các phép toán bit nhanh của CPU

### 3. Lookup Table (Tối ưu cho Follow-up)

**Ý tưởng**: Tạo bảng tra cứu (lookup table) để đảo ngược từng byte (8 bit) một cách nhanh chóng.

**Cách hoạt động**:
- Tạo bảng tra cứu cho tất cả 256 giá trị có thể của 1 byte (0-255)
- Chia số 32-bit thành 4 byte
- Đảo ngược từng byte bằng cách tra cứu trong bảng
- Sắp xếp lại các byte đã đảo ngược

**Độ phức tạp**:
- Thời gian: O(1) - chỉ cần 4 lần tra cứu
- Không gian: O(256) = O(1) - lưu bảng tra cứu

**Ưu điểm**: Rất nhanh khi hàm được gọi nhiều lần (giải quyết Follow-up question)

**Nhược điểm**: Cần khởi tạo bảng tra cứu, phức tạp hơn trong triển khai

## Giải thích chi tiết thuật toán được chọn

### Thuật toán: Bit Manipulation

Chúng ta sẽ duyệt qua từng bit của số đầu vào và xây dựng kết quả từ phải sang trái.

**Các bước chi tiết:**

1. **Khởi tạo**: `result = 0`

2. **Vòng lặp 32 lần** (cho mỗi bit):
   - Lấy bit cuối cùng của `n`: `bit = n & 1`
   - Dịch chuyển `result` sang trái 1 bit và thêm `bit` vào: `result = (result << 1) | bit`
   - Dịch chuyển `n` sang phải 1 bit: `n >>= 1`

3. **Trả về**: `result`

**Ví dụ từng bước với n = 43261596:**

```
n = 43261596 = 00000010100101000001111010011100 (32 bit)
result = 0

Bước 0: bit = n & 1 = 0
        result = (0 << 1) | 0 = 0
        n = n >> 1 = 00000001010010100000111101001110

Bước 1: bit = n & 1 = 0
        result = (0 << 1) | 0 = 0
        n = n >> 1 = 00000000101001010000011110100111

Bước 2: bit = n & 1 = 1
        result = (0 << 1) | 1 = 1
        n = n >> 1 = 00000000010100101000001111010011

Bước 3: bit = n & 1 = 1
        result = (1 << 1) | 1 = 11
        n = n >> 1 = 00000000001010010100000111101001

... (tiếp tục cho đến bước 31)

Kết quả cuối cùng: 00111001011110000010100101000000 = 964176192
```

**Giải thích các phép toán bit:**

- `n & 1`: Lấy bit cuối cùng (LSB). Nếu bit cuối là 1 thì kết quả là 1, nếu là 0 thì kết quả là 0.
- `result << 1`: Dịch chuyển `result` sang trái 1 bit, tương đương nhân với 2.
- `result | bit`: Thêm bit mới vào vị trí cuối cùng của `result`.
- `n >>= 1`: Dịch chuyển `n` sang phải 1 bit, loại bỏ bit cuối cùng đã xử lý.

## Giải thích code

### Hàm chính

```12:32:internal/bit_manipulation/reverse_bits/reverse_bits.go
func reverseBits(n uint32) uint32 {
	var result uint32

	// Duyệt qua tất cả 32 bit
	for i := 0; i < 32; i++ {
		// Lấy bit cuối cùng (LSB) của n bằng cách AND với 1
		bit := n & 1

		// Dịch chuyển kết quả sang trái 1 bit và thêm bit mới vào vị trí cuối
		result = (result << 1) | bit

		// Dịch chuyển n sang phải 1 bit để xử lý bit tiếp theo
		n >>= 1
	}

	return result
}
```

**Giải thích từng phần:**

1. **Khởi tạo biến `result`**: Biến này sẽ chứa kết quả sau khi đảo ngược các bit.

2. **Vòng lặp `for i := 0; i < 32; i++`**: Duyệt qua đúng 32 bit của số nguyên 32-bit.

3. **`bit := n & 1`**: 
   - Phép toán AND với 1 sẽ trả về bit cuối cùng của `n`
   - Nếu bit cuối là 1 → `bit = 1`
   - Nếu bit cuối là 0 → `bit = 0`

4. **`result = (result << 1) | bit`**:
   - `result << 1`: Dịch chuyển `result` sang trái 1 bit (tạo chỗ trống ở cuối)
   - `| bit`: Thêm bit mới vào vị trí cuối cùng
   - Ví dụ: nếu `result = 101` và `bit = 1`, thì `result` mới = `1011`

5. **`n >>= 1`**: Dịch chuyển `n` sang phải 1 bit để loại bỏ bit đã xử lý và chuẩn bị cho bit tiếp theo.

## Phân tích độ phức tạp

### Độ phức tạp thời gian: O(1)

Mặc dù có vòng lặp, nhưng vòng lặp này luôn chạy đúng 32 lần (số bit cố định của số nguyên 32-bit), không phụ thuộc vào giá trị của `n`. Do đó, độ phức tạp thời gian là O(1).

### Độ phức tạp không gian: O(1)

Chỉ sử dụng một số biến cố định (`result`, `bit`, `i`), không sử dụng cấu trúc dữ liệu phụ thuộc vào kích thước đầu vào. Do đó, độ phức tạp không gian là O(1).

## Trả lời Follow-up questions

### Follow-up: Nếu hàm này được gọi nhiều lần, làm thế nào để tối ưu?

**Giải pháp: Sử dụng Lookup Table**

Khi hàm được gọi nhiều lần, ta có thể tối ưu bằng cách sử dụng bảng tra cứu (lookup table) để đảo ngược từng byte (8 bit) một cách nhanh chóng.

**Cách triển khai:**

1. **Tạo bảng tra cứu**: Tạo một mảng 256 phần tử, mỗi phần tử chứa giá trị đảo ngược của byte tương ứng (0-255).

2. **Chia số thành 4 byte**: Một số 32-bit có thể được chia thành 4 byte:
   - Byte 0: bit 0-7 (LSB)
   - Byte 1: bit 8-15
   - Byte 2: bit 16-23
   - Byte 3: bit 24-31 (MSB)

3. **Đảo ngược từng byte**: Sử dụng bảng tra cứu để đảo ngược từng byte.

4. **Sắp xếp lại**: Đảo ngược thứ tự các byte và kết hợp lại.

**Ví dụ:**

```go
// Khởi tạo lookup table (chỉ cần làm 1 lần)
var lookupTable [256]uint32

func init() {
    for i := 0; i < 256; i++ {
        // Đảo ngược byte i
        reversed := uint32(0)
        for j := 0; j < 8; j++ {
            bit := (i >> j) & 1
            reversed = (reversed << 1) | bit
        }
        lookupTable[i] = reversed
    }
}

func reverseBitsOptimized(n uint32) uint32 {
    // Chia thành 4 byte và đảo ngược từng byte
    byte0 := lookupTable[(n >> 0) & 0xFF]  // bit 0-7
    byte1 := lookupTable[(n >> 8) & 0xFF]  // bit 8-15
    byte2 := lookupTable[(n >> 16) & 0xFF] // bit 16-23
    byte3 := lookupTable[(n >> 24) & 0xFF] // bit 24-31
    
    // Kết hợp lại với thứ tự đảo ngược
    return (byte0 << 24) | (byte1 << 16) | (byte2 << 8) | byte3
}
```

**Độ phức tạp:**
- Thời gian: O(1) - chỉ cần 4 lần tra cứu trong bảng
- Không gian: O(256) = O(1) - lưu bảng tra cứu

**Ưu điểm:**
- Nhanh hơn phương pháp bit manipulation khi được gọi nhiều lần
- Bảng tra cứu chỉ cần khởi tạo một lần

**Nhược điểm:**
- Cần thêm bộ nhớ cho bảng tra cứu
- Code phức tạp hơn một chút

**Khi nào nên dùng:**
- Khi hàm được gọi rất nhiều lần (hàng triệu lần)
- Khi hiệu suất là yếu tố quan trọng nhất
- Khi có đủ bộ nhớ để lưu bảng tra cứu

**Khi nào không cần:**
- Khi hàm chỉ được gọi vài lần
- Khi code đơn giản và dễ đọc là ưu tiên
- Khi bộ nhớ bị hạn chế
