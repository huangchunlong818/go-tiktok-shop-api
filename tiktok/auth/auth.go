package auth

import (
	"context"
	"encoding/json"
	"github.com/huangchunlong818/go-tiktok-shop-api/tiktok/common/common"
	"github.com/huangchunlong818/go-tiktok-shop-api/tiktok/common/config"
)

//tiktok shop 授权

type TiktokShopAuth struct {
	config                   *config.Config
	*common.TiktokShopCommon // 嵌入 common.TiktokShopCommon
}

var newServer *TiktokShopAuth

type AuthClientInterface interface {
	ReloadToken(ctx context.Context, refreshToken string) GetTokenByAuthCodeRsp
	GetTokenByAuthCode(ctx context.Context, authCode string) GetTokenByAuthCodeRsp
	GetAuthUrl(country string) string
	GetTokenByAuthCodeApi() common.GetApiConfig
	ReloadTokenUrl() common.GetApiConfig
}

// getNewService 是一个私有函数，用于返回 tiktokShopAuths 实例
func GetNewService(config *config.Config) AuthClientInterface {
	if newServer == nil {
		newServer = &TiktokShopAuth{
			config:           config,
			TiktokShopCommon: common.GetNewService(config),
		}
	}
	return newServer
}

// 根据reftoken 刷新令牌
func (a *TiktokShopAuth) ReloadToken(ctx context.Context, refreshToken string) GetTokenByAuthCodeRsp {
	result := GetTokenByAuthCodeRsp{}
	if refreshToken == "" {
		result.Code = common.ErrCode
		result.Message = "refresh_token cannot be empty"
		return result
	}

	query := map[string]string{
		"refresh_token": refreshToken,
		"grant_type":    "refresh_token",
		"app_secret":    a.config.App.Secret,
	}
	//请求接口
	r := a.SendTiktokApi(ctx, a.ReloadTokenUrl(), query, nil, nil)
	result = GetTokenByAuthCodeRsp{
		Code:     r.Code,
		Message:  r.Message,
		HttpCode: r.HttpCode,
	}
	if !a.IsSuccess(r) {
		return result
	}

	//解析数据
	err := json.Unmarshal(r.Data, &result)
	if err != nil {
		r.Code = common.ErrCode
		r.Message = "ReloadToken response error " + err.Error()
		return result
	}

	return result
}

// 根据授权码获取token和 reftoken
func (a *TiktokShopAuth) GetTokenByAuthCode(ctx context.Context, authCode string) GetTokenByAuthCodeRsp {
	result := GetTokenByAuthCodeRsp{}
	if authCode == "" {
		result.Code = common.ErrCode
		result.Message = "authCode cannot be empty"
		return result
	}

	query := map[string]string{
		"auth_code":  authCode,
		"grant_type": "authorized_code",
		"app_secret": a.config.App.Secret,
	}
	//请求接口
	r := a.SendTiktokApi(ctx, a.GetTokenByAuthCodeApi(), query, nil, nil)
	result = GetTokenByAuthCodeRsp{
		Code:     r.Code,
		Message:  r.Message,
		HttpCode: r.HttpCode,
	}
	if !a.IsSuccess(r) {
		return result
	}

	//解析数据
	err := json.Unmarshal(r.Data, &result)
	if err != nil {
		r.Code = common.ErrCode
		r.Message = "GetTokenByAuthCode response error " + err.Error()
		return result
	}

	return result
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
func (a *TiktokShopAuth) GetTokenByAuthCodeApi() common.GetApiConfig {
	api := "/api/v2/token/get" //请求API PATH

	return common.GetApiConfig{
		ContentType: "application/json",           //请求头content-type 类型
		Method:      "get",                        //请求方法类型
		Api:         api,                          //请求API PATH地址不带域名
		FullApi:     a.config.AuthApiDomain + api, //请求的API 完整地址，带域名
	}
}

// 获取刷新token地址
func (a *TiktokShopAuth) ReloadTokenUrl() common.GetApiConfig {
	api := "/api/v2/token/refresh" //请求API PATH

	return common.GetApiConfig{
		ContentType: "application/json",           //请求头content-type 类型
		Method:      "get",                        //请求方法类型
		Api:         api,                          //请求API PATH地址不带域名
		FullApi:     a.config.AuthApiDomain + api, //请求的API 完整地址，带域名
	}
}
