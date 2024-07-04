package product

import (
	"context"
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
	//断言分页token
	if next, ok := r.Data["next_page_token"].(string); !ok {
		r.Code = common.ErrCode
		r.Message = "GetProducts next_page_token response error"
		return result
	} else {
		result.Data.NextPageToken = next
	}

	//断言总数
	if total, ok := r.Data["total_count"].(float64); !ok {
		r.Code = common.ErrCode
		r.Message = "GetProducts total_count response error"
		return result
	} else {
		result.Data.TotalCount = int(total)
	}

	//断言列表
	products, err := b.CheckSliceAny(r.Data["products"])
	if err != nil {
		r.Code = common.ErrCode
		r.Message = "GetProducts products" + err.Error()
		return result
	}
	if len(products) < 1 {
		return result
	}

	//获取具体产品
	for _, product := range products {
		if tmp, err := b.CheckMapStringAny(product); err == nil && tmp != nil {
			//product_sync_fail_reasons
			productSyncFailReasons, err := b.ChangeAnyToStringSlice(tmp["product_sync_fail_reasons"])
			if err != nil {
				r.Code = common.ErrCode
				r.Message = "GetProducts product_sync_fail_reasons" + err.Error()
				return result
			}

			//sales_regions
			salesRegions, err := b.ChangeAnyToStringSlice(tmp["sales_regions"])
			if err != nil {
				r.Code = common.ErrCode
				r.Message = "GetProducts sales_regions" + err.Error()
				return result
			}

			//skus
			var skus []Skus
			if skusTmp, err := b.CheckSliceAny(tmp["skus"]); err == nil {
				for _, sku := range skusTmp {
					if skusString, err := b.CheckMapStringAny(sku); err == nil {
						price, _ := b.CheckMapStringAny(skusString["price"])
						var inventorys []Inventory
						inventory, _ := b.CheckSliceAny(skusString["inventory"])
						if len(inventory) > 0 {
							for _, value := range inventory {
								if tmpInventory, err := b.CheckMapStringAny(value); err == nil {
									inventorys = append(inventorys, Inventory{
										Quantity:    int(tmpInventory["quantity"].(float64)),
										WarehouseId: tmpInventory["warehouse_id"].(string),
									})
								}
							}
						}
						skus = append(skus, Skus{
							Id:        skusString["id"].(string),
							Inventory: inventorys,
							Price: Price{
								Currency:          b.CheckString(price["currency"]),
								SalePrice:         b.CheckString(price["sale_price"]),
								TaxExclusivePrice: b.CheckString(price["tax_exclusive_price"]),
							},
							SellerSku: skusString["seller_sku"].(string),
						})
					}
				}
			}

			result.Data.Products = append(result.Data.Products, Products{
				CreateTime:             int(tmp["create_time"].(float64)),
				Id:                     tmp["id"].(string),
				IsNotForSale:           tmp["is_not_for_sale"].(bool),
				ProductSyncFailReasons: productSyncFailReasons,
				SalesRegions:           salesRegions,
				Skus:                   skus,
				Status:                 tmp["status"].(string),
				Title:                  tmp["title"].(string),
				UpdateTime:             int(tmp["update_time"].(float64)),
			})
		}
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
