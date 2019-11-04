// pages/index/main.js
//获取应用实例
const app = getApp()

function json2Form(json) {
  var str = [];
  for (var p in json) {
    str.push(encodeURIComponent(p) + "=" + encodeURIComponent(json[p]));
  }
  return str.join("&");
}

Page({
  data: {
    userID:'',
    userMile:{},
    nickname:{},
    image:{},
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
        openid:app.globalData.openid,
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
  onShow: function () {
    //初始化会员编号
    var strQueryUrL = `https://ziweitec.com/queryMember?openid=${app.globalData.openid}`
    wx.request({
      url: strQueryUrL,
      success: res => {
        //this.globalData.openid = res.data.openid
        this.setData({
          userID: res.data,
          userMile:''
        })
      }
    })
  },
  onHide: function () {
    console.log("onHide: function () ")
    this.setData({
      userID:'',
      userMile: ''
    })

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
    //GET方法
    /*app.request(`https://ziweitec.com/liquanle?no=${this.data.userID}&mile=${this.data.userMile}`)
      .then(res => {
        console.log("res", res)
      }).catch(err => {
        console.error(err)
      }).finally(() => {
        wx.hideLoading()
      })*/
    //POST方式

    if (app.globalData.userInfo == null){
      this.setData({
        nickname : "unknown",
        image : "unknown"
      })
    }else{
      this.setData({
        nickname: app.globalData.userInfo.nickName,
        image: app.globalData.userInfo.avatarUrl
      })
    }

    let dataval = {
      userID: this.data.userID,
      mile: this.data.userMile,
      openid: app.globalData.openid,
      nickname: this.data.nickname,
      image: this.data.image
    };
    wx.request({
      url: 'https://ziweitec.com/liquanle',
      method: 'POST',
      data: dataval,
      header: {
        "Content-Type": "application/x-www-form-urlencoded"
      },
    })

    wx.showToast({
      title: this.data.userID + '号成功打卡' + this.data.userMile +'公里！',
      icon: 'none',
      duration: 6000//持续的时间
    })
  }
})
