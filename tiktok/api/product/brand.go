package product

import (
	"context"
	"errors"
	"fmt"
	"github.com/huangchunlong818/go-tiktok-shop-api/tiktok/common/common"
)

//tiktok shop brand

// 获取店铺品牌
func (b *TiktokProduct) GetBrands(ctx context.Context, token string, query map[string]string) (result BrandsRsp, err error) {
	//请求接口
	r, err := b.common.SendTiktokApi(ctx, b.GetBrandsConfig(token), query, nil)
	if err != nil {
		return
	}

	//断言分页token
	if next, ok := r["next_page_token"].(string); !ok {
		return result, errors.New("GetBrandsApi next_page_token response error")
	} else {
		result.NextPageToken = next
	}

	//断言总数
	if total, ok := r["total_count"].(float64); !ok {
		return result, errors.New("GetBrandsApi total_count response error")
	} else {
		result.TotalCount = int(total)
	}

	//断言品牌列表
	brands, err := b.common.CheckSliceAny(r["brands"])
	if err != nil {
		err = errors.New("GetBrandsApi brands" + err.Error())
		return
	}
	if len(brands) < 1 {
		return
	}

	//获取具体品牌
	for _, brand := range brands {
		if tmp, err := b.common.CheckMapStringAny(brand); err == nil && tmp != nil {
			result.Brands = append(result.Brands, Brands{
				AuthorizedStatus: tmp["authorized_status"].(string),
				BrandStatus:      tmp["brand_status"].(string),
				Id:               tmp["id"].(string),
				Name:             tmp["name"].(string),
				IsT1Brand:        tmp["is_t1_brand"].(bool),
			})
		}
	}

	return
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
