// Bắt đầu từ phiên bản 1.18, Go đã hỗ trợ thêm cho
// _generics_, also known as _type parameters_
// (tham số kiểu dữ liêu).

package main

import "fmt"

// Một ví dụ về hàm generic, `MapKeys` nhận
// một map với bất kì kiểu nào và trả về một slices các
// key của nó. Hàm này có hai type parameters - `K` và
// `V`; `K` có ràng buộc `comparable` (so sánh),
// nghĩa là chúng ta có thể so sánh các giá trị của
// kiểu này với toán tử `==` và`!=` . Điều này là
//
//	để sử dụng key trong map của Go.
//
// `V` có ràng buộc `any` (bất kì), nghĩa là nó không
// bị hạn chết bởi bất kì đièu gì (`any` là một
// định danh khác cho `interface{}`).
func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

// Một ví dụ về type (kiểu) generics là `List`, `List`
// là một danh sách liên kết đơn với các giá trị của
// bất kỳ kiểu dữ liệu nào.
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

// Chúng ta có thể định nghĩa các phương thức trên
// generic types giống như các kiểu thông thường,
// nhưng chúng ta phải giữ các type
// parameters đúng vị trí. Kiểu này là `List[T]`,
// không phải `List`.
func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	var m = map[int]string{1: "2", 2: "4", 4: "8"}

	// Khi gọi các hàm generic, chúng ta thường
	//có thể dựa vào _type inference_. Lưu ý rằng
	// chúng ta không cần phải  xác định kiểu cho
	// `K` and `V` khi gọi `MapKeys` - trình
	// biên dịch sẽ tự động suy ra chúng
	fmt.Println("keys:", MapKeys(m))

	// ... tuy nhiên, chúng ta cũng có thể xác định
	// chúng một cách rõ ràng.

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.GetAll())
}
