package logistics

import (
	"context"
	"github.com/huangchunlong818/go-tiktok-shop-api/tiktok/common/common"
	"github.com/huangchunlong818/go-tiktok-shop-api/tiktok/common/config"
)

type TiktokLogistics struct {
	config *config.Config
	*common.TiktokShopCommon
}

var newServer *TiktokLogistics

type LogisticsApiClientInterface interface {
	GetWarehouses(ctx context.Context, token string, query map[string]string) WarehousesResultRsp
	GetWarehousesConfig(token string) common.GetApiConfig
}

// GetNewService 获取实例
func GetNewService(config *config.Config) LogisticsApiClientInterface {
	if newServer == nil {
		newServer = &TiktokLogistics{
			config:           config,
			TiktokShopCommon: common.GetNewService(config),
		}
	}
	return newServer
}
