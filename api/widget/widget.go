package widget

import (
	"context"
	"errors"
	"tiktok-shop-api/tiktok/common"
)

//tiktok shop

type TiktokWidget struct {
}

var newServer *TiktokWidget

// 获取实例
func GetNewService() *TiktokWidget {
	if newServer == nil {
		newServer = &TiktokWidget{}
	}
	return newServer
}

// 获取所有授权店铺
func (s *TiktokWidget) GetWidgetToken(ctx context.Context, token string) (result Token, err error) {
	//请求接口
	r, err := common.SendTiktokApi(ctx, GetWidgetToken(token), nil, nil)

	//断言所有店铺数据 是一个 any切片
	data, err := common.CheckMapStringAny(r["widget_token"])
	if err != nil {
		err = errors.New("GetWidgetToken widget_token " + err.Error())
	}
	if data != nil {
		result = Token{
			Token:    data["token"].(string),
			ExpireAt: int64(data["expire_at"].(float64)),
		}
	}

	return
}
