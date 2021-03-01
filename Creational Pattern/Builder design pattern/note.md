## Mô tả
Việc tạo ra 1 biến có thể chỉ đơn giản bằng cách tạo ra 1 {} rồi để cho biến giá trị zero, hoặc tạo ra 1 object mà cần gọi vài API, kiểm tra vài trạng thái và tạo object cho các field của nó. Bạn cũng có thể tạo 1 object mà chứa nhiều objects khác

Cùng lúc đó bạn có thể phải tạo nhiều kiểu objects khác
For example, you'll use almost the same technique to build a car as you would build a bus,
except that they'll be of different sizes and number of seats, so why don't we reuse the
construction process? This is where the Builder pattern comes to the rescue.

## Nhiệm vụ
- Tạo 1 object từng bước 1 bằng cách điền các trường và tạo các object con trong nó
- Tái sử dụng phương thức tạo giữa nhiều object với nhau

## Tổng kết
Builder Pattern giúp chúng ta maintain 1 số lượng lớn sản phẩm mà chỉ cần sử dụng 1 kiểu chạy duy nhất

Đồng thời khi định nghĩa 1 `builder pattern` người mới đến đọc sẽ dễ hiểu luồng hơn. Nhưng bạn phải chắc chắn rằng bạn có 1 interface ổn định ít thay đổi để tạo ra `builder pattern`