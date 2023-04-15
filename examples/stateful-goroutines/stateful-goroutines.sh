# Running our program shows that the goroutine-based
# state management example completes about 80,000
# total operations.
# Khi chạy chương trình chỉ ra rằng ví dụ của sự quản
# lý trạng thái dựa trên goroutine hoàn thành tổng cộng 
# khoảng 80,000 phép tính. 
$ go run stateful-goroutines.go
readOps: 71708
writeOps: 7177

# Đối với trường hợp cụ thể này, hướng tiếp cận dựa trên
# goroutine liên quan nhiều hơn đến cái dựa trên mutex.
# Nó có thể hữu ích trong một số trường hợp cụ thể, ví dụ
# khi bạn có nhiều channel tham gia hoặc khi quản lý 
# nhiều mutex như vậy sẽ gây lỗi. Bạn nên sử dụng bất kì 
# phương pháp nào cảm thấy tự nhiên nhất, đặc biệt là 
# đối với sự hiểu biết đúng đắn của chương trình.