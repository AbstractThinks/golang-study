go build -gcflags "-N -l" -o demo demo.go



(gdb) l <-------------------- l命令相当于list，从第一行开始例出原码。
(gdb) <-------------------- 直接回车表示，重复上一次命令
(gdb) break number <-------------------- 设置断点，在源程序第number行处。
(gdb) r <--------------------- 运行程序，run命令简写
(gdb) info breakpoints <-------------------- 查看所有断点信息。
(gdb) n <--------------------- 单条语句执行，next命令简写。
(gdb) c <--------------------- 继续运行程序，continue命令简写。
(gdb) p i <--------------------- 打印变量i的值，print命令简写。
(gdb) bt <--------------------- 查看函数堆栈。
(gdb) finish <--------------------- 退出函数。
(gdb) delete 1 <--------------------- 删除一个断点
(gdb) delete 1-10 <--------------------- 删除多个断点
(gdb) clear func_name         //删除函数的所有断点
(gdb) clear filename.go:func_name   //删除文件：函数的所有断点
(gdb) clear 12                  //删除行号的所有断点
(gdb) clear filename.go:12           //删除文件：行号的所有断点
