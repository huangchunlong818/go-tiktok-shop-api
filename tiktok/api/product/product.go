package product

import (
	"context"
	"go-tiktok-shop-api/tiktok/common/common"
	"go-tiktok-shop-api/tiktok/common/config"
)

type TiktokProduct struct {
	config *config.Config
	common *common.TiktokShopCommon
}

var newServer *TiktokProduct

type ProductApiClientInterface interface {
	//品牌
	GetBrands(ctx context.Context, token string, query map[string]string) (result BrandsRsp, err error)
	GetBrandsConfig(token string) common.GetApiConfig

	//分类
	GetCateRule(ctx context.Context, token string, cateId string, query map[string]string) (result CateRuleRsp, err error)
	GetCate(ctx context.Context, token string, query map[string]string) (result CateRsp, err error)
	GetCateConfig(token string) common.GetApiConfig
	GetCateRuleConfig(token string, cateId string) common.GetApiConfig
}

// 获取实例
func GetNewService(config *config.Config) ProductApiClientInterface {
	if newServer == nil {
		newServer = &TiktokProduct{
			config: config,
			common: common.GetNewService(config),
		}
	}
	return newServer
}
