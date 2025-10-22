# 代码介绍
用于获取微信小程序AccessToken，登陆获取openid，获取用户手机号

## 安装
```bash
go get github.com/xiongxiubo/wx_gdk_go
```

## 使用
```go
import (
    "github.com/xiongxiubo/wx_gdk_go"
)
// 初始化
wx := wx_gdk_go.CreateClient("appid", "appsecret")

// 获取AccessToken
accessToken, err := wx.GetAccessToken()
if err != nil {
    log.Fatal(err)
}

// 获取openid
openid, err := wx.GetOpenid("code")
if err != nil {
    log.Fatal(err)
}

// 获取用户手机号
phone, err := wx.GetPhoneNumber("AccessToken","code")
if err != nil {
    log.Fatal(err)
}
```
