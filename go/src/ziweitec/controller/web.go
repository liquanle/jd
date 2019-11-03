/*
 * web.go
 * author: lql
 * email: 6188806#qq.com
 * date: 2019/11/3
 */
package controller

import (
	"../lib"
)

type Web struct {
	*lib.WebEngine
}

var web = &Web{WebEngine:lib.Web}

func init() {
}

//这是用来初始化的
func Start()  {
	web.Start()
}