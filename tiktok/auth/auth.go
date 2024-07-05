package auth

import (
	"context"
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
	}
	//请求接口
	r := a.SendTiktokApi(ctx, a.ReloadTokenUrl(), query, nil)
	result = GetTokenByAuthCodeRsp{
		Code:     r.Code,
		Message:  r.Message,
		HttpCode: r.HttpCode,
	}
	if !a.IsSuccess(r) {
		return result
	}

	return a.DoAuthData(r, result)
}

// 处理授权和刷新token返回结果
func (a *TiktokShopAuth) DoAuthData(r common.ComApiRsp, result GetTokenByAuthCodeRsp) GetTokenByAuthCodeRsp {
	data, err := a.ChangeAnyToStringSlice(r.Data["granted_scopes"])
	if err != nil {
		result.Code = common.ErrCode
		result.Message = err.Error()
		return result
	}
	result.Data = GetTokenByAuthCodeData{
		AccessToken:          r.Data["access_token"].(string),
		AccessTokenExpireIn:  int64(r.Data["access_token_expire_in"].(float64)),
		RefreshToken:         r.Data["refresh_token"].(string),
		RefreshTokenExpireIn: int64(r.Data["refresh_token_expire_in"].(float64)),
		OpenID:               r.Data["open_id"].(string),
		SellerName:           r.Data["seller_name"].(string),
		SellerBaseRegion:     r.Data["seller_base_region"].(string),
		UserType:             int(r.Data["user_type"].(float64)),
		GrantedScopes:        data,
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
	}
	//请求接口
	r := a.SendTiktokApi(ctx, a.GetTokenByAuthCodeApi(), query, nil)
	result = GetTokenByAuthCodeRsp{
		Code:     r.Code,
		Message:  r.Message,
		HttpCode: r.HttpCode,
	}
	if !a.IsSuccess(r) {
		return result
	}

	return a.DoAuthData(r, result)
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
