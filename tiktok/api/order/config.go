package order

type HandlingDuration struct {
	Days string `json:"days"`
	Type string `json:"type"`
}

type CombinedListingSkus struct {
	ProductID string `json:"product_id"`
	SellerSku string `json:"seller_sku"`
	SkuCount  int    `json:"sku_count"`
	SkuID     string `json:"sku_id"`
}

type ItemTax struct {
	TaxAmount string `json:"tax_amount"`
	TaxRate   string `json:"tax_rate"`
	TaxType   string `json:"tax_type"`
}

type LineItems struct {
	BuyerServiceFee      string                `json:"buyer_service_fee"`
	CancelReason         string                `json:"cancel_reason"`
	CancelUser           string                `json:"cancel_user"`
	CombinedListingSkus  []CombinedListingSkus `json:"combined_listing_skus"`
	Currency             string                `json:"currency"`
	DisplayStatus        string                `json:"display_status"`
	HandlingDurationDays string                `json:"handling_duration_days"`
	ID                   string                `json:"id"`
	IsGift               bool                  `json:"is_gift"`
	ItemTax              []ItemTax             `json:"item_tax"`
	OriginalPrice        string                `json:"original_price"`
	PackageID            string                `json:"package_id"`
	PackageStatus        string                `json:"package_status"`
	PlatformDiscount     string                `json:"platform_discount"`
	ProductID            string                `json:"product_id"`
	ProductName          string                `json:"product_name"`
	RetailDeliveryFee    string                `json:"retail_delivery_fee"`
	RtsTime              int                   `json:"rts_time"`
	SalePrice            string                `json:"sale_price"`
	SellerDiscount       string                `json:"seller_discount"`
	SellerSku            string                `json:"seller_sku"`
	ShippingProviderID   string                `json:"shipping_provider_id"`
	ShippingProviderName string                `json:"shipping_provider_name"`
	SkuID                string                `json:"sku_id"`
	SkuImage             string                `json:"sku_image"`
	SkuName              string                `json:"sku_name"`
	SkuType              string                `json:"sku_type"`
	SmallOrderFee        string                `json:"small_order_fee"`
	TrackingNumber       string                `json:"tracking_number"`
}

type Packages struct {
	Id string `json:"id"`
}

type Payment struct {
	BuyerServiceFee             string `json:"buyer_service_fee"`
	Currency                    string `json:"currency"`
	HandlingFee                 string `json:"handling_fee"`
	ItemInsuranceFee            string `json:"item_insurance_fee"`
	OriginalShippingFee         string `json:"original_shipping_fee"`
	OriginalTotalProductPrice   string `json:"original_total_product_price"`
	PlatformDiscount            string `json:"platform_discount"`
	ProductTax                  string `json:"product_tax"`
	RetailDeliveryFee           string `json:"retail_delivery_fee"`
	SellerDiscount              string `json:"seller_discount"`
	ShippingFee                 string `json:"shipping_fee"`
	ShippingFeePlatformDiscount string `json:"shipping_fee_platform_discount"`
	ShippingFeeSellerDiscount   string `json:"shipping_fee_seller_discount"`
	ShippingFeeTax              string `json:"shipping_fee_tax"`
	ShippingInsuranceFee        string `json:"shipping_insurance_fee"`
	SmallOrderFee               string `json:"small_order_fee"`
	SubTotal                    string `json:"sub_total"`
	Tax                         string `json:"tax"`
	TotalAmount                 string `json:"total_amount"`
}

type DeliveryPreferences struct {
	DropOffLocation string `json:"drop_off_location"`
}

type DistrictInfo struct {
	AddressLevel     string `json:"address_level"`
	AddressLevelName string `json:"address_level_name"`
	AddressName      string `json:"address_name"`
}

type RecipientAddress struct {
	AddressDetail string         `json:"address_detail"`
	AddressLine1  string         `json:"address_line1"`
	AddressLine2  string         `json:"address_line2"`
	AddressLine3  string         `json:"address_line3"`
	AddressLine4  string         `json:"address_line4"`
	DistrictInfo  []DistrictInfo `json:"district_info"`
	FirstName     string         `json:"first_name"`
	FullAddress   string         `json:"full_address"`
	LastName      string         `json:"last_name"`
	Name          string         `json:"name"`
	PhoneNumber   string         `json:"phone_number"`
	PostalCode    string         `json:"postal_code"`
	RegionCode    string         `json:"region_code"`
}

type Order struct {
	BuyerEmail                         string           `json:"buyer_email"`
	BuyerMessage                       string           `json:"buyer_message"`
	CancelOrderSLATime                 int              `json:"cancel_order_sla_time"`
	CancelReason                       string           `json:"cancel_reason"`
	CancelTime                         int              `json:"cancel_time"`
	CancellationInitiator              string           `json:"cancellation_initiator"`
	CollectionDueTime                  int              `json:"collection_due_time"`
	CollectionTime                     int              `json:"collection_time"`
	CommercePlatform                   string           `json:"commerce_platform"`
	Cpf                                string           `json:"cpf"`
	CreateTime                         int              `json:"create_time"`
	DeliveryDueTime                    int              `json:"delivery_due_time"`
	DeliveryOptionID                   string           `json:"delivery_option_id"`
	DeliveryOptionName                 string           `json:"delivery_option_name"`
	DeliveryOptionRequiredDeliveryTime int              `json:"delivery_option_required_delivery_time"`
	DeliverySLATime                    int              `json:"delivery_sla_time"`
	DeliveryTime                       int              `json:"delivery_time"`
	DeliveryType                       string           `json:"delivery_type"`
	FastDispatchSLATime                int              `json:"fast_dispatch_sla_time"`
	FulfillmentType                    string           `json:"fulfillment_type"`
	HandlingDuration                   HandlingDuration `json:"handling_duration"`
	HasUpdatedRecipientAddress         bool             `json:"has_updated_recipient_address"`
	ID                                 string           `json:"id"`
	IsBuyerRequestCancel               bool             `json:"is_buyer_request_cancel"`
	IsCod                              bool             `json:"is_cod"`
	IsOnHoldOrder                      bool             `json:"is_on_hold_order"`
	IsReplacementOrder                 bool             `json:"is_replacement_order"`
	IsSampleOrder                      bool             `json:"is_sample_order"`
	LineItems                          []LineItems      `json:"line_items"`
	NeedUploadInvoice                  string           `json:"need_upload_invoice"`
	OrderType                          string           `json:"order_type"`
	Packages                           []Packages       `json:"packages"`
	PaidTime                           int              `json:"paid_time"`
	Payment                            Payment          `json:"payment"`
	PaymentMethodName                  string           `json:"payment_method_name"`
	PickUpCutOffTime                   int              `json:"pick_up_cut_off_time"`
	RecipientAddress                   RecipientAddress `json:"recipient_address"`
	ReleaseDate                        int              `json:"release_date"`
	ReplacedOrderID                    string           `json:"replaced_order_id"`
	RequestCancelTime                  int              `json:"request_cancel_time"`
	RtsSLATime                         int              `json:"rts_sla_time"`
	RtsTime                            int              `json:"rts_time"`
	SellerNote                         string           `json:"seller_note"`
	ShippingDueTime                    int              `json:"shipping_due_time"`
	ShippingProvider                   string           `json:"shipping_provider"`
	ShippingProviderID                 string           `json:"shipping_provider_id"`
	ShippingType                       string           `json:"shipping_type"`
	SplitOrCombineTag                  string           `json:"split_or_combine_tag"`
	Status                             string           `json:"status"`
	TrackingNumber                     string           `json:"tracking_number"`
	TtsSLATime                         int              `json:"tts_sla_time"`
	UpdateTime                         int              `json:"update_time"`
	UserID                             string           `json:"user_id"`
	WarehouseID                        string           `json:"warehouse_id"`
}

type OrdersRsp struct {
	Orders []Order `json:"orders"`
}

type OrderResultRsp struct {
	Code     int       `json:"code"`     //逻辑状态码
	Message  string    `json:"message"`  //错误信息
	Data     OrdersRsp `json:"data"`     //数据
	HttpCode int       `json:"httpCode"` //请求tiktok的HTTP状态码
}
