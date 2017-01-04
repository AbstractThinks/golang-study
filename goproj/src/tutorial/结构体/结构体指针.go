package main

import "fmt"

type Vertex struct {
	X int
	Y int
}
var (
	v1 = Vertex{1, 2}  // 类型为 Vertex
	v2 = Vertex{X: 1}  // Y:0 被省略
	v3 = Vertex{}      // X:0 和 Y:0
	p1  = &Vertex{1, 2} // 类型为 *Vertex
)
func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
        fmt.Println(v1, p1, v2, v3) //{1 2} &{1 2} {1 0} {0 0}
        fmt.Println(v1, *p1, v2, v3) //{1 2} {1 2} {1 0} {0 0}
}
