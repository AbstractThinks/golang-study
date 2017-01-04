package main
// 采用别名的方式访问引入的包
/*
import . "fmt" 亦可使用.代替
则在函数中可以省略包名, 不推荐使用
*/
import system "fmt"

// 指定返回类型
//当两个或多个连续的函数命名参数是同一类型，则除了最后一个类型之外，其他都可以省略。
//例 func add(x, y int)  int
func add(x int, y int)  int{
        return x + y
}
// 多值返回
func swap(x, y string) (string, string) {
        return y , x
}
//Go 的返回值可以被命名，并且就像在函数体开头声明的变量那样使用。
func split(sum int) (x, y int) {
        x = sum * 4 / 9
        y = sum - x
        return
}
func main() {
        system.Println(add(44, 13))
        a, b := swap("world", "hello")
        system.Println(a+b)
        system.Println(a, b)
        system.Println(split(17))
}
