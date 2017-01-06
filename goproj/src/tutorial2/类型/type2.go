package main
import (
        "fmt"
)

type Books struct {
   title string
   author string
   subject string
   book_id int
}

/*
   引用类型包括slice、map和channel
*/
func main() {
        a := new([]int)   //new 计算类型大小，为其分配零值内存，返回指针
        b := make([]int, 3)  //make会被编译器翻译成具体的创建函数，由其分配内存和初始化成员结构并返回对象

        fmt.Println(*a)
        fmt.Println(b)

        var tran_b byte = 100
        var res_n string = string(tran_b)
        fmt.Println(tran_b)
        fmt.Println(res_n)   // d    ASCII码表100对应小写字母d


                var Book1 Books        /* 声明 Book1 为 Books 类型 */
                var Book2 Books        /* 声明 Book2 为 Books 类型 */

                /* book 1 描述 */
                Book1.title = "Go 语言"
                Book1.author = "www.runoob.com"
                Book1.subject = "Go 语言教程"
                Book1.book_id = 6495407

                /* book 2 描述 */
                Book2.title = "Python 教程"
                Book2.author = "www.runoob.com"
                Book2.subject = "Python 语言教程"
                Book2.book_id = 6495700

                /* 打印 Book1 信息 */
                printBook(Book1)

                /* 打印 Book2 信息 */
                printBook(Book2)

                fmt.Println(&Book1)

                printBook_pointer(&Book1)
                printBook_pointer(&Book2)

}

func printBook( book Books ) {
   fmt.Printf( "Book title : %s\n", book.title);
   fmt.Printf( "Book author : %s\n", book.author);
   fmt.Printf( "Book subject : %s\n", book.subject);
   fmt.Printf( "Book book_id : %d\n", book.book_id);
}
func printBook_pointer(book *Books ) {
   fmt.Printf( "Book title : %s\n", book.title);
   fmt.Printf( "Book author : %s\n", book.author);
   fmt.Printf( "Book subject : %s\n", book.subject);
   fmt.Printf( "Book book_id : %d\n", book.book_id);
}
