/*
基本的 for 循环包含三个由分号分开的组成部分：

初始化语句：在第一次循环执行前被执行
循环条件表达式：每轮迭代开始前被求值
后置语句：每轮迭代后被执行


*/

package main

import "fmt"

func main() {
	sum_for := 0

	for i := 0; i < 10; i++ {
		sum_for += i
	}

        sum_while := 1
	for sum_while < 1000 {
		sum_while += sum_while
	}

	fmt.Println(sum_for)
	fmt.Println(sum_while)
}
