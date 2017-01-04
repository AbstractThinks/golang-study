package main

import "fmt"

func main() {

//延迟调用的参数会立刻生成，但是在上层函数返回前函数都不会被调用。

        fmt.Println("counting")

        defer fmt.Println("world")
	for i := 0; i < 10; i++ {
                //延迟的函数调用被压入一个栈中。当函数返回时， 会按照后进先出的顺序调用被延迟的函数调用。
                //结果hello 9 8 7 6 5 4 3 2 1 0 hello
		defer fmt.Println(i)
	}

	fmt.Println("done")

	fmt.Println("hello")
}
