package main

import (
	"fmt"
)

type Stringer interface {
	Print() string
}

type User struct {
	id   int
	name string
}

func (self *User) Print() string {
	return fmt.Sprintf("user %d, %s", self.id, self.name)
}

type struct1 struct {
	x int
	y float32
	z string
}

func main() {
	t1 := new(int)
	fmt.Println("t1=", t1)   //t1= 0xc0420341c0
	fmt.Println("*t1=", *t1) //*t1= 0

	var t2 *string = new(string)
	fmt.Println("t2=", t2)   //t2= 0xc042034230
	fmt.Println("*t2=", *t2) //*t2=

	t3 := new(struct1)
	fmt.Println("t3=", t3)   //t3= &{0 0 }
	fmt.Println("*t3=", *t3) //*t3= {0 0 }

	var t4 Stringer = new(User)
	fmt.Println("t4=", t4)         //t4= &{0 }
	fmt.Println("t4=", t4.Print()) //t4= user 0,

	//fmt.Println("t4.id=", *t4.id) 报错
	// fmt.Println("*t4=", *t4)  报错
}
