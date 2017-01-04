package main

import "fmt"

func main() {
        //类型 [n]T 是一个有 n 个类型为 T 的值的数组
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1]) //Hello World
	fmt.Println(a)  //[Hello World]
}
