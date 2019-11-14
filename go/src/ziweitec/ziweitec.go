package main

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

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

//全局变量
var coder = base64.NewEncoding("LQL163GOODLUCKOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_")

//从打卡记录中同步提取会员注册信息
func addMember(db *sql.DB, userid, oid string) {
	var curTIme = time.Now().Format("2006-01-02 15:04:05")

	fmt.Printf("\n———————提取会员信息%s—————————\n", curTIme)
	//查询数据
	strQueryWhere := fmt.Sprintf("SELECT userID FROM member where openid = '%s'", oid)
	fmt.Println("strQueryWhere = ", strQueryWhere)
	rows, err := db.Query(strQueryWhere)
	checkErr(err)

	if rows.Next() {
		fmt.Println("记录已存在!")
		rows.Close()
		return
	} else {
		//插入数据
		gStmt, err := db.Prepare("INSERT INTO member(userID, openid, time) values(?,?,?)")
		checkErr(err)

		var curTime = time.Now().Format("2006-01-02 15:04:05")

		_, err = gStmt.Exec(userid, oid, curTime)
		checkErr(err)
		rows.Close()
		gStmt.Close()
		fmt.Printf("userid = %s\nopenid = %s\n", userid, oid)
	}
}

//处理打卡信息
func daka(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Println("GET请求，返回")
		return
	} else {
		//fmt.Println("POST请求，返回")
		r.ParseForm() //解析参数，默认是不会解析的

		var curTIme = time.Now().Format("2006-01-02 15:04:05")
		fmt.Printf("\n———————打卡%s—————————\n", curTIme)
		gDb, err := sql.Open("sqlite3", "./run.db")
		checkErr(err)

		//先获取值
		nickName, userID, openid, image, mile := r.Form.Get("nickname"), r.Form.Get("userID"), r.Form.Get("openid"), r.Form.Get("image"), r.Form.Get("mile")
		var score int
		fMile, err := strconv.ParseFloat(mile, 10)
		if fMile >= 3 && fMile < 5 {
			score = 2
		} else if fMile >= 5 && fMile < 10 {
			score = 3
		} else if fMile >= 10 && fMile < 15 {
			score = 4
		} else if fMile >= 15 && fMile < 21.0975 {
			score = 5
		} else if fMile >= 21.0975 && fMile < 30 {
			score = 7
		} else if fMile >= 30 && fMile < 42.195 {
			score = 8
		} else if fMile >= 42.195 {
			score = 9
		}

		//先获取值
		year, mon, day := time.Now().Date()
		//hour, min, sec := now.Clock()

		dtSml := fmt.Sprintf("%d-%02d-%02d 00:00:00", year, mon, day)
		dtBig := fmt.Sprintf("%d-%02d-%02d 00:00:00", year, mon, day+1)

		//先判断当天是否已打过卡
		//select * from runRec where time > '2019-11-03 00:00:00' and time < '2019-11-05 00:00:00' and userID = '477'
		preSql := fmt.Sprintf("select userID from runRec where time > '%s' and time < '%s' and userID = '%s'", dtSml, dtBig, userID)
		fmt.Println(preSql)
		rows, err := gDb.Query(preSql)
		checkErr(err)

		if rows.Next() {
			rows.Close()
			gDb.Close()
			w.Write([]byte("today_is_exist"))
			fmt.Println("今天已经打过卡了")
			return
		}

		//插入数据
		gStmt, err := gDb.Prepare("INSERT INTO runRec(nickname, userID, openid, mile, score, time, image) values(?,?,?,?,?,?,?)")
		checkErr(err)

		//fmt.Println("gStmt 前")
		var curTime = time.Now().Format("2006-01-02 15:04:05")
		// res, err := gStmt.Exec("477", "wxno", "5", curTime)

		_, err = gStmt.Exec(nickName, userID, openid, mile, score, curTime, image)
		checkErr(err)

		fmt.Printf("userid = %s\nnickname = %s\nopenid = %s\nimage = %s\nmile = %s\n", userID, nickName, openid, image, mile)
		//从打卡记录中提取会员信息
		addMember(gDb, userID, openid)

		gDb.Close()
	}
}

//查询会员信息
func queryMember(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//fmt.Println("queryMember")
		r.ParseForm() //解析参数，默认是不会解析的

		gDb, err := sql.Open("sqlite3", "./run.db")
		checkErr(err)

		var curTIme = time.Now().Format("2006-01-02 15:04:05")

		fmt.Printf("\n———————查询userID%s—————————\n", curTIme)
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

			expiration := time.Now()
			expiration = expiration.AddDate(1, 0, 0)

			//可以使用SetCookie代码set与add,比较方便
			ckUserID := http.Cookie{
				Name:    "userID",
				Value:   userIDval,
				Expires: expiration,
			}

			//设置头部必须在设置body之前，否则不起作用
			http.SetCookie(w, &ckUserID)
			w.Write([]byte(string(userIDval)))

			fmt.Println(userIDval)
		}

		gDb.Close()
	} else {
		fmt.Println("post请求，返回")
		return
	}
}

//得到要查询的所有的打卡记录
func queryRunRecord(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Println("queryRunRecord")
		r.ParseForm() //解析参数，默认是不会解析的

		gDb, err := sql.Open("sqlite3", "./run.db")
		checkErr(err)

		var curTIme = time.Now().Format("2006-01-02 15:04:05")

		fmt.Printf("\n———————queryrunrec%s—————————\n", curTIme)
		//查询数据
		openid := r.Form.Get("openid")
		fmt.Println("openid = ", openid)
		strQueryWhere := fmt.Sprintf("SELECT userID, openid, mile, score, time FROM runRec where openid = '%s'", openid)
		fmt.Println("strQueryWhere = ", strQueryWhere)
		rows, err := gDb.Query(strQueryWhere)
		checkErr(err)

		var runRecs []runRecord
		for rows.Next() {
			var rr runRecord
			var tv time.Time

			err := rows.Scan(&rr.UserID, &rr.Openid, &rr.Mile, &rr.Score, &tv)
			rr.Time = tv.Format("2006-01-02 15:04:05")
			checkErr(err)

			fmt.Printf("userID = %s mile = %.1f score = %d time = %s\n", rr.UserID, rr.Mile, rr.Score, rr.Time)
			fmt.Println("time : ", rr.Time)

			//添加到数组中
			runRecs = append(runRecs, rr)
			fmt.Println()
		}
		rows.Close()

		data, err := json.Marshal(runRecs)
		if err != nil {
			panic(err)
		}

		//获取记录条数
		nRCount := len(runRecs)
		ckstr := strconv.Itoa(nRCount)

		expiration := time.Now()
		expiration = expiration.AddDate(1, 0, 0)

		//可以使用SetCookie代码set与add,比较方便
		ckrr := http.Cookie{
			Name:    "rrCount",
			Value:   ckstr,
			Expires: expiration,
		}

		//设置头部必须在设置body之前，否则不起作用
		http.SetCookie(w, &ckrr)
		fmt.Println(runRecs)

		byteRet := []byte(data)

		w.Write(byteRet)
		w.WriteHeader(200)

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

		var curTIme = time.Now().Format("2006-01-02 15:04:05")
		fmt.Printf("\n———————获取openid%s—————————\n", curTIme)
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
	fmt.Println("service is started, current version:1.1.3, modified date 2019-11-11")
	//初始化数据库
	initFun()

	http.HandleFunc("/daka", daka)                  //打卡处理逻辑
	http.HandleFunc("/getOpenid", getOpenId)        //得到openID
	http.HandleFunc("/queryMember", queryMember)    //通过查询openid得到userID
	http.HandleFunc("/queryRunRec", queryRunRecord) //通过查询openid得到userID

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
