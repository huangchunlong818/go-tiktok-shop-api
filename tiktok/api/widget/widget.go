package widget

import (
	"context"
	"errors"
	"fmt"
	"tiktokShop/tiktok/common/common"
	"tiktokShop/tiktok/common/config"
)

//tiktok shop

type TiktokWidget struct {
	config *config.Config
	common *common.TiktokShopCommon
}

var newServer *TiktokWidget

type WidgetApiClientInterface interface {
	GetWidgetToken(ctx context.Context, token string) (result Token, err error)
	GetWidgetTokenConfig(token string) common.GetApiConfig
}

// 获取实例
func GetNewService(config *config.Config) WidgetApiClientInterface {
	if newServer == nil {
		newServer = &TiktokWidget{
			config: config,
			common: common.GetNewService(config),
		}
	}
	return newServer
}

// 获取所有授权店铺
func (s *TiktokWidget) GetWidgetToken(ctx context.Context, token string) (result Token, err error) {
	//请求接口
	r, err := s.common.SendTiktokApi(ctx, s.GetWidgetTokenConfig(token), nil, nil)

	//断言所有店铺数据 是一个 any切片
	if data, err := s.common.CheckMapStringAny(r["widget_token"]); err != nil {
		err = errors.New("GetWidgetToken widget_token " + err.Error())
	} else {
		if data != nil {
			result = Token{
				Token:    data["token"].(string),
				ExpireAt: int64(data["expire_at"].(float64)),
			}
		}
	}

	return
}

// 获取小部件token
func (s *TiktokWidget) GetWidgetTokenConfig(token string) common.GetApiConfig { //请求方式
	api := fmt.Sprintf("/authorization/%s/widget_token", s.config.WidgetVersion) //请求API PATH

	return common.GetApiConfig{
		ContentType: "application/json",         //请求头content-type 类型
		Method:      "get",                      //请求方法类型
		Api:         api,                        //请求API PATH地址不带域名
		FullApi:     s.config.TkApiDomain + api, //请求的API 完整地址，带域名
		Token:       token,
	}
}
