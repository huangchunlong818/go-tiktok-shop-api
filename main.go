package main

import (
	"context"
	"fmt"
	tiktokShop "github.com/huangchunlong818/go-tiktok-shop-api/tiktok"
	"github.com/huangchunlong818/go-tiktok-shop-api/tiktok/common/config"
)

func main() {
	ctx := context.Background()

	tmp := config.WithApp(config.AppConfig{
		AppId:  "xxxx",
		AppKey: "xxxx",
		Secret: "xxxx",
	})
	//获取实例对象
	shopClient := tiktokShop.NewTiktokShopClient(tmp)

	// 授权相关
	authClient := shopClient.GetAuthClient()
	authUrl := authClient.GetAuthUrl("us")
	fmt.Println("授权相关请求：", authUrl)

	//授权API相关
	authApiClient := shopClient.GetAuthApiClient()
	r := authApiClient.GetAuthorizedShops(ctx, "ttttttttttt")
	fmt.Println("授权API相关请求：", r)

	//获取小部件相关API实例
	widgetClient := shopClient.GetWidgetApiClient()
	token := widgetClient.GetWidgetToken(ctx, "ttttttttttttt")
	fmt.Println("小部件相关API请求：", token)

	// 产品相关--品牌
	productClient := shopClient.GetProductApiClient()
	brands := productClient.GetBrands(ctx, "tttt", nil)
	fmt.Println("产品相关API品牌请求：", brands)

	// 产品相关--发布商品校验规则
	prerequisitesQuery := map[string]string{
		"shop_cipher": "xxxxx",
	}
	prerequisites := productClient.GetPrerequisites(ctx, "token", prerequisitesQuery)
	fmt.Println("发布商品校验规则：", prerequisites)

	// 产品相关--详情
	productQuery := map[string]string{
		"shop_cipher":                 "xxxxx",
		"return_under_review_version": "0", // true: "1";false: "0"
	}
	product := productClient.GetProduct(ctx, "token", "productId", productQuery)
	fmt.Println("产品详情：", product)

	// 产品相关--产品列表
	query := map[string]string{
		"shop_cipher": "xxxx",
		"page_size":   "10",
	}
	body := map[string]any{}
	products := productClient.GetProducts(ctx, "xxxxx", query, body)
	fmt.Println("产品请求：", products)
}
