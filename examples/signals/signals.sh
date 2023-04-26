# Khi ta chạy chương trình này, nó sẽ chờ 
# một tín hiệu. Bằng cách gõ `ctrl-C` 
# (terminal sẽ hiển thị là `^C`)
# ta có thể gửi một tín hiệu `SIGINT`,
# khiến cho chương trình in ra `interrupt` 
# và sau đó thoát.
$ go run signals.go
awaiting signal
^C
interrupt
exiting
