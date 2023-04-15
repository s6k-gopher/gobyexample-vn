# Chương trình được thực thi sẽ 
# hiển thị 5 tác vụ được thực thi bởi
# các worker khác nhau. Chương trình 
# chỉ mất khoảng 2 giây để thực thi, mặc dù
# nó thực hiện các tác vụ trong khoảng
# tổng cộng là 5 giây, vì có 3 worker chạy song song.
$ time go run worker-pools.go 
worker 1 started  job 1
worker 2 started  job 2
worker 3 started  job 3
worker 1 finished job 1
worker 1 started  job 4
worker 2 finished job 2
worker 2 started  job 5
worker 3 finished job 3
worker 1 finished job 4
worker 2 finished job 5

real	0m2.358s