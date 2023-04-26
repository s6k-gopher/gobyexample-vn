// Go hỗ trợ các công cụ tích hợp sẵn cho việc tạo nội dung động hoặc hiển thị
// nội dung tùy chỉnh ra cho người dùng với package `text/template`. Một package tương tự
// tên là `html/template` cũng cung cấp các API tương tự nhưng có thêm các
// tính năng bảo mật và khuyến khích được sử dụng để tạo nội dung có định dạng HTML.
package main

import (
	"os"
	"text/template"
)

func main() {

	// Ta có thể tạo ra một template mới và truyền vào nội dung của nó
	// từ một chuỗi.
	// Template là một sự kết hợp của văn bản tĩnh và các "hành động" (action) được gói trong dấu
	// `{{...}}`, thứ thường được sử dụng để chèn nội dung động.
	t1 := template.New("t1")
	t1, err := t1.Parse("Value is {{.}}\n")
	if err != nil {
		panic(err)
	}

	// Ngoài ra, ta có thể sử dụng hàm `template.Must` để
	// panic trong trường hợp `Parse` trả về một lỗi. Điều này
	// đặc biệt hữu ích cho các template được khởi tạo
	// với phạm vi toàn cục (global scope).
	t1 = template.Must(t1.Parse("Value: {{.}}\n"))

	// Bằng việc "thực thi" template, ta tạo ra nội dung của nó
	// với các giá trị cụ thể của các action. Action được đặt trong `{{.}}`
	// được thay thế bằng giá trị được truyền vào `Execute` dưới dạng tham số.
	t1.Execute(os.Stdout, "some text")
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, []string{
		"Go",
		"Rust",
		"C++",
		"C#",
	})

	// Hàm trợ giúp mà ta sẽ sử dụng bên dưới.
	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}

	// Nếu dữ liệu là một struct, ta có thể sử dụng action `{{.FieldName}}` để truy cập
	// các trường của struct đó. Các trường phải được export (được khai báo với chữ cái đầu tiên viết hoa)
	// để có thể truy cập bởi template khi template đang được thực thi.
	t2 := Create("t2", "Name: {{.Name}}\n")

	t2.Execute(os.Stdout, struct {
		Name string
	}{"Jane Doe"})

	// Tương tự với maps; với maps thì không có hạn chế về chữ hoa/chữ thường
	// đối với tên khóa (key names).

	t2.Execute(os.Stdout, map[string]string{
		"Name": "Mickey Mouse",
	})

	// if/else hỗ trợ việc thực thi có điều kiện trong template. Một giá trị được coi là
	// false nếu nó là giá trị mặc định của một kiểu dữ liệu, ví dụ như 0, một chuỗi rỗng,
	// con trỏ nil, v.v.
	// Ví dụ dưới đây minh họa một tính năng khác
	// của template: loại bỏ các khoảng trắng bằng cách sử dụng `-` trong các action.
	t3 := Create("t3",
		"{{if . -}} yes {{else -}} no {{end}}\n")
	t3.Execute(os.Stdout, "not empty")
	t3.Execute(os.Stdout, "")

	// Những khối vòng lặp bằng range (range blocks) cho phép ta lặp qua các
	// slices, arrays, maps hoặc channels. Trong range block, phần `{{.}}` được gán cho
	// giá trị phần tử hiện tại của vòng lặp.
	t4 := Create("t4",
		"Range: {{range .}}{{.}} {{end}}\n")
	t4.Execute(os.Stdout,
		[]string{
			"Go",
			"Rust",
			"C++",
			"C#",
		})
}
