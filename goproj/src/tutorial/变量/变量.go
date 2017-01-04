package main

import "fmt"

var c, python, java bool
var a, b int = 1, 2

func main () {
        var i int

        // := 结构不能使用在函数外。
        k := 3

        // 局部变量, 退出后回收内存
        c, python, java = true, false,true
        fmt.Println(i, c, python, java)
        fmt.Println(a, b)
        fmt.Println(k)
}
