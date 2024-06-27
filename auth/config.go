package auth

import tiktok "tiktok-shop-api"

// 获取授权基础连接
func GetAuthUrl(country string) string {
	if country == "us" {
		//美国
		return tiktok.UsAuthUrl()
	}
	//非美国
	return tiktok.OtherAuthUrl()
}

// 根据用户授权码，获取token和reftoken API地址
func GetTokenByAuthCodeApi() string {
	return tiktok.AuthApiDomain() + "/api/v2/token/get"
}

// 根据用户授权码，获取token和reftoken API地址
func ReloadToken() string {
	return tiktok.AuthApiDomain() + "/api/v2/token/refresh"
}

// 整体相应结构
type GetTokenByAuthCodeRsp struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    GetTokenByAuthCodeData `json:"data"`
}

// 定义一个结构体来表示 data 部分的数据
type GetTokenByAuthCodeData struct {
	AccessToken          string   `json:"access_token"`
	AccessTokenExpireIn  int64    `json:"access_token_expire_in"`
	RefreshToken         string   `json:"refresh_token"`
	RefreshTokenExpireIn int64    `json:"refresh_token_expire_in"`
	OpenID               string   `json:"open_id"`
	SellerName           string   `json:"seller_name"`
	SellerBaseRegion     string   `json:"seller_base_region"`
	UserType             int      `json:"user_type"`
	GrantedScopes        []string `json:"granted_scopes"`
}
