// Go hỗ trợ _embedding_ của structs and interfaces
// để thể hiện liền mạch hơn __thành phần_ của kiểu.
// Điều này không nên nhầm với `//go:embed`, đây là
// lệnh Go được giới thiệu trực tiếp trong bản Go 1.16+
// để nhúng các tệp và thư mục vào tệp nhị phân của ứng dụng.

package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// Một `container` được _embeds_ (nhúng) a `base`. Một embedding giống như
// một trường không có tên.
type container struct {
	base
	str string
}

func main() {

	// Khi tạo structs với literals,chúng ta phải
	// khởi tạo embedding rõ ràng; ở dây
	// kiểu embedded struct đóng vai trò là tên trường.
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	//Chúng ta có thể truy cập trực tiếp vào các trường cơ sở của `co`,
	// ví dụ `co.num`.
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	// Ngoài ra, chúng ta có thể định nghĩa đường dẫn đầy đủ
	// bằng cách sử dụng tên của kiểu embedded struct.
	fmt.Println("also num:", co.base.num)

	// Vì `container` embeds `base`, các phương thức
	// của `base` cũng trở thành phương thức của `container`.
	// Ở đây, chúng ta gọi một phương thức được nhúng từ `base`
	// trực tiếp từ `co`.
	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	// Embedding structs với các phương thức có thể được sử dụng
	// để cấp cho hiện thực của interface dựa trên structs khác.
	// Điều này cho chúng ta thấy một `container` hiện thực
	// `describer` interface bời vì nó embeds `base`.
	var d describer = co
	fmt.Println("describer:", d.describe())
}
