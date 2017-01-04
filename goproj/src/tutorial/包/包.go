//每个 Go 程序都是由包组成的。
//程序运行的入口是包 main 。
package main

//导入了包 "fmt" 和 "math/rand" 。
/*
注意：若导入了未使用的包编译会报错
*/
import (
        "fmt"
        "math/rand"
)
//在 Go 中，首字母大写的名称是被导出的。
func main() {
        fmt.Println("My favorite number is ",  rand.Intn(10))
}
