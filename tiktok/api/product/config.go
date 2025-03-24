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
	Cod                   Cod                     `json:"cod"`
	Epr                   Epr                     `json:"epr"`
	PackageDimension      PackageDimension        `json:"package_dimension"`
	ProductCertifications []ProductCertifications `json:"product_certifications"`
	SizeChart             SizeChart               `json:"size_chart"`
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
	Id                    string                 `json:"id"`
	IsRequired            bool                   `json:"is_required"`
	Name                  string                 `json:"name"`
	SampleImageUrl        string                 `json:"sample_image_url"`
	RequirementConditions []RequirementCondition `json:"requirement_conditions"`
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
	Cate []Cate `json:"categories"`
}

type Cate struct {
	Id                 string   `json:"id"`                  //分类ID
	IsLeaf             bool     `json:"is_leaf"`             //该类别是否为叶类别。仅支持使用叶类别创建和编辑产品
	LocalName          string   `json:"local_name"`          //商店经营所在国家/地区的类别名称
	ParentId           string   `json:"parent_id"`           //父类别ID，一级类别的父类别ID为“0”
	PermissionStatuses []string `json:"permission_statuses"` //卖家对类别的权限 1.AVAILABLE：您拥有该类别的权限，可以在该类别下创建产品。2. INVITE_ONLY ：该类别是邀请类别，您不能选择类别创建产品。请放心客户经理或店铺支持团队以访问此类别或选择其他类别的权限。3、NON_MAIN_CATEGORY：该类目不在卖家店铺主类目范围内，卖家无权使用。建议商家联系AM进行处理
}

// 图片优化结构
type OptimizedImagesResultRsp struct {
	Code     int                 `json:"code"`     //逻辑状态码
	Message  string              `json:"message"`  //错误信息
	Data     OptimizedImagesData `json:"data"`     //数据
	HttpCode int                 `json:"httpCode"` //请求tiktok的HTTP状态码
}

type OptimizedImagesData struct {
	Images []Images `json:"images"`
}

type OptimizedImages struct {
	Height         int    `json:"height"`
	OptimizeStatus string `json:"optimize_status"`
	OptimizedUri   string `json:"optimized_uri"`
	OptimizedUrl   string `json:"optimized_url"`
	OriginalUri    string `json:"original_uri"`
	OriginalUrl    string `json:"original_url"`
	Width          int    `json:"width"`
}

// 文件上传结构
type FileUploadResultRsp struct {
	Code     int            `json:"code"`     //逻辑状态码
	Message  string         `json:"message"`  //错误信息
	Data     FileUploadData `json:"data"`     //数据
	HttpCode int            `json:"httpCode"` //请求tiktok的HTTP状态码
}

type FileUploadData struct {
	Format string `json:"format"`
	Id     string `json:"id"`
	Name   string `json:"name"`
	Url    string `json:"url"`
}

// 图片上传结构
type ImageUploadResultRsp struct {
	Code     int             `json:"code"`     //逻辑状态码
	Message  string          `json:"message"`  //错误信息
	Data     ImageUploadData `json:"data"`     //数据
	HttpCode int             `json:"httpCode"` //请求tiktok的HTTP状态码
}

type ImageUploadData struct {
	Height  int    `json:"height"`
	Uri     string `json:"uri"`
	Url     string `json:"url"`
	UseCase string `json:"use_case"`
	Width   int    `json:"width"`
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

// CateAttrsResultRsp 分类属性
type CateAttrsResultRsp struct {
	Code     int          `json:"code"`     //逻辑状态码
	Message  string       `json:"message"`  //错误信息
	Data     CateAttrsRsp `json:"data"`     //数据
	HttpCode int          `json:"httpCode"` //请求tiktok的HTTP状态码
}

type CateAttrsRsp struct {
	CateAttrs []CateAttr `json:"attributes"`
}

type CateAttr struct {
	Id                    string                 `json:"id"`
	IsCustomizable        bool                   `json:"is_customizable"`
	IsMultipleSelection   bool                   `json:"is_multiple_selection"`
	IsRequired            bool                   `json:"is_requried"`
	Name                  string                 `json:"name"`
	Type                  string                 `json:"type"`
	Values                []Values               `json:"values"`
	RequirementConditions []RequirementCondition `json:"requirement_conditions"`
	ValueDataFormat       string                 `json:"value_data_format"`
}

type RequirementCondition struct {
	AttributeId      string `json:"attribute_id"`
	AttributeValueId string `json:"attribute_value_id"`
	ConditionType    string `json:"condition_type"`
}

type Values struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// PrerequisitesResultRsp 店铺的商品规则以及是否满足上架商品的条件
type PrerequisitesResultRsp struct {
	Code     int              `json:"code"`     //逻辑状态码
	Message  string           `json:"message"`  //错误信息
	Data     PrerequisitesRsp `json:"data"`     //数据
	HttpCode int              `json:"httpCode"` //请求tiktok的HTTP状态码
}

type PrerequisitesRsp struct {
	CheckResults []CheckResults `json:"check_results"`
}

type CheckResults struct {
	CheckItem   string   `json:"check_item"`
	FailReasons []string `json:"fail_reasons"`
	IsFailed    bool     `json:"is_failed"`
}

// ProductResultRsp 产品详情响应
type ProductResultRsp struct {
	Code     int     `json:"code"`     //逻辑状态码
	Message  string  `json:"message"`  //错误信息
	Data     Product `json:"data"`     //数据
	HttpCode int     `json:"httpCode"` //请求tiktok的HTTP状态码
}

// CategoryChain 分类链
type CategoryChain struct {
	Id        string `json:"id"`
	ParentId  string `json:"parent_id"`
	LocalName string `json:"local_name"`
	IsLeaf    bool   `json:"is_leaf"`
}

type Brand struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Image struct {
	Height    int      `json:"height"`
	Width     int      `json:"width"`
	ThumbUrls []string `json:"thumb_urls"`
	Uri       string   `json:"uri"`
	Urls      []string `json:"urls"`
}

type File struct {
	Id     string   `json:"id"`
	Urls   []string `json:"urls"`
	Name   string   `json:"name"`
	Format string   `json:"format"`
}

type Video struct {
	Id       string `json:"id"`
	CoverUrl string `json:"cover_url"`
	Format   string `json:"format"`
	Url      string `json:"url"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	Size     int    `json:"size"`
}

type PackageDimensions struct {
	Length string `json:"length"`
	Width  string `json:"width"`
	Height string `json:"height"`
	Unit   string `json:"unit"`
}

type PackageWeight struct {
	Value string `json:"value"`
	Unit  string `json:"unit"`
}

type SkuPrice struct {
	TaxExclusivePrice string `json:"tax_exclusive_price"`
	SalePrice         string `json:"sale_price"`
	Currency          string `json:"currency"`
	UnitPrice         string `json:"unit_price"`
}

type SkuInventory struct {
	WarehouseId string `json:"warehouse_id"`
	Quantity    int    `json:"quantity"`
}

type SkuIdentifierCode struct {
	Code string `json:"code"`
	Type string `json:"type"`
}

type CombinedSku struct {
	ProductId string `json:"product_id"`
	SkuId     string `json:"sku_id"`
	SkuCount  int    `json:"sku_count"`
}

type ReplicateSource struct {
	ProductId string `json:"product_id"`
	ShopId    string `json:"shop_id"`
	SkuId     string `json:"sku_id"`
}

type GlobalListingPolicy struct {
	PriceSync       bool            `json:"price_sync"`
	InventoryType   string          `json:"inventory_type"`
	ReplicateSource ReplicateSource `json:"replicate_Source"`
}

type SalesAttribute struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	ValueId   string `json:"value_id"`
	ValueName string `json:"value_name"`
	SkuImg    Image  `json:"sku_img"`
}

type Sku struct {
	Id                  string              `json:"id"`
	SellerSku           string              `json:"seller_sku"`
	Price               SkuPrice            `json:"price"`
	Inventory           []SkuInventory      `json:"inventory"`
	IdentifierCode      SkuIdentifierCode   `json:"identifier_code"`
	SalesAttributes     []SalesAttribute    `json:"sales_attributes"`
	ExternalSkuId       string              `json:"external_sku_id"`
	CombinedSkus        []CombinedSku       `json:"combined_skus"`
	GlobalListingPolicy GlobalListingPolicy `json:"global_listing_policy"`
	SkuUnitCount        string              `json:"sku_unit_count"`
}

type Certification struct {
	Id     string  `json:"id"`
	Title  string  `json:"title"`
	Files  []File  `json:"files"`
	Images []Image `json:"images"`
}

type Template struct {
	Id string `json:"id"`
}

type ProductSizeChart struct {
	Image    Image    `json:"image"`
	Template Template `json:"template"`
}

type Value struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type ProductAttribute struct {
	Id     string  `json:"id"`
	Name   string  `json:"name"`
	Values []Value `json:"values"`
}

type AuditFailedReason struct {
	Position    string   `json:"position"`
	Reasons     []string `json:"reasons"`
	Suggestions []string `json:"suggestions"`
}

type DeliveryOption struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	IsAvailable bool   `json:"is_available"`
}

type Manufacturer struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}

type Product struct {
	Id                 string              `json:"id"`
	Status             string              `json:"status"`
	Title              string              `json:"title"`
	CategoryChains     []CategoryChain     `json:"category_chains"`
	Brand              Brand               `json:"brand"`
	MainImages         []Image             `json:"main_images"`
	Video              Video               `json:"video"`
	Description        string              `json:"description"`
	PackageDimensions  PackageDimensions   `json:"package_dimensions"`
	PackageWeight      PackageWeight       `json:"package_weight"`
	Skus               []Sku               `json:"skus"`
	Certifications     []Certification     `json:"certifications"`
	SizeChart          ProductSizeChart    `json:"size_chart"`
	IsCodAllowed       bool                `json:"is_cod_allowed"`
	ProductAttributes  []ProductAttribute  `json:"product_attributes"`
	AuditFailedReasons []AuditFailedReason `json:"audit_failed_reasons"`
	UpdateTime         int                 `json:"update_time"`
	CreateTime         int                 `json:"create_time"`
	DeliveryOptions    []DeliveryOption    `json:"delivery_options"`
	ExternalProductId  string              `json:"external_product_id"`
	ProductTypes       []string            `json:"product_types"`
	Manufacturer       Manufacturer        `json:"manufacturer"`
	IsNotForSale       bool                `json:"is_not_for_sale"`
}

// Check Product Listing
type DiagnosisResults struct {
	Code       string `json:"code"`
	HowToSolve string `json:"how_to_solve"`
}

type Images struct {
	Height       int    `json:"height"`
	OptimizedUri string `json:"optimized_uri"`
	OptimizedUrl string `json:"optimized_url"`
	Uri          string `json:"uri"`
	Url          string `json:"url"`
	Width        int    `json:"width"`
}

type SeoWords struct {
	Text string `json:"text"`
}

type SmartTexts struct {
	Text string `json:"text"`
}

type Suggestions struct {
	Images     []Images     `json:"images"`
	SeoWords   []SeoWords   `json:"seo_words"`
	SmartTexts []SmartTexts `json:"smart_texts"`
}

type Diagnoses struct {
	DiagnosisResults []DiagnosisResults `json:"diagnosis_results"`
	Field            string             `json:"field"`
	Suggestions      Suggestions        `json:"suggestions"`
}

type FailReasons struct {
	Message string `json:"message"`
}

type Warnings struct {
	Message string `json:"message"`
}

type CheckProductListingRsp struct {
	CheckResult string        `json:"check_result"`
	Diagnoses   []Diagnoses   `json:"diagnoses"`
	FailReasons []FailReasons `json:"fail_reasons"`
	Warnings    Warnings      `json:"warnings"`
}

type CheckProductListingResultRsp struct {
	Code     int                    `json:"code"`     //逻辑状态码
	Message  string                 `json:"message"`  //错误信息
	Data     CheckProductListingRsp `json:"data"`     //数据
	HttpCode int                    `json:"httpCode"` //请求tiktok的HTTP状态码
}

type SalesAttributes struct {
	Id      string `json:"id"`
	ValueId string `json:"value_id"`
}

type CreateProductSkus struct {
	ExternalSkuId   string            `json:"external_sku_id"`
	Id              string            `json:"id"`
	SalesAttributes []SalesAttributes `json:"sales_attributes"`
	SellerSku       string            `json:"seller_sku"`
}

type WarningMessages struct {
	Message string `json:"message"`
}

type CreateProductRsp struct {
	ProductId string              `json:"product_id"`
	Skus      []CreateProductSkus `json:"skus"`
	Warnings  []WarningMessages   `json:"warnings"`
}

type CreateProductResultRsp struct {
	Code     int              `json:"code"`     //逻辑状态码
	Message  string           `json:"message"`  //错误信息
	Data     CreateProductRsp `json:"data"`     //数据
	HttpCode int              `json:"httpCode"` //请求tiktok的HTTP状态码
}

type PartialEditProductRsp struct {
	ProductId string              `json:"product_id"`
	Skus      []CreateProductSkus `json:"skus"`
}

type PartialEditProductResultRsp struct {
	Code     int                   `json:"code"`     //逻辑状态码
	Message  string                `json:"message"`  //错误信息
	Data     PartialEditProductRsp `json:"data"`     //数据
	HttpCode int                   `json:"httpCode"` //请求tiktok的HTTP状态码
}

type UpdateProductPriceRsp struct {
}

type UpdateProductPriceResultRsp struct {
	Code     int                   `json:"code"`     //逻辑状态码
	Message  string                `json:"message"`  //错误信息
	Data     UpdateProductPriceRsp `json:"data"`     //数据
	HttpCode int                   `json:"httpCode"` //请求tiktok的HTTP状态码
}

type DeleteProductsRsp struct {
	Errors []Errors `json:"errors"`
}

type Detail struct {
	ProductId string `json:"product_id"`
}

type Errors struct {
	Code    int    `json:"code"`
	Detail  Detail `json:"detail"`
	Message string `json:"message"`
}

type DeleteProductsResultRsp struct {
	Code     int               `json:"code"`     //逻辑状态码
	Message  string            `json:"message"`  //错误信息
	Data     DeleteProductsRsp `json:"data"`     //数据
	HttpCode int               `json:"httpCode"` //请求tiktok的HTTP状态码
}

type DeactivateProductsRsp struct {
	Errors []Errors `json:"errors"`
}

type DeactivateProductsResultRsp struct {
	Code     int                   `json:"code"`     //逻辑状态码
	Message  string                `json:"message"`  //错误信息
	Data     DeactivateProductsRsp `json:"data"`     //数据
	HttpCode int                   `json:"httpCode"` //请求tiktok的HTTP状态码
}

type ActivateProductsRsp struct {
	Errors []Errors `json:"errors"`
}

type ActivateProductsExtraErrors struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ActivateProductsDetail struct {
	ExtraErrors []ActivateProductsExtraErrors `json:"extra_errors"`
	ProductId   string                        `json:"product_id"`
}

type ActivateProductsErrors struct {
	Code    int                    `json:"code"`
	Detail  ActivateProductsDetail `json:"detail"`
	Message string                 `json:"message"`
}

type ActivateProductsResultRsp struct {
	Code     int                 `json:"code"`     //逻辑状态码
	Message  string              `json:"message"`  //错误信息
	Data     ActivateProductsRsp `json:"data"`     //数据
	HttpCode int                 `json:"httpCode"` //请求tiktok的HTTP状态码
}

type RecommendCate struct {
	Id                 string   `json:"id"`
	Name               string   `json:"name"`
	Level              int      `json:"level"`
	IsLeaf             bool     `json:"is_leaf"`
	PermissionStatuses []string `json:"permission_statuses"`
}

type RecommendCateRsp struct {
	LeafCategoryId string          `json:"leaf_category_id"`
	Categories     []RecommendCate `json:"categories"`
}

type RecommendCateResultRsp struct {
	Code     int              `json:"code"`     //逻辑状态码
	Message  string           `json:"message"`  //错误信息
	Data     RecommendCateRsp `json:"data"`     //数据
	HttpCode int              `json:"httpCode"` //请求tiktok的HTTP状态码
}
