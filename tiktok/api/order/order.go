package order

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/huangchunlong818/go-tiktok-shop-api/tiktok/common/common"
	"github.com/huangchunlong818/go-tiktok-shop-api/tiktok/common/config"
)

type TiktokOrder struct {
	config *config.Config
	*common.TiktokShopCommon
}

var newServer *TiktokOrder

type OrderApiClientInterface interface {
	GetOrders(ctx context.Context, token string, query map[string]string) OrderResultRsp
	GetOrdersConfig(token string) common.GetApiConfig
}

// GetOrderService 获取实例
func GetOrderService(config *config.Config) OrderApiClientInterface {
	if newServer == nil {
		newServer = &TiktokOrder{
			config:           config,
			TiktokShopCommon: common.GetNewService(config),
		}
	}
	return newServer
}

// GetOrders 获取订单详情（可批量）
func (to *TiktokOrder) GetOrders(ctx context.Context, token string, query map[string]string) OrderResultRsp {
	//请求接口
	r := to.SendTiktokApi(ctx, to.GetOrdersConfig(token), query, nil, nil)
	result := OrderResultRsp{
		Code:     r.Code,
		Message:  r.Message,
		HttpCode: r.HttpCode,
	}
	if !to.IsSuccess(r) {
		return result
	}

	//解析数据
	err := json.Unmarshal(r.Data, &result)
	if err != nil {
		r.Code = common.ErrCode
		r.Message = "GetOrders response error " + err.Error()
		return result
	}

	return result
}

func (to *TiktokOrder) GetOrdersConfig(token string) common.GetApiConfig { //请求方式
	api := fmt.Sprintf("/order/%s/orders", to.config.Version) //请求API PATH

	return common.GetApiConfig{
		ContentType: "application/json",          //请求头content-type 类型
		Method:      "get",                       //请求方法类型
		Api:         api,                         //请求API PATH地址不带域名
		FullApi:     to.config.TkApiDomain + api, //请求的API 完整地址，带域名
		Token:       token,
	}
}
