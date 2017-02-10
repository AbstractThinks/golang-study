package main

import (
	"fmt"
)

type Stringer interface {
	String() string
}

type Printer interface {
	Stringer
	Print()
}
type User struct {
	id   int
	name string
}

func (self *User) String() string {
	fmt.Println(*self) //{1 Tom}
	fmt.Println(self.id)
	return fmt.Sprintf("user %d, %s", self.id, self.name)
}

func (self *User) Print() {
	fmt.Println(self.String())
}

/*
type User2 struct {
        id int
        name string
}

func (self *User2) Print() {
        fmt.Println(self.String())
}
var t3 Printer = &User2{2, "Jesse"}     //报错，只实现接口一个方法不算继承了接口，必须实现接口中所有的方法才算继承了接口
t3.Print()
*/

func main() {
	fmt.Println("hello world")

	var t Printer = &User{1, "Tom"}
	var t2 Printer = new(User)

	t.Print()
	t2.Print()

	fmt.Println("test:", t2)

}

func intertest(printer Printer) {
	fmt.Println("type test")

}
