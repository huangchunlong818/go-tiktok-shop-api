package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/huangchunlong818/go-tiktok-shop-api/tiktok/common/common"
	"github.com/huangchunlong818/go-tiktok-shop-api/tiktok/common/config"
)

//tiktok shop

type TiktokShop struct {
	config *config.Config
	common *common.TiktokShopCommon
}

var newServer *TiktokShop

type AuthApiClientInterface interface {
	GetAuthorizedShopsApi(token string) common.GetApiConfig
	GetAuthorizedShops(ctx context.Context, token string) (result []Shops, err error)
}

// 获取实例
func GetNewService(config *config.Config) AuthApiClientInterface {
	if newServer == nil {
		newServer = &TiktokShop{
			config: config,
			common: common.GetNewService(config),
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
func (s *TiktokShop) GetAuthorizedShops(ctx context.Context, token string) (result []Shops, err error) {
	//请求接口
	r, err := s.common.SendTiktokApi(ctx, s.GetAuthorizedShopsApi(token), nil, nil)
	if err != nil {
		return
	}

	//断言所有店铺数据 是一个 any切片
	data, err := s.common.CheckSliceAny(r["shops"])
	if err != nil {
		err = errors.New("GetAuthorizedShopsApi shops " + err.Error())
		return
	}
	if len(data) < 1 {
		return
	}
	for _, now := range data {
		//断言单个授权店铺， 是 map[string]any
		if tmp, err := s.common.CheckMapStringAny(now); err == nil && tmp != nil {
			result = append(result, Shops{
				Cipher:     tmp["cipher"].(string),
				Code:       tmp["code"].(string),
				Id:         tmp["id"].(string),
				Name:       tmp["name"].(string),
				Region:     tmp["region"].(string),
				SellerType: tmp["seller_type"].(string),
			})
		}
	}

	return
}
