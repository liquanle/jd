// sqllite project main.go
package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	// strHitCard := "CREATE TABLE `runRec` ( " +
	// 	"`uid` INTEGER PRIMARY KEY AUTOINCREMENT," +
	// 	"`会员` VARCHAR(64) NULL," +
	// 	"`微信号` VARCHAR(64) NULL," +
	// 	"`距离` VARCHAR(64) NULL," +
	// 	"`时间` DATE NULL" +
	// 	");"
	strRunRec := "CREATE TABLE `runRec` ( " +
		"`uid` INTEGER PRIMARY KEY AUTOINCREMENT," +
		"`nickname` VARCHAR(64) NULL," +
		"`userID` VARCHAR(64) NULL," +
		"`openid` VARCHAR(64) NULL," +
		"`mile` VARCHAR(64) NULL," +
		"`time` DATE NULL," +
		"`image` VARCHAR(64) NULL" +
		");"

	gDb, err := sql.Open("sqlite3", "./run.db")
	res, err := gDb.Exec(strRunRec)
	checkErr(err)

	strInfo := "CREATE TABLE `userinfo` ( " +
		"`uid` INTEGER PRIMARY KEY AUTOINCREMENT," +
		"`username` VARCHAR(64) NULL," +
		"`departname` VARCHAR(64) NULL," +
		"`created` DATE NULL" +
		");"

	strInfo2 := "CREATE TABLE `userdeatail` (" +
		"`uid` INT(10) NULL," +
		"`intro` TEXT NULL," +
		"`profile` TEXT NULL," +
		"PRIMARY KEY (`uid`)" +
		");"

	db, err := sql.Open("sqlite3", "./daka.db")
	db.Exec(strInfo)
	db.Exec(strInfo2)
	checkErr(err)

	strExitTable := "SELECT count(*) from sqlite_master where type='table' and name = ?"
	stmt, err := db.Prepare(strExitTable)
	res, err = stmt.Exec("userinfo")
	fmt.Println(res)

	//插入数据
	// stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
	// checkErr(err)

	var curTime = time.Now().Format("2006-01-02 15:04:05")
	// res, err := stmt.Exec("李全乐", "研发部门", curTime)
	// checkErr(err)
	fmt.Println(curTime)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
	//更新数据
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("李四光", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	//查询数据
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	//删除数据
	// stmt, err = db.Prepare("delete from userinfo where uid=?")
	// checkErr(err)

	// res, err = stmt.Exec(id)
	// checkErr(err)

	// affect, err = res.RowsAffected()
	// checkErr(err)

	fmt.Println(affect)

	db.Close()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
