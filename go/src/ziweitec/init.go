// init
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"lql.com/tool/file"
)

//数据库初始化
func initFun() {
	fmt.Println("Initing database!")
	strRunRec := "CREATE TABLE `runRec` ( " +
		"`uid` INTEGER PRIMARY KEY AUTOINCREMENT," +
		"`nickname` VARCHAR(64) NULL," +
		"`userID` VARCHAR(64) NULL," +
		"`openid` VARCHAR(64) NULL," +
		"`mile` FLOAT NULL," +
		"`score` INT NULL," +
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
	if bExist {
		fmt.Println("database is existed!")
	} else {
		fmt.Println("database is not existed!")
	}

	gDb, err := sql.Open("sqlite3", "./run.db")
	if err == nil {
		fmt.Println("dataBase[run.db] is created！")
	}

	if !bExist {
		fmt.Println("Tables is created!")
		_, err = gDb.Exec(strRunRec)
		checkErr(err)

		_, err = gDb.Exec(strMember)
		checkErr(err)
	}

	gDb.Close()
}
