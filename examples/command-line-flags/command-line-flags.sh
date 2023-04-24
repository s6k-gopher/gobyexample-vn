# Để thử nghiệm chương trình các cờ dòng lệnh,
# trước tiên là biên dịch nó và chạy tệp nhị phân
# kết quả trực tiếp.
$ go build command-line-flags.go

# Hãy thử chương trình đã xây dựng bằng cách cung cấp
# giá trị cho tất cả các cờ
$ ./command-line-flags -word=opt -numb=7 -fork -svar=flag
word: opt
numb: 7
fork: true
svar: flag
tail: []

# Lưu ý rằng nếu bạn bỏ qua các cờ chúng sẽ tự động
# lấy giá trị mặc định của chúng. 
$ ./command-line-flags -word=opt
word: opt
numb: 42
fork: false
svar: bar
tail: []

# Đối số vị trí theo sau có thể được cung cấp sau
# bất kỳ cờ nào.
$ ./command-line-flags -word=opt a1 a2 a3
word: opt
...
tail: [a1 a2 a3]

# Lưu ý rằng package (gói) `flag` yêu cầu tất cả
# các cờ xuất hiện trước các đối số vị trí (nếu không
# các cờ sẽ được hiểu là các đối số vị trí).
$ ./command-line-flags -word=opt a1 a2 a3 -numb=7
word: opt
numb: 42
fork: false
svar: bar
tail: [a1 a2 a3 -numb=7]

# Sử dụng cờ `-h` hoặc `--help` để nhân văn bản
# trợ giúp được tạo tự động cho chương trình dòng lệnh.
$ ./command-line-flags -h
Usage of ./command-line-flags:
  -fork=false: a bool
  -numb=42: an int
  -svar="bar": a string var
  -word="foo": a string

# Nếu bạn cung cấp một cờ mà không được chỉ định
# cho package `flag`, chương trình sẽ in thông báo
# lỗi và hiển thị văn bản trợ giúp
$ ./command-line-flags -wat
flag provided but not defined: -wat
Usage of ./command-line-flags:
...