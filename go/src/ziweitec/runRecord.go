// runRecord
package main

//跑步记录结构,字段首字母必须大写，否则导不出去
type runRecord struct {
	UserID string
	Openid string
	Mile   float32
	Score  int
	Time   string
	Image  string
}
