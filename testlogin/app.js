//app.js
App({
  onLaunch: function () {
    // 展示本地存储能力
    var logs = wx.getStorageSync('logs') || []
    logs.unshift(Date.now())
    wx.setStorageSync('logs', logs)

    // 登录
    wx.login({
      success: res => {
        // 发送 res.code 到后台换取 openId, sessionKey, unionId
        const logger = wx.getLogManager({ level: 1 })
        logger.log({ str: res.code }, 'basic log', 100, [1, 2, 3])
        logger.info({ str: res.code  }, 'info log', 100, [1, 2, 3])
        logger.debug({ str: res.code  }, 'debug log', 100, [1, 2, 3])
        logger.warn({ str: res.code  }, 'warn log', 100, [1, 2, 3])

        var APPID = 'wxa1b1f1fe3051de10'
        var secret = '7b14195be09a4db9a8ab38afd4aa9fd7'

        var strUrl = `https://ziweitec.com/getOpenid?appid=${APPID}&secret=${secret}&js_code=${res.code}&grant_type=authorization_code`
        wx.request({
          url: strUrl,
          success: function (res) {
            console.log(res.data)
            console.info("这是info")
            console.error("这是error")
            console.warn("这是warn")
          }
        })
        
      }
    })
    // 获取用户信息
    wx.getSetting({
      success: res => {
        if (res.authSetting['scope.userInfo']) {
          // 已经授权，可以直接调用 getUserInfo 获取头像昵称，不会弹框
          wx.getUserInfo({
            success: res => {
              // 可以将 res 发送给后台解码出 unionId
              this.globalData.userInfo = res.userInfo

              // 由于 getUserInfo 是网络请求，可能会在 Page.onLoad 之后才返回
              // 所以此处加入 callback 以防止这种情况
              if (this.userInfoReadyCallback) {
                this.userInfoReadyCallback(res)
              }
            },
            fail: res => {
              log.info("getUserInfo", res)
            }
          })
        }
      },
      fail: res => {
        log.info("getSetting", res)
      },
      complete :res=>{
        log.info("getsetting_complete", res)
      }
    })
  },
  globalData: {
    userInfo: null
  }
})