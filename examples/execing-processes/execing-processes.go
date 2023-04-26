// Ở ví dụ trước chúng ta đã xem qua [khởi tạo các
// process bên ngoài](spawning-processes). Chúng ta làm
// như thế này khi cần một process bên ngoài có thể
// truy cập từ một process Go đang chạy. Đôi khi chúng ta
// chỉ muốn hoàn toàn thay thế process Go hiện tại bằng
// một process khác (có thể không phải process của Go).
// Để làm điều này chúng ta sẽ sử dụng hàm
// <a href="https://en.wikipedia.org/wiki/Exec_(operating_system)"><code>exec</code></a>
// của Go.

package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {

	// Ở trong ví dụ này chúng ta sẽ thực thi lệnh `ls`. Go
	// cần một đường dẫn tuyệt đối đến file binary mà chúng ta
	// muốn thực thi, vì vậy chúng ta sẽ sử dụng hàm `exec.LookPath`
	// để tìm nó (có thể là `/bin/ls`).
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	// `Exec` nhận đối số dưới dạng một slice (không phải
	// một string). Chúng ta sẽ truyền một số đối số thông
	// dụng cho `ls`. Lưu ý rằng đối số đầu tiên phải là tên
	// của chương trình mà ta muốn thực thi.
	args := []string{"ls", "-a", "-l", "-h"}

	// `Exec` cũng cần một tập hợp các [biến môi trường](environment-variables)
	// để sử dụng. Ở đây chúng ta chỉ cung cấp môi trường
	// hiện tại.
	env := os.Environ()

	// Đây mới chính là lệnh `syscall.Exec` thực sự. Nếu lệnh này
	// thành công, thì quá trình thực thi của chương trình sẽ
	// kết thúc ở đây và được thay thế bởi process `/bin/ls -a -l -h`.
	// Nếu có lỗi xảy ra, chúng ta sẽ nhận được một giá trị trả về.
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
