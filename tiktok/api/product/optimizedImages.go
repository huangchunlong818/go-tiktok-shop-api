package product

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/huangchunlong818/go-tiktok-shop-api/tiktok/common/common"
)

//tiktok shop product optimized Images

// OptimizedImages 图片优化
func (b *TiktokProduct) OptimizedImages(ctx context.Context, token string, query map[string]string, body map[string]any) OptimizedImagesResultRsp {
	//请求接口
	r := b.SendTiktokApi(ctx, b.GetOptimizedImagesConfig(token), query, body, nil)
	result := OptimizedImagesResultRsp{
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
		r.Message = "OptimizedImages response error " + err.Error()
		return result
	}

	return result
}

// GetOptimizedImagesConfig 图片优化版本
func (b *TiktokProduct) GetOptimizedImagesConfig(token string) common.GetApiConfig { //请求方式
	api := fmt.Sprintf("/product/%s/images/optimize", b.config.OptimizedImagesVersion) //请求API PATH

	return common.GetApiConfig{
		ContentType: "application/json",         //请求头content-type 类型
		Method:      "post",                     //请求方法类型
		Api:         api,                        //请求API PATH地址不带域名
		FullApi:     b.config.TkApiDomain + api, //请求的API 完整地址，带域名
		Token:       token,
	}
}
