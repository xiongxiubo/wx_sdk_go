package wxgdkgo

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type Client struct {
	AppID     string
	AppSecret string
}

func CreateClient(appID, appSecret string) *Client {
	return &Client{
		AppID:     appID,
		AppSecret: appSecret,
	}
}

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

func (c *Client) GetAccessToken() (AccessTokenResponse, error) {
	// 实现获取access_token的逻辑
	url := "https://api.weixin.qq.com/cgi-bin/token"
	// 发送GET请求到url，携带params参数
	// 解析响应JSON，提取access_token字段
	// 返回access_token和nil错误
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return AccessTokenResponse{}, err
	}
	q := req.URL.Query()
	q.Add("grant_type", "client_credential")
	q.Add("appid", c.AppID)
	q.Add("secret", c.AppSecret)
	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req)
	if err != nil {
		return AccessTokenResponse{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return AccessTokenResponse{}, err
	}
	var result AccessTokenResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return AccessTokenResponse{}, err
	}

	return result, nil
}

type WxLoginResponse struct {
	UnionID    string `json:"unionid"`
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
}

func (c *Client) WxLogin(js_code string) (WxLoginResponse, error) {
	// 实现微信登录的逻辑
	url := "https://api.weixin.qq.com/sns/jscode2session"
	// 发送GET请求到url，携带params参数
	// 解析响应JSON，提取access_token字段
	// 返回access_token和nil错误
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return WxLoginResponse{}, err
	}
	q := req.URL.Query()
	q.Add("appid", c.AppID)
	q.Add("secret", c.AppSecret)
	q.Add("js_code", js_code)
	q.Add("grant_type", "authorization_code")
	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req)
	if err != nil {
		return WxLoginResponse{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return WxLoginResponse{}, err
	}
	var result WxLoginResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return WxLoginResponse{}, err
	}
	return result, nil
}

// 获取用户手机号
type WxGetPhoneNumberResponse struct {
	PhoneInfo string `json:"phone_info"`
	ErrCode   int    `json:"errcode"`
	ErrMsg    string `json:"errmsg"`
}
type WxPhoneInfo struct {
	PhoneNumber     string `json:"phoneNumber"`
	CountryCode     string `json:"countryCode"`
	PurePhoneNumber string `json:"purePhoneNumber"`
}

func (c *Client) WxGetPhoneNumber(accessToken, code string) (WxGetPhoneNumberResponse, error) {
	// 实现获取用户手机号的逻辑
	url := "https://api.weixin.qq.com/wxa/business/getuserphonenumber"
	// 发送POST请求到url，携带params参数
	// 解析响应JSON，提取phone_info字段
	// 返回phone_info和nil错误
	params := map[string]string{
		"code": code,
	}
	paramsBytes, err := json.Marshal(params)
	if err != nil {
		return WxGetPhoneNumberResponse{}, err
	}
	client := http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(paramsBytes))
	if err != nil {
		return WxGetPhoneNumberResponse{}, err
	}
	q := req.URL.Query()
	q.Add("access_token", accessToken)
	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req)
	if err != nil {
		return WxGetPhoneNumberResponse{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return WxGetPhoneNumberResponse{}, err
	}
	var result WxGetPhoneNumberResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return WxGetPhoneNumberResponse{}, err
	}
	return result, nil
}
