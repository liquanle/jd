// testhttp project main.go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	client2 := &http.Client{}
	strUrl := "https://api.weixin.qq.com/sns/jscode2session?appid=wxa1b1f1fe3051de10&secret=7b14195be09a4db9a8ab38afd4aa9fd7&js_code=0810dpha1NzsmN1P2Qha15n9ha10dph-&grant_type=authorization_code"
	resp, err := client2.Get(strUrl)
	fmt.Printf("%x\n", &client2)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
}
