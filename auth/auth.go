package auth

import (
	"errors"
	"strconv"
	tiktok "tiktok-shop-api"
	"tiktok-shop-api/tiktok/common"
	"time"
)

//tiktok shop 授权

type TiktokShopAuth struct {
}

var tiktokShopAuth *TiktokShopAuth

// 获取实例
func GetNewService() *TiktokShopAuth {
	if tiktokShopAuth == nil {
		tiktokShopAuth = &TiktokShopAuth{}
	}
	return tiktokShopAuth
}

// 根据reftoken 刷新令牌
func (a *TiktokShopAuth) ReloadToken(refreshToken string) (result GetTokenByAuthCodeData, err error) {
	if refreshToken == "" {
		return result, errors.New("refresh_token cannot be empty")
	}

	return a.DoToken("refresh_token", refreshToken, ReloadToken(), "refresh_token")
}

// 根据授权码获取token和 reftoken
func (a *TiktokShopAuth) GetTokenByAuthCode(authCode string) (result GetTokenByAuthCodeData, err error) {
	if authCode == "" {
		return result, errors.New("auth_code cannot be empty")
	}

	return a.DoToken("auth_code", authCode, GetTokenByAuthCodeApi(), "authorized_code")
}

var client = resty.New()

// 操作token
func (a *TiktokShopAuth) DoToken(paramKey string, paramValue string, api string, grantType string) (result GetTokenByAuthCodeData, err error) {
	//定义响应体
	var (
		res    GetTokenByAuthCodeRsp
		errRsp common.ComErrorResponse
	)

	//请求tiktok
	client.SetTimeout(10 * time.Second)
	resp, err := client.R().
		SetQueryParam("app_key", tiktok.AppKey()).
		SetQueryParam("app_secret", tiktok.Secret()).
		SetQueryParam(paramKey, paramValue).
		SetQueryParam("grant_type", grantType).
		SetResult(&res).
		SetError(&errRsp).
		Get(api)
	if err != nil {
		return result, err
	}
	if resp.IsSuccess() {
		//这里也有可能是失败的
		if res.Code > 0 {
			err = errors.New("错误信息：" + res.Message + "，错误代码：" + strconv.FormatInt(int64(res.Code), 10))
		} else {
			result = res.Data
		}
	} else {
		err = errors.New("错误信息：" + errRsp.Message + "，错误代码：" + strconv.FormatInt(int64(errRsp.Code), 10))
	}

	return result, err
}
