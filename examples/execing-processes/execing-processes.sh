# Khi chạy chương trình của chúng ta, 
# process sẽ được thay thế bởi `ls`.
$ go run execing-processes.go
total 16
drwxr-xr-x  4 mark 136B Oct 3 16:29 .
drwxr-xr-x 91 mark 3.0K Oct 3 12:50 ..
-rw-r--r--  1 mark 1.3K Oct 3 16:28 execing-processes.go

# Lưu ý rằng Go không cung cấp một hàm `fork`
# Unix. Thường thì điều này không phải là vấn đề,
# vì khi bắt đầu các goroutine, tạo các process, 
# và exec các process đều đáp ứng được hầu hết 
# các trường hợp sử dụng của `fork`.