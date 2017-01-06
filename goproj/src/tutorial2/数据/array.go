package main

import (
        "fmt"
)

func main() {
        a := [3] int{1, 2}                          //未初始化元素值为0
        b := [...] int{1, 2, 3, 5}               // 通过初始化值确定数组长度
        c := [5] int{2:100, 4:200}         //使用索引号初始化元素
        d := [...] struct {
                name string
                age uint8
        }{
                {"user1", 10},
                {"user2", 20},          //逗号不能省略掉，会引起编译错误
        }
        fmt.Println(a)          //[1 2 0]
        fmt.Println(b)          //[1 2 3 5]
        fmt.Println(c)          //[0 0 100 0 200]
        fmt.Println(d)          //[{user1 10} {user2 20}]


        // 切片slice
        s1 := []int{0, 1, 2, 3, 8:100}
        s2 := make([]int, 6, 8)
        s3 := make([]int, 6)   //省略cap,相当于cap=len

        fmt.Println(s1, len(s1), cap(s1))       //[0 1 2 3 0 0 0 0 100] 9 9
        fmt.Println(s2, len(s2), cap(s2))       //[0 0 0 0 0 0] 6 8
        fmt.Println(s3, len(s3), cap(s3))       //[0 0 0 0 0 0] 6 6


        m := map[string] int {"a" : 1}

        m1 := map[int] struct {
                name string
                age int
        }{
                1: {"user1", 10},
                2: {"user2", 20},
        }

        fmt.Println(m)
        fmt.Println(m1)
}
