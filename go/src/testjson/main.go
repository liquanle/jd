// testjson project main.go
package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"

	"github.com/bitly/go-simplejson"
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
	jsonfmt := `{"flag":1, "data":{"dev":{"id":"1904230004","deviceVersion":"V2.1","freq1":"1440","freq2":"1440","connectFreq":"1440","softVersion":"V2.04","energy":"100","signal":"51","tem":"9","clockStatus":"0","save_con":"13","read_con":"11","IO1_Status":"1","IO2_Status":"0"},"datas":[{"channel":"3","value":86,"time":"20191111171903","status":"0"},{"channel":"4","value":35,"time":"20191111171903","status":"0"}]}}`
	fmt.Println(jsonfmt)
	bytes := []byte(jsonfmt)

	jss, err := simplejson.NewJson(bytes)

	if err == nil {
		flag1 := jss.Get("flag").MustInt()
		fmt.Println("flag:", flag1)

		deviceid := jss.Get("data").Get("dev").Get("id").MustString()
		fmt.Println("deviceid:", deviceid)

	} else {
		fmt.Println("解析错误！:")
	}

	mapSer, err := jss.Get("data").Get("datas").Array()
	for _, row := range mapSer {
		if each_map, ok := row.(map[string]interface{}); ok {
			fmt.Println(reflect.TypeOf(each_map["channel"]))

			if serN, ok := each_map["channel"].(string); ok {
				fmt.Println("channel:", serN)
			}

			fmt.Println(reflect.TypeOf(each_map["value"]))
			if number, ok := each_map["value"].(json.Number); ok {

				value, error := strconv.ParseInt(string(number), 10, 0)

				if error == nil {
					fmt.Println("value:", value)
				} else {
					fmt.Println(error)
				}
			}

		}

	}
}
