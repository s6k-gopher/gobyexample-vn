// Trong suốt quá trình thực thi chương trình, ta thường muốn tạo
// dữ liệu tạm mà không cần thiết sau khi chương trình kết thúc.
// *Tập tin và thư mục tạm thời* rất hữu ích cho mục đích này
// vì về lâu dài chúng không khiến cho hệ thống tập tin
// chứa nhiều tập tin rác.

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// Cách dễ nhất để tạo một tập tin tạm thời là bằng cách
	// gọi `os.CreateTemp`. Nó tạo một tập tin *và*
	// mở nó để đọc và ghi. Chúng ta dùng `""` làm tham số đầu tiên,
	// vì vậy `os.CreateTemp` sẽ tạo tập tin trong vị trí mặc định
	// của hệ điều hành.
	f, err := os.CreateTemp("", "sample")
	check(err)

	// Hiển thị tên của tập tin tạm thời. Trên các hệ điều hành
	// dựa trên Unix, thư mục có thể là `/tmp`. Tên tập tin bắt đầu
	// bằng tiền tố được truyền vào làm tham số thứ hai của `os.CreateTemp`
	// và các tham số còn lại được chọn tự động để đảm bảo rằng
	// việc thực thi đồng thời sẽ luôn tạo ra các tên tập tin khác nhau.
	fmt.Println("Temp file name:", f.Name())

	// Sau khi chúng ta hoàn thành, tập tin tạm thời sẽ được xóa.
	// Hệ điều hành có thể sẽ tự động xóa các tập tin tạm thời sau
	// một khoảng thời gian nhất định, nhưng tốt hơn là chúng ta
	// nên làm điều này một cách tường minh.
	defer os.Remove(f.Name())

	// Ta có thể ghi một số dữ liệu vào tập tin.
	_, err = f.Write([]byte{1, 2, 3, 4})
	check(err)

	// Nếu chúng ta có ý định ghi nhiều tập tin tạm thời,
	// ta nên ưu tiên tạo một *thư mục tạm thời*.
	// Các tham số của `os.MkdirTemp` cũng tương tự như `os.CreateTemp`,
	// nhưng nó trả về một tên thư mục thay vì một tập tin đang mở.
	dname, err := os.MkdirTemp("", "sampledir")
	check(err)
	fmt.Println("Temp dir name:", dname)

	defer os.RemoveAll(dname)

	// Giờ thì ta có thể tạo tên tập tin tạm thời bằng cách
	// thêm tiền tố vào tên thư mục tạm thời.
	fname := filepath.Join(dname, "file1")
	err = os.WriteFile(fname, []byte{1, 2}, 0666)
	check(err)
}
