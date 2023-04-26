# Ta mong đợi sẽ nhận được chính xác 50,000 phép tính.
# Nếu chúng ta sử dụng phép tính không đảm bảo
# tính toàn vẹn dữ liệu như `ops++` để tăng biến đếm,
# ta có thể nhận được một kết quả khác, thay đổi giữa
# các lần chạy, bởi vì các goroutines sẽ can thiệp lẫn
# nhau. Hơn nữa, ta sẽ dính các lỗi về chạy đua dữ liệu
# (data race) khi sử dụng flag `-race`.
$ go run atomic-counters.go
ops: 50000

# Tiếp theo, chúng ta sẽ xem xét mutex, một công cụ 
# khác để quản lý state.