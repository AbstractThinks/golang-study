package main

import (
        "fmt"
        "time"
)
func spinner(delay time.Duration) {
        for {
                for _, r := range `-\|/` {
                        fmt.Printf("\r%c", r)
                        time.Sleep(delay)
                }
        }
}
func fib(x int) int {
        if x < 2 {
                return x
        }
        return fib(x - 1) + fib(x -2)
}
func main() {
        fmt.Println("hello world ")
        /*
        main goroutine将计算菲波那契数列的第45个元素值。
        由于计算函数使用低效的递归，所以会运行相当长时间，
        在此期间我们想让用户看到一个可见的标识来表明程序依然在正常运行，
        所以来做一个动画的小图标：
        */
        go spinner(100 * time.Millisecond)
        const n = 45
        fibN := fib(n)
        fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
        /*
        主函数返回时，所有的goroutine都会被直接打断，程序退出。
        除了从主函数退出或者直接终止程序之外，没有其它的编程方法能够让一个goroutine来打断另一个的执行，
        但是通过goroutine之间的通信来让一个goroutine请求其它的goroutine，
        并被请求的goroutine自行结束执行来实现打断另一个的执行这个目的。
        */
}
