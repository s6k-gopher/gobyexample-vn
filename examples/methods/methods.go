// Go hỗ trợ _phương thức_ được khai báo trên các kiểu `struct`.
package main

import "fmt"

type rect struct {
	width, height int
}

// Phương thức `area` có  _kiểu tham số đầu vào (receiver type)_ là `*rect`.
func (r *rect) area() int {
	return r.width * r.height
}

// Phương thức có thể được khai báo với receiver type là con trỏ
// hoặc giá trị. Dưới đây là một ví dụ về receiver type là giá trị.
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	// Ở đây ta gọi 2 phương thức đã được định nghĩa cho kiểu `struct` của ta.
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	// Go tự động chuyển đổi giữa giá trị và con trỏ khi
	// phương thức được gọi. Bạn có thể muốn sử dụng receiver
	// type là con trỏ để tránh việc sao chép khi phương thức
	// được gọi đến hoặc để cho phép phương thức thay đổi giá
	// trị của struct.
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
