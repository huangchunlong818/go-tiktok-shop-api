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
		AppId:  "7373895977494873902",
		AppKey: "6ck2qip0j8tni",
		Secret: "812d148e09ce85a8fefd563db92ecddcc10d8eb0",
	})
	//获取实例对象
	shopClient := tiktokShop.NewTiktokShopClient(tmp)

	// 授权相关
	authClient := shopClient.GetAuthClient()
	authUrl := authClient.GetAuthUrl("us")
	fmt.Println("授权相关请求：", authUrl)

	//授权API相关
	authApiClient := shopClient.GetAuthApiClient()
	shops, err := authApiClient.GetAuthorizedShops(ctx, "ttttttttttt")
	fmt.Println("授权API相关请求：", shops, err)

	//获取小部件相关API实例
	widgetClient := shopClient.GetWidgetApiClient()
	token, err := widgetClient.GetWidgetToken(ctx, "ttttttttttttt")
	fmt.Println("小部件相关API请求：", token, err)

	// 产品相关--品牌
	productClient := shopClient.GetProductApiClient()
	brands, err := productClient.GetBrands(ctx, "tttt", nil)
	fmt.Println("产品相关API品牌请求：", brands, err)

	// 产品相关--分类
	cate, err := productClient.GetCate(ctx, "tttt", nil)
	fmt.Println("产品相关API分类请求：", cate, err)
	cateRule, err := productClient.GetCateRule(ctx, "tttt", "123", nil)
	fmt.Println("产品相关API分类规则请求：", cateRule, err)
	//////请求产品
	query := map[string]string{
		"shop_cipher": "TTP_rFn99gAAAAAb_884OIFnjDUlwzNICXz1",
		"page_size":   "10",
	}
	body := map[string]any{}
	products, err := productClient.GetProducts(ctx, "TTP_ddi9XwAAAACYzgCfZsjGkD_X-fg7OFdW5M3e5X3IULLkcgnpWHrMXzyXtq8EtuP1kAsZs8o2w0cDv2keaUfZPa5TirnGuIxf8AtzO5tNXoXPZ6jdK5Io5WUKZBYMukLN5NNs-jX-aC8lJZ1cn-1xlbSvAQugaTnB_XoEnDhQv41pctrsq2Aqgw", query, body)
	fmt.Println("产品请求：", products, err)
}
