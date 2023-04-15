# Khi chúng ta chạy chương trình này, chúng ta sẽ thấy
# đầu tiên là kết quả của lần gọi đồng bộ, sau đó là
# kết quả của hai goroutines. Kết quả của các 
# goroutines' có thể được xen kẽ lẫn nhau vì các 
# goroutine đang được thực thi song song bởi 
# Go runtime.
$ go run goroutines.go
direct : 0
direct : 1
direct : 2
goroutine : 0
going
goroutine : 1
goroutine : 2
done

# Tiếp theo, chúng ta sẽ xem xét một phần bổ sung 
# cho goroutine trong các concurrent Go program 
#(chương trình Go  đồng thời): channels (kênh).