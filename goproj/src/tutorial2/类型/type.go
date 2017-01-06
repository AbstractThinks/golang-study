package main

//别名访问fmt
import (
        system "fmt"
)

// 多变量定义
var  x, y, z int

//x = "Hi! jesse"
// 多变量赋值，可省略类型，由值类型推导变量类型
var s, n = "abc", 123
var (
        a int
        b float32
        c string
)
var str = "Hi! Jesse"
//数组声明   [size]类型{value}
var array1 = [5]int{1, 2, 3}    //创建数组(声明长度)
var array2 = [...]int{6, 7, 8}    //创建数组(不声明长度)
var array3 = [5]int{}   //创建数组(声明长度)，但不赋值

//常量在声明时必须赋值
const c_x, c_y int = 1, 2
const c_s = "hello world"
const (
        c_a, c_b = 10, 100
        c_c bool = false
        c_c1    //如果不提供类型和初始化值那么视作与上一常量相同
)

const (
        Sunday = iota    //0
        Monday              //1
        Tuesday             //2
        Wednesday       //3
        Thursday           //4
        Friday                 //5
        Saturday            //6
)

const (
        A = iota           //0
        B                      //1
        C = "c"             //c
        D                      //c    与上一行相同
        E = iota            //4    显示恢复
        F                       //5
)

func main() {
        system.Println("hello world")

        // i := 0 局部定义的变量若未使用会引起编译报错
        //可使用 _ := 0规避
        const c_x = "xxx" //局部定义的常量未使用不会引起编译报错

        system.Println(n, s)   //123 abc    函数外变量
        n, s := "hello, world!", 0x1234  //这种赋值方式只能在函数体内
        system.Println(n, s)   //hello, world! 4660   函数内变量

        system.Println(x)  //未赋值的int类型默认值为0
        system.Println(b)  //未赋值的float类型默认值为0
        system.Println(c)  //未赋值的fstring类型默认值为""
        system.Println(array3)   //未赋值的数组类型默认值为[0, 0, 0, 0, 0]

        _, func_test_str := test()
        system.Println(func_test_str)   //abc

}
//函数定义
// func 函数名() 返回值类型 {}
// func 函数名() {}
func test() (int, string) {
        return 1, "abc"
}
