// Go hỗ trợ rất tốt cho việc định dạng chuỗi trong
// hàm `printf`. Dưới đây là một số ví dụ về các
// tác vụ định dạng chuỗi phổ biến.

package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {

	// Go cung cấp một số "động từ" được thiết kế để
	// định dạng các giá trị chung trong Go. Ví dụ, đoạn dưới đây
	// in ra một đối tượng của struct `point` được định nghĩa ở trên.
	p := point{1, 2}
	fmt.Printf("struct1: %v\n", p)

	// Nếu giá trị là một struct, biến thể `%+v` sẽ
	// in ra thêm tên các trường của struct.
	fmt.Printf("struct2: %+v\n", p)

	// Biến thể `%#v` sẽ in ra một cú pháp Go đại diện
	// cho giá trị, ví dụ, một đoạn mã nguồn để
	// tạo ra giá trị đó.
	fmt.Printf("struct3: %#v\n", p)

	// Để in ra kiểu của một giá trị, sử dụng `%T`.
	fmt.Printf("type: %T\n", p)

	// Định dạng giá trị boolean thì làm như sau.
	fmt.Printf("bool: %t\n", true)

	// Có rất nhiều tùy chọn để định dạng số nguyên.
	// Sử dụng `%d` để định dạng chuẩn, dùng hệ cơ số 10.
	fmt.Printf("int: %d\n", 123)

	// Dưới đây sẽ in ra dạng nhị phân của số nguyên.
	fmt.Printf("bin: %b\n", 14)

	// Dưới đây sẽ in ra kí tự tương ứng với
	// giá trị số nguyên được truyền vào.
	fmt.Printf("char: %c\n", 33)

	// `%x` hỗ trợ mã hóa hex.
	fmt.Printf("hex: %x\n", 456)

	// Có rất nhiều tùy chọn để định dạng số thực.
	// Để định dạng số thực theo cách cơ bản, sử dụng `%f`.
	fmt.Printf("float1: %f\n", 78.9)

	// `%e` và `%E` định dạng số thực theo (một số
	// phiên bản khác nhau của) ký hiệu khoa học.
	fmt.Printf("float2: %e\n", 123400000.0)
	fmt.Printf("float3: %E\n", 123400000.0)

	// Để in ra một chuỗi một cách cơ bản, sử dụng `%s`.
	fmt.Printf("str1: %s\n", "\"string\"")

	// Để đặt chuỗi trong dấu ngoặc kép, sử dụng `%q`.
	fmt.Printf("str2: %q\n", "\"string\"")

	// Đối với các giá trị nguyên ở phía trên, `%x` sẽ
	// in ra chuỗi ở cơ số thập lục phân, với hai kí tự
	// đầu ra cho mỗi byte của giá trị đầu vào.
	fmt.Printf("str3: %x\n", "hex this")

	// Để in ra một giá trị con trỏ đại diện cho biến đó, sử dụng `%p`.
	fmt.Printf("pointer: %p\n", &p)

	// Khi định dạng số, bạn sẽ thường muốn kiểm soát
	// độ rộng và số các chữ số thập phân của kết quả.
	// Để chỉ định độ rộng của một số nguyên, sử dụng một
	// số sau dấu `%` trong biểu thức. Mặc định, kết quả
	// sẽ được căn phải và đệm bằng khoảng trắng.
	fmt.Printf("width1: |%6d|%6d|\n", 12, 345)

	// Bạn cũng có thể chỉ định độ rộng của số thực được in ra,
	// tuy nhiên thường bạn sẽ muốn giới hạn số chữ số thập phân
	// cùng lúc với độ rộng với chung 1 cú pháp duy nhất
	fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)

	// Để căn trái, sử dụng dấu `-`.
	fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)

	// Bạn cũng có thể muốn kiểm soát độ rộng khi định dạng
	// chuỗi, đặc biệt là để đảm bảo rằng chúng căn chỉnh
	// dưới dạng bảng khi in ra. Để căn phải với một độ rộng
	// một cách cơ bản ta làm như sau.
	fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")

	// Để căn trái, sử dụng dấu `-` tương tự như với số.
	fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")

	// Đến giờ chúng ta đã được xem về `Printf`, hàm in
	// chuỗi đã được định dạng ra `os.Stdout`. `Sprintf` định
	// dạng và trả về một chuỗi mà không in ra bất cứ đâu.
	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s)

	// Bạn có thể định dạng và in ra `io.Writers` khác ngoài
	// `os.Stdout` bằng cách sử dụng `Fprintf`.
	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}
