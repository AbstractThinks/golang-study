package main

import (
        "fmt"
)

type Node struct {
        _ int
        id int
        data *byte
        next *Node
}

type File struct {
        name string
        size int
        attr struct {
                perm int
                owner int
        }
}
func main() {

        fmt.Println("hello world")

        n1 := Node{
                id: 1,
                data: nil,
        }
        n2 := Node{
                id: 2,
                data: nil,
                next: &n1,
        }

        fmt.Println(n1)
        fmt.Println(n2)
        fmt.Println(*(n2.next))

        f := File{
                name: "test.txt",
                size: 1025,
        }
        f2 := File{
                name: "test.txt",
                size: 1025,
                //attr: {0755, 1}  error
        }
        var attr_f2 = struct {
                perm int
                owner int
        }{2, 0755}
        f2.attr = attr_f2
        fmt.Println(f)
        fmt.Println(f2)
}
