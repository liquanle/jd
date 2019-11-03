package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"lql.com/tool/file"

	"github.com/bitly/go-simplejson"
	_ "github.com/mattn/go-sqlite3"
)

// var gDb *sql.DB     //数据库
// var gStmt *sql.Stmt //操作实例

// type resOpenID struct {
// 	session_key string
// 	openid      string
// }

// type resOpenIDslice struct {
// 	resOpenIDs []resOpenID
// }

//处理打卡信息
func dakaservice(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Println("GET请求，返回")
		return
	} else {
		//fmt.Println("POST请求，返回")
		r.ParseForm() //解析参数，默认是不会解析的

		gDb, err := sql.Open("sqlite3", "./run.db")
		checkErr(err)

		var curTIme = time.Now().Format("2006-01-02 15:04:05")

		fmt.Printf("————————————%s——————————————\n", curTIme)
		//fmt.Printf("the current time:%s \n", curTIme)
		//fmt.Println("method:", r.Method) //获取请求的方法

		//插入数据
		gStmt, err := gDb.Prepare("INSERT INTO runRec(nickname, userID, openid, mile, time, image) values(?,?,?,?,?,?)")
		checkErr(err)

		//fmt.Println("gStmt 前")
		var curTime = time.Now().Format("2006-01-02 15:04:05")
		// res, err := gStmt.Exec("477", "wxno", "5", curTime)
		nickName, userID, openid, image, mile := r.Form.Get("nickname"), r.Form.Get("userID"), r.Form.Get("openid"), r.Form.Get("image"), r.Form.Get("mile")
		_, err = gStmt.Exec(nickName, userID, openid, mile, curTime, image)
		checkErr(err)
		// affect, err := res.RowsAffected()
		// checkErr(err)

		//fmt.Println(affect)

		//strOut := fmt.Sprintf("%s hit %s mile!\n", userID, mile)
		//fmt.Fprintf(w, strOut) //这个写入到w的是输出到客户端的
		//fmt.Println(strOut)
		fmt.Printf("userid = %s\nnickname = %s\nopenid = %s\nimage = %s\nmile = %s", userID, nickName, openid, image, mile)
		gDb.Close()
	}
}

//查询会员信息
func queryMember(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Println("queryMember")
		r.ParseForm() //解析参数，默认是不会解析的

		gDb, err := sql.Open("sqlite3", "./run.db")
		checkErr(err)

		var curTIme = time.Now().Format("2006-01-02 15:04:05")

		fmt.Printf("————————————%s——————————————\n", curTIme)
		//查询数据
		openid := r.Form.Get("openid")
		fmt.Println("openid = ", openid)
		strQueryWhere := fmt.Sprintf("SELECT userID FROM member where openid = '%s'", openid)
		fmt.Println("strQueryWhere = ", strQueryWhere)
		rows, err := gDb.Query(strQueryWhere)
		checkErr(err)

		for rows.Next() {
			userIDval := ""

			err := rows.Scan(&userIDval)
			checkErr(err)

			fmt.Println(userIDval)
			w.Write([]byte(string(userIDval)))
		}

		gDb.Close()
	} else {
		fmt.Println("post请求，返回")
		return
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

/*
	通过微信接口调用得到OpenId
*/
func getOpenId(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	r.ParseForm()                    //解析参数，默认是不会解析的
	if r.Method == "GET" {
		strAppID := r.Form.Get("appid")
		strSecret := r.Form.Get("secret")
		strJs_code := r.Form.Get("js_code")
		strGrant_type := r.Form.Get("grant_type")

		//fmt.Println(strAppID + "\n" + strSecret + "\n" + strJs_code + "\n" + strGrant_type)

		client2 := &http.Client{}

		strUrl := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=%s", strAppID, strSecret, strJs_code, strGrant_type)
		//fmt.Println(strUrl)
		resp, err := client2.Get(strUrl)

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(string(body))

		// json.Unmarshal([]byte(string(body)), &s)

		js, err := simplejson.NewJson([]byte(string(body)))

		sk := js.Get("session_key").MustString()
		od := js.Get("openid").MustString()
		fmt.Println("session_key:", sk)
		fmt.Println("openid:", od)

		//这是返回信息，返回到调用方，微信小程序
		w.Write([]byte(string(body)))
		//w.WriteHeader(404)
	} else {
		//请求的是登陆数据，那么执行登陆的逻辑判断
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])

	}
}

func main() {
	//初始化数据库
	initFun()

	fmt.Println("开始运行！")
	http.HandleFunc("/", dakaservice)            //设置访问的路由
	http.HandleFunc("/getOpenid", getOpenId)     //得到openID
	http.HandleFunc("/queryMember", queryMember) //通过查询openid得到userID
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
		"`nickname` VARCHAR(64) NULL," +
		"`userID` VARCHAR(64) NULL," +
		"`openid` VARCHAR(64) NULL," +
		"`mile` VARCHAR(64) NULL," +
		"`time` DATE NULL," +
		"`image` VARCHAR(256) NULL" +
		");"

	strMember := "CREATE TABLE `member` ( " +
		"`uid` INTEGER PRIMARY KEY AUTOINCREMENT," +
		"`nickname` VARCHAR(64) NULL," +
		"`userID` VARCHAR(64) NULL," +
		"`openid` VARCHAR(64) NULL," +
		"`time` DATE NULL," +
		"`image` VARCHAR(256) NULL" +
		");"

	bExist, err := file.PathExists("./run.db")
	gDb, err := sql.Open("sqlite3", "./run.db")

	if !bExist {
		_, err = gDb.Exec(strRunRec)
		checkErr(err)

		_, err = gDb.Exec(strMember)
		checkErr(err)
	}

	gDb.Close()
}
