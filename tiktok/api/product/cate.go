package product

import (
	"context"
	"errors"
	"fmt"
	"github.com/huangchunlong818/go-tiktok-shop-api/tiktok/common/common"
)

// tiktok shop cate
// 获取店铺分类规则
func (b *TiktokProduct) GetCateRule(ctx context.Context, token string, cateId string, query map[string]string) (result CateRuleRsp, err error) {
	//请求接口
	r, err := b.common.SendTiktokApi(ctx, b.GetCateRuleConfig(token, cateId), query, nil)
	if err != nil {
		return
	}

	//断言cod
	cod, err := b.common.CheckMapStringAny(r["cod"])
	if err != nil {
		return result, errors.New("GetCateRule cod" + err.Error())
	}
	if cod != nil {
		result.Cod = &Cod{IsSupported: cod["is_supported"].(bool)}
	}

	//断言epr
	epr, err := b.common.CheckMapStringAny(r["epr"])
	if err != nil {
		return result, errors.New("GetCateRule epr" + err.Error())
	}
	if epr != nil {
		result.Epr = &Epr{IsRequired: epr["is_required"].(bool)}
	}

	//断言package
	packages, err := b.common.CheckMapStringAny(r["package_dimension"])
	if err != nil {
		return result, errors.New("GetCateRule package_dimension" + err.Error())
	}
	if packages != nil {
		result.PackageDimension = &PackageDimension{IsRequired: packages["is_required"].(bool)}
	}

	//断言size_chart
	sizeChart, err := b.common.CheckMapStringAny(r["size_chart"])
	if err != nil {
		return result, errors.New("GetCateRule size_chart" + err.Error())
	}
	if sizeChart != nil {
		result.SizeChart = &SizeChart{
			IsRequired:  sizeChart["is_required"].(bool),
			IsSupported: sizeChart["is_supported"].(bool),
		}
	}

	//断言product_certifications
	products, err := b.common.CheckSliceAny(r["product_certifications"])
	if err != nil {
		return result, errors.New("GetCateRule product_certifications" + err.Error())
	}
	if len(products) < 1 {
		return
	}

	//获取具体ProductCertifications
	for _, product := range products {
		if tmp, err := b.common.CheckMapStringAny(product); err == nil && tmp != nil {
			result.ProductCertifications = append(result.ProductCertifications, ProductCertifications{
				Id:             tmp["id"].(string),
				Name:           tmp["name"].(string),
				IsRequired:     tmp["is_required"].(bool),
				SampleImageUrl: tmp["sample_image_url"].(string),
			})
		}
	}

	return
}

// 获取店铺分类
func (b *TiktokProduct) GetCate(ctx context.Context, token string, query map[string]string) (result CateRsp, err error) {
	//请求接口
	r, err := b.common.SendTiktokApi(ctx, b.GetCateConfig(token), query, nil)
	if err != nil {
		return
	}

	//断言
	cate, err := b.common.CheckSliceAny(r["categories"])
	if err != nil {
		return result, errors.New("GetCateRule product_certifications" + err.Error())
	}
	if len(cate) < 1 {
		return
	}

	//封装分类
	for _, cat := range cate {
		if tmp, err := b.common.CheckMapStringAny(cat); err != nil && tmp != nil {
			//组装PermissionStatuses
			var permissArr []string
			if permissions, err := b.common.CheckSliceAny(tmp["permission_statuses"]); err != nil && permissions != nil {
				for _, permission := range permissions {
					if permissionStr, ok := permission.(string); ok {
						permissArr = append(permissArr, permissionStr)
					}
				}
			}
			result.Cate = append(result.Cate, Cate{
				Id:                 tmp["id"].(string),
				IsLeaf:             tmp["is_leaf"].(bool),
				LocalName:          tmp["local_name"].(string),
				ParentId:           tmp["parent_id"].(string),
				PermissionStatuses: permissArr,
			})
		}
	}

	return
}

// 分类
func (b *TiktokProduct) GetCateConfig(token string) common.GetApiConfig { //请求方式
	api := fmt.Sprintf("/product/%s/categories", b.config.Version) //请求API PATH

	return common.GetApiConfig{
		ContentType: "application/json",         //请求头content-type 类型
		Method:      "get",                      //请求方法类型
		Api:         api,                        //请求API PATH地址不带域名
		FullApi:     b.config.TkApiDomain + api, //请求的API 完整地址，带域名
		Token:       token,
	}
}

// 分类规则
func (b *TiktokProduct) GetCateRuleConfig(token string, cateId string) common.GetApiConfig { //请求方式
	api := fmt.Sprintf("/product/%s/categories/%s/rules", b.config.Version, cateId) //请求API PATH

	return common.GetApiConfig{
		ContentType: "application/json",         //请求头content-type 类型
		Method:      "get",                      //请求方法类型
		Api:         api,                        //请求API PATH地址不带域名
		FullApi:     b.config.TkApiDomain + api, //请求的API 完整地址，带域名
		Token:       token,
	}
}
