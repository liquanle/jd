package main

import (
    "fmt"
    "net/http"
	"html/template"
    "strings"
	"time"
    "log"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()  //解析参数，默认是不会解析的
	var goos = time.Now().Format("2006-01-02 15:04:05")
    
	fmt.Printf("——————————————————%s————————————————————\n", goos)
	fmt.Printf("the current time:%s \n", goos)
	fmt.Println("method:", r.Method) //获取请求的方法

    //fmt.Println(r.Form)  //这些信息是输出到服务器端的打印信息
    fmt.Println("path", r.URL.Path)
    fmt.Println("scheme", r.URL.Scheme)
    fmt.Println(r.Form["url_long"])
    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }
    fmt.Fprintf(w, "Hello everyone!") //这个写入到w的是输出到客户端的
}

func login(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method) //获取请求的方法
	r.ParseForm()  //解析参数，默认是不会解析的
    if r.Method == "GET" {
        t, _ := template.ParseFiles("login.gtpl")
        t.Execute(w, nil)
    } else {
        //请求的是登陆数据，那么执行登陆的逻辑判断
        fmt.Println("username:", r.Form["username"])
        fmt.Println("password:", r.Form["password"])
    }
}

func main() {
	fmt.Println("开始运行！")
<<<<<<< HEAD
    http.HandleFunc("/", sayhelloName) //设置访问的路由
=======
    http.HandleFunc("/liquanle", sayhelloName) //设置访问的路由
>>>>>>> 279377104914d4c93565ca50ab75bec345fa4ea4
	//http.HandleFunc("/login", login)   //设置登录路由
    err := http.ListenAndServe(":8080", nil) //设置监听的端口
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}