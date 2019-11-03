//app.js
let app = require("./sim.js/index.js")

App(Object.assign(app,{
  onLaunch: function () {
    // 展示本地存储能力
    var logs = wx.getStorageSync('logs') || []
    logs.unshift(Date.now())
    wx.setStorageSync('logs', logs)

    // 登录
    wx.login({
      success: res => {
        // 发送 res.code 到后台换取 openId, sessionKey, unionId
     
        var APPID = 'wxa1b1f1fe3051de10'
        var secret = '7b14195be09a4db9a8ab38afd4aa9fd7'

        var strUrl = `https://ziweitec.com/getOpenid?appid=${APPID}&secret=${secret}&js_code=${res.code}&grant_type=authorization_code`
        wx.request({
          url: strUrl,
          success: res=> {
            this.globalData.openid = res.data.openid
            console.log(res.data.openid)
            console.log("wx.login")
            wx.navigateTo({
              url: './main',
            })
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
              console.info("getUserInfo", res)
            }
          })
        }else{
          // wx.showModal({
          //   title: '用户未授权',
          //   content: '拒绝授权将不能体验小程序完整功能，点击确定开启授权',
          //   success: (res) => {
          //     console.log(res)
          //     if (res.confirm) {
          //       wx.navigateTo({
          //         url: './index',
          //       })
          //       wx.authorize({
          //         scope: 'scope.userInfo',
          //         scope: 'scope.userLocation'
          //       })
          //       //wx.openSetting({})
          //     }
          //   }
          // })
        }
      },
      fail: res => {
        console.info("getSetting", res)
      },
      complete: res => {
        console.info("getsetting_complete", res)
      }
    })
  },
  globalData: {
    userInfo: null,
    openid:''
  }
}))