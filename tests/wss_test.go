package tests

import (
	"net/http"
	"testing"
)

// BenchmarkGetRequest-8 300000 4351 ns/op 32 B/op 2 allocs/op
// -8表示8个CPU线程执行；300000表示总共执行了30万次；4531ns/op，表示每次执行耗时4531纳秒；32B/op表示每次执行分配了32字节内存；2 allocs/op表示每次执行分配了2次对象。
func BenchmarkGetRequest(b *testing.B) {
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			get()
		}
	})
}

func get() {
	_, err := http.Get("http://127.0.0.1:9080/report")
	if err != nil {
		return
	}
}
