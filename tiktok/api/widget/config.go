package widget

type Token struct {
	Token    string `json:"token"`     //小部件token
	ExpireAt int64  `json:"expire_at"` //过期时间
}

// 整体相应结构
type GetTokenRsp struct {
	Code     int    `json:"code"`     //逻辑状态码
	Message  string `json:"message"`  //错误信息
	Data     Token  `json:"data"`     //数据
	HttpCode int    `json:"httpCode"` //请求tiktok的HTTP状态码
}
