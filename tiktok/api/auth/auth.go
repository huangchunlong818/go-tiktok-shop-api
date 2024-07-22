package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/huangchunlong818/go-tiktok-shop-api/tiktok/common/common"
	"github.com/huangchunlong818/go-tiktok-shop-api/tiktok/common/config"
)

//tiktok shop

type TiktokShop struct {
	config                   *config.Config
	*common.TiktokShopCommon // 嵌入 common.TiktokShopCommon
}

var newServer *TiktokShop

type AuthApiClientInterface interface {
	GetAuthorizedShopsApi(token string) common.GetApiConfig
	GetAuthorizedShops(ctx context.Context, token string) ShopsRsp
}

// 获取实例
func GetNewService(config *config.Config) AuthApiClientInterface {
	if newServer == nil {
		newServer = &TiktokShop{
			config:           config,
			TiktokShopCommon: common.GetNewService(config),
		}
	}
	return newServer
}

// 获取卖家授权的所有店铺
func (s *TiktokShop) GetAuthorizedShopsApi(token string) common.GetApiConfig { //请求方式
	api := fmt.Sprintf("/authorization/%s/shops", s.config.Version) //请求API PATH

	return common.GetApiConfig{
		ContentType: "application/json",         //请求头content-type 类型
		Method:      "get",                      //请求方法类型
		Api:         api,                        //请求API PATH地址不带域名
		FullApi:     s.config.TkApiDomain + api, //请求的API 完整地址，带域名
		Token:       token,                      //请求的token
	}
}

// 获取所有授权店铺
func (s *TiktokShop) GetAuthorizedShops(ctx context.Context, token string) ShopsRsp {
	//请求接口
	r := s.SendTiktokApi(ctx, s.GetAuthorizedShopsApi(token), nil, nil, nil)
	result := ShopsRsp{
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
		r.Message = "GetAuthorizedShops response error " + err.Error()
		return result
	}

	return result
}
