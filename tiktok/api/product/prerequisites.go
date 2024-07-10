package product

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/huangchunlong818/go-tiktok-shop-api/tiktok/common/common"
)

// GetPrerequisites 获取店铺的商品规则以及是否满足上架商品的条件
func (b *TiktokProduct) GetPrerequisites(ctx context.Context, token string, query map[string]string) PrerequisitesResultRsp {
	//请求接口
	r := b.SendTiktokApi(ctx, b.GetPrerequisitesConfig(token), query, nil)
	result := PrerequisitesResultRsp{
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
		r.Message = "GetPrerequisites response error " + err.Error()
		return result
	}

	return result
}

func (b *TiktokProduct) GetPrerequisitesConfig(token string) common.GetApiConfig { //请求方式
	api := fmt.Sprintf("/product/%s/prerequisites", b.config.PrerequisitesVersion) //请求API PATH

	return common.GetApiConfig{
		ContentType: "application/json",         //请求头content-type 类型
		Method:      "get",                      //请求方法类型
		Api:         api,                        //请求API PATH地址不带域名
		FullApi:     b.config.TkApiDomain + api, //请求的API 完整地址，带域名
		Token:       token,
	}
}
