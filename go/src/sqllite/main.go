// sqllite project main.go
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	strHitCard := "CREATE TABLE `runRec` ( " +
		"`uid` INTEGER PRIMARY KEY AUTOINCREMENT," +
		"`会员` VARCHAR(64) NULL," +
		"`微信号` VARCHAR(64) NULL," +
		"`距离` VARCHAR(64) NULL," +
		"`时间` DATE NULL" +
		");"

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

	//插入数据
	stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
	checkErr(err)

	var curTime = time.Now().Format("2006-01-02 15:04:05")
	res, err := stmt.Exec("李全乐", "研发部门", curTime)
	checkErr(err)

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
