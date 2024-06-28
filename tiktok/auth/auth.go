package auth

import (
	"errors"
	"github.com/go-resty/resty/v2"
	"github.com/huangchunlong818/go-tiktok-shop-api/tiktok/common/common"
	"github.com/huangchunlong818/go-tiktok-shop-api/tiktok/common/config"
	"strconv"
	"time"
)

//tiktok shop 授权

type TiktokShopAuth struct {
	config *config.Config
}

var newServer *TiktokShopAuth

type AuthClientInterface interface {
	ReloadToken(refreshToken string) (result GetTokenByAuthCodeData, err error)
	GetTokenByAuthCode(authCode string) (result GetTokenByAuthCodeData, err error)
	GetAuthUrl(country string) string
	GetTokenByAuthCodeApi() string
	ReloadTokenUrl() string
}

// getNewService 是一个私有函数，用于返回 tiktokShopAuths 实例
func GetNewService(config *config.Config) AuthClientInterface {
	if newServer == nil {
		newServer = &TiktokShopAuth{
			config: config,
		}
	}
	return newServer
}

// 根据reftoken 刷新令牌
func (a *TiktokShopAuth) ReloadToken(refreshToken string) (result GetTokenByAuthCodeData, err error) {
	if refreshToken == "" {
		return result, errors.New("refresh_token cannot be empty")
	}

	return a.DoToken("refresh_token", refreshToken, a.ReloadTokenUrl(), "refresh_token")
}

// 根据授权码获取token和 reftoken
func (a *TiktokShopAuth) GetTokenByAuthCode(authCode string) (result GetTokenByAuthCodeData, err error) {
	if authCode == "" {
		return result, errors.New("auth_code cannot be empty")
	}

	return a.DoToken("auth_code", authCode, a.GetTokenByAuthCodeApi(), "authorized_code")
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
		SetQueryParam("app_key", a.config.App.AppKey).
		SetQueryParam("app_secret", a.config.App.Secret).
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

// 获取授权基础连接
func (a *TiktokShopAuth) GetAuthUrl(country string) string {
	if country == "us" {
		//美国
		return a.config.UsAuthUrl
	}
	//非美国
	return a.config.OtherAuthUrl
}

// 获取token和reftoken API地址
func (a *TiktokShopAuth) GetTokenByAuthCodeApi() string {
	return a.config.AuthApiDomain + "/api/v2/token/get"
}

// 获取刷新token地址
func (a *TiktokShopAuth) ReloadTokenUrl() string {
	return a.config.AuthApiDomain + "/api/v2/token/refresh"
}
