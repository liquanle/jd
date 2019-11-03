/*
 * 文件暂无名
 * author: liyi
 * email: 9830131#qq.com
 * date: 2017/7/12
 */
package controller

import (
	"../lib"
	"gopkg.in/kataras/iris.v6"
	"fmt"
	"math"
)

type (
	History struct {
		sim.Model `xorm:"extends"`
		UserId int64 `json:"user_id"`
		Record        string `xorm:"char(50)" json:"record"`
	}
	HistoryPage struct {
		List       []History `json:"list"`
		Size       int `json:"size"`
		Page       int `json:"page"`   //自1开始计数
		Count      int `json:"count"`
		TotalPage  int `json:"total_page"`
	}
)

func init() {
	const URL_BASE = ""
	const MODEL_NAME = "history"

	if web.DB == nil {
		return
	}
	if web.Config.Debug {
		if exist,_ := web.DB.IsTableExist(MODEL_NAME); !exist {
			web.DB.Sync2(new(History))
		}
	}

	/*
	使用curl测试接口
	POST   /history     创建
	curl -X POST -d '{"title":"x"}' http://localhost:4000/history

	GET    /history/1 查看
	curl http://localhost:4000/history/1

	GET    /historys
	curl http://localhost:4000/historys

	DELETE /history/1 删除
	curl -X DELETE http://localhost:4000/history/1

	PUT    /history/1 更新
	curl -X PUT -d '{"title":"xxx"}' http://localhost:4000/history/1
	*/

	var sub = web.Router
	if URL_BASE != "" {
		sub = web.Party(URL_BASE)
	}
	// 拉取单条记录
	sub.Get(fmt.Sprintf("/%s/:id", MODEL_NAME), func(c *iris.Context) {
		var r sim.Result

		var id, err = c.ParamInt64("id")
		if err != nil {
			r.Code = -1
			r.Message = "id unvalid"
			c.JSON(200, r)
			return
		}

		var data History
		if _, err := web.DB.ID(id).Get(&data); err == nil {
			if data.ID > 0 {
				r.Code = 1
				r.Message = "success"
				r.Data = data
			}else{
				r.Code = -3
				r.Message = "not found"
			}
		}else{
			r.Code = -2
			r.Message = "select err"
		}

		c.JSON(200, r)
	})

	// 分页拉取记录,如果没有指定size,拉取所有限1000条以内
	sub.Get(fmt.Sprintf("/%ss", MODEL_NAME), func(c *iris.Context) {
		var r sim.Result
		var page,_ = c.URLParamInt("page")
		var size,_ = c.URLParamInt("size")

		if page == 0 {
			page = 1
		}
		if size == 0 {
			size = 1000
		}
		var data = HistoryPage{Page:page,Size:size}
		var offset = (data.Page - 1) * data.Size
		var my = web.GetWeappUser(c)

		if err := web.DB.Where("user_id = ?", my.ID).Desc("id").Limit(data.Size, offset).Find(&data.List); err == nil {
			if total, err := web.DB.Where("user_id = ?", my.ID).Count(new(History)); err == nil {
				data.Count = int(total)
				data.TotalPage = int(math.Ceil(float64(total) / float64(data.Size)))
			}else{
				sim.Debug("select count err", err.Error())
			}
			//for k,v := range data.List{
			//	web.DB.GetOneById(&v.User,v.UserId)
			//	data.List[k].User = v.User
			//}
			r.Code = 1
			r.Data = data
		}else{
			sim.Debug("select page err", err.Error())
		}

		c.JSON(200, r)
	})

	// 新增单条记录
	sub.Post(fmt.Sprintf("/%s", MODEL_NAME), func(c *iris.Context) {
		var r sim.Result

		var one History
		if err := c.ReadJSON(&one); err != nil {
			sim.Debug("read data err", err.Error())
			r.Code = -1
			r.Message = "read data err"
			c.JSON(200, r)
			return
		}

		var my = web.GetWeappUser(c)
		one.UserId = my.ID

		if affected, err := web.DB.Insert(&one); err == nil {
			if affected > 0 {
				r.Code = 1
				r.Data = one.ID
			}
		}else{
			sim.Debug("post new bean err",err.Error())
		}

		c.JSON(200, r)
	})

	// 删除单条记录
	sub.Delete(fmt.Sprintf("/%s/:id", MODEL_NAME), func(c *iris.Context) {
		var r sim.Result

		var id, err = c.ParamInt64("id")
		if err != nil {
			r.Code = -1
			r.Message = "id unvalid"
			c.JSON(200, r)
			return
		}

		var one History
		if _, err := web.DB.ID(id).Get(&one); err != nil {
			r.Code = -2
			r.Message = "not found"
			c.JSON(200, r)
			return
		}

		if affected, err := web.DB.Id(id).Delete(&one); err == nil {
			if affected > 0 {
				r.Code = 1
			}
		}

		c.JSON(200, r)
	})

	// 更新单条记录
	sub.Put(fmt.Sprintf("/%s/:id", MODEL_NAME), func(c *iris.Context) {
		var r sim.Result

		var id, err = c.ParamInt64("id")
		if err != nil {
			r.Code = -1
			r.Message = "id unvalid"
			c.JSON(200, r)
			return
		}

		var one History
		if _, err := web.DB.ID(id).Get(&one); err != nil {
			r.Code = -2
			r.Message = "not found"
			c.JSON(200, r)
			return
		}

		if err := c.ReadJSON(&one); err != nil {
			r.Code = -3
			r.Message = "read data err"
			c.JSON(200, r)
			return
		}

		if _, err := web.DB.Id(id).Update(&one); err == nil {
			//如果当前内容与原内容一样,affected会返回0
			r.Code = 1
			r.Data = one.ID
		}

		c.JSON(200, r)
	})
}