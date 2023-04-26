// _Closing_ Đóng một channel nghĩa là sẽ không có giá trị
// nào được gửi vào channel nữa. Điều này có thể dùng để
// xác định việc hoàn thành thao tác nhận giá trị
// từ channel.

package main

import "fmt"

// Trong ví dụ này, chúng ta sẽ sử dụng một channel `jobs`
// để truyền `job` cần thực hiện từ `main()` đến một
// `worker`. Khi không còn `job` nào cho `worker` thực
// hiện thì chúng ta sẽ dùng `close` để đóng
// channel `jobs`.
func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	// Dưới đây là một ví dụ về dùng vòng lặp trong `worker`
	// nhận giá trị từ `jobs` với `j, more := <-jobs`.
	// Trong trường hợp không còn `job` mới, giá trị `more`
	// sẽ là `false` nếu `jobs` đã được `close` và tất cả
	// các giá trị trong channel đã được nhận. Chúng ta sử
	// dụng điều này để thông báo cho `done` khi `worker`
	// đã hoàn thành việc nhận tất cả `job`.
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// Có 3 `jobs` được gửi đến worker thông qua channel
	// `jobs`, sau đó channel `jobs` sẽ được đóng.
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	// Chúng ta chờ worker thực hiện công việc thông qua
	// phương pháp [synchronization](channel-synchronization)
	// chúng ta đã có ở ví dụ trước.
	<-done
}
