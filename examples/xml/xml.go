// Go hỗ trợ sẵn cho việc mã hóa và giải mã XML và
// các định dạng tương tự XML với package `encoding.xml`.
package main

import (
	"encoding/xml"
	"fmt"
)

// Kiểu Plant sẽ được ánh xạ sang định dạng XML. Tương tự như
// các ví dụ về JSON, các tag của các trường sẽ chứa các chỉ
// dẫn cho encoder và decoder. Ở đây ta sẽ sử dụng một số tính
// năng đặc biệt của package XML: trường `XMLName` sẽ quy định
// tên của phần tử XML đại diện cho struct này;
// `id,attr` có nghĩa là trường `Id` là một _thuộc tính_ XML
// chứ không phải là một phần tử lồng nhau.
type Plant struct {
	XMLName xml.Name `xml:"plant"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v",
		p.Id, p.Name, p.Origin)
}

func main() {
	coffee := &Plant{Id: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}

	// Tạo ra biểu diễn XML của kiểu Plant; sử dụng
	// `MarshalIndent` để tạo ra một đầu ra dễ đọc hơn.
	out, _ := xml.MarshalIndent(coffee, " ", "  ")
	fmt.Println(string(out))

	// Để thêm một tiêu đề XML chung vào đầu ra, ta có thể
	// thêm nó vào một cách trực tiếp.
	fmt.Println(xml.Header + string(out))

	// Sử dụng `Unmarshal` để phân tích cú pháp một luồng bytes
	// với đinh dạng XML thành một cấu trúc dữ liệu. Nếu XML
	// bị lỗi cú pháp hoặc không thể ánh xạ vào kiểu Plant,
	// một mô tả về lỗi sẽ được trả về.
	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p)

	tomato := &Plant{Id: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}

	// Tag `parent>child>plant` cho biết encoder sẽ lồng
	// tất cả các `plant` vào trong `<parent><child>...`
	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants  []*Plant `xml:"parent>child>plant"`
	}

	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}

	out, _ = xml.MarshalIndent(nesting, " ", "  ")
	fmt.Println(string(out))
}
