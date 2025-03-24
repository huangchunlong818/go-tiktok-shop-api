package product

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/huangchunlong818/go-tiktok-shop-api/tiktok/common/common"
)

// tiktok shop cate
// 获取店铺分类规则
func (b *TiktokProduct) GetCateRule(ctx context.Context, token string, cateId string, query map[string]string) CateRuleResultRsp {
	//请求接口
	r := b.SendTiktokApi(ctx, b.GetCateRuleConfig(token, cateId), query, nil, nil)
	result := CateRuleResultRsp{
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
		r.Message = "GetCateRule response error " + err.Error()
		return result
	}

	return result
}

// 获取店铺分类
func (b *TiktokProduct) GetCate(ctx context.Context, token string, query map[string]string) CateResultRsp {
	//请求接口
	r := b.SendTiktokApi(ctx, b.GetCateConfig(token), query, nil, nil)
	result := CateResultRsp{
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
		r.Message = "GetCateRule response error " + err.Error()
		return result
	}

	return result
}

// GetCateAttrs 分类属性
func (b *TiktokProduct) GetCateAttrs(ctx context.Context, token string, cateId string, query map[string]string) CateAttrsResultRsp {
	//请求接口
	r := b.SendTiktokApi(ctx, b.GetCateAttrsConfig(token, cateId), query, nil, nil)
	result := CateAttrsResultRsp{
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
		r.Message = "GetCateAttrs response error " + err.Error()
		return result
	}

	return result
}

// GetRecommendCate 获取推荐分类
func (b *TiktokProduct) GetRecommendCate(ctx context.Context, token string, query map[string]string, body map[string]any) RecommendCateResultRsp {
	//请求接口
	r := b.SendTiktokApi(ctx, b.GetRecommendCateConfig(ctx, token), query, body, nil)
	result := RecommendCateResultRsp{
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
		r.Message = "Recommend Category response error " + err.Error()
		return result
	}

	return result
}

// 分类
func (b *TiktokProduct) GetCateConfig(token string) common.GetApiConfig { //请求方式
	api := fmt.Sprintf("/product/%s/categories", b.config.Version) //请求API PATH

	return common.GetApiConfig{
		ContentType: "application/json",         //请求头content-type 类型
		Method:      "get",                      //请求方法类型
		Api:         api,                        //请求API PATH地址不带域名
		FullApi:     b.config.TkApiDomain + api, //请求的API 完整地址，带域名
		Token:       token,
	}
}

// GetCateAttrsConfig 分类属性
func (b *TiktokProduct) GetCateAttrsConfig(token string, cateId string) common.GetApiConfig { //请求方式
	api := fmt.Sprintf("/product/%s/categories/%s/attributes", b.config.Version, cateId) //请求API PATH

	return common.GetApiConfig{
		ContentType: "application/json",         //请求头content-type 类型
		Method:      "get",                      //请求方法类型
		Api:         api,                        //请求API PATH地址不带域名
		FullApi:     b.config.TkApiDomain + api, //请求的API 完整地址，带域名
		Token:       token,
	}
}

// 分类规则
func (b *TiktokProduct) GetCateRuleConfig(token string, cateId string) common.GetApiConfig { //请求方式
	api := fmt.Sprintf("/product/%s/categories/%s/rules", b.config.Version, cateId) //请求API PATH

	return common.GetApiConfig{
		ContentType: "application/json",         //请求头content-type 类型
		Method:      "get",                      //请求方法类型
		Api:         api,                        //请求API PATH地址不带域名
		FullApi:     b.config.TkApiDomain + api, //请求的API 完整地址，带域名
		Token:       token,
	}
}

// 推荐分类
func (b *TiktokProduct) GetRecommendCateConfig(ctx context.Context, token string) common.GetApiConfig { //请求方式
	api := fmt.Sprintf("/product/%s/categories/recommend", b.config.Version) //请求API PATH

	return common.GetApiConfig{
		ContentType: "application/json",         //请求头content-type 类型
		Method:      "post",                     //请求方法类型
		Api:         api,                        //请求API PATH地址不带域名
		FullApi:     b.config.TkApiDomain + api, //请求的API 完整地址，带域名
		Token:       token,
	}
}
