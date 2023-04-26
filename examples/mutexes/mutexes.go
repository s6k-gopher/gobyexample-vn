// Trong ví dụ trước, chúng ta đã thấy cách quản lý
// trạng thái bộ đếm đơn giản bằng cách sử dụng [atomic operations](atomic-counters).
// Đối với trạng thái phức tạp hơn, chúng ta có thể sử dụng một [_mutex_](https://en.wikipedia.org/wiki/Mutual_exclusion)
// để truy cập dữ liệu một cách an toàn giữa nhiều goroutines.
package main

import (
	"fmt"
	"sync"
)

// Struct Container chứa một map các counters; vì chúng ta muốn
// cập nhật nó đồng thời từ nhiều goroutines, chúng ta
// thêm một `Mutex` để đồng bộ hóa truy cập vào nó.
// Lưu ý rằng không được copy mutexes, nên nếu
// `struct` này được truyền đi chỗ khác, ta nên sử dụng
// con trỏ.
type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	// Khoá mutex trước khi truy cập `counters`; giải phóng
	// nó ở cuối hàm bằng cách sử dụng một câu lệnh [defer](defer).
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func main() {
	c := Container{
		// Lưu ý rằng ta có thể sử dụng giá trị mặc định là không của mutex, do đó
		// không cần phải khởi tạo nó.
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	// Hàm sau đây làm tăng giá trị của một biến đếm
	// trong một vòng lặp.
	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}

	// Thực thi đồng thời nhiều goroutines; lưu ý rằng
	// chúng đều truy cập vào cùng một `Container`,
	// và hai trong số chúng truy cập vào cùng một biến đếm.
	wg.Add(3)
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)

	// Chờ cho các goroutines chạy xong.
	wg.Wait()
	fmt.Println(c.counters)
}
