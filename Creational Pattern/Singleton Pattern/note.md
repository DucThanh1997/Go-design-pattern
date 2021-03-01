## Mô tả
Là 1 đơn vị đơn lẻ của object, và được chắc chắn rằng nó không có cái bị trùng

Trong lần gọi đầu tiên, nó sẽ được tạo và được tái sử dụng trong tất cả các phần của app mà cần 1 hành động cụ thể

Nó được sử dụng trong các trường hợp
    + Khi bạn sử dụng 1 connection đến với database vào mỗi query
    + Khi bạn sử dụng SSH vào server để làm vài thứ mà ko muốn mở lại vào mỗi lần connect 
    + Nếu bạn muốn limit access tới 1 vài biến hoặc 1 khoảng không gian, bạn sẽ dùng Singleton
    + Nếu bạn muốn limit số lần gọi ở trong 1 vài nơi


## Nhiệm vụ
- Chúng ta cần 1 single, và shared value, của 1 kiểu nhất định
- Chúng ta cần giới hạn việc tạo 1 object của 1 vài kiểu trong suốt quá trình chạy 

