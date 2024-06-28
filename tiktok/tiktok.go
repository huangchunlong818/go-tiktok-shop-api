package tiktokShop

//共用tiktok配置
import (
	apiAuth "tiktokShop/tiktok/api/auth"
	"tiktokShop/tiktok/api/product"
	"tiktokShop/tiktok/api/widget"
	"tiktokShop/tiktok/auth"
	"tiktokShop/tiktok/common/config"
)

type AuthClient auth.AuthClientInterface                //授权服务接口
type AuthApiClient apiAuth.AuthApiClientInterface       //授权API服务接口
type WidgetApiClient widget.WidgetApiClientInterface    //小部件API服务接口
type ProductApiClient product.ProductApiClientInterface //小部件API服务接口

// 获取授权相关实例 无sign签名请求
func GetAuthClient(options ...config.Option) AuthClient {
	c := config.NewTiktokShopClient().SetConfig(options...)
	return auth.GetNewService(c.GetConfig())
}

// 获取授权相关API实例 有sign签名请求
func GetAuthApiClient(options ...config.Option) AuthApiClient {
	c := config.NewTiktokShopClient().SetConfig(options...)
	return apiAuth.GetNewService(c.GetConfig())
}

// 获取小部件相关API实例 有sign签名请求
func GetWidgetApiClient(options ...config.Option) WidgetApiClient {
	c := config.NewTiktokShopClient().SetConfig(options...)
	return widget.GetNewService(c.GetConfig())
}

// 获取产品相关API实例 有sign签名请求
func GetProductApiClient(options ...config.Option) ProductApiClient {
	c := config.NewTiktokShopClient().SetConfig(options...)
	return product.GetNewService(c.GetConfig())
}
