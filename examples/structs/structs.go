// _Struct_ của Go là tập hợp các trường dữ liệu được khai báo kiểu.
// Chúng được sử dụng để nhóm các dữ liệu lại với nhau
// để tạo thành các bản ghi.
package main

import "fmt"

// Kiểu struct `person` này có các trường `name` và `age`.
type person struct {
	name string
	age  int
}

// `newPerson` tạo một struct `person` mới với tên được truyền vào.
func newPerson(name string) *person {
	// Bạn có thể trả về một cách an toàn một con trỏ của biến cục bộ
	// vì biến cục bộ sẽ tồn tại trong phạm vi của hàm.
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {

	// Cú pháp này tạo một struct mới.
	fmt.Println(person{"Bob", 20})

	// Bạn có thể đặt tên cho các trường khi khởi tạo một struct.
	fmt.Println(person{name: "Alice", age: 30})

	// Các trường bị bỏ qua sẽ được gán giá trị mặc định.
	fmt.Println(person{name: "Fred"})

	// Kí tự `&` trước tên struct sẽ trả về một con trỏ đến struct đó.
	fmt.Println(&person{name: "Ann", age: 40})

	// Một cách viết phổ biến là đóng gói việc tạo một struct mới trong các hàm khởi tạo.
	fmt.Println(newPerson("Jon"))

	// Truy cập vào các trường của struct bằng dấu chấm.
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// Bạn có thể sử dụng dấu chấm với con trỏ struct - các
	// con trỏ sẽ được tự động giải tham chiếu.
	sp := &s
	fmt.Println(sp.age)

	// Các struct có thể thay đổi được.
	sp.age = 51
	fmt.Println(sp.age)
}
