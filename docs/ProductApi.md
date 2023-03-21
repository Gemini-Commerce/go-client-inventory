# \ProductApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProductAddMediaGalleryEntry**](ProductApi.md#ProductAddMediaGalleryEntry) | **Post** /product.Product/AddMediaGalleryEntry | 
[**ProductBulkCreateAttribute**](ProductApi.md#ProductBulkCreateAttribute) | **Post** /product.Product/BulkCreateAttribute | 
[**ProductBulkUpdate**](ProductApi.md#ProductBulkUpdate) | **Post** /product.Product/BulkUpdate | 
[**ProductCreateAttributeOptions**](ProductApi.md#ProductCreateAttributeOptions) | **Post** /product.Product/CreateAttributeOptions | 
[**ProductCreateEntity**](ProductApi.md#ProductCreateEntity) | **Post** /product.Product/CreateEntity | 
[**ProductCreateOptionsList**](ProductApi.md#ProductCreateOptionsList) | **Post** /product.Product/CreateOptionsList | 
[**ProductCreateProduct**](ProductApi.md#ProductCreateProduct) | **Post** /product.Product/CreateProduct | 
[**ProductDelete**](ProductApi.md#ProductDelete) | **Post** /product.Product/Delete | 
[**ProductGetAttributeOption**](ProductApi.md#ProductGetAttributeOption) | **Post** /product.Product/GetAttributeOption | 
[**ProductGetAttributeOptionByCode**](ProductApi.md#ProductGetAttributeOptionByCode) | **Post** /product.Product/GetAttributeOptionByCode | 
[**ProductGetAttributeOptions**](ProductApi.md#ProductGetAttributeOptions) | **Post** /product.Product/GetAttributeOptions | 
[**ProductGetEntity**](ProductApi.md#ProductGetEntity) | **Post** /product.Product/GetEntity | 
[**ProductGetOptionsList**](ProductApi.md#ProductGetOptionsList) | **Post** /product.Product/GetOptionsList | 
[**ProductGetProduct**](ProductApi.md#ProductGetProduct) | **Post** /product.Product/GetProduct | 
[**ProductGetProductByCode**](ProductApi.md#ProductGetProductByCode) | **Post** /product.Product/GetProductByCode | 
[**ProductGetProductByUrlKey**](ProductApi.md#ProductGetProductByUrlKey) | **Post** /product.Product/GetProductByUrlKey | 
[**ProductListAttributeOptions**](ProductApi.md#ProductListAttributeOptions) | **Post** /product.Product/ListAttributeOptions | 
[**ProductListEntities**](ProductApi.md#ProductListEntities) | **Post** /product.Product/ListEntities | 
[**ProductListOptionsLists**](ProductApi.md#ProductListOptionsLists) | **Post** /product.Product/ListOptionsLists | 
[**ProductListProducts**](ProductApi.md#ProductListProducts) | **Post** /product.Product/ListProducts | 
[**ProductListProductsByIds**](ProductApi.md#ProductListProductsByIds) | **Post** /product.Product/ListProductsByIds | 
[**ProductRemoveMediaGalleryEntry**](ProductApi.md#ProductRemoveMediaGalleryEntry) | **Post** /product.Product/RemoveMediaGalleryEntry | 
[**ProductUpdateAttributeOptions**](ProductApi.md#ProductUpdateAttributeOptions) | **Post** /product.Product/UpdateAttributeOptions | 
[**ProductUpdateMediaGalleryEntry**](ProductApi.md#ProductUpdateMediaGalleryEntry) | **Post** /product.Product/UpdateMediaGalleryEntry | 
[**ProductUpdateOptionsList**](ProductApi.md#ProductUpdateOptionsList) | **Post** /product.Product/UpdateOptionsList | 
[**ProductUpdateProduct**](ProductApi.md#ProductUpdateProduct) | **Post** /product.Product/UpdateProduct | 



## ProductAddMediaGalleryEntry

> ProductAddMediaGalleryEntryResponse ProductAddMediaGalleryEntry(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewProductAddMediaGalleryEntryRequest() // ProductAddMediaGalleryEntryRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductApi.ProductAddMediaGalleryEntry(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.ProductAddMediaGalleryEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductAddMediaGalleryEntry`: ProductAddMediaGalleryEntryResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.ProductAddMediaGalleryEntry`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductAddMediaGalleryEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProductAddMediaGalleryEntryRequest**](ProductAddMediaGalleryEntryRequest.md) |  | 

### Return type

[**ProductAddMediaGalleryEntryResponse**](ProductAddMediaGalleryEntryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductBulkCreateAttribute

> EntitymanagerBulkCreateAttributeResponse ProductBulkCreateAttribute(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewEntitymanagerBulkCreateAttributeRequest() // EntitymanagerBulkCreateAttributeRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductApi.ProductBulkCreateAttribute(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.ProductBulkCreateAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductBulkCreateAttribute`: EntitymanagerBulkCreateAttributeResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.ProductBulkCreateAttribute`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductBulkCreateAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**EntitymanagerBulkCreateAttributeRequest**](EntitymanagerBulkCreateAttributeRequest.md) |  | 

### Return type

[**EntitymanagerBulkCreateAttributeResponse**](EntitymanagerBulkCreateAttributeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductBulkUpdate

> ProductBulkUpdateResponse ProductBulkUpdate(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewProductBulkUpdateRequest() // ProductBulkUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductApi.ProductBulkUpdate(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.ProductBulkUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductBulkUpdate`: ProductBulkUpdateResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.ProductBulkUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductBulkUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProductBulkUpdateRequest**](ProductBulkUpdateRequest.md) |  | 

### Return type

[**ProductBulkUpdateResponse**](ProductBulkUpdateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductCreateAttributeOptions

> EntitymanagerCreateAttributeOptionsResponse ProductCreateAttributeOptions(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewEntitymanagerCreateAttributeOptionsRequest() // EntitymanagerCreateAttributeOptionsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductApi.ProductCreateAttributeOptions(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.ProductCreateAttributeOptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductCreateAttributeOptions`: EntitymanagerCreateAttributeOptionsResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.ProductCreateAttributeOptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductCreateAttributeOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**EntitymanagerCreateAttributeOptionsRequest**](EntitymanagerCreateAttributeOptionsRequest.md) |  | 

### Return type

[**EntitymanagerCreateAttributeOptionsResponse**](EntitymanagerCreateAttributeOptionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductCreateEntity

> EntitymanagerCreateEntityResponse ProductCreateEntity(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewEntitymanagerEntity() // EntitymanagerEntity | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductApi.ProductCreateEntity(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.ProductCreateEntity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductCreateEntity`: EntitymanagerCreateEntityResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.ProductCreateEntity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductCreateEntityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**EntitymanagerEntity**](EntitymanagerEntity.md) |  | 

### Return type

[**EntitymanagerCreateEntityResponse**](EntitymanagerCreateEntityResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductCreateOptionsList

> EntitymanagerCreateOptionsListResponse ProductCreateOptionsList(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewEntitymanagerCreateOptionsListRequest() // EntitymanagerCreateOptionsListRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductApi.ProductCreateOptionsList(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.ProductCreateOptionsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductCreateOptionsList`: EntitymanagerCreateOptionsListResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.ProductCreateOptionsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductCreateOptionsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**EntitymanagerCreateOptionsListRequest**](EntitymanagerCreateOptionsListRequest.md) |  | 

### Return type

[**EntitymanagerCreateOptionsListResponse**](EntitymanagerCreateOptionsListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductCreateProduct

> ProductCreateProductResponse ProductCreateProduct(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewProductCreateProductRequest() // ProductCreateProductRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductApi.ProductCreateProduct(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.ProductCreateProduct``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductCreateProduct`: ProductCreateProductResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.ProductCreateProduct`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductCreateProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProductCreateProductRequest**](ProductCreateProductRequest.md) |  | 

### Return type

[**ProductCreateProductResponse**](ProductCreateProductResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductDelete

> ProductDeleteResponse ProductDelete(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewProductDeleteRequest() // ProductDeleteRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductApi.ProductDelete(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.ProductDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductDelete`: ProductDeleteResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.ProductDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProductDeleteRequest**](ProductDeleteRequest.md) |  | 

### Return type

[**ProductDeleteResponse**](ProductDeleteResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductGetAttributeOption

> EntitymanagerGetAttributeOptionResponse ProductGetAttributeOption(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewEntitymanagerGetAttributeOptionRequest() // EntitymanagerGetAttributeOptionRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductApi.ProductGetAttributeOption(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.ProductGetAttributeOption``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductGetAttributeOption`: EntitymanagerGetAttributeOptionResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.ProductGetAttributeOption`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductGetAttributeOptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**EntitymanagerGetAttributeOptionRequest**](EntitymanagerGetAttributeOptionRequest.md) |  | 

### Return type

[**EntitymanagerGetAttributeOptionResponse**](EntitymanagerGetAttributeOptionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductGetAttributeOptionByCode

> EntitymanagerGetAttributeOptionByCodeResponse ProductGetAttributeOptionByCode(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewEntitymanagerGetAttributeOptionByCodeRequest() // EntitymanagerGetAttributeOptionByCodeRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductApi.ProductGetAttributeOptionByCode(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.ProductGetAttributeOptionByCode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductGetAttributeOptionByCode`: EntitymanagerGetAttributeOptionByCodeResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.ProductGetAttributeOptionByCode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductGetAttributeOptionByCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**EntitymanagerGetAttributeOptionByCodeRequest**](EntitymanagerGetAttributeOptionByCodeRequest.md) |  | 

### Return type

[**EntitymanagerGetAttributeOptionByCodeResponse**](EntitymanagerGetAttributeOptionByCodeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductGetAttributeOptions

> EntitymanagerGetAttributeOptionsResponse ProductGetAttributeOptions(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewEntitymanagerGetAttributeOptionsRequest() // EntitymanagerGetAttributeOptionsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductApi.ProductGetAttributeOptions(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.ProductGetAttributeOptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductGetAttributeOptions`: EntitymanagerGetAttributeOptionsResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.ProductGetAttributeOptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductGetAttributeOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**EntitymanagerGetAttributeOptionsRequest**](EntitymanagerGetAttributeOptionsRequest.md) |  | 

### Return type

[**EntitymanagerGetAttributeOptionsResponse**](EntitymanagerGetAttributeOptionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductGetEntity

> EntitymanagerEntity ProductGetEntity(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewEntitymanagerEntityRequest() // EntitymanagerEntityRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductApi.ProductGetEntity(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.ProductGetEntity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductGetEntity`: EntitymanagerEntity
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.ProductGetEntity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductGetEntityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**EntitymanagerEntityRequest**](EntitymanagerEntityRequest.md) |  | 

### Return type

[**EntitymanagerEntity**](EntitymanagerEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductGetOptionsList

> EntitymanagerGetOptionsListResponse ProductGetOptionsList(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewEntitymanagerGetOptionsListRequest() // EntitymanagerGetOptionsListRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductApi.ProductGetOptionsList(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.ProductGetOptionsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductGetOptionsList`: EntitymanagerGetOptionsListResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.ProductGetOptionsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductGetOptionsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**EntitymanagerGetOptionsListRequest**](EntitymanagerGetOptionsListRequest.md) |  | 

### Return type

[**EntitymanagerGetOptionsListResponse**](EntitymanagerGetOptionsListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductGetProduct

> ProductGetProductResponse ProductGetProduct(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewProductGetProductRequest() // ProductGetProductRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductApi.ProductGetProduct(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.ProductGetProduct``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductGetProduct`: ProductGetProductResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.ProductGetProduct`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductGetProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProductGetProductRequest**](ProductGetProductRequest.md) |  | 

### Return type

[**ProductGetProductResponse**](ProductGetProductResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductGetProductByCode

> ProductGetProductByCodeResponse ProductGetProductByCode(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewProductGetProductByCodeRequest() // ProductGetProductByCodeRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductApi.ProductGetProductByCode(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.ProductGetProductByCode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductGetProductByCode`: ProductGetProductByCodeResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.ProductGetProductByCode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductGetProductByCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProductGetProductByCodeRequest**](ProductGetProductByCodeRequest.md) |  | 

### Return type

[**ProductGetProductByCodeResponse**](ProductGetProductByCodeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductGetProductByUrlKey

> ProductGetProductByUrlKeyResponse ProductGetProductByUrlKey(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewProductGetProductByUrlKeyRequest() // ProductGetProductByUrlKeyRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductApi.ProductGetProductByUrlKey(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.ProductGetProductByUrlKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductGetProductByUrlKey`: ProductGetProductByUrlKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.ProductGetProductByUrlKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductGetProductByUrlKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProductGetProductByUrlKeyRequest**](ProductGetProductByUrlKeyRequest.md) |  | 

### Return type

[**ProductGetProductByUrlKeyResponse**](ProductGetProductByUrlKeyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductListAttributeOptions

> EntitymanagerListAttributeOptionsResponse ProductListAttributeOptions(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewEntitymanagerListAttributeOptionsRequest() // EntitymanagerListAttributeOptionsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductApi.ProductListAttributeOptions(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.ProductListAttributeOptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductListAttributeOptions`: EntitymanagerListAttributeOptionsResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.ProductListAttributeOptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductListAttributeOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**EntitymanagerListAttributeOptionsRequest**](EntitymanagerListAttributeOptionsRequest.md) |  | 

### Return type

[**EntitymanagerListAttributeOptionsResponse**](EntitymanagerListAttributeOptionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductListEntities

> EntitymanagerListEntitiesResponse ProductListEntities(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewEntitymanagerListEntitiesRequest() // EntitymanagerListEntitiesRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductApi.ProductListEntities(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.ProductListEntities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductListEntities`: EntitymanagerListEntitiesResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.ProductListEntities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductListEntitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**EntitymanagerListEntitiesRequest**](EntitymanagerListEntitiesRequest.md) |  | 

### Return type

[**EntitymanagerListEntitiesResponse**](EntitymanagerListEntitiesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductListOptionsLists

> EntitymanagerListOptionsListsResponse ProductListOptionsLists(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewEntitymanagerListOptionsListsRequest() // EntitymanagerListOptionsListsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductApi.ProductListOptionsLists(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.ProductListOptionsLists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductListOptionsLists`: EntitymanagerListOptionsListsResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.ProductListOptionsLists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductListOptionsListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**EntitymanagerListOptionsListsRequest**](EntitymanagerListOptionsListsRequest.md) |  | 

### Return type

[**EntitymanagerListOptionsListsResponse**](EntitymanagerListOptionsListsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductListProducts

> ProductListProductsResponse ProductListProducts(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewProductListProductsRequest() // ProductListProductsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductApi.ProductListProducts(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.ProductListProducts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductListProducts`: ProductListProductsResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.ProductListProducts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductListProductsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProductListProductsRequest**](ProductListProductsRequest.md) |  | 

### Return type

[**ProductListProductsResponse**](ProductListProductsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductListProductsByIds

> ProductListProductsByIdsResponse ProductListProductsByIds(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewProductListProductsByIdsRequest() // ProductListProductsByIdsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductApi.ProductListProductsByIds(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.ProductListProductsByIds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductListProductsByIds`: ProductListProductsByIdsResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.ProductListProductsByIds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductListProductsByIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProductListProductsByIdsRequest**](ProductListProductsByIdsRequest.md) |  | 

### Return type

[**ProductListProductsByIdsResponse**](ProductListProductsByIdsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductRemoveMediaGalleryEntry

> map[string]interface{} ProductRemoveMediaGalleryEntry(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewProductRemoveMediaGalleryEntryRequest() // ProductRemoveMediaGalleryEntryRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductApi.ProductRemoveMediaGalleryEntry(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.ProductRemoveMediaGalleryEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductRemoveMediaGalleryEntry`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.ProductRemoveMediaGalleryEntry`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductRemoveMediaGalleryEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProductRemoveMediaGalleryEntryRequest**](ProductRemoveMediaGalleryEntryRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductUpdateAttributeOptions

> EntitymanagerUpdateAttributeOptionsResponse ProductUpdateAttributeOptions(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewEntitymanagerUpdateAttributeOptionsRequest() // EntitymanagerUpdateAttributeOptionsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductApi.ProductUpdateAttributeOptions(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.ProductUpdateAttributeOptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductUpdateAttributeOptions`: EntitymanagerUpdateAttributeOptionsResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.ProductUpdateAttributeOptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductUpdateAttributeOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**EntitymanagerUpdateAttributeOptionsRequest**](EntitymanagerUpdateAttributeOptionsRequest.md) |  | 

### Return type

[**EntitymanagerUpdateAttributeOptionsResponse**](EntitymanagerUpdateAttributeOptionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductUpdateMediaGalleryEntry

> map[string]interface{} ProductUpdateMediaGalleryEntry(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewProductUpdateMediaGalleryEntryRequest() // ProductUpdateMediaGalleryEntryRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductApi.ProductUpdateMediaGalleryEntry(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.ProductUpdateMediaGalleryEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductUpdateMediaGalleryEntry`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.ProductUpdateMediaGalleryEntry`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductUpdateMediaGalleryEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProductUpdateMediaGalleryEntryRequest**](ProductUpdateMediaGalleryEntryRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductUpdateOptionsList

> EntitymanagerUpdateOptionsListResponse ProductUpdateOptionsList(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewEntitymanagerUpdateOptionsListRequest() // EntitymanagerUpdateOptionsListRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductApi.ProductUpdateOptionsList(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.ProductUpdateOptionsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductUpdateOptionsList`: EntitymanagerUpdateOptionsListResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.ProductUpdateOptionsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductUpdateOptionsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**EntitymanagerUpdateOptionsListRequest**](EntitymanagerUpdateOptionsListRequest.md) |  | 

### Return type

[**EntitymanagerUpdateOptionsListResponse**](EntitymanagerUpdateOptionsListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductUpdateProduct

> ProductUpdateProductResponse ProductUpdateProduct(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewProductUpdateProductRequest() // ProductUpdateProductRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductApi.ProductUpdateProduct(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.ProductUpdateProduct``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductUpdateProduct`: ProductUpdateProductResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.ProductUpdateProduct`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductUpdateProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProductUpdateProductRequest**](ProductUpdateProductRequest.md) |  | 

### Return type

[**ProductUpdateProductResponse**](ProductUpdateProductResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

