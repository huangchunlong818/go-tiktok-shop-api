package logistics

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/huangchunlong818/go-tiktok-shop-api/tiktok/common/common"
)

// GetWarehouses 获取仓库
func (tl *TiktokLogistics) GetWarehouses(ctx context.Context, token string, query map[string]string) WarehousesResultRsp {
	//请求接口
	r := tl.SendTiktokApi(ctx, tl.GetWarehousesConfig(token), query, nil, nil)
	result := WarehousesResultRsp{
		Code:     r.Code,
		Message:  r.Message,
		HttpCode: r.HttpCode,
	}
	if !tl.IsSuccess(r) {
		return result
	}

	//解析数据
	err := json.Unmarshal(r.Data, &result)
	if err != nil {
		r.Code = common.ErrCode
		r.Message = "GetWarehouses response error " + err.Error()
		return result
	}

	return result
}

func (tl *TiktokLogistics) GetWarehousesConfig(token string) common.GetApiConfig { //请求方式
	api := fmt.Sprintf("/logistics/%s/warehouses", tl.config.Version) //请求API PATH

	return common.GetApiConfig{
		ContentType: "application/json",          //请求头content-type 类型
		Method:      "get",                       //请求方法类型
		Api:         api,                         //请求API PATH地址不带域名
		FullApi:     tl.config.TkApiDomain + api, //请求的API 完整地址，带域名
		Token:       token,
	}
}
