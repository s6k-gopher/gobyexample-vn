# Chạy chương trình sẽ cho thấy rằng các biến đếm 
# được cập nhật như mong đợi.
$ go run mutexes.go
map[a:20000 b:10000]

# Tiếp theo chúng ta sẽ xem xét 
# việc thực hiện cùng một tác vụ
# quản lý trạng thái sử dụng goroutines và channels.