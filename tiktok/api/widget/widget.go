package widget

import (
	"context"
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

	//断言所有店铺数据 是一个 any切片
	if data, err := s.CheckMapStringAny(r.Data["widget_token"]); err != nil {
		r.Code = common.ErrCode
		r.Message = "GetWidgetToken widget_token " + err.Error()
		return result
	} else {
		if data != nil {
			result.Data = Token{
				Token:    data["token"].(string),
				ExpireAt: int64(data["expire_at"].(float64)),
			}
		}
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
