/*
 * weapp-controller.go
 * author: lql
 * email: 6188806#qq.com
 * date: 2019/11/3
 */
package lib

import (
	"gopkg.in/kataras/iris.v6"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type (
	WeappUser struct {
		Model  `xorm:"extends"`
		OpenID    string `xorm:"char(32)" json:"openId"`
		UnionID   string `xorm:"char(32)" json:"unionId"`
		Nickname  string `xorm:"char(50)" json:"nickName"`
		Gender    int    `xorm:"tinyint" json:"gender"`
		City      string `xorm:"char(20)" json:"city"`
		Province  string `xorm:"char(20)" json:"province"`
		Country   string `xorm:"char(20)" json:"country"`
		AvatarURL string `xorm:"tinytext" json:"avatarUrl"`
		Language  string `xorm:"char(10)" json:"language"`
		Watermark struct {
					  Timestamp int64  `json:"timestamp"`
					  AppID     string `json:"appid"`
				  } `xorm:"json" json:"watermark"`
	}
	// 微信json2session返回的数据结构
	WeappSession struct {
		SessionKey string `json:"session_key"`
		ExpiresIn  int `json:"expires_in"`
		Openid     string `json:"openid"`
	}
	WeappController struct {
		Web              *WebEngine
		AppId, AppSecret string
	}
)

func (this *WeappController) Init() {
	this.Web.Get("/weapp/login", func(c *iris.Context) {
		var r Result
		var code, encryptedData, iv = this.getLoginArgsFromHeader(c.Request.Header)
		var session = this.retrieveSession(code)
		//Debug("weapp login args",code, encryptedData, iv, session)

		if user, err := this.decryptEncryptedData(encryptedData, iv, session.SessionKey); err == nil {
			if exist, err := this.Web.DB.Where("open_id = ?", user.OpenID).NoAutoCondition().Get(&user); err == nil {
				if !exist {
					this.Web.DB.InsertOne(&user)
				}
				r.Data = user
				r.Code = 1
			}else{
				r.Code = -2
			}
		}else{
			r.Code = -1
		}

		c.JSON(200, r)
	})
}

// 从请求头中获取小程序登陆的参数
func (this *WeappController) getLoginArgsFromHeader(header http.Header) (code, encryptedData, iv string) {
	code = header.Get("X-WX-Code")
	encryptedData = header.Get("X-WX-Encrypted-Data")
	iv = header.Get("X-WX-IV")
	return
}

// 拉取小程序登陆必须的session,依code拉取
func (this *WeappController) retrieveSession(code string) (weappSession WeappSession) {
	var url = fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", this.AppId, this.AppSecret, code);
	var res = HttpGet(url, nil, nil)
	//Debug("retrieveSession res",res)
	//map[session_key:vroPndZ1jeWdxkVAo5V05A== expires_in:7200 openid:o-hrq0EVYOTJHX9MWqk-LF-_KL0o]
	ParseJson(&weappSession, res)
	return
}

// 解密小程序加密信息
func (this *WeappController) decryptEncryptedData(encryptedData, iv, sessionKey string) (user WeappUser, err error) {
	aesKey, err := base64.StdEncoding.DecodeString(sessionKey)
	if err != nil {
		return
	}
	cipherText, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return
	}
	ivBytes, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return
	}
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return
	}
	mode := cipher.NewCBCDecrypter(block, ivBytes)
	mode.CryptBlocks(cipherText, cipherText)
	cipherText, err = this.pkcs7Unpad(cipherText, block.BlockSize())
	if err != nil {
		return
	}
	err = json.Unmarshal(cipherText, &user)
	if err != nil {
		return
	}
	if user.Watermark.AppID != this.AppId {
		err = errors.New("app id not match")
		return
	}
	return
}

// pkcs7Unpad returns slice of the original data without padding
func (this *WeappController) pkcs7Unpad(data []byte, blockSize int) ([]byte, error) {
	if blockSize <= 0 {
		return nil, errors.New("invalid block size")
	}
	if len(data) % blockSize != 0 || len(data) == 0 {
		return nil, errors.New("invalid PKCS7 data")
	}
	c := data[len(data) - 1]
	n := int(c)
	if n == 0 || n > len(data) {
		return nil, errors.New("invalid padding on input")
	}
	for i := 0; i < n; i++ {
		if data[len(data) - n + i] != c {
			return nil, errors.New("invalid padding on input")
		}
	}
	return data[:len(data) - n], nil
}