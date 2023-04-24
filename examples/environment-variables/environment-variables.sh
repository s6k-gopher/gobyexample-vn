# Chạy chương trình sẽ cho thấy chúng ta lấy được
# giá trị cho `FOO` mà chúng ta thiết lập, nhưng
# `BAR` là rỗng.
$ go run environment-variables.go
FOO: 1
BAR: 

# Danh sách các khoá trong môi trường sẽ phụ thuộc 
# vào máy của bạn.
TERM_PROGRAM
PATH
SHELL
...
FOO

# Nếu chúng ta thiết lập `BAR` trong môi trường trước, 
# chương trình chạy sẽ lấy giá trị đó.
$ BAR=2 go run environment-variables.go
FOO: 1
BAR: 2
...