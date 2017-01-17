```go

$GOPATH/src/
    libproj1/foo/
            – foo1.go
    app1
            – main.go

```

```go

//foo1.go
package foo

import "fmt"

func Foo1() {
    fmt.Println("Foo1")
}

```

```go

// main.go
package main

import (
    "libproj1/foo"
)

func main() {
    foo.Foo1()
}

```

在使用第三方包的时候，当源码和.a均已安装的情况下，编译器链接的是源码。


```go

$GOPATH/src/
    libproj1/foo/
            – foo1.go
    app1
            – main.go

```go

//foo1.go
package bar

import "fmt"

func Foo1() {
    fmt.Println("Foo1")
}

```

```go

// main.go
package main

import (
    "libproj1/foo"
    // import后面的最后一个元素应该是路径，就是目录，并非包名。
    // "libproj1/bar"   //error
)

func main() {
    foo.Foo1()
}

```

