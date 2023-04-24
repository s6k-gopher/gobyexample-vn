// Đôi lúc, chúng ta cần tạo ra các tiến trình
// không phải là code Go.
package main

import (
	"fmt"
	"io"
	"os/exec"
)

func main() {

	// Chúng ta sẽ bắt đầu với một lệnh đơn giản không
	// có tham số và chỉ in ra một thông báo đơn giản.
	// Hàm `exec.Command` sẽ tạo ra một đối tượng
	// để thực thi tiến trình bên ngoài giúp chúng ta.
	dateCmd := exec.Command("date")

	// `Output` sẽ chạy lệnh, đợi cho đến khi nó hoàn
	// thành và thu thập kết quả của nó. Nếu không có lỗi,
	// `dateOut` sẽ chứa các thông tin ngày giờ.
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	// `Output` và các phương thức khác của `Command` sẽ
	// trả về `*exec.Error` nếu có lỗi khi thực thi lệnh
	// (vd: đường dẫn sai), và `*exec.ExitError` nếu lệnh
	// chạy nhưng trả về một mã lỗi khác 0.
	_, err = exec.Command("date", "-x").Output()
	if err != nil {
		switch e := err.(type) {
		case *exec.Error:
			fmt.Println("failed executing:", err)
		case *exec.ExitError:
			fmt.Println("command exit rc =", e.ExitCode())
		default:
			panic(err)
		}
	}

	// Tiếp theo chúng ta sẽ xem một ví dụ phức tạp hơn
	// thực thi lệnh `grep` để tìm các dòng chứa từ
	// "hello"
	grepCmd := exec.Command("grep", "hello")

	// Ví dụ dưới đây chúng ta sẽ dùng StdinPipe và
	// StdoutPipe để truyền dữ liệu vào và lấy dữ liệu
	// Truyền dữ liệu vào grepIn để ghi vào tiến trình
	// grep, và lấy dữ liệu từ grepOut để đọc kết quả
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := io.ReadAll(grepOut)
	grepCmd.Wait()

	// Chúng ta bỏ qua kiểm tra lỗi trong ví dụ dưới đây,
	// nhưng bạn có thể sử dụng cú pháp `if err != nil`
	// cho tất cả chúng. Chúng ta cũng chỉ thu thập kết quả
	// `StdoutPipe`, nhưng bạn có thể thu thập `StderrPipe`
	// theo cùng một cách.
	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	// Chú ý khi thực thi lệnh chúng ta cần cung cấp
	// một mảng các tham số, thay vì chỉ cần truyền vào
	// một chuỗi lệnh. Nếu bạn muốn thực thi một đầy đủ
	// lệnh, bạn có thể sử dụng `-c` của `bash`
	// để thực thi.
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}
