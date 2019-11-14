// sxjgpost project main.go
package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"net/http"
	"time"
)

/*
	jf: 井盖格式字符串
	jd: 井盖id
	jgs:井盖状态
	ct: 时间
	idx: for循环索引
*/
func postJGMsg(jf, jgID, jgs, ct string, idx int) {
	fmt.Println("")
	fmt.Println("######start#######", jgID)

	//井盖状态 【1】代表打开  【0】代表正常
	jsonStr := fmt.Sprintf(jf, jgID, jgs, ct, jgs, ct)
	fmt.Println(jsonStr)

	resp, error := http.Post("http://112.31.239.184:8100/iot/HDDTService", "application/x-www-form-urlencoded", strings.NewReader(jsonStr))
	if error != nil {
		fmt.Println(error)
	}
	//
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	fmt.Println("#####end#######", jgID)

	//time.Sleep(time.Duration(5-idx) * time.Second)
}

//主进程
func mainfun() {
	//jgids := []string{"1904230001", "1904230003", "1904230004", "1904230005", "1904230006", "1904230007", "1904230008", "1904230010"}
	jgids := []string{"1904230001", "1904230003", "1904230004", "1904230005"}
	//jgids := []string{"1904230004"}
	//成功
	curtime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println("curtime =", curtime)

	now := time.Now()
	year, mon, day := now.Date()
	hour, min, sec := now.Clock()

	curtime = fmt.Sprintf("%d%d%d%02d%02d%02d", year, mon, day, hour, min, sec)
	fmt.Println("curtime =", curtime)
	jsonfmt := `flag=0&data={"dev":{"id":"%s","deviceVersion":"V2.1","freq1":"1440","freq2":"1440","connectFreq":"1440","softVersion":"V2.04","energy":"100","signal":"51","tem":"9","clockStatus":"0","save_con":"13","read_con":"11","IO1_Status":"1","IO2_Status":"0"},"datas":[{"channel":"3","value":"%s","time":"%s","status":"0"},{"channel":"4","value":"%s","time":"%s","status":"0"}]}`

	var jgStatu = "0"
	for i, e := range jgids {
		//启动多线程
		go postJGMsg(jsonfmt, e, jgStatu, curtime, i)
		//postJGMsg(jsonfmt, e, jgStatu, curtime, i)
	}
}

func main() {
	fmt.Println("sxjgpost program!")

	//定时任务开始
	for {
		mainfun()
		now := time.Now()
		// 计算下一个零点
		next := now.Add(time.Hour * 24)
		next = time.Date(next.Year(), next.Month(), next.Day(), 1, 2, 3, 0, next.Location())
		t := time.NewTimer(next.Sub(now))
		<-t.C
	}
	//定时任务结束

	time.Sleep(time.Duration(5) * time.Second)
	fmt.Println("#主线程结束#")
}
