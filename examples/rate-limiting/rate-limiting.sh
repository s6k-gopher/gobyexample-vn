# Khi thực thi chương trình, 
# chúng ta thấy rằng các yêu cầu 
# của batch đầu tiên được xử lý 
# một lần mỗi 200 mili giây như mong muốn.
$ go run rate-limiting.go
request 1 2012-10-19 00:38:18.687438 +0000 UTC
request 2 2012-10-19 00:38:18.887471 +0000 UTC
request 3 2012-10-19 00:38:19.087238 +0000 UTC
request 4 2012-10-19 00:38:19.287338 +0000 UTC
request 5 2012-10-19 00:38:19.487331 +0000 UTC

# Đối với batch request thứ hai, 
# chúng ta xử lý 3 request đầu tiên
# ngay lập tức nhờ vào khả năng 
# tăng tốc của rate limiting, sau đó xử lý 2 request
# còn lại với khoảng thời gian trễ ~200ms mỗi request.
request 1 2012-10-19 00:38:20.487578 +0000 UTC
request 2 2012-10-19 00:38:20.487645 +0000 UTC
request 3 2012-10-19 00:38:20.487676 +0000 UTC
request 4 2012-10-19 00:38:20.687483 +0000 UTC
request 5 2012-10-19 00:38:20.887542 +0000 UTC