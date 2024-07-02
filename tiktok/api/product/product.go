package product

import (
	"context"
	"errors"
	"fmt"
	"github.com/huangchunlong818/go-tiktok-shop-api/tiktok/common/common"
	"github.com/huangchunlong818/go-tiktok-shop-api/tiktok/common/config"
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
	GetProducts(ctx context.Context, token string, query map[string]string, body map[string]any) (result ProductsRsp, err error)
}

// 获取产品，搜索产品
func (b *TiktokProduct) GetProducts(ctx context.Context, token string, query map[string]string, body map[string]any) (result ProductsRsp, err error) {
	//请求接口
	r, err := b.common.SendTiktokApi(ctx, b.GetProductsConfig(token), query, body)
	if err != nil {
		return
	}

	//断言分页token
	if next, ok := r["next_page_token"].(string); !ok {
		return result, errors.New("GetProducts next_page_token response error")
	} else {
		result.NextPageToken = next
	}

	//断言总数
	if total, ok := r["total_count"].(float64); !ok {
		return result, errors.New("GetProducts total_count response error")
	} else {
		result.TotalCount = int(total)
	}

	//断言列表
	products, err := b.common.CheckSliceAny(r["products"])
	if err != nil {
		err = errors.New("GetProducts products" + err.Error())
		return result, err
	}
	if len(products) < 1 {
		return result, nil
	}

	//获取具体产品
	for _, product := range products {
		if tmp, err := b.common.CheckMapStringAny(product); err == nil && tmp != nil {
			//product_sync_fail_reasons
			productSyncFailReasons, err := b.common.ChangeAnyToStringSlice(tmp["product_sync_fail_reasons"])
			if err != nil {
				return result, err
			}

			//sales_regions
			salesRegions, err := b.common.ChangeAnyToStringSlice(tmp["sales_regions"])
			if err != nil {
				return result, err
			}

			//skus
			var skus []Skus
			if skusTmp, err := b.common.CheckSliceAny(tmp["skus"]); err == nil {
				for _, sku := range skusTmp {
					if skusString, err := b.common.CheckMapStringAny(sku); err == nil {
						price, _ := b.common.CheckMapStringAny(skusString["price"])
						var inventorys []Inventory
						inventory, _ := b.common.CheckSliceAny(tmp)
						if len(inventory) > 0 {
							for _, value := range inventory {
								if tmpInventory, err := b.common.CheckMapStringAny(value); err != nil {
									inventorys = append(inventorys, Inventory{
										Quantity:    tmpInventory["quantity"].(int),
										WarehouseId: tmpInventory["warehouse_id"].(string),
									})
								}
							}
						}
						skus = append(skus, Skus{
							Id:        skusString["id"].(string),
							Inventory: inventorys,
							Price: Price{
								Currency:          b.common.CheckString(price["currency"]),
								SalePrice:         b.common.CheckString(price["sale_price"]),
								TaxExclusivePrice: b.common.CheckString(price["tax_exclusive_price"]),
							},
							SellerSku: skusString["seller_sku"].(string),
						})
					}
				}
			}

			result.Products = append(result.Products, Products{
				CreateTime:             int(tmp["create_time"].(float64)),
				Id:                     tmp["id"].(string),
				IsNotForSale:           tmp["is_not_for_sale"].(bool),
				ProductSyncFailReasons: productSyncFailReasons,
				SalesRegions:           salesRegions,
				Skus:                   nil,
				Status:                 tmp["status"].(string),
				Title:                  tmp["title"].(string),
				UpdateTime:             int(tmp["update_time"].(float64)),
			})
		}
	}

	return
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
			config: config,
			common: common.GetNewService(config),
		}
	}
	return newServer
}
