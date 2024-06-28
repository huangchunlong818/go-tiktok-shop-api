package product

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
