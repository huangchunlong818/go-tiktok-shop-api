package product

import (
	"context"
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

	//断言分页token
	if next, ok := r.Data["next_page_token"].(string); !ok {
		r.Code = common.ErrCode
		r.Message = "GetBrandsApi next_page_token response error"
		return result
	} else {
		result.Data.NextPageToken = next
	}

	//断言总数
	if total, ok := r.Data["total_count"].(float64); !ok {
		r.Code = common.ErrCode
		r.Message = "GetBrandsApi total_count response error"
		return result
	} else {
		result.Data.TotalCount = int(total)
	}

	//断言品牌列表
	brands, err := b.CheckSliceAny(r.Data["brands"])
	if err != nil {
		r.Code = common.ErrCode
		r.Message = "GetBrandsApi brands" + err.Error()
		return result
	}
	if len(brands) < 1 {
		return result
	}

	//获取具体品牌
	for _, brand := range brands {
		if tmp, err := b.CheckMapStringAny(brand); err == nil && tmp != nil {
			result.Data.Brands = append(result.Data.Brands, Brands{
				AuthorizedStatus: tmp["authorized_status"].(string),
				BrandStatus:      tmp["brand_status"].(string),
				Id:               tmp["id"].(string),
				Name:             tmp["name"].(string),
				IsT1Brand:        tmp["is_t1_brand"].(bool),
			})
		}
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
