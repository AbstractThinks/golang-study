package main

import (
        "fmt"
        "html/template"
        "log"
        "net/http"
        "strings"
)

func main() {
        http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
                r.ParseForm()    //解析url传递的参数，对于POST则解析相应包的主体(request body)
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
        http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
                fmt.Println("method: ", r.Method)
                r.ParseForm()
                fmt.Println("r.Form is ",r.Form)   //r.Form url信息
                fmt.Println("path is ", r.URL.Path)
                fmt.Println("scheme", r.URL.Scheme)
                if r.Method == "GET" {
                        t, _ := template.ParseFiles("form.html")
                        t.Execute(w, nil)
                }
                if r.Method == "POST" {
                        fmt.Println("username: ", r.Form["username"])
                        fmt.Println("password: ", r.Form["password"])
                        if "1" == r.Form["username"][0] && "1" == r.Form["password"][0] {
                                http.Redirect(w,r, "/hello", http.StatusFound)
                        } else {
                                http.Redirect(w,r, "/error", http.StatusFound)
                        }
                }
        })
        
        http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
                fmt.Println("method: ", r.Method)
                r.ParseForm()
                fmt.Println("r.Form is ",r.Form)   //r.Form url信息
                fmt.Println("path is ", r.URL.Path)
                fmt.Println("scheme", r.URL.Scheme)
                if r.Method == "GET" {
                        t, _ := template.ParseFiles("hello.html")
                        t.Execute(w, nil)
                }

        })
        http.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
                fmt.Println("method: ", r.Method)
                r.ParseForm()
                fmt.Println("r.Form is ",r.Form)   //r.Form url信息
                fmt.Println("path is ", r.URL.Path)
                fmt.Println("scheme", r.URL.Scheme)
                if r.Method == "GET" {
                        t, _ := template.ParseFiles("error.html")
                        t.Execute(w, nil)
                }

        })
        err := http.ListenAndServe(":9090", nil)
        if err != nil {
                log.Fatal("ListenAndServe: ", err)
        }
}
