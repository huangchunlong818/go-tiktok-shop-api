package main

import (
	"context"
	"fmt"
	tiktokShop "github.com/huangchunlong818/go-tiktok-shop-api/tiktok"
	"github.com/huangchunlong818/go-tiktok-shop-api/tiktok/common/config"
)

func main() {
	ctx := context.Background()
	// 授权相关
	authClient := tiktokShop.GetAuthClient(config.WithApp(config.AppConfig{
		AppId:  "123",
		AppKey: "345",
		Secret: "678",
	}))
	authUrl := authClient.GetAuthUrl("us")
	fmt.Println("授权相关请求：", authUrl)

	//授权API相关
	authApiClient := tiktokShop.GetAuthApiClient(config.WithApp(config.AppConfig{
		AppId:  "123",
		AppKey: "345",
		Secret: "678",
	}))
	shops, err := authApiClient.GetAuthorizedShops(ctx, "ttttttttttt")
	fmt.Println("授权API相关请求：", shops, err)

	//获取小部件相关API实例
	widgetClient := tiktokShop.GetWidgetApiClient(config.WithApp(config.AppConfig{
		AppId:  "123",
		AppKey: "345",
		Secret: "678",
	}))
	token, err := widgetClient.GetWidgetToken(ctx, "ttttttttttttt")
	fmt.Println("小部件相关API请求：", token, err)

	// 产品相关--品牌
	productClient := tiktokShop.GetProductApiClient(config.WithApp(config.AppConfig{
		AppId:  "123",
		AppKey: "345",
		Secret: "678",
	}))
	brands, err := productClient.GetBrands(ctx, "tttt", nil)
	fmt.Println("产品相关API品牌请求：", brands, err)

	// 产品相关--分类
	cate, err := productClient.GetCate(ctx, "tttt", nil)
	fmt.Println("产品相关API分类请求：", cate, err)
	cateRule, err := productClient.GetCateRule(ctx, "tttt", "123", nil)
	fmt.Println("产品相关API分类规则请求：", cateRule, err)
}
