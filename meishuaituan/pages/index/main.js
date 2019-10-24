// pages/index/main.js
//获取应用实例
const app = getApp()

Page({
  data: {
    userID:{},
    userMile:{},
    tiptext: '请跑团的每位会员成员填写会员编号及日总跑量，每天12点前完成一日跑量填写，感谢配合！',
    userInfo: {},
    hasUserInfo: false,
    canIUse: wx.canIUse('button.open-type.getUserInfo')
  },
  //事件处理函数
  bindViewTap: function () {
    wx.navigateTo({
      url: '../logs/logs'
    })
  },
  onLoad: function () {
    if (app.globalData.userInfo) {
      this.setData({
        userInfo: app.globalData.userInfo,
        hasUserInfo: true
      })
    } else if (this.data.canIUse) {
      // 由于 getUserInfo 是网络请求，可能会在 Page.onLoad 之后才返回
      // 所以此处加入 callback 以防止这种情况
      app.userInfoReadyCallback = res => {
        this.setData({
          userInfo: res.userInfo,
          hasUserInfo: true
        })
      }
    } else {
      // 在没有 open-type=getUserInfo 版本的兼容处理
      wx.getUserInfo({
        success: res => {
          app.globalData.userInfo = res.userInfo
          this.setData({
            userInfo: res.userInfo,
            hasUserInfo: true
          })
        }
      })
    }
  },
  getUserInfo: function (e) {
    console.log(e)
    app.globalData.userInfo = e.detail.userInfo
    this.setData({
      userInfo: e.detail.userInfo,
      hasUserInfo: true
    })
  },

  //获取用户输入的用户名
  userIDInput: function (e) {
    this.setData({
      userID: e.detail.value
    })
  },

  //获取用户输入的公里数
  userMileInput: function (e) {
    this.setData({
      userMile: e.detail.value
    })
  },

  btnCommit:function()
  {
    //增加访问request接口
    let app = getApp()
    
    wx.showLoading({ title: '加载中' })

    app.request(`http://liquanle.com:8099/liquanle?no=${this.data.userID}&mile=${this.data.userMile}`)
      .then(res => {
        console.log("res", res)

      }).catch(err => {
        console.error(err)
      }).finally(() => {
        wx.hideLoading()
      })

    wx.showToast({
      title: this.data.userID + '号成功打卡' + this.data.userMile +'公里！',
      icon: 'none',
      duration: 3000//持续的时间
    })
  }
})
