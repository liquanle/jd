// testjson project main.go
package main

import (
	"encoding/json"
	"fmt"
)

type chnItem struct {
	channel string
	value   string
	time    string
	status  string
}

type deviceItem struct {
	id            string
	deviceVersion string
	freq1         string
	freq2         string
	connectFreq   string
	softVersion   string
	energy        string
	signal        string
	tem           string
	clockStatus   string
	save_con      string
	read_con      string
	IO1_Status    string
	IO2_Status    string
}

type dataItem struct {
	device deviceItem
	datas  []chnItem
}

type jgItem struct {
	flag int
	di   dataItem
}

func main() {
	jsonfmt := `{"flag":"1", "data":{"dev":{"id":"192828","deviceVersion":"V2.1","freq1":"1440","freq2":"1440","connectFreq":"1440","softVersion":"V2.04","energy":"100","signal":"51","tem":"9","clockStatus":"0","save_con":"13","read_con":"11","IO1_Status":"1","IO2_Status":"0"},"datas":[{"channel":"3","value":"0","time":"20191111171903","status":"0"},{"channel":"4","value":"0","time":"20191111171903","status":"0"}]}}`
	fmt.Println(jsonfmt)
	bytes := []byte(jsonfmt)

	//1.Unmarshal的第一个参数是json字符串，第二个参数是接受json解析的数据结构。
	//第二个参数必须是指针，否则无法接收解析的数据，如stu仍为空对象StuRead{}
	//2.可以直接stu:=new(StuRead),此时的stu自身就是指针
	stu := jgItem{}
	err := json.Unmarshal(bytes, &stu)

	//解析失败会报错，如json字符串格式不对，缺"号，缺}等。
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(stu)
}
