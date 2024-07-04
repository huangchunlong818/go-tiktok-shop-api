package product

// 分类规则响应结构体
type CateRuleResultRsp struct {
	Code     int         `json:"code"`     //逻辑状态码
	Message  string      `json:"message"`  //错误信息
	Data     CateRuleRsp `json:"data"`     //数据
	HttpCode int         `json:"httpCode"` //请求tiktok的HTTP状态码
}

// 分类规则相应  注释查看proto
type CateRuleRsp struct {
	Cod                   *Cod                    `json:"cod"`
	Epr                   *Epr                    `json:"epr"`
	PackageDimension      *PackageDimension       `json:"package_dimension"`
	ProductCertifications []ProductCertifications `json:"product_certifications"`
	SizeChart             *SizeChart              `json:"size_chart"`
}

type Cod struct {
	IsSupported bool `json:"is_supported"`
}

type Epr struct {
	IsRequired bool `json:"is_required"`
}

type PackageDimension struct {
	IsRequired bool `json:"is_required"`
}

type ProductCertifications struct {
	Id             string `json:"id"`
	IsRequired     bool   `json:"is_required"`
	Name           string `json:"name"`
	SampleImageUrl string `json:"sample_image_url"`
}

type SizeChart struct {
	IsRequired  bool `json:"is_required"`
	IsSupported bool `json:"is_supported"`
}

// 分类响应结构体
type CateResultRsp struct {
	Code     int     `json:"code"`     //逻辑状态码
	Message  string  `json:"message"`  //错误信息
	Data     CateRsp `json:"data"`     //数据
	HttpCode int     `json:"httpCode"` //请求tiktok的HTTP状态码
}

// 分类响应
type CateRsp struct {
	Cate []Cate `json:"cate"`
}

type Cate struct {
	Id                 string   `json:"id"`                  //分类ID
	IsLeaf             bool     `json:"is_leaf"`             //该类别是否为叶类别。仅支持使用叶类别创建和编辑产品
	LocalName          string   `json:"local_name"`          //商店经营所在国家/地区的类别名称
	ParentId           string   `json:"parent_id"`           //父类别ID，一级类别的父类别ID为“0”
	PermissionStatuses []string `json:"permission_statuses"` //卖家对类别的权限 1.AVAILABLE：您拥有该类别的权限，可以在该类别下创建产品。2. INVITE_ONLY ：该类别是邀请类别，您不能选择类别创建产品。请放心客户经理或店铺支持团队以访问此类别或选择其他类别的权限。3、NON_MAIN_CATEGORY：该类目不在卖家店铺主类目范围内，卖家无权使用。建议商家联系AM进行处理
}

// 品牌响应结构体
type BrandsResultRsp struct {
	Code     int       `json:"code"`     //逻辑状态码
	Message  string    `json:"message"`  //错误信息
	Data     BrandsRsp `json:"data"`     //数据
	HttpCode int       `json:"httpCode"` //请求tiktok的HTTP状态码
}

// 品牌响应
type BrandsRsp struct {
	Brands        []Brands `json:"brands"`
	NextPageToken string   `json:"next_page_token"`
	TotalCount    int      `json:"total_count"`
}

type Brands struct {
	AuthorizedStatus string `json:"authorized_status"` //品牌授权情况：未经授权 UNAUTHORIEZD   授权 AUTHORIZED
	BrandStatus      string `json:"brand_status"`      //返回品牌的可用状态。如果品牌不符合 NICE 分类，则会被标记为不可用。可用的 AVAILABLE   不可用 UNAVAILABLE
	Id               string `json:"id"`                //brand id
	IsT1Brand        bool   `json:"is_t1_brand"`       //是否T1品牌
	Name             string `json:"name"`              //brand name
}

// 产品相应
type ProductsResultRsp struct {
	Code     int         `json:"code"`     //逻辑状态码
	Message  string      `json:"message"`  //错误信息
	Data     ProductsRsp `json:"data"`     //数据
	HttpCode int         `json:"httpCode"` //请求tiktok的HTTP状态码
}

// 产品相应
type ProductsRsp struct {
	NextPageToken string     `json:"next_page_token"`
	Products      []Products `json:"products"`
	TotalCount    int        `json:"total_count"`
}

type Inventory struct {
	Quantity    int    `json:"quantity"`
	WarehouseId string `json:"warehouse_id"`
}

type Price struct {
	Currency          string `json:"currency"`
	SalePrice         string `json:"sale_price"`
	TaxExclusivePrice string `json:"tax_exclusive_price"`
}

type Skus struct {
	Id        string      `json:"id"`
	Inventory []Inventory `json:"inventory"`
	Price     Price       `json:"price"`
	SellerSku string      `json:"seller_sku"`
}

type Products struct {
	CreateTime             int      `json:"create_time"`
	Id                     string   `json:"id"`
	IsNotForSale           bool     `json:"is_not_for_sale"`
	ProductSyncFailReasons []string `json:"product_sync_fail_reasons"`
	SalesRegions           []string `json:"sales_regions"`
	Skus                   []Skus   `json:"skus"`
	Status                 string   `json:"status"`
	Title                  string   `json:"title"`
	UpdateTime             int      `json:"update_time"`
}
