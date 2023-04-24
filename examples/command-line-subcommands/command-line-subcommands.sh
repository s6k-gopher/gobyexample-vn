$ go build command-line-subcommands.go 

# Đầu tiên kích hoạt lệnh con foo.
$ ./command-line-subcommands foo -enable -name=joe a1 a2
subcommand 'foo'
  enable: true
  name: joe
  tail: [a1 a2]

# Bây giờ thử bar.
$ ./command-line-subcommands bar -level 8 a1
subcommand 'bar'
  level: 8
  tail: [a1]

# Nhưng bar sẽ không chấp nhận các cờ của foo.
$ ./command-line-subcommands bar -enable a1
flag provided but not defined: -enable
Usage of bar:
  -level int
    	level

# Tiếp theo, chúng ta sẽ xem xét biến môi trường,
# một cách phổ biến khác để tham số hóa 
# các chương trình.