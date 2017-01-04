package main

import "fmt"

func main() {

	i, j := 42, 2701

	p := &i         // point to i  把变量i的地址赋值给变量p
	fmt.Println(*p) // read i through the pointer 打印以变量p的值为地址的存储内容   42
	*p = 21         // set i through the pointer 将值赋值给p的值为地址的内存
	fmt.Println(i)  // see the new value of i   21

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j   73
}
