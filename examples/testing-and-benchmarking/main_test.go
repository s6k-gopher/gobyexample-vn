// Unit test là một phần quan trọng khi viết Go.
// Package `testing` cung cấp những công cụ chúng ta cần
// để viết unit test and lệnh `go test` để chạy test.

// Để trình diễn, đoạn code này nằm trong package `main`,
// nhưng nó có thể nằm trong bất cứ package nào. Code
// testing thường nằm trong cùng package với code nó
// test.

package main

import (
	"fmt"
	"testing"
)

// Chúng ta sẽ test implementation đơn giản của số
// nguyên nhỏ nhất. Thông thường, code chúng ta test
// sẽ nằm trong source file được đặt tên như là
// `intutils.go`, và file test cho nó thường được đặt
// tên `inutils_test.go`.
func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Test dược tạo bằng viết một hàm với tên bắt đầu với
// `Test`.
func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		// `t.Error*` sẽ report test fail nhưng tiếp tục
		// chạy test. `t.Fatal*` sẽ report test fail
		// nhưng sẽ dừng test ngay lập tức.
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}

// Viết test có thể lặp đi lặp lại, bởi thế người ta
// hay dùng *table-driven style*, khi đầu vào test và
// đầu ra mong đợi được list trong một bảng và một
// vòng lặp đơn đi qua chúng và thực hiện logic test.
func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	for _, tt := range tests {
		// t.Run cho phép chạy "subtest", mỗi cái cho
		// một mục của bảng. Chúng được hiển thị riêng
		// rẽ khi thực hiện `go test -v`.
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

// Benchmark test thường nằm trong các file `_test.go`
// và tên bắt đầu với `Benchmark`. `testing` runner
// thực thi mỗi hàm benchmark một vài lần, tăng `bn.N`
// trên mỗi lần chạy cho đến khi đo được chính xác.
func BenchmarkIntMin(b *testing.B) {
	// Thường benchmark chạy một hàm mà chúng ta
	// benchmark trong vòng lặp `b.N` lần.
	for i := 0; i < b.N; i++ {
		IntMin(1, 2)
	}
}
