// pages/index/runrec.js
//获取应用实例
const app = getApp()

Page({

  /**
   * 页面的初始数据
   */
  data: {
    rrdatas:{}
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (options) {

  },

  /**
   * 生命周期函数--监听页面初次渲染完成
   */
  onReady: function () {

  },

  /**
   * 生命周期函数--监听页面显示
   */
  onShow: function () {
    let oid = app.globalData.openid;
    oid = 'oTteQ4pCKlyScD7JwEzxdf9MTTtQ'
    var strUrl = `https://ziweitec.com/queryRunRec?openid=${oid}`;
    wx.request({
      url: strUrl,
      method:"GET",
      success:res=>{
        this.setData({
          rrdatas:res.data
        })
        //var a = base64_decode(res.data)
        //var a = eval('(' + res.data + ')');
        for (var i = 0; i < res.data.length; i++){
          console.log("Mile" + res.data[i]["Mile"]);
        }
      }
    })
  },

  /**
   * 生命周期函数--监听页面隐藏
   */
  onHide: function () {

  },

  /**
   * 生命周期函数--监听页面卸载
   */
  onUnload: function () {

  },

  /**
   * 页面相关事件处理函数--监听用户下拉动作
   */
  onPullDownRefresh: function () {

  },

  /**
   * 页面上拉触底事件的处理函数
   */
  onReachBottom: function () {

  },

  /**
   * 用户点击右上角分享
   */
  onShareAppMessage: function () {

  }
})