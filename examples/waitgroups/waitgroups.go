// Để đợi nhiều goroutine kết thúc, ta có thể sử dụng
// *wait group*.
package main

import (
	"fmt"
	"sync"
	"time"
)

// Đây sẽ là hàm mà ta sẽ chạy trong mỗi goroutine.
func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	// Dừng một khoảng để giả lập một tác vụ ngốn rất nhiều thời gian.
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {

	// WaitGroup này được dùng để đợi cho tất cả các goroutine
	// được khởi tạo ở đây chạy xong. Lưu ý: nếu một WaitGroup được
	// truyền vào các hàm, nó nên được truyền bằng cách *sử dụng con trỏ*.
	var wg sync.WaitGroup

	// Khởi chạy một vài goroutine và tăng giá trị số lượng của WaitGroup
	// cho mỗi goroutine.
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		// Tránh việc sử dụng lại cùng một giá trị `i` trong mỗi closure của goroutine.
		// Đọc [FAQ](https://golang.org/doc/faq#closures_and_goroutines)
		i := i

		// Bọc hàm gọi worker trong một closure để đảm bảo rằng
		// WaitGroup sẽ được thông báo khi một worker kết thúc. Bằng cách này
		// worker không cần phải quan tâm đến các kiểu nguyên thuỷ của concurrency
		// (concurrency primitives) trong quá trình thực thi.
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	// Chặn cho đến khi giá trị số lượng của WaitGroup trở về 0, tức là
	// tất cả các worker đã thông báo rằng chúng đã chạy xong.
	wg.Wait()

	// Chú ý rằng: phương pháp này không có cách nào
	// để truyền lỗi trực tiếp từ các worker. Để sử dụng
	// trong các trường hợp nâng cao hơn, hãy xem xét sử dụng
	// [errgroup package](https://pkg.go.dev/golang.org/x/sync/errgroup).
}
