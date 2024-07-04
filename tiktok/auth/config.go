package auth

// 整体相应结构
type GetTokenByAuthCodeRsp struct {
	Code     int                    `json:"code"`     //逻辑状态码
	Message  string                 `json:"message"`  //错误信息
	Data     GetTokenByAuthCodeData `json:"data"`     //数据
	HttpCode int                    `json:"httpCode"` //请求tiktok的HTTP状态码
}

// 定义一个结构体来表示 data 部分的数据
type GetTokenByAuthCodeData struct {
	AccessToken          string   `json:"access_token"`
	AccessTokenExpireIn  int64    `json:"access_token_expire_in"`
	RefreshToken         string   `json:"refresh_token"`
	RefreshTokenExpireIn int64    `json:"refresh_token_expire_in"`
	OpenID               string   `json:"open_id"`
	SellerName           string   `json:"seller_name"`
	SellerBaseRegion     string   `json:"seller_base_region"`
	UserType             int      `json:"user_type"`
	GrantedScopes        []string `json:"granted_scopes"`
}
