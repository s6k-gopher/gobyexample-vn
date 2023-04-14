# Chúng ta nhận giá trị `"one"` và sau đó `"two"` như
# mong đợi.
$ time go run select.go 
received one
received two

# Luư ý rằng tổng thời gian thực thi chỉ xấp xỉ 2 giây
# kể từ khi cả giây 1 và giây 2 `Sleeps` thực thi đồng
# thời.
real	0m2.245s
