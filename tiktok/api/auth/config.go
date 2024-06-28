package auth

type Shops struct {
	Cipher     string `json:"cipher"`      //请求接口需要用到的
	Code       string `json:"code"`        //商店代码
	Id         string `json:"id"`          //商店ID
	Name       string `json:"name"`        //商店名称
	Region     string `json:"region"`      //商店所在区域，国家代码
	SellerType string `json:"seller_type"` //商店类型？1代表跨境店铺CROSS_BORDER， 2代表本地店铺LOCAL
}
