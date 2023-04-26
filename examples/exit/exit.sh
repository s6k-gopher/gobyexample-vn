# Nếu bạn chạy `exit.go` bằng `go run`,
# việc thoát chương trình sẽ được thực thi
# và được in ra bởi `go`.
$ go run exit.go
exit status 3

# Bằng cách biên dịch và thực thi một file binary
# bạn có thể thấy được status trong terminal.
$ go build exit.go
$ ./exit
$ echo $?
3

# Lưu ý rằng ký tự `!` trong chương trình của chúng ta 
# không được in ra.