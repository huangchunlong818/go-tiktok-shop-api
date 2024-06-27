package product

import (
	"context"
	"errors"
	"tiktok-shop-api/tiktok/common"
)

//tiktok shop

type TiktokProductBrand struct {
}

var newBrandServer *TiktokProductBrand

// 获取实例
func GetNewBrandService() *TiktokProductBrand {
	if newBrandServer == nil {
		newBrandServer = &TiktokProductBrand{}
	}
	return newBrandServer
}

// 获取店铺品牌
func (b *TiktokProductBrand) GetBrands(ctx context.Context, token string, query map[string]string) (result BrandsRsp, err error) {
	//请求接口
	r, err := common.SendTiktokApi(ctx, GetBrands(token), query, nil)
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
	brands, err := common.CheckSliceAny(r["brands"])
	if err != nil {
		err = errors.New("GetBrandsApi brands" + err.Error())
		return
	}
	if len(brands) < 1 {
		return
	}

	//获取具体品牌
	for _, brand := range brands {
		if tmp, err := common.CheckMapStringAny(brand); err == nil && tmp != nil {
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
