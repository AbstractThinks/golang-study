
/* 定义接口 */
/*
type interface_name interface {
   method_name1 [return_type]
   method_name2 [return_type]
   method_name3 [return_type]
   ...
   method_namen [return_type]
}

*/


/* 定义结构体 */
/*
type struct_name struct {
  ...
}
*/


/* 实现接口方法 */
/*
func (struct_name_variable struct_name) method_name1() [return_type] {
  ...
}
func (struct_name_variable struct_name) method_namen() [return_type] {
  ...
}
*/



package main

import (
    "fmt"
)

type Phone interface {
    call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
    fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
    fmt.Println("I am iPhone, I can call you!")
}

func main() {
    var phone Phone

    phone = new(NokiaPhone)
    fmt.Println(phone)     //地址
    phone.call()

    phone = new(IPhone)
    phone.call()

}
