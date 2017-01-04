package main

import (
	"fmt"
	"strings"
)

func main() {
        //[]T 是一个元素类型为 T 的 slice
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("s ==", s)  //s == [2 3 5 7 11 13]

        // 代表从 1 开始
        fmt.Println("s[1:4] ==", s[1:4])   //[3 5 7]

	// 省略下标代表从 0 开始
	fmt.Println("s[:3] ==", s[:3])   //[2 3 5]

	// 省略上标代表到 len(s) 结束
	fmt.Println("s[4:] ==", s[4:])  // [11 13]


        //len(s) 返回 slice s 的长度
	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d] == %d\n", i, s[i])
	}
        /*
        s[0] == 2
        s[1] == 3
        s[2] == 5
        s[3] == 7
        s[4] == 11
        s[5] == 13
        */


        //slice 可以包含任意的类型，包括另一个 slice。

        // Create a tic-tac-toe board.
	game := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	game[0][0] = "X"
	game[2][2] = "O"
	game[2][0] = "X"
	game[1][0] = "O"
	game[0][2] = "X"

	printBoard(game)


        a := make([]int, 5)
	printSlice("a", a)    //a len=5 cap=5 [0 0 0 0 0]
	b := make([]int, 0, 5)
	printSlice("b", b)    //b len=0 cap=5 []
	c := b[:2]
	printSlice("c", c)    //c len=2 cap=5 [0 0]
	d := c[2:5]
	printSlice("d", d)   //d len=3 cap=3 [0 0 0]


        //slice 的零值是 nil 。
        //一个 nil 的 slice 的长度和容量是 0。
        var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")
	}

        //向slice添加元素
        var a_append []int
	printSlice("a_append", a_append)

	// append works on nil slices.
	a_append = append(a_append, 0)
	printSlice("a_append", a_append)

	// the slice grows as needed.
	a_append = append(a_append, 1)
	printSlice("a_append", a_append)

	// we can add more than one element at a time.
	a_append = append(a_append, 2, 3, 4)
	printSlice("a_append", a_append)
}

func printBoard(s [][]string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%s\n", strings.Join(s[i], " "))
	}
}


func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
