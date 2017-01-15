package main

import (
        "fmt"
        "net/http"
        "strings"
        "log"
)
//func sayhelloName(w http)
func main() {
        //http.HandlerFunc("/", sayhelloName)     //设置访问路由
        http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
                r.ParseForm() //解析参数
                fmt.Println("r.Form is ",r.Form)   //r.Form url信息
                fmt.Println("path is ", r.URL.Path)
                fmt.Println("scheme", r.URL.Scheme)
                fmt.Println(r.Form["url_long"])
                for k, v := range r.Form {
                        fmt.Println("key: ", k)
                        fmt.Println("val: ", strings.Join(v, ""))
                }
                fmt.Fprintf(w, "Hello astaxie")   //输出到客户端
        })
        err := http.ListenAndServe(":9090", nil)        //设置监听端口
        if err != nil {
                log.Fatalf("ListenAndServe: ", err)
        }
}


 //http://localhost:9090/?url_long=111&url_long=222
 /*
 r.Form is  map[url_long:[111 222]]         fmt.Println("r.Form is ",r.Form)
 path is  /                                                     fmt.Println("path is ", r.URL.Path)
 scheme                                                      fmt.Println("scheme", r.URL.Scheme)
 [111 222]                                                    fmt.Println(r.Form["url_long"])
 key:  url_long
 val:  111222
 */

// http://localhost:9090/?url_long=111&url_long2=222
/*
r.Form is  map[url_long:[111] url_long2:[222]]            fmt.Println("r.Form is ",r.Form)
path is  /                                                                            fmt.Println("path is ", r.URL.Path)
scheme                                                                             fmt.Println("scheme", r.URL.Scheme)
[111]                                                                                   fmt.Println(r.Form["url_long"])
key:  url_long2
val:  222
key:  url_long
val:  111
*/
