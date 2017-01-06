package main

import (
        "fmt"
        "runtime"
)


func say(s string) {
                for i := 0; i < 5; i++ {
                        runtime.Gosched()   //让cpu把时间片让给别人，下次某个时间继续恢复执行该goroutine
                        fmt.Println(s)
                }
}

// goroutines是Go实行并发程序的一种方式
//语法是：go function_name()
func main() {
        go say("world")
        say("hello")

}
