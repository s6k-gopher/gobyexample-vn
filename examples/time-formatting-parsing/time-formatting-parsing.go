// Go hỗ trợ định dạng và phân tích thời gian
// thông qua các layout dựa trên patterns.
package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	// Dưới đây là một ví dụ cơ bản về định dạng thời gian
	// theo RFC3339, sử dụng hằng số layout tương ứng.
	t := time.Now()
	p(t.Format(time.RFC3339))

	// Phân tích thời gian sử dụng cùng các giá trị layout
	// với `Format`.
	t1, e := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	p(t1)

	// `Format` và `Parse` sử dụng các layout dựa trên ví dụ.
	// Thông thường bạn sẽ sử dụng một hằng số từ package `time` cho
	// các layout này, nhưng bạn cũng có thể cung cấp các layout
	// tùy chỉnh. Các layout phải sử dụng thời gian tham chiếu
	// `Mon Jan 2 15:04:05 MST 2006` để hiển thị pattern dùng để
	// định dạng/ phân tích một đối tượng time/chuỗi thời gian cho trước.
	// Thời gian được dùng làm ví dụ phải chính xác như sau: năm 2006,
	// 15 cho giờ, Monday cho ngày trong tuần, v.v.
	p(t.Format("3:04PM"))
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	p(t2)

	// Để biểu diễn theo dạng số, bạn cũng có thể sử dụng
	// chuỗi định dạng chuẩn với các thành phần được trích xuất
	// từ giá trị thời gian.
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	// `Parse` sẽ trả về một lỗi khi đầu vào bị lỗi
	// để giải thích vấn đề khi phân tích thời gian.
	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	p(e)
}
