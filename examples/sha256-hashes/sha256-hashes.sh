# Chạy chương trình tính toán giá trị băm và in nó
# dưới dạng hex có thể đọc được bởi con người.
$ go run sha256-hashes.go
sha256 this string
1af1dfa857bf1d8814fe1af8983c18080019922e557f15a8a...


# Bạn có thể tính toán các băm khác sử dụng một
# mẫu tương tự như mẫu được hiển thị ở trên. Ví dụ,
# để tính toán các giá trị băm SHA512, import 
# `crypto/sha512` và sử dụng `sha512.New()`.

# Lưu ý rằng nếu bạn cần các giá trị băm an toàn 
# bằng mật mã, bạn nên nghiên cứu kỹ
# [sức mạnh băm](https://en.wikipedia.org/wiki/Cryptographic_hash_function)!
