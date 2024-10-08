package config

//共用tiktok配置
import (
	"fmt"
)

// 配置
type Config struct {
	Version                string    //版本，通用大版本
	WidgetVersion          string    //版本，widget小部件 版本
	PrerequisitesVersion   string    //版本，获取店铺的商品规则以及是否满足上架商品的条件
	ProductVersion         string    //版本，产品 版本
	OptimizedImagesVersion string    //版本，图片优化默认版本
	App                    AppConfig //tiktok shop 应用 app配置
	UsAuthUrl              string    //美国授权地址
	OtherAuthUrl           string    //美国外的授权地址
	AuthApiDomain          string    //授权接口域名
	TkApiDomain            string    //tiktok shop api 操作域名
}

type AppConfig struct {
	AppId  string //应用ID
	AppKey string //应用 App key
	Secret string //App secret
}

// 实现了链式操作的客户端
type TiktokShopClient struct {
	configs *Config
}

// 获取实例
func NewTiktokShopClient() *TiktokShopClient {
	return &TiktokShopClient{
		configs: &Config{},
	}
}

// 设置配置
func (t *TiktokShopClient) SetConfig(options ...Option) *TiktokShopClient {
	for _, option := range options {
		option(t.configs)
	}

	//检查是否设置了APP
	if t.configs.App.AppId == "" || t.configs.App.AppKey == "" || t.configs.App.Secret == "" {
		return nil
	}

	// 检查是否设置了 Version，如果没有则设置为默认值
	if t.configs.Version == "" {
		t.configs.Version = DefaultApiVersion()
	}

	// 图片优化默认版本
	if t.configs.OptimizedImagesVersion == "" {
		t.configs.OptimizedImagesVersion = DefaultOptimizedImagesVersion()
	}

	// 检查是否设置了 WidgetVersion，如果没有则设置为默认值
	if t.configs.WidgetVersion == "" {
		t.configs.WidgetVersion = DefaultWidgetVersion()
	}
	// 检查是否设置了 PrerequisitesVersion，如果没有则设置为默认值
	if t.configs.PrerequisitesVersion == "" {
		t.configs.PrerequisitesVersion = DefaultPrerequisitesVersion()
	}
	// 检查是否设置了 UsAuthUrl，如果没有则设置为默认值
	if t.configs.UsAuthUrl == "" {
		t.configs.UsAuthUrl = DefaultUsAuthUrl(t.configs.App.AppId)
	}
	// 检查是否设置了 OtherAuthUrl，如果没有则设置为默认值
	if t.configs.OtherAuthUrl == "" {
		t.configs.OtherAuthUrl = DefaultOtherAuthUrl(t.configs.App.AppId)
	}
	// 检查是否设置了 AuthApiDomain，如果没有则设置为默认值
	if t.configs.AuthApiDomain == "" {
		t.configs.AuthApiDomain = DefaultAuthApiDomain()
	}
	// 检查是否设置了 UsAuthUrl，如果没有则设置为默认值
	if t.configs.TkApiDomain == "" {
		t.configs.TkApiDomain = DefaultTkApiDomain()
	}
	//检查是否设置了产品版本，如果没有则用默认值
	if t.configs.ProductVersion == "" {
		t.configs.ProductVersion = DefaultProductVersion()
	}

	return t
}

// GetConfig 返回配置
func (t *TiktokShopClient) GetConfig() *Config {
	return t.configs
}

type Option func(*Config)

func WithVersion(version string) Option {
	return func(config *Config) {
		config.Version = version
	}
}
func WithWidgetVersion(widgetVersion string) Option {
	return func(config *Config) {
		config.WidgetVersion = widgetVersion
	}
}
func WithProductVersion(productVersion string) Option {
	return func(config *Config) {
		config.ProductVersion = productVersion
	}
}
func WithUsAuthUrl(usAuthUrl string) Option {
	return func(config *Config) {
		config.UsAuthUrl = usAuthUrl
	}
}
func WithOtherAuthUrl(otherAuthUrl string) Option {
	return func(config *Config) {
		config.OtherAuthUrl = otherAuthUrl
	}
}
func WithAuthApiDomain(authApiDomain string) Option {
	return func(config *Config) {
		config.AuthApiDomain = authApiDomain
	}
}
func WithTkApiDomain(tkApiDomain string) Option {
	return func(config *Config) {
		config.TkApiDomain = tkApiDomain
	}
}
func WithApp(app AppConfig) Option {
	return func(config *Config) {
		config.App = app
	}
}

// tiktok shop api 版本 通用API版本，还有其他版本 默认值
func DefaultApiVersion() string {
	return "202309"
}

// tiktok shop api widget 小部件版本 默认值
func DefaultWidgetVersion() string {
	return "202401"
}

// tiktok shop api Check Listing Prerequisites 获取店铺的商品规则以及是否满足上架商品的条件
func DefaultPrerequisitesVersion() string {
	return "202312"
}

// tiktok shop api product产品版本 默认值
func DefaultProductVersion() string {
	return "202312"
}

// 图片优化默认版本
func DefaultOptimizedImagesVersion() string {
	return "202404"
}

// tiktok shop 美国授权地址 默认值
func DefaultUsAuthUrl(appId string) string {
	return fmt.Sprintf("https://services.us.tiktokshop.com/open/authorize?service_id=%s", appId)
}

// tiktok shop 除美国外其他国家授权地址 默认值
func DefaultOtherAuthUrl(appId string) string {
	return fmt.Sprintf("https://services.tiktokshop.com/open/authorize?service_id=%s", appId)
}

// tiktok shop  auth操作域名 授权相关 获取token和刷新token 默认值
func DefaultAuthApiDomain() string {
	return "https://auth.tiktok-shops.com"
}

// tiktok shop api 操作域名  API操作相关 默认值
func DefaultTkApiDomain() string {
	return "https://open-api.tiktokglobalshop.com"
}
