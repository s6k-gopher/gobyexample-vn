// _Interfaces_ là tên tập hợp của các
// phương thức đặc trưng.
package main

import (
	"fmt"
	"math"
)

// Đây là một interface cơ bản cho các hình học.
type geometry interface {
	area() float64
	perim() float64
}

// Ví dụ của chúng ta sẽ thực hiện interface này trên
// kiểu  dữ liệu `rect` and `circle`.
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

// Để hiện thực một interface trong Go, chúng ta chỉ cần
// thực hiện tất cả các phương thức trong giao diện. Ở đây,
// chúng ta hiện thực `geometry` trên kiểu `rect`s.
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// Hiện thực cho kiểu `circle`s.
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// Nếu một biến có kiểu interface, thì chúng ta có thể gọi
// phương thức trong interface. Đây là
// hàm tổng quát `measure` có thể tận dụng tính năng này
// để hoạt động trên bất kì `geometry`.
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	// Cả hai kiểu cấu trúc `circle` và `rect`
	// đều hiện thực geometry interface nên chúng ta có thể
	// sử dụng phiên bản của các cấu trúc này như là
	// đối số `measure`.
	measure(r)
	measure(c)
}
