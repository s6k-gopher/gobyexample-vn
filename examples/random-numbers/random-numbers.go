// Go's `math/rand` package (gói) cung cấp tạo
// [số giả ngẫu nhiên](https://en.wikipedia.org/wiki/Pseudorandom_number_generator).

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// Ví dụ, `rand.Intn` trả về một `int` (số nguyên) ngẫu nhiên n,
	// `0 <= n < 100`.
	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()

	// `rand.Float64` trả về một `float64` (số thực) ngẫu nhiên `f`,
	// `0.0 <= f < 1.0`.
	fmt.Println(rand.Float64())

	// Điều này có thể được sử dụng để tạo ra các số thực
	// ngẫu nhiên trong các phạm vi khác, ví dụ  `5.0 <= f' < 10.0`.
	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	// Bộ tạo số mặc định là xác định được, vì thế mặc định nó sẽ
	// tạo ra cùng một chuỗi số mỗi lần. Để tạo ra các chuỗi khác nhau,
	// hãy cung cấp cho nó một seed thay đổi.  Lưu ý rằng điều này
	// không an toàn để sử dụng cho các số ngẫu nhiên mà bạn dự định
	// là bí mật; hãy sử dụng `crypto/rand` cho những số đó.
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	// Kết quả gọi `rand.Rand` giống như nhũng hàm
	// trong `rand` package.
	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100))
	fmt.Println()

	// Nếu bạn khởi tạo một nguồn với cùng một số, nó
	// sẽ tạo ra cùng một chuỗi số ngẫu nhiên.
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
	fmt.Println()
	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), ",")
	fmt.Print(r3.Intn(100))
}
