/*
 * web-engine.go
 * author: lql
 * email: 6188806#qq.com
 * date: 2019/11/3
 */
package lib

import (
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
	"gopkg.in/kataras/iris.v6/adaptors/cors"
	"fmt"
	"gopkg.in/kataras/iris.v6/adaptors/view"
	"gopkg.in/kataras/iris.v6/adaptors/sessions"
	"os"
	"strconv"
)

type WebEngine struct {
	*iris.Framework
	Config Config
	DB     *DB
	Weapp *WeappController
}

var Web = &WebEngine{}

func init() {
	fmt.Println("web engine init")
	Web.Init()
}

// 获取当前登陆的微信小程序用户
func (this *WebEngine) GetWeappUser(c *iris.Context) (u WeappUser) {
	const USER_KEY = "WEAPPUSER"
	var id = c.RequestHeader("X-WX-Id")
	Debug("id",id)

	if c.Get(USER_KEY) != nil {
		u = c.Get(USER_KEY).(WeappUser)
	}else if id, err := strconv.ParseInt(id,10,64); err == nil {
		this.DB.ID(id).Get(&u)
		c.Set(USER_KEY,u)
	}

	return u
}

// 从配置文件config.ini读取相关的配置项
func (this *WebEngine) Init() {
	var confileFile = "./config.ini"
	if _, err := os.Stat(confileFile); os.IsNotExist(err) {
		fmt.Println("config.ini文件未找到")
		panic(err)
	}

	// 无论是以gin热编译,还是直接启动,都能从程序执行的当前目录找到config.ini
	if err := ReadTomlConfig(&this.Config, confileFile); err == nil {
		this.Log("已读取配置文件")
		//this.Log("config", ToMapObject(this.Config))
	}else{
		panic(err)
	}

	this.Framework = iris.New(iris.Configuration{Gzip: true, Charset: "UTF-8"})
	this.Framework.Adapt(httprouter.New())
	if this.Config.Debug {
		this.Framework.Adapt(iris.DevLogger())
	}

	// 开始CORS跨域支持
	if this.Config.Cors.Enable {
		this.Framework.Adapt(cors.New(cors.Options{
			AllowedOrigins: []string{"*"},
			AllowCredentials:true,
		}))
	}

	// 初始化数据库
	if this.Config.Mysql.Enable {
		this.DB = &DB{ShowSql:this.Config.Debug,DataSource:this.Config.Mysql.DataSource, DriverName:DB_DRIVER_MYSQL}
		this.DB.Init()
	}else if this.Config.Sqlite3.Enable {
		Debug("init sqlite3 db")
		this.DB = &DB{ShowSql:this.Config.Debug,DataSource:this.Config.Sqlite3.Filepath, DriverName:DB_DRIVER_SQLITE3}
		this.DB.Init()
	}

	// 启用小程序用户自动登陆
	if this.DB != nil && this.Config.Weapp.Enable {
		if exist,_ := this.DB.IsTableExist(new(WeappUser)); !exist {
			this.DB.Sync2(new(WeappUser))
		}
		this.Weapp = &WeappController{Web:this,AppId:this.Config.Weapp.AppId,AppSecret:this.Config.Weapp.AppSecret}
		this.Weapp.Init()
	}

	// 初始化模板
	if this.Config.Html.Enable {
		this.Framework.StaticWeb("/static", this.Config.Html.StaticDir)
		var djangoAdapt = view.Django(this.Config.Html.TemplateDir, ".html")
		//原来的Config.IsDevelopment转移到了这里,设置为true,表示模板文件热加载
		djangoAdapt.Reload(this.Config.Debug)
		this.Framework.Adapt(djangoAdapt)
	}

	// 初始化session
	if this.Config.Session.Enable {
		var sessionAdapt = sessions.New(sessions.Config{Cookie: this.Config.Session.Key})
		this.Framework.Adapt(sessionAdapt)
	}

	this.Any("/hi", func(c *iris.Context) {
		if user := this.GetWeappUser(c); user.ID > 0 {
			c.WriteString(fmt.Sprintf("hi,%s",user.Nickname))
			return
		}
		c.WriteString("hi,lib.go")
	})
}

func (this *WebEngine) Start() {
	this.Framework.Listen(this.Config.Addr)
}

func (this *WebEngine) Log(v ...interface{}) {
	if this.Config.Debug {
		Debug(v...)
	}
}