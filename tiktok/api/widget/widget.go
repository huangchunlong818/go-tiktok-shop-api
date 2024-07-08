package widget

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/huangchunlong818/go-tiktok-shop-api/tiktok/common/common"
	"github.com/huangchunlong818/go-tiktok-shop-api/tiktok/common/config"
)

//tiktok shop

type TiktokWidget struct {
	config *config.Config
	*common.TiktokShopCommon
}

var newServer *TiktokWidget

type WidgetApiClientInterface interface {
	GetWidgetToken(ctx context.Context, token string) GetTokenRsp
	GetWidgetTokenConfig(token string) common.GetApiConfig
}

// 获取实例
func GetNewService(config *config.Config) WidgetApiClientInterface {
	if newServer == nil {
		newServer = &TiktokWidget{
			config:           config,
			TiktokShopCommon: common.GetNewService(config),
		}
	}
	return newServer
}

// 获取所有授权店铺
func (s *TiktokWidget) GetWidgetToken(ctx context.Context, token string) GetTokenRsp {
	//请求接口
	r := s.SendTiktokApi(ctx, s.GetWidgetTokenConfig(token), nil, nil)
	result := GetTokenRsp{
		Code:     r.Code,
		Message:  r.Message,
		HttpCode: r.HttpCode,
	}
	if !s.IsSuccess(r) {
		return result
	}

	//解析数据
	err := json.Unmarshal(r.Data, &result)
	if err != nil {
		r.Code = common.ErrCode
		r.Message = "GetWidgetToken response error " + err.Error()
		return result
	}

	return result
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
