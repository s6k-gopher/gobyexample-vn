# `zeroval` không thay đổi `i` trong `main`, nhưng
# `zeroptr` thì có vì nó có một tham chiếu đến
# địa chỉ bộ nhớ của biến `i` đó.
$ go run pointers.go
initial: 1
zeroval: 1
zeroptr: 0
pointer: 0x42131100