// Đọc và ghi tập tin là các nhiệm vụ cơ bản cần thiết
// cho nhiều chương trình Go. Đầu tiên, chúng ta sẽ xem
// một số ví dụ về đọc tập tin.

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Đọc tập tin yêu cầu kiểm tra hầu hết các lần gọi hàm
// để tìm lỗi. Hàm hỗ trợ này sẽ giúp việc tìm lỗi của
// chúng ta dễ dàng hơn.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// Có lẽ nhiệm vụ cơ bản nhất của đọc tập tin là
	// đọc toàn bộ nội dung tập tin.
	dat, err := os.ReadFile("/tmp/dat")
	check(err)
	fmt.Print(string(dat))

	// Bạn thường muốn kiểm soát hơn về cách và những phần
	// của tập tin được đọc. Về những nhiệm vụ này, trước hết
	// bạn cần `mở` tập tin để lấy giá trị `os.File`.
	f, err := os.Open("/tmp/dat")
	check(err)

	// Đọc một số byte từ đầu tập tin. Cho phép đọc tối đa
	//  5 byte nhưng cũng ghi nhận số byte thực sự đã được đọc.
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	// Bạn cũng có thể sử dụng `Seek` để tìm vị trí trong tập tin
	// và `đọc` từ đó.
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	// Package (gói) `io`  cung cấp một số hàm có thể hữu ích
	// cho việc đọc tập tin. Ví dụ, các tác vụ đọc như trên
	// có thể được hiện thực một cách đáng tin cậy hơn
	// với hàm `ReadAtLeast`.
	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// Không có chức năng quay lại tích hợp sẵn,
	// nhưng hàm`Seek(0, 0)` có thể thực hiện điều này.
	_, err = f.Seek(0, 0)
	check(err)

	// Package  `bufio` triển khai một bộ đệm đọc có thể
	// hữu ích vừa vì tính hiệu quả với nhiều lần đọc nhỏ,
	// vừa vì các phương thức đọc bổ sung mà nó cung cấp
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	// Đóng tập tin khi bạn hoàn thành (thông thường
	// điều này sẽ được lên kế hoạch ngay sau khi `mở`
	// tập tin với `defer`).
	f.Close()
}
