// Trường hợp chúng ta muốn thực thi code Go tại một thời
// điểm trong tương lai, hoặc lặp lại sau một khoảng thời
// gian cố định. Go cung cấp _timer_ and _ticker_ hỗ trợ
// thực hiện các việc trên dễ dàng hơn. Chúng ta bắt đầu
// với timers trong ví dụ này
// (tickers sẽ được giới thiệu ở phần tiếp theo)

package main

import (
	"fmt"
	"time"
)

func main() {

	// Timers thể hiện một sự kiện thực thi trong tương lai.
	// Bạn quyết định timer sẽ chờ trong bao lâu, và timer sẽ
	// tạo ra channel và thực thi sau đúng thời gian chờ.
	// Ví dụ timer1 sẽ chờ 2 giây.
	timer1 := time.NewTimer(2 * time.Second)

	// `<-timer1.C` sẽ buộc channel `C` phải chờ cho đến khi
	// hết thời gian chờ của timer1 (2 giây) thì mới thực thi
	<-timer1.C
	fmt.Println("Timer 1 fired")

	// Nếu bạn chỉ muốn chờ một khoảng thời gian, bạn có thể
	// sử dụng `time.Sleep`. Việc sử dụng timer sẽ tiện hơn
	// vì bạn có thể hủy timer trước khi nó bắt đầu chạy.
	// Xem ví dụ dưới đây để hiểu rõ hơn.
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	// Cho dù chờ đủ thời gian 2 giây
	// `time.Sleep(2 * time.Second)`
	// Kết quả vẫn là "Timer 2 stopped"
	// vì timer đã được hủy bởi `timer2.Stop()`
	// trước khi đến giờ thực thi.
	time.Sleep(2 * time.Second)
}
