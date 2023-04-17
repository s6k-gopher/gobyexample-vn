# Sử dụng những lệnh này để chạy example.
# (Lưu ý: bởi vì giới hạn của go payground, 
# example này chỉ được chọn trên máy cá nhân).
$ mkdir -p folder
$ echo "hello go" > folder/single_file.txt
$ echo "123" > folder/file1.hash
$ echo "456" > folder/file2.hash

$ go run embed-directive.go
hello go
hello go
123
456
