// Trong ví dụ này, chúng ta sẽ xem xét cách triển khai
// một _worker pool_ sử dụng goroutines và channels.

package main

import (
	"fmt"
	"time"
)

// Dưới đây là một worker mà ta sẽ thực thi nhiều
// instance cùng lúc. Các worker này sẽ nhận
// tác vụ từ channel `jobs` và gửi kết quả tương ứng
// đến channel `results`. Ta sẽ dừng một giây mỗi tác vụ
// để giả lập một tác vụ ngốn rất nhiều thời gian.
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {

	// Để sử dụng pool worker của chúng ta, ta cần gửi
	// tác vụ cho chúng và thu thập kết quả. Ta sẽ tạo
	// 2 channel để thực hiện việc này.
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Dưới đây ta khởi tạo 3 worker, đầu tiên chúng sẽ
	// bị chặn vì chưa có tác vụ nào.
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Ở đây ta gửi 5 `tác vụ` và sau đó `đóng` channel
	// đó để chỉ ra rằng đây là tất cả các tác vụ mà ta có.
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Cuối cùng, ta thu thập tất cả các kết quả xử lý các tác vụ.
	// Điều này cũng đảm bảo rằng các goroutine worker đã hoàn thành.
	// Một cách khác để đợi kết quả từ nhiều
	// goroutine là sử dụng [WaitGroup](waitgroups).
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
