package widget

import (
	"fmt"
	tiktok "tiktok-shop-api"
	"tiktok-shop-api/common"
)

type Token struct {
	Token    string `json:"token"`     //小部件token
	ExpireAt int64  `json:"expire_at"` //过期时间
}

// 获取小部件token
func GetWidgetToken(token string) common.GetApiConfig { //请求方式
	api := fmt.Sprintf("/authorization/%s/widget_token", tiktok.WidgetVersion()) //请求API PATH

	return common.GetApiConfig{
		ContentType: "application/json",         //请求头content-type 类型
		Method:      "get",                      //请求方法类型
		Api:         api,                        //请求API PATH地址不带域名
		FullApi:     tiktok.TkApiDomain() + api, //请求的API 完整地址，带域名
		Token:       token,
	}
}
