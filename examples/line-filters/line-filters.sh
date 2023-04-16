# Để thử line filter chúng ta, trước tiên hãy tạo một 
# tập tin với vài dòng chữ thường.
$ echo 'hello'   > /tmp/lines
$ echo 'filter' >> /tmp/lines

# Sau đó sử dụng line filter để lấy các dòng chữ in hoa.
$ cat /tmp/lines | go run line-filters.go
HELLO
FILTER