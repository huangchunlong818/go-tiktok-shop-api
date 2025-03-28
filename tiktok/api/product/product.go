package product

import (
	"context"
	"encoding/json"
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
	GetCateAttrs(ctx context.Context, token string, cateId string, query map[string]string) CateAttrsResultRsp
	GetCateConfig(token string) common.GetApiConfig
	GetCateRuleConfig(token string, cateId string) common.GetApiConfig
	GetCateAttrsConfig(token string, cateId string) common.GetApiConfig

	// 获取推荐分类
	GetRecommendCateConfig(ctx context.Context, token string) common.GetApiConfig
	GetRecommendCate(ctx context.Context, token string, query map[string]string, body map[string]any) RecommendCateResultRsp
	
	// 获取店铺的商品规则以及是否满足上架商品的条件
	GetPrerequisites(ctx context.Context, token string, query map[string]string) PrerequisitesResultRsp
	GetPrerequisitesConfig(token string) common.GetApiConfig

	//产品
	GetProducts(ctx context.Context, token string, query map[string]string, body map[string]any) ProductsResultRsp
	GetProductsConfig(token string) common.GetApiConfig

	//产品详情
	GetProduct(ctx context.Context, token string, productId string, query map[string]string) ProductResultRsp
	GetProductConfig(token string, productId string) common.GetApiConfig

	// 创建产品
	CreateProduct(ctx context.Context, token string, query map[string]string, body map[string]any) CreateProductResultRsp
	CreateProductConfig(token string) common.GetApiConfig

	// 删除产品
	DeleteProducts(ctx context.Context, token string, query map[string]string, body map[string]any) DeleteProductsResultRsp
	DeleteProductsConfig(token string) common.GetApiConfig

	// 产品下架
	DeactivateProducts(ctx context.Context, token string, query map[string]string, body map[string]any) DeactivateProductsResultRsp
	DeactivateProductsConfig(token string) common.GetApiConfig

	// 产品上架
	ActivateProducts(ctx context.Context, token string, query map[string]string, body map[string]any) ActivateProductsResultRsp
	ActivateProductsConfig(token string) common.GetApiConfig

	// 修改部分产品信息
	PartialEditProduct(ctx context.Context, token string, productId string, query map[string]string, body map[string]any) PartialEditProductResultRsp
	PartialEditProductConfig(token string, productId string) common.GetApiConfig

	//修改产品价格
	UpdateProductPrice(ctx context.Context, token string, productId string, query map[string]string, body map[string]any) UpdateProductPriceResultRsp
	UpdateProductPriceConfig(token string, productId string) common.GetApiConfig

	// 产品图片上传
	ImageUpload(ctx context.Context, token string, body map[string]any, filePath string) ImageUploadResultRsp
	GetImageUploadConfig(token string) common.GetApiConfig

	// 产品图片压缩
	OptimizedImages(ctx context.Context, token string, query map[string]string, body map[string]any) OptimizedImagesResultRsp
	GetOptimizedImagesConfig(token string) common.GetApiConfig

	// 产品附件上传
	FileUpload(ctx context.Context, token string, query map[string]string, body map[string]any, filePath string) FileUploadResultRsp
	GetFileUploadConfig(token string) common.GetApiConfig

	// 检测产品发布字段
	CheckProductListing(ctx context.Context, token string, query map[string]string, body map[string]any) CheckProductListingResultRsp
	CheckProductListingConfig(token string) common.GetApiConfig
}

// 获取产品，搜索产品
func (b *TiktokProduct) GetProducts(ctx context.Context, token string, query map[string]string, body map[string]any) ProductsResultRsp {
	//请求接口
	r := b.SendTiktokApi(ctx, b.GetProductsConfig(token), query, body, nil)
	result := ProductsResultRsp{
		Code:     r.Code,
		Message:  r.Message,
		HttpCode: r.HttpCode,
	}
	if !b.IsSuccess(r) {
		return result
	}

	//解析数据
	err := json.Unmarshal(r.Data, &result)
	if err != nil {
		r.Code = common.ErrCode
		r.Message = "GetBrandsApi response error " + err.Error()
		return result
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

// GetProduct 获取产品详情
func (b *TiktokProduct) GetProduct(ctx context.Context, token string, productId string, query map[string]string) ProductResultRsp {
	//请求接口
	r := b.SendTiktokApi(ctx, b.GetProductConfig(token, productId), query, nil, nil)
	result := ProductResultRsp{
		Code:     r.Code,
		Message:  r.Message,
		HttpCode: r.HttpCode,
	}
	if !b.IsSuccess(r) {
		return result
	}

	//解析数据
	err := json.Unmarshal(r.Data, &result)
	if err != nil {
		r.Code = common.ErrCode
		r.Message = "GetProductConfig response error " + err.Error()
		return result
	}

	return result
}

// GetProductConfig 产品详情
func (b *TiktokProduct) GetProductConfig(token string, productId string) common.GetApiConfig { //请求方式
	api := fmt.Sprintf("/product/%s/products/%s", b.config.Version, productId) //请求API PATH

	return common.GetApiConfig{
		ContentType: "application/json",         //请求头content-type 类型
		Method:      "get",                      //请求方法类型
		Api:         api,                        //请求API PATH地址不带域名
		FullApi:     b.config.TkApiDomain + api, //请求的API 完整地址，带域名
		Token:       token,
	}
}

// CreateProduct 创建产品
func (b *TiktokProduct) CreateProduct(ctx context.Context, token string, query map[string]string, body map[string]any) CreateProductResultRsp {
	//请求接口
	r := b.SendTiktokApi(ctx, b.CreateProductConfig(token), query, body, nil)
	result := CreateProductResultRsp{
		Code:     r.Code,
		Message:  r.Message,
		HttpCode: r.HttpCode,
	}
	if !b.IsSuccess(r) {
		return result
	}

	//解析数据
	err := json.Unmarshal(r.Data, &result)
	if err != nil {
		r.Code = common.ErrCode
		r.Message = "CreateProduct response error " + err.Error()
		return result
	}

	return result
}

func (b *TiktokProduct) CreateProductConfig(token string) common.GetApiConfig { //请求方式
	api := fmt.Sprintf("/product/%s/products", b.config.Version) //请求API PATH

	return common.GetApiConfig{
		ContentType: "application/json",         //请求头content-type 类型
		Method:      "post",                     //请求方法类型
		Api:         api,                        //请求API PATH地址不带域名
		FullApi:     b.config.TkApiDomain + api, //请求的API 完整地址，带域名
		Token:       token,
	}
}

// PartialEditProduct 修改产品的部分信息
func (b *TiktokProduct) PartialEditProduct(ctx context.Context, token string, productId string, query map[string]string, body map[string]any) PartialEditProductResultRsp {
	//请求接口
	r := b.SendTiktokApi(ctx, b.PartialEditProductConfig(token, productId), query, body, nil)
	result := PartialEditProductResultRsp{
		Code:     r.Code,
		Message:  r.Message,
		HttpCode: r.HttpCode,
	}
	if !b.IsSuccess(r) {
		return result
	}

	//解析数据
	err := json.Unmarshal(r.Data, &result)
	if err != nil {
		r.Code = common.ErrCode
		r.Message = "PartialEditProduct response error " + err.Error()
		return result
	}

	return result
}

func (b *TiktokProduct) PartialEditProductConfig(token string, productId string) common.GetApiConfig { //请求方式
	api := fmt.Sprintf("/product/%s/products/%s/partial_edit", b.config.Version, productId) //请求API PATH

	return common.GetApiConfig{
		ContentType: "application/json",         //请求头content-type 类型
		Method:      "post",                     //请求方法类型
		Api:         api,                        //请求API PATH地址不带域名
		FullApi:     b.config.TkApiDomain + api, //请求的API 完整地址，带域名
		Token:       token,
	}
}

// DeleteProducts 删除产品
func (b *TiktokProduct) DeleteProducts(ctx context.Context, token string, query map[string]string, body map[string]any) DeleteProductsResultRsp {
	//请求接口
	r := b.SendTiktokApi(ctx, b.DeleteProductsConfig(token), query, body, nil)
	result := DeleteProductsResultRsp{
		Code:     r.Code,
		Message:  r.Message,
		HttpCode: r.HttpCode,
	}
	if !b.IsSuccess(r) {
		return result
	}

	//解析数据
	err := json.Unmarshal(r.Data, &result)
	if err != nil {
		r.Code = common.ErrCode
		r.Message = "DeleteProducts response error " + err.Error()
		return result
	}

	return result
}

func (b *TiktokProduct) DeleteProductsConfig(token string) common.GetApiConfig { //请求方式
	api := fmt.Sprintf("/product/%s/products", b.config.Version) //请求API PATH

	return common.GetApiConfig{
		ContentType: "application/json",         //请求头content-type 类型
		Method:      "delete",                   //请求方法类型
		Api:         api,                        //请求API PATH地址不带域名
		FullApi:     b.config.TkApiDomain + api, //请求的API 完整地址，带域名
		Token:       token,
	}
}

// DeactivateProducts 产品下架
func (b *TiktokProduct) DeactivateProducts(ctx context.Context, token string, query map[string]string, body map[string]any) DeactivateProductsResultRsp {
	//请求接口
	r := b.SendTiktokApi(ctx, b.DeactivateProductsConfig(token), query, body, nil)
	result := DeactivateProductsResultRsp{
		Code:     r.Code,
		Message:  r.Message,
		HttpCode: r.HttpCode,
	}
	if !b.IsSuccess(r) {
		return result
	}

	//解析数据
	err := json.Unmarshal(r.Data, &result)
	if err != nil {
		r.Code = common.ErrCode
		r.Message = "DeactivateProducts response error " + err.Error()
		return result
	}

	return result
}

func (b *TiktokProduct) DeactivateProductsConfig(token string) common.GetApiConfig { //请求方式
	api := fmt.Sprintf("/product/%s/products/deactivate", b.config.Version) //请求API PATH

	return common.GetApiConfig{
		ContentType: "application/json",         //请求头content-type 类型
		Method:      "post",                     //请求方法类型
		Api:         api,                        //请求API PATH地址不带域名
		FullApi:     b.config.TkApiDomain + api, //请求的API 完整地址，带域名
		Token:       token,
	}
}

// ActivateProducts 商品上架
func (b *TiktokProduct) ActivateProducts(ctx context.Context, token string, query map[string]string, body map[string]any) ActivateProductsResultRsp {
	//请求接口
	r := b.SendTiktokApi(ctx, b.ActivateProductsConfig(token), query, body, nil)
	result := ActivateProductsResultRsp{
		Code:     r.Code,
		Message:  r.Message,
		HttpCode: r.HttpCode,
	}
	if !b.IsSuccess(r) {
		return result
	}

	//解析数据
	err := json.Unmarshal(r.Data, &result)
	if err != nil {
		r.Code = common.ErrCode
		r.Message = "ActivateProducts response error " + err.Error()
		return result
	}

	return result
}

func (b *TiktokProduct) ActivateProductsConfig(token string) common.GetApiConfig { //请求方式
	api := fmt.Sprintf("/product/%s/products/activate", b.config.Version) //请求API PATH

	return common.GetApiConfig{
		ContentType: "application/json",         //请求头content-type 类型
		Method:      "post",                     //请求方法类型
		Api:         api,                        //请求API PATH地址不带域名
		FullApi:     b.config.TkApiDomain + api, //请求的API 完整地址，带域名
		Token:       token,
	}
}

// UpdateProductPrice 修改产品价格
func (b *TiktokProduct) UpdateProductPrice(ctx context.Context, token string, productId string, query map[string]string, body map[string]any) UpdateProductPriceResultRsp {
	//请求接口
	r := b.SendTiktokApi(ctx, b.UpdateProductPriceConfig(token, productId), query, body, nil)
	result := UpdateProductPriceResultRsp{
		Code:     r.Code,
		Message:  r.Message,
		HttpCode: r.HttpCode,
	}
	if !b.IsSuccess(r) {
		return result
	}

	//解析数据
	err := json.Unmarshal(r.Data, &result)
	if err != nil {
		r.Code = common.ErrCode
		r.Message = "UpdateProductPrice response error " + err.Error()
		return result
	}

	return result
}

func (b *TiktokProduct) UpdateProductPriceConfig(token string, productId string) common.GetApiConfig { //请求方式
	api := fmt.Sprintf("/product/%s/products/%s/prices/update", b.config.Version, productId) //请求API PATH

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
