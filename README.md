<h1>Tiktok Shop Api - Golang</h1>

## Api Version
[Tiktok Shop Api 202309](https://partner.tiktokshop.com/docv2/page/64f198f74830a5028854c106)

## Install
```
# Project introduction
go get github.com/huangchunlong818/go-tiktok-shop-api

# Renew
go get -u github.com/huangchunlong818/go-tiktok-shop-api

# Specifying a version
go get -u github.com/huangchunlong818/go-tiktok-shop-api@v0.2.4
```

## Development Environment
go version go1.20.5

## Example
### Initialization
```go
import (
    "context"
    "github.com/huangchunlong818/go-tiktok-shop-api/tiktok/common/config"
)

ctx := context.Background()

config := config.WithApp(config.AppConfig{
    AppId:  "xxxx",
    AppKey: "xxxx",
    Secret: "xxxx",
})

// Get instance object
shopClient := tiktokShop.NewTiktokShopClient(config)

var query map[string]string
var body map[string]any
```

### Authorization related
```go
// Get auth url
authClient := shopClient.GetAuthClient()
authUrl := authClient.GetAuthUrl("us")

authApiClient := shopClient.GetAuthApiClient()
getAuthorizedShops := authApiClient.GetAuthorizedShops(ctx, "Access Token")
```

### Products related
```go
// Get shop brands
productApiClient := shopClient.GetProductApiClient()
brands := productApiClient.GetBrands(ctx, "Access Token", query)

// Upload Image
productApiClient := shopClient.GetProductApiClient()
imageUpload := productApiClient.ImageUpload(ctx, "Access Token", body, "image path")
```

### Logistics related
```go
// Get shop warehouses
widgetApiClient := shopClient.GetWidgetApiClient()
warehouses := widgetApiClient.GetWarehouses(ctx, "Access Token", query)
```

## Supported APIs
### Authorization
1. [Get Authorized Shops - GetAuthorizedShops](https://partner.tiktokshop.com/docv2/page/6507ead7b99d5302be949ba9?external_id=6507ead7b99d5302be949ba9#Back%20To%20Top)
### Product
1. [Get Brands - GetBrands](https://partner.tiktokshop.com/docv2/page/6503075656e2bb0289dd5d01)
2. [Get Category Rules - GetCateRule](https://partner.tiktokshop.com/docv2/page/6509c0febace3e02b74594a9?external_id=6509c0febace3e02b74594a9#Back%20To%20Top)
3. [Get Categories - GetCate](https://partner.tiktokshop.com/docv2/page/6509c89d0fcef602bf1acd9b?external_id=6509c89d0fcef602bf1acd9b)
4. [Get Attributes - GetCateAttrs](https://partner.tiktokshop.com/docv2/page/6509c5784a0bb702c0561cc8?external_id=6509c5784a0bb702c0561cc8)
5. [Check Listing Prerequisites - GetPrerequisites](https://partner.tiktokshop.com/docv2/page/6571ae94c5524602c081d0bb?external_id=6571ae94c5524602c081d0bb)
6. [Search Products - GetProducts](https://partner.tiktokshop.com/docv2/page/65854ffb8f559302d8a6acda?external_id=65854ffb8f559302d8a6acda)
7. [Get Product - GetProduct](https://partner.tiktokshop.com/docv2/page/6509d85b4a0bb702c057fdda?external_id=6509d85b4a0bb702c057fdda)
8. [Create Product - CreateProduct](https://partner.tiktokshop.com/docv2/page/6502fc8da57708028b42b18a?external_id=6502fc8da57708028b42b18a)
9. [Delete Products - DeleteProducts](https://partner.tiktokshop.com/docv2/page/6503079ebb2a4d028d515acf?external_id=6503079ebb2a4d028d515acf)
10. [Deactivate Products - DeactivateProducts](https://partner.tiktokshop.com/docv2/page/6509de450fcef602bf1d087c?external_id=6509de450fcef602bf1d087c)
11. [Activate Product - ActivateProducts](https://partner.tiktokshop.com/docv2/page/650306ff5a12ff0294eab4a9?external_id=650306ff5a12ff0294eab4a9)
12. [Partial Edit Product - PartialEditProduct](https://partner.tiktokshop.com/docv2/page/650a98d74a0bb702c06c3289?external_id=650a98d74a0bb702c06c3289)
13. [Update Price - UpdateProductPrice](https://partner.tiktokshop.com/docv2/page/650307de5a12ff0294eac8b0?external_id=650307de5a12ff0294eac8b0)
14. [Upload Product Image - ImageUpload](https://partner.tiktokshop.com/docv2/page/6509df95defece02be598a22?external_id=6509df95defece02be598a22)
15. [Optimized Images - OptimizedImages](https://partner.tiktokshop.com/docv2/page/665692b35d39dc02deb49a97?external_id=665692b35d39dc02deb49a97)
16. [Upload Product File - FileUpload](https://partner.tiktokshop.com/docv2/page/6509dffdc16ffe02b8dc10c5?external_id=6509dffdc16ffe02b8dc10c5)
17. [Check Product Listing - CheckProductListing](https://partner.tiktokshop.com/docv2/page/650a0ee8f1fd3102b91c6493?external_id=650a0ee8f1fd3102b91c6493)
### Logistics
1. [Get Warehouse List - GetWarehouses](https://partner.tiktokshop.com/docv2/page/650aa418defece02be6e66b6?external_id=650aa418defece02be6e66b6)

## More content to come
