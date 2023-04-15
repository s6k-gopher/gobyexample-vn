# Chạy chương trình sẽ gây ra panic, in một thông báo
# lỗi và dấu vết goroutine, và thoát với một trạng 
# thái non-zero.


# Khi panic đầu tiên trong `main` được bắn ra, chương
# trình thoát mà không đạt đến phần còn lại của code.
# Nếu bạn muốn để xem chương trình hãy thửu tạo một
# tệp tạm thời, comment lại cái panic đầu tiên.

$ go run panic.go
panic: a problem

goroutine 1 [running]:
main.main()
    /.../panic.go:12 +0x47
...
exit status 2

# Lưu ý rằng không giống như một số ngôn ngữ sử dụng
# exception để xử lý nhiều lỗi, trong Go nó là idiomatic 
# để sử dụng các giá trị trả về chỉ báo lỗi bất cứ
# khi nào có thể.