// Một chuỗi trong Go là một read-only slice của các byte. Ngôn ngữ Go
// và thư viện chuẩn trong Go xử lý chuỗi một cách rất độc đáo - coi nó
// như một container chứa các ký tự được mã hóa theo [UTF-8](https://en.wikipedia.org/wiki/UTF-8).
// Trong các ngôn ngữ khác, chuỗi được tạo thành từ các "ký tự".
// Trong Go, khái niệm ký tự được gọi là `rune` - nó là một số nguyên
// biểu diễn một code point của Unicode.
// [Bài viết trên blog Go](https://go.dev/blog/strings) là một bài viết
// khá hay để giới thiệu về chủ đề này.
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	// `s` là một `string` được gán một giá trị literal
	// biểu diễn từ "hello" trong tiếng Thái Lan. Chuỗi
	// literal trong Go được mã hóa theo UTF-8.
	const s = "สวัสดี"

	// Bởi vì chuỗi tương đương với `[]byte`, nên đoạn code
	// dưới đây sẽ in ra độ dài của các bytes chứa trong chuỗi `s`
	fmt.Println("Len:", len(s))

	// Truy cập vào một chuỗi sẽ cho ra các giá trị byte tại
	// từng chỉ số. Vòng lặp dưới đây sẽ in ra các giá trị hex
	// của các byte mà tạo thành các code point trong chuỗi `s`.
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	// Để đếm xem có bao nhiêu _runes_ bên trong một chuỗi, chúng
	// ta có thể sử dụng package `utf8`. Lưu ý rằng thời gian chạy
	// của `RuneCountInString` phụ thuộc vào kích thước của chuỗi,
	// bởi vì nó phải giải mã từng UTF-8 rune tuần tự. Một số ký tự
	// trong tiếng Thái Lan được biểu diễn bởi nhiều điểm mã UTF-8,
	// cho nên kết quả của việc đếm này có thể gây bất ngờ.
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	// Vòng lặp `range` xử lý chuỗi một cách rất đặc biệt và
	// giải mã từng `rune` cùng với vị trí của nó trong chuỗi.
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	// Ta có thể đạt được kết quả tương tự bằng cách sử dụng
	// hàm `utf8.DecodeRuneInString`.
	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		// Đoạn code dưới đây minh họa việc truyền một giá trị `rune`
		// vào một hàm.
		examineRune(runeValue)
	}
}

func examineRune(r rune) {

	// Các giá trị được chứa trong dấu nháy đơn là _rune literals_.
	// Chúng ta có thể so sánh trực tiếp một giá trị `rune` với một rune literal
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
