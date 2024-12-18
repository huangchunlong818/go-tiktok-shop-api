package tiktokShop

//共用tiktok配置
import (
	apiAuth "github.com/huangchunlong818/go-tiktok-shop-api/tiktok/api/auth"
	"github.com/huangchunlong818/go-tiktok-shop-api/tiktok/api/logistics"
	"github.com/huangchunlong818/go-tiktok-shop-api/tiktok/api/order"
	"github.com/huangchunlong818/go-tiktok-shop-api/tiktok/api/product"
	"github.com/huangchunlong818/go-tiktok-shop-api/tiktok/api/widget"
	"github.com/huangchunlong818/go-tiktok-shop-api/tiktok/auth"
	"github.com/huangchunlong818/go-tiktok-shop-api/tiktok/common/config"
)

type AuthClient auth.AuthClientInterface                      //授权服务接口
type AuthApiClient apiAuth.AuthApiClientInterface             //授权API服务接口
type WidgetApiClient widget.WidgetApiClientInterface          //小部件API服务接口
type ProductApiClient product.ProductApiClientInterface       //小部件API服务接口
type LogisticsApiClient logistics.LogisticsApiClientInterface //物流信息API服务接口
type OrderApiClient order.OrderApiClientInterface             //订单API服务接口

type TiktokShop struct {
	client *config.TiktokShopClient
}

var newClient *TiktokShop

func NewTiktokShopClient(options ...config.Option) *TiktokShop {
	if newClient == nil {
		newClient = &TiktokShop{
			client: config.NewTiktokShopClient().SetConfig(options...),
		}
	}
	return newClient
}

// 设置配置
func (tk *TiktokShop) SetOptions(options ...config.Option) {
	config.NewTiktokShopClient().SetConfig(options...)
}

// 获取授权相关实例 无sign签名请求
func (tk *TiktokShop) GetAuthClient() AuthClient {
	return auth.GetNewService(tk.client.GetConfig())
}

// 获取授权相关API实例 有sign签名请求
func (tk *TiktokShop) GetAuthApiClient() AuthApiClient {
	return apiAuth.GetNewService(tk.client.GetConfig())
}

// 获取小部件相关API实例 有sign签名请求
func (tk *TiktokShop) GetWidgetApiClient() WidgetApiClient {
	return widget.GetNewService(tk.client.GetConfig())
}

// 获取产品相关API实例 有sign签名请求
func (tk *TiktokShop) GetProductApiClient() ProductApiClient {
	return product.GetNewService(tk.client.GetConfig())
}

// 获取物流资源相关API实例 有sign签名请求
func (tk *TiktokShop) GetLogisticsApiClient() LogisticsApiClient {
	return logistics.GetNewService(tk.client.GetConfig())
}

// 获取订单资源相关API实例 有sign签名请求
func (tk *TiktokShop) GetOrderApiClient() OrderApiClient {
	return order.GetOrderService(tk.client.GetConfig())
}
