# Chuỗi được mã hóa thành các giá trị khác nhau với các
# mã hóa base64 chuẩn và URL (đuôi `+` vs `-`)
# nhưng chúng đều được giải mã thành  chuỗi ban đầu
# như mong muốn
$ go run base64-encoding.go
YWJjMTIzIT8kKiYoKSctPUB+
abc123!?$*&()'-=@~

YWJjMTIzIT8kKiYoKSctPUB-
abc123!?$*&()'-=@~