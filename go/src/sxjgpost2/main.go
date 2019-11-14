// sxjgpost project main.go
package main

import (
	"bytes"
	"fmt"

	"net/http"
	"time"
)

func main() {
	fmt.Println("sxjgpost program!")

	sxclient := &http.Client{}
	curtime := time.Now().Format("2006-01-02 15:04:05")

	jgids := []string{"1904230001", "1904230003", "1904230004", "1904230005"}

	for _, ele := range jgids {
		jsonfmt := `{flag:1 data:{"dev":{"id":"%s","deviceVersion":"V2.1","freq1":"1440","freq2":"1440","connectFreq":"1440","softVersion":"V2.04","energy":"100","signal":"51","tem":"9","clockStatus":"0","save_con":"13","read_con":"11","IO1_Status":"1","IO2_Status":"0"},"datas":[{"channel":"3","value":"0","time":"%s","status":"0"},{"channel":"4","value":"0","time":"%s","status":"0"}]}}`
		jsonStr := fmt.Sprintf(jsonfmt, ele, curtime, curtime)
		fmt.Println(jsonStr)

		jsoncmd := []byte(jsonStr)
		url := "http://112.31.239.184:8100/iot/HDDTService"
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsoncmd))
		//req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		resp, err := sxclient.Do(req)
		if err != nil {
			// handle error
			fmt.Println(resp)
		}

		defer resp.Body.Close()
		//body, err := ioutil.ReadAll(resp.Body)
		// if err != nil {
		// 	fmt.Println(err)
		// }

		statuscode := resp.StatusCode

		hea := resp.Header

		//fmt.Println(string(body))
		fmt.Println(statuscode)
		fmt.Println(hea)
		fmt.Println("__________________________")
	}
}
