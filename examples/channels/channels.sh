# Khi chúng ra chạy chương trình gói tin `"ping"`
# được truyền thành công từ một goroutine đến một
# goroutine khác thông qua channel của chúng ta.

$ go run channels.go 
ping

# Mặc định gửi và nhận bị chặn cho đến khi cả 
# channel gửi và nhận sẵn sàng. Thuộc tính này cho phép
# chúng ta đợi gói tin `"ping"` ở cuối chương trình mà 
# không phải sử dụng bất kì sự đồng bộ hóa nào khác. 
