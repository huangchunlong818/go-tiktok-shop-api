package auth

import (
	"fmt"
	tiktok "tiktok-shop-api"
	"tiktok-shop-api/common"
)

type Shops struct {
	Cipher     string `json:"cipher"`      //请求接口需要用到的
	Code       string `json:"code"`        //商店代码
	Id         string `json:"id"`          //商店ID
	Name       string `json:"name"`        //商店名称
	Region     string `json:"region"`      //商店所在区域，国家代码
	SellerType string `json:"seller_type"` //商店类型？1代表跨境店铺CROSS_BORDER， 2代表本地店铺LOCAL
}

// 获取卖家授权的所有店铺
func GetAuthorizedShopsApi(token string) common.GetApiConfig { //请求方式
	api := fmt.Sprintf("/authorization/%s/shops", tiktok.ApiVersion()) //请求API PATH

	return common.GetApiConfig{
		ContentType: "application/json",         //请求头content-type 类型
		Method:      "get",                      //请求方法类型
		Api:         api,                        //请求API PATH地址不带域名
		FullApi:     tiktok.TkApiDomain() + api, //请求的API 完整地址，带域名
		Token:       token,                      //请求的token
	}
}
