package product

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/huangchunlong818/go-tiktok-shop-api/tiktok/common/common"
)

// CheckProductListing 检测商品字段
func (b *TiktokProduct) CheckProductListing(ctx context.Context, token string, query map[string]string, body map[string]any) CheckProductListingResultRsp {
	//请求接口
	r := b.SendTiktokApi(ctx, b.CheckProductListingConfig(token), query, body, nil)
	result := CheckProductListingResultRsp{
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
		r.Message = "CheckProductListing response error " + err.Error()
		return result
	}

	return result
}

func (b *TiktokProduct) CheckProductListingConfig(token string) common.GetApiConfig { //请求方式
	api := fmt.Sprintf("/product/%s/products/listing_check", b.config.Version) //请求API PATH

	return common.GetApiConfig{
		ContentType: "application/json",         //请求头content-type 类型
		Method:      "post",                     //请求方法类型
		Api:         api,                        //请求API PATH地址不带域名
		FullApi:     b.config.TkApiDomain + api, //请求的API 完整地址，带域名
		Token:       token,
	}
}
