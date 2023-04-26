// Cơ chế mặc định để quản lý state trong Go là giao tiếp
// qua channels. Chúng ta đã thấy điều này ở ví dụ với
// [worker pools](worker-pools). Ngoài ra còn có một vài
// lựa chọn khác để quản lý state. Ở đây chúng ta sẽ
// xem xét việc sử dụng package `sync/atomic` cho _các biến đếm
// nguyên tử (atomic counters)_ được truy cập bởi nhiều goroutines.

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	// Ta sẽ dùng một số nguyên không dấu để biểu diễn
	// biến đếm (luôn dương).
	var ops uint64

	// Một WaitGroup sẽ giúp chúng ta đợi cho tất cả
	// các goroutine hoàn thành việc thực thi.
	var wg sync.WaitGroup

	// Ta sẽ khởi tạo 50 goroutine mà mỗi goroutine sẽ
	// tăng biến đếm lên 1000 lần.
	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {
				// Để tự động tăng biến đếm, ta sử dụng
				// `AddUint64`, truyền vào địa chỉ vùng nhớ
				// của biến đếm `ops` với
				// cú pháp `&`.
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}

	// Chờ cho tất cả các goroutine hoàn thành.
	wg.Wait()

	// Bây giờ ta có thể an toàn truy cập vào `ops` vì ta
	// biết rằng không có goroutine nào khác đang ghi vào nó.
	// Việc đọc các biến đếm nguyên tử một cách an toàn trong
	// khi chúng đang được cập nhật cũng khả thi, sử dụng các
	// hàm như `atomic.LoadUint64`.
	fmt.Println("ops:", ops)
}
