package product

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/huangchunlong818/go-tiktok-shop-api/tiktok/common/common"
)

//tiktok shop product file upload

// FileUpload 文件上传（PDF）
func (b *TiktokProduct) FileUpload(ctx context.Context, token string, body map[string]any, filePath string) FileUploadResultRsp {
	//请求接口
	r := b.SendTiktokApi(ctx, b.GetFileUploadConfig(token), nil, body, map[string]string{
		"data": filePath,
	})
	result := FileUploadResultRsp{
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
		r.Message = "GetBrands response error " + err.Error()
		return result
	}

	return result
}

// GetFileUploadConfig 文件上传
func (b *TiktokProduct) GetFileUploadConfig(token string) common.GetApiConfig { //请求方式
	api := fmt.Sprintf("/product/%s/files/upload", b.config.Version) //请求API PATH

	return common.GetApiConfig{
		ContentType: "multipart/form-data",      //请求头content-type 类型
		Method:      "post",                     //请求方法类型
		Api:         api,                        //请求API PATH地址不带域名
		FullApi:     b.config.TkApiDomain + api, //请求的API 完整地址，带域名
		Token:       token,
	}
}
