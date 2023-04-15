# Lưu ý rằng dù slice khác với array, nhưng được 
# thể hiện cùng dạng với array khi sử dụng `fmt.Println`.
$ go run slices.go
emp: [  ]
set: [a b c]
get: c
len: 3
apd: [a b c d e f]
cpy: [a b c d e f]
sl1: [c d e]
sl2: [a b c d e]
sl3: [c d e f]
dcl: [g h i]
2d:  [[0] [1 2] [2 3 4]]

# Hãy xem qua blog này của team Go (https://go.dev/blog/slices-intro)
# để biết thêm nhiều chi tiết về thiết kế và 
# thực hiện mảng trong Go.

# Bây giờ chúng ta đã xem qua arrays và slices, 
# chúng ta sẽ xem qua cấu trúc dữ liệu 
# quan trọng khác của Go: maps.