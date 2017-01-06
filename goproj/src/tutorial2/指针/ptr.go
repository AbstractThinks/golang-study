package main

import (
        "fmt"
)

type data struct{
        a int
        name string
        nickaname = "test"
}

func main() {
        fmt.Println("hello world")

        var d = data{1234, "jesse"}  //初始化必须实现所有字段
        var p *data
        p = &d

        fmt.Println(p)
        fmt.Println(p.a)     //直接用"."访问目标成员
}
