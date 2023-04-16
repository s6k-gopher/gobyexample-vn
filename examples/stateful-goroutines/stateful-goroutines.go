// Trong ví dụ trước chúng ta sử dụng explicit locking
// với [mutexes](mutexes) để đồng bộ hóa truy cập chia
// sẻ trạng thái giữa nhiều goroutine. Một lựa chọn khác
// là sử dụng tính năng đồng bộ hóa built-in của
// goroutine và channel để đạt được kết quả tương tự.
// Hướng tiếp cận dựa trên channel này phù hợp với ý
// tưởng chia sẻ bộ nhớ của Go bằng cách giao tiếp và sở
// hữu từng phần dữ liệu bởi chính xác một goroutine.

package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// Trong ví dụ này trạng thái sẽ được sử hữu bởi một
// goroutine duy nhất. Điều này sẽ đảm bảo rằng dữ liệu
// sẻ không bao giờ bị hỏng hoặc truy cập đồng thời. Để
// đọc hoặc khi trạng thái đó, những goroutine khác sẽ
// gửi thông tin đến goroutine đang sử hữu và nhận trả
// lời tương ứng. Những `readop` và `writeOp` `struct`
// này gói gọn những yêu cầu đó và một cách để goroutine
// sở hữu để đáp ứng.
type readOp struct {
	key  int
	resp chan int
}
type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {

	// Như trước đó, chúng ta sẽ đếm số lượng toán tử
	// được thực thi.
	var readOps uint64
	var writeOps uint64

	// Các channel `reads` và `writes` sẽ được sử dụng
	// bởi các goroutine khác để đưa ra các yêu cầu đọc
	// và ghi một cách tương ứng.
	reads := make(chan readOp)
	writes := make(chan writeOp)

	// Đây là goroutine mà chiếm quyền `state`, là một
	// map như trong ví dụ trước nhưng bây giờ là riêng
	// tư với stateful goroutine. Goroutine này lặp đi lặp
	// lại chọn trên channel `reads` and `writes`, đáp ứng
	//  yêu cầu khi chúng đến. Một response được thực
	// hiện bằng cách thực hiện đầu tiên yêu càu hoạt
	// động và sau đó  gửi giá trị trên channel response
	// `resp` để biểu thị thành công (và giá trị mong muốn
	//  trong trường hợp `reads`).
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// Chạy 100 goroutine để thực hiện tháo tác reads đến
	// goroutine sở hữu trạng thái qua channel `reads`.
	// MỖi lần đọc yêu cầu xây dựng một `readOp`, gửi
	// nó qua channel `reads`, and sau đó nhận về kết quả
	// trên channel `resp` được cung cấp.
	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Chúng ta cho chạy 10 toán tử writes, sử dụng
	// cách tiếp cận tương tự.
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Để goroutine hoạt động trong một giây.
	time.Sleep(time.Second)

	// Cuối cùng, capture và báo cáo số lượng toán tử.
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
}
