// [_Command-line flags_ (cờ dòng lệnh)](https://en.wikipedia.org/wiki/Command-line_interface#Command-line_option)
// là cách thông thường để chỉ định các tuỳ chọn cho các
// chương chình dòng lệnh. Ví dụ, trong lệnh `wc -l`, `-l`
// là một cờ dòng lệnh.

package main

// Go cung cấp một package (gói) `flag` hỗ trợ cơ bản cho phân tích
// cú pháp các command-line flag. Chúng ta sẽ sử dụng package để
// hiện thực chương trình dòng lệnh ví dụ của mình.
import (
	"flag"
	"fmt"
)

func main() {

	// Các khai báo cờ cơ bản có sẵn cho các tùy chọn chuỗi
	// số nguyên, và luận lí. Ở đây chúng ta khai báo một
	// cờ chuỗi `word` với giá trị mặc định `"foo"` và một
	// mô tả ngắn. Hàm `flag.String` nãy trả về một
	// con trỏ chuỗi (không phải là một giá trị chuỗi);
	// chúng ta sẽ thấy cách sử dụng con trỏ này bên dưới.
	wordPtr := flag.String("word", "foo", "a string")

	// Đây là các khai báo cờ `numb` and `fork`, sử dụng một
	// cách tiếp cận tương tự với cờ `word`.
	numbPtr := flag.Int("numb", 42, "an int")
	forkPtr := flag.Bool("fork", false, "a bool")

	// Cũng có thể khai báo một tùy chọn sử dụng một biến
	// đã được khai báo ở một nơi khác trong chương trình.
	// Lưu ý rằng chúng ta cần truyền vào một con trỏ tới
	// hàm khai báo cờ.
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	// Sau khi tất cả các cờ đã được khai báo, gọi `flag.Parse()`
	// để thực thi phân tích cú pháp dòng lệnh.
	flag.Parse()

	// Ở đây chúng ta chỉ đơn giản là in ra các tùy chọn
	// đã được phân tích cú pháp và bất kỳ đối số vị trí
	// dư thừa nào. Chú ý rằng chúng ta cần huỷ tham chiếu
	// các con trỏ để lấy giá trị tùy chọn thực sự, ví dụ `*wordPtr`
	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *forkPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}
