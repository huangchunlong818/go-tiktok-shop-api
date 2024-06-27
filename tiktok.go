package tiktok

//共用tiktok配置
import (
	"fmt"
	"tiktok-shop-api/global"
)

// tiktok shop api 版本 通用API版本，还有其他版本
func ApiVersion() string {
	return global.Config.Tiktok.Version
}

// tiktok shop api widget 小部件版本
func WidgetVersion() string {
	return global.Config.Tiktok.WidgetVersion
}

// tiktok 应用 app  id
func AppId() string {
	return global.Config.Tiktok.AppId
}

// tiktok 应用 App key
func AppKey() string {
	return global.Config.Tiktok.AppKey
}

// tiktok 应用 App secret
func Secret() string {
	return global.Config.Tiktok.AppSecret
}

// tiktok shop 美国授权地址
func UsAuthUrl() string {
	return fmt.Sprintf("https://services.us.tiktokshop.com/open/authorize?service_id=%s", AppId())
}

// tiktok shop 除美国外其他国家授权地址
func OtherAuthUrl() string {
	return fmt.Sprintf("https://services.tiktokshop.com/open/authorize?service_id=%s", AppId())
}

// tiktok shop  auth操作域名 授权相关 获取token和刷新token
func AuthApiDomain() string {
	return "https://auth.tiktok-shops.com"
}

// tiktok shop api 操作域名  API操作相关
func TkApiDomain() string {
	return "https://open-api.tiktokglobalshop.com"
}
