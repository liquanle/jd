package main

import (
	"encoding/json"
	_ "encoding/json"
	"fmt"
	"reflect"
	"strconv"

	"github.com/bitly/go-simplejson"
)

type Server struct {
	ServerName string
	ServerIP   string
}

type Serverslice struct {
	Servers []Server
}

type ok struct {
	flag    string
	servero Serverslice
}

func main() {
	//var s = ok{}
	str := `{"flag":"1", 
			  "servers":[
					{"serverName":"Shanghai_VPN","price":28000},
					{"serverName":"Beijing_VPN","price":9280}
				]}`
	//json.Unmarshal([]byte(str), &s)
	jss, err := simplejson.NewJson([]byte(string(str)))

	if err == nil {
		flag1 := jss.Get("flag").MustString()

		fmt.Println("flag:", flag1)
		//fmt.Println("openid:", od)
		//fmt.Println(jss)
	} else {
		fmt.Println("解析错误！:")
	}

	mapSer, err := jss.Get("servers").Array()
	for _, row := range mapSer {
		if each_map, ok := row.(map[string]interface{}); ok {
			fmt.Println(reflect.TypeOf(each_map["serverName"]))

			if serN, ok := each_map["serverName"].(string); ok {
				fmt.Println(serN)
			}

			fmt.Println(reflect.TypeOf(each_map["price"]))
			if number, ok := each_map["price"].(json.Number); ok {

				nPrice, error := strconv.ParseInt(string(number), 10, 0)

				if error == nil {
					fmt.Println("number:", nPrice)
				} else {
					fmt.Println(error)
				}
			}

		}

	}
}
