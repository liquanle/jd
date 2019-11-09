// get project main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	_ "strings"
	"time"
)

func good(w http.ResponseWriter, r *http.Request) {
	expiration := time.Now()
	expiration = expiration.AddDate(1, 0, 0)
	c1 := http.Cookie{
		Name:    "lql",
		Value:   "332",
		Path:    "/",
		Domain:  "localhost",
		Expires: expiration,
	}

	w.Write([]byte("http://mp.weixin.qq.com"))
	//w.Header().Set("Location", "http://mp.weixin.qq.com")
	w.Header().Set("Set-Cookie", c1.String())
	//w.Header().Set("set-Cookie", c1.String())
	//http.SetCookie(w, &c1)
	//w.WriteHeader(200)
}

func tt(w http.ResponseWriter, r *http.Request) {

	expiration := time.Now()
	expiration = expiration.AddDate(1, 0, 0)

	c1 := http.Cookie{Name: "name", Value: "liquanle", Expires: expiration}
	c2 := http.Cookie{Name: "sex", Value: "male", Expires: expiration}
	c3 := http.Cookie{Name: "mark", Value: "98", Expires: expiration}

	http.SetCookie(w, &c1)

	c4 := http.Cookie{
		Name:    "address",
		Value:   "baoding",
		Expires: expiration,
	}

	c5 := http.Cookie{
		Name:     "role",
		Value:    "manager",
		HttpOnly: true,
	}

	// c3 := http.Cookie{
	// 	Name:    "mark",
	// 	Value:   "96",
	// 	Path:    "/",
	// 	Domain:  "localhost",
	// 	Expires: expiration,
	// }

	//w.Header().Set("Set-cookie", c2.String())
	w.Header().Add("Set-cookie", c2.String())
	w.Header().Add("Set-cookie", c3.String())
	w.Header().Add("Set-cookie", c4.String())
	w.Header().Add("Set-cookie", c5.String())

	w.Write([]byte("tt"))
}

func two(w http.ResponseWriter, r *http.Request) {

	expiration := time.Now()
	expiration = expiration.AddDate(1, 0, 0)

	c1 := http.Cookie{Name: "name", Value: "liquanle", Expires: expiration}

	w.Header().Set("Set-cookie", c1.String())

}

//使用两个setcookie
func ss(w http.ResponseWriter, r *http.Request) {

	expiration := time.Now()
	expiration = expiration.AddDate(1, 0, 0)

	c1 := http.Cookie{Name: "booname", Value: "c++", Expires: expiration}
	c2 := http.Cookie{Name: "author", Value: "luxun", Expires: expiration}
	c3 := http.Cookie{Name: "price", Value: "120", Expires: expiration}

	//可以使用SetCookie代码set与add,比较方便
	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)
	http.SetCookie(w, &c3)
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	for _, c := range r.Cookies() {
		nm := c.Name + "\t" + c.Value
		fmt.Println(nm)
	}

}

func main() {
	fmt.Println("开始")
	http.HandleFunc("/tt", tt)
	http.HandleFunc("/good", good)
	http.HandleFunc("/two", two)
	http.HandleFunc("/ss", ss)
	http.HandleFunc("/getCookie", getCookie)
	err := http.ListenAndServe(":8091", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
