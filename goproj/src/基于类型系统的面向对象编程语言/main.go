// http://studygolang.com/articles/2544

package main

import (
	"fmt"
)

/**
 *
 * go语言中，可以给任意类型增加方法（除指针类型）
 *
 *
 */
type Integer int

/**
 *
 * 面向对象编程和面向过程变成对比
 *
 *	func (a Integer) Less(b Integer) bool {  // 面向对象
 *	    return a < b
 *	}
 *
 *	func Integer_Less(a Integer, b Integer) bool {  // 面向过程
 *	    return a < b
 *	}
 *
 *	a.Less(2)  // 面向对象
 *	Integer_Less(a, 2)  // 面向过程
 *
 */

/**
 *
 * 面向过程的写法
 * func Integer_Less(a Integer, b Integer) bool {
 * 		var a Integer = 1
 * 		if Integer_Less(a, 2) {
 * 			fmt.Println(a, "Less 2")
 * 		}
 * }
 *
 *
 * 面向对象的写法
 *
 */
func (a Integer) Less(b Integer) bool {
	return a < b
}

func main() {
	var a Integer = 1
	if a.Less(2) {
		fmt.Println(a, "Less 2")
	}
}
