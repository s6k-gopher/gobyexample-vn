# Để thử nghiệm với tham số command-line, tốt nhất là 
# build binary với `go build` trước.
$ go build command-line-arguments.go
$ ./command-line-arguments a b c d
[./command-line-arguments a b c d]       
[a b c d]
c
# Sau đó, chúng ta sẽ đi sâu hơn vào xử lý command-line
# với flag.
