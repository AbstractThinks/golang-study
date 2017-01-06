package main


import (
        "fmt"
)

type Data struct {
        x int
}

func (self Data) ValueTest() {
        fmt.Println(self)
}

func (self *Data) PointerTest() {
        fmt.Println(self)
        fmt.Printf("addr is : %p\n", self)     //addr is : 0x1170c06c
}


func main() {
        fmt.Println("hello world")

        test := Data{2}
        test.ValueTest()
        test.PointerTest()
}
