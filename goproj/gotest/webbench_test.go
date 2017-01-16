/**
 *
 *	压力测试格式
 *	func BenchmarkXXX(b *testing.B){}
 * 	go test 不会默认执行压力测试函数，执行压力测试需带上参数 -test.bench
 * 	例:
 * 		go test -test.bench=".*" 测试全部的压力测试函数
 * 		go test -tets.bench="test_name"
 *
 */

package gotest

import (
	"testing"
)

func Benchmark_Division(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Division(4, 5)
	}
}

// func Benchmark_TimeConsumingFunction(b *testing.B) {
// 	b.StopTimer()

// 	b.StartTimer()
// 	for i :- 0; i < b.N; i++ {
// 		Division(4, 5)
// 	}
// }
