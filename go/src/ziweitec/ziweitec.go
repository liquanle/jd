package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// var gDb *sql.DB     //数据库
// var gStmt *sql.Stmt //操作实例

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Println("GET请求，返回")
		return
	} else {
		fmt.Println("POST请求，返回")
		r.ParseForm() //解析参数，默认是不会解析的

		strRunRec := "CREATE TABLE `runRec` ( " +
			"`uid` INTEGER PRIMARY KEY AUTOINCREMENT," +
			"`会员` VARCHAR(64) NULL," +
			"`微信号` VARCHAR(64) NULL," +
			"`距离` VARCHAR(64) NULL," +
			"`时间` DATE NULL" +
			");"

		gDb, err := sql.Open("sqlite3", "./run.db")
		gDb.Exec(strRunRec)
		checkErr(err)

		var curTIme = time.Now().Format("2006-01-02 15:04:05")

		fmt.Printf("——————————————————%s————————————————————\n", curTIme)
		fmt.Printf("the current time:%s \n", curTIme)
		fmt.Println("method:", r.Method) //获取请求的方法

		//fmt.Println(r.Form)  //这些信息是输出到服务器端的打印信息
		//fmt.Println("path", r.URL.Path)
		//fmt.Println("scheme", r.URL.Scheme)
		//fmt.Println(r.Form["url_long"])

		//插入数据
		gStmt, err := gDb.Prepare("INSERT INTO runRec(会员, 微信号, 距离, 时间) values(?,?,?,?)")
		checkErr(err)

		fmt.Println("gStmt 前")
		var curTime = time.Now().Format("2006-01-02 15:04:05")
		// res, err := gStmt.Exec("477", "wxno", "5", curTime)
		res, err := gStmt.Exec(r.Form.Get("no"), "wxno", r.Form.Get("mile"), curTime)
		checkErr(err)
		affect, err := res.RowsAffected()
		checkErr(err)

		fmt.Println(affect)

		strOut := fmt.Sprintf("%s hit %s mile!\n", r.Form.Get("no"), r.Form.Get("mile"))
		//fmt.Fprintf(w, strOut) //这个写入到w的是输出到客户端的
		fmt.Println(strOut)
		gDb.Close()
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	r.ParseForm()                    //解析参数，默认是不会解析的
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
	//初始化数据库
	//initFun()

	fmt.Println("开始运行！")
	http.HandleFunc("/", sayhelloName) //设置访问的路由
	//http.HandleFunc("/login", login)   //设置登录路由
	//err := http.ListenAndServe(":8080", nil)
	err := http.ListenAndServeTLS(":443", "2988657_ziweitec.com.pem", "2988657_ziweitec.com.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func initFun() {
	strRunRec := "CREATE TABLE `runRec` ( " +
		"`uid` INTEGER PRIMARY KEY AUTOINCREMENT," +
		"`会员` VARCHAR(64) NULL," +
		"`微信号` VARCHAR(64) NULL," +
		"`距离` VARCHAR(64) NULL," +
		"`时间` DATE NULL" +
		");"

	gDb, err := sql.Open("sqlite3", "./run.db")
	gDb.Exec(strRunRec)
	checkErr(err)
	// gStmt, err := gDb.Prepare("INSERT INTO runRec(会员, 微信号, 距离, 时间) values(?,?,?,?)")
	// checkErr(err)

	// fmt.Println("gStmt 前")
	// var curTime = time.Now().Format("2006-01-02 15:04:05")
	// res, err := gStmt.Exec("lql", "wxno", "30", curTime)
	// checkErr(err)
	// affect, err := res.RowsAffected()
	// checkErr(err)

	// fmt.Println(affect)
}
