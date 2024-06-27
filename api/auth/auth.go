package auth

import (
	"context"
	"errors"
	"tiktok-shop-api/common"
)

//tiktok shop

type TiktokShop struct {
}

var newServer *TiktokShop

// 获取实例
func GetNewService() *TiktokShop {
	if newServer == nil {
		newServer = &TiktokShop{}
	}
	return newServer
}

// 获取所有授权店铺
func (s *TiktokShop) GetAuthorizedShops(ctx context.Context, token string) (result []Shops, err error) {
	//请求接口
	r, err := common.SendTiktokApi(ctx, GetAuthorizedShopsApi(token), nil, nil)
	if err != nil {
		return
	}

	//断言所有店铺数据 是一个 any切片
	data, err := common.CheckSliceAny(r["shops"])
	if err != nil {
		err = errors.New("GetAuthorizedShopsApi shops " + err.Error())
		return
	}
	if len(data) < 1 {
		return
	}
	for _, now := range data {
		//断言单个授权店铺， 是 map[string]any
		if tmp, err := common.CheckMapStringAny(now); err == nil && tmp != nil {
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
