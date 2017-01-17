### 变量

```go

var 变量名字 类型 = 表达式

其中“类型”或“= 表达式”两个部分可以省略其中的一个。如果省略的是类型信息，那么将根据初始化表达式来推导变量的类型信息。如果初始化表达式被省略，那么将用零值初始化该变量

```

### 简短变量声明

```go

名字 := 表达式

只能在函数内部使用

```

### new函数

```go

表达式new(T)将创建一个T类型的匿名变量，初始化为T类型的零值，然后返回变量地址，返回的指针类型为*T

p := new(int)   // p, *int 类型, 指向匿名的 int 变量
fmt.Println(*p) // "0"
*p = 2          // 设置 int 匿名变量的值为 2
fmt.Println(*p) // "2"

```


### 赋值

```go

x = 1                       // 命名变量的赋值
*p = true                   // 通过指针间接赋值
person.name = "bob"         // 结构体字段赋值
count[x] = count[x] * scale // 数组、slice或map的元素赋值
v := 1
v++    // 等价方式 v = v + 1；v 变成 2
v--    // 等价方式 v = v - 1；v 变成 1
medals := []string{"gold", "silver", "bronze"}
medals[0] = "gold" 
medals[1] = "silver" 
medals[2] = "bronze"
x, y = y, x
a[i], a[j] = a[j], a[i]

```
