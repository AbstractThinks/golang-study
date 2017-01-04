/*
除非以 fallthrough 语句结束，否则分支会自动终止。

*/

package main

import (
	"fmt"
	"runtime"
        "time"
)

func main() {

	fmt.Println("Go runs on ")

	switch os := runtime.GOOS; os {
        	case "darwin":
        		fmt.Println("OS X.")
        	case "linux":
        		fmt.Println("Linux.")
                case "windows":
        		fmt.Println("Linux.")
                        fallthrough
        	default:
        		// freebsd, openbsd,
        		// plan9, windows...
        		fmt.Println(os)
	}

// 没有条件的 switch 同 switch true 一样。
        t := time.Now()
	switch {
        	case t.Hour() < 12:
        		fmt.Println("Good morning!")
        	case t.Hour() < 17:
        		fmt.Println("Good afternoon.")
        	default:
        		fmt.Println("Good evening.")
	}
}
