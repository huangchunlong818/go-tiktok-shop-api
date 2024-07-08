package product

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/huangchunlong818/go-tiktok-shop-api/tiktok/common/common"
)

//tiktok shop brand

// 获取店铺品牌
func (b *TiktokProduct) GetBrands(ctx context.Context, token string, query map[string]string) BrandsResultRsp {
	//请求接口
	r := b.SendTiktokApi(ctx, b.GetBrandsConfig(token), query, nil)
	result := BrandsResultRsp{
		Code:     r.Code,
		Message:  r.Message,
		HttpCode: r.HttpCode,
	}
	if !b.IsSuccess(r) {
		return result
	}

	//解析数据
	err := json.Unmarshal(r.Data, &result)
	if err != nil {
		r.Code = common.ErrCode
		r.Message = "GetBrands response error " + err.Error()
		return result
	}

	return result
}

// 品牌
func (b *TiktokProduct) GetBrandsConfig(token string) common.GetApiConfig { //请求方式
	api := fmt.Sprintf("/product/%s/brands", b.config.Version) //请求API PATH

	return common.GetApiConfig{
		ContentType: "application/json",         //请求头content-type 类型
		Method:      "get",                      //请求方法类型
		Api:         api,                        //请求API PATH地址不带域名
		FullApi:     b.config.TkApiDomain + api, //请求的API 完整地址，带域名
		Token:       token,
	}
}
