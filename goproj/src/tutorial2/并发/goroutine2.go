package main

import (
        "fmt"
)

func sum(a []int, c chan int) {
        sum := 0
        for _, v := range a {
                sum += v
        }
        c <- sum      //向通道c写入数据值sum
}
func main() {
                fmt.Println("hello world")
                a := []int{1, 3, 5, 2, 2}
                c := make(chan int)
                go sum(a[:], c)
                x := <- c
                fmt.Println(x)

                c2 := make(chan int,2)     //首先声明定义一个缓存大小为2的特定类型为int的通道
                c2 <- 10
                c2 <- 20
                fmt.Println(<-c2)
                fmt.Println(<-c2)

                var c3 chan int
                c3 <- 10
                fmt.Println(<-c3)
}
