package product

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/huangchunlong818/go-tiktok-shop-api/tiktok/common/common"
	"github.com/huangchunlong818/go-tiktok-shop-api/tiktok/common/config"
)

type TiktokProduct struct {
	config *config.Config
	*common.TiktokShopCommon
}

var newServer *TiktokProduct

type ProductApiClientInterface interface {
	//品牌
	GetBrands(ctx context.Context, token string, query map[string]string) BrandsResultRsp
	GetBrandsConfig(token string) common.GetApiConfig

	//分类
	GetCateRule(ctx context.Context, token string, cateId string, query map[string]string) CateRuleResultRsp
	GetCate(ctx context.Context, token string, query map[string]string) CateResultRsp
	GetCateConfig(token string) common.GetApiConfig
	GetCateRuleConfig(token string, cateId string) common.GetApiConfig

	//产品
	GetProducts(ctx context.Context, token string, query map[string]string, body map[string]any) ProductsResultRsp
	GetProductsConfig(token string) common.GetApiConfig
}

// 获取产品，搜索产品
func (b *TiktokProduct) GetProducts(ctx context.Context, token string, query map[string]string, body map[string]any) ProductsResultRsp {
	//请求接口
	r := b.SendTiktokApi(ctx, b.GetProductsConfig(token), query, body)
	result := ProductsResultRsp{
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
		r.Message = "GetBrandsApi response error " + err.Error()
		return result
	}

	return result
}

// 产品
func (b *TiktokProduct) GetProductsConfig(token string) common.GetApiConfig { //请求方式
	api := fmt.Sprintf("/product/%s/products/search", b.config.ProductVersion) //请求API PATH

	return common.GetApiConfig{
		ContentType: "application/json",         //请求头content-type 类型
		Method:      "post",                     //请求方法类型
		Api:         api,                        //请求API PATH地址不带域名
		FullApi:     b.config.TkApiDomain + api, //请求的API 完整地址，带域名
		Token:       token,
	}
}

// 获取实例
func GetNewService(config *config.Config) ProductApiClientInterface {
	if newServer == nil {
		newServer = &TiktokProduct{
			config:           config,
			TiktokShopCommon: common.GetNewService(config),
		}
	}
	return newServer
}
