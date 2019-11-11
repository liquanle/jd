package main

import (
	//"encoding/json"
	"fmt"

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
	str := `{"flag":"1", "servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	//json.Unmarshal([]byte(str), &s)
	js, err := simplejson.NewJson([]byte(string(str)))

	if err == nil {
		flag1 := js.Get("flag").MustString()

		fmt.Println("flag:", flag1)
		//fmt.Println("openid:", od)
		fmt.Println(js)
	}

}
