# Prompt Template - Giải bài toán LeetCode bằng Go

## Cách sử dụng

Copy nội dung bài toán từ LeetCode (bao gồm mô tả, ví dụ, ràng buộc, function signature), sau đó gửi kèm prompt bên dưới.

---

## Prompt

Trong folder `<tên_folder>`, hãy viết code và test cho bài toán dưới, viết file .md mô tả bài toán và chi tiết cách giải.

### Yêu cầu:

1. **File code chính** (`<tên_file>.go`):
   - Package name phù hợp với folder cha
   - Comment giải thích thuật toán bằng tiếng Việt
   - Giữ đúng function signature theo đề bài LeetCode
   - Code clean, tối ưu, có comment giải thích từng bước

2. **File test** (`<tên_file>_test.go`):
   - Sử dụng table-driven tests
   - Bao gồm các test case: ví dụ từ đề bài, trường hợp biên (edge cases), số âm, số lớn, mảng rỗng, phần tử trùng lặp, v.v.
   - Helper functions nếu cần (so sánh slices, tạo linked list, v.v.)

3. **File README.md**:
   - Mô tả bài toán bằng tiếng Việt
   - Ví dụ minh họa
   - Ràng buộc
   - Phân tích các cách tiếp cận (từ brute force đến tối ưu), kèm độ phức tạp
   - Giải thích chi tiết thuật toán được chọn với ví dụ từng bước
   - Giải thích code (trích dẫn từng đoạn code quan trọng)
   - Phân tích độ phức tạp thời gian và không gian
   - Trả lời Follow-up questions (nếu có)

### Bài toán:

```
<Paste nội dung bài toán và function signature từ LeetCode vào đây>
```
