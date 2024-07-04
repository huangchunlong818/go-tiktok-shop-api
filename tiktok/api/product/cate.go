package product

import (
	"context"
	"fmt"
	"github.com/huangchunlong818/go-tiktok-shop-api/tiktok/common/common"
)

// tiktok shop cate
// 获取店铺分类规则
func (b *TiktokProduct) GetCateRule(ctx context.Context, token string, cateId string, query map[string]string) CateRuleResultRsp {
	//请求接口
	r := b.SendTiktokApi(ctx, b.GetCateRuleConfig(token, cateId), query, nil)
	result := CateRuleResultRsp{
		Code:     r.Code,
		Message:  r.Message,
		HttpCode: r.HttpCode,
	}
	if !b.IsSuccess(r) {
		return result
	}
	//断言cod
	cod, err := b.CheckMapStringAny(r.Data["cod"])
	if err != nil {
		r.Code = common.ErrCode
		r.Message = "GetCateRule cod" + err.Error()
		return result
	}
	if cod != nil {
		result.Data.Cod = &Cod{IsSupported: cod["is_supported"].(bool)}
	}

	//断言epr
	epr, err := b.CheckMapStringAny(r.Data["epr"])
	if err != nil {
		r.Code = common.ErrCode
		r.Message = "GetCateRule epr" + err.Error()
		return result
	}
	if epr != nil {
		result.Data.Epr = &Epr{IsRequired: epr["is_required"].(bool)}
	}

	//断言package
	packages, err := b.CheckMapStringAny(r.Data["package_dimension"])
	if err != nil {
		r.Code = common.ErrCode
		r.Message = "GetCateRule package_dimension" + err.Error()
		return result
	}
	if packages != nil {
		result.Data.PackageDimension = &PackageDimension{IsRequired: packages["is_required"].(bool)}
	}

	//断言size_chart
	sizeChart, err := b.CheckMapStringAny(r.Data["size_chart"])
	if err != nil {
		r.Code = common.ErrCode
		r.Message = "GetCateRule size_chart" + err.Error()
		return result
	}
	if sizeChart != nil {
		result.Data.SizeChart = &SizeChart{
			IsRequired:  sizeChart["is_required"].(bool),
			IsSupported: sizeChart["is_supported"].(bool),
		}
	}

	//断言product_certifications
	products, err := b.CheckSliceAny(r.Data["product_certifications"])
	if err != nil {
		r.Code = common.ErrCode
		r.Message = "GetCateRule product_certifications" + err.Error()
		return result
	}
	if len(products) < 1 {
		return result
	}

	//获取具体ProductCertifications
	for _, product := range products {
		if tmp, err := b.CheckMapStringAny(product); err == nil && tmp != nil {
			result.Data.ProductCertifications = append(result.Data.ProductCertifications, ProductCertifications{
				Id:             tmp["id"].(string),
				Name:           tmp["name"].(string),
				IsRequired:     tmp["is_required"].(bool),
				SampleImageUrl: tmp["sample_image_url"].(string),
			})
		}
	}

	return result
}

// 获取店铺分类
func (b *TiktokProduct) GetCate(ctx context.Context, token string, query map[string]string) CateResultRsp {
	//请求接口
	r := b.SendTiktokApi(ctx, b.GetCateConfig(token), query, nil)
	result := CateResultRsp{
		Code:     r.Code,
		Message:  r.Message,
		HttpCode: r.HttpCode,
	}
	if !b.IsSuccess(r) {
		return result
	}
	//断言
	cate, err := b.CheckSliceAny(r.Data["categories"])
	if err != nil {
		r.Code = common.ErrCode
		r.Message = "GetCateRule product_certifications" + err.Error()
		return result
	}
	if len(cate) < 1 {
		return result
	}

	//封装分类
	for _, cat := range cate {
		if tmp, err := b.CheckMapStringAny(cat); err != nil && tmp != nil {
			//组装PermissionStatuses
			var permissArr []string
			if permissions, err := b.CheckSliceAny(tmp["permission_statuses"]); err != nil && permissions != nil {
				for _, permission := range permissions {
					if permissionStr, ok := permission.(string); ok {
						permissArr = append(permissArr, permissionStr)
					}
				}
			}
			result.Data.Cate = append(result.Data.Cate, Cate{
				Id:                 tmp["id"].(string),
				IsLeaf:             tmp["is_leaf"].(bool),
				LocalName:          tmp["local_name"].(string),
				ParentId:           tmp["parent_id"].(string),
				PermissionStatuses: permissArr,
			})
		}
	}

	return result
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
