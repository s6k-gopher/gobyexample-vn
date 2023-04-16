# Hãy thử chạy code (mã) ghi tệp tin.
$ go run writing-files.go 
wrote 5 bytes
wrote 7 bytes
wrote 9 bytes

# Sau đó kiểm tra nội dung của các tập tin đã viết.
$ cat /tmp/dat1
hello
go
$ cat /tmp/dat2
some
writes
buffered

# Tiếp theo, chúng ta sẽ xem xét áp dụng một số 
# ý tưởng I/O tệp mà chúng ta vừa thấy cho các
# luồng `sdtin` và `stdout`.