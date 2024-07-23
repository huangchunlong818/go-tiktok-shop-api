package product

type Address struct {
	City          string `json:"city"`
	ContactPerson string `json:"contact_person"`
	Distict       string `json:"distict"`
	FullAddress   string `json:"full_address"`
	PhoneNumber   string `json:"phone_number"`
	PostalCode    string `json:"postal_code"`
	Region        string `json:"region"`
	RegionCode    string `json:"region_code"`
	State         string `json:"state"`
	Town          string `json:"town"`
}

type Warehouses struct {
	Address      Address `json:"address"`
	EffectStatus string  `json:"effect_status"`
	Id           string  `json:"id"`
	IsDefault    bool    `json:"is_default"`
	Name         string  `json:"name"`
	SubType      string  `json:"sub_type"`
	Type         string  `json:"type"`
}

type WarehouseRsp struct {
	Warehouses []Warehouses `json:"warehouses"`
}

type WarehousesResultRsp struct {
	Code     int          `json:"code"`     //逻辑状态码
	Message  string       `json:"message"`  //错误信息
	Data     WarehouseRsp `json:"data"`     //数据
	HttpCode int          `json:"httpCode"` //请求tiktok的HTTP状态码
}
