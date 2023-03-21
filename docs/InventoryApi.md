# \InventoryApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InventoryAdjustQty**](InventoryApi.md#InventoryAdjustQty) | **Post** /inventory.Inventory/AdjustQty | 
[**InventoryAdjustQtyCommitted**](InventoryApi.md#InventoryAdjustQtyCommitted) | **Post** /inventory.Inventory/AdjustQtyCommitted | adjust operations are intended for increment/decrement quantities fields by a certain amount
[**InventoryCreateStockItem**](InventoryApi.md#InventoryCreateStockItem) | **Post** /inventory.Inventory/CreateStockItem | 
[**InventoryGetQtySalable**](InventoryApi.md#InventoryGetQtySalable) | **Post** /inventory.Inventory/GetQtySalable | 
[**InventoryGetStockItem**](InventoryApi.md#InventoryGetStockItem) | **Post** /inventory.Inventory/GetStockItem | 
[**InventoryListStockItems**](InventoryApi.md#InventoryListStockItems) | **Post** /inventory.Inventory/ListStockItems | 
[**InventoryListStockItemsBySkus**](InventoryApi.md#InventoryListStockItemsBySkus) | **Post** /inventory.Inventory/ListStockItemsBySkus | 
[**InventoryRebalanceStockQtys**](InventoryApi.md#InventoryRebalanceStockQtys) | **Post** /inventory.Inventory/RebalanceStockQtys | rebalance operation is intended for apply or discard a certain amount of qty_committed to qty and old qty_committed
[**InventoryUpdateStockItem**](InventoryApi.md#InventoryUpdateStockItem) | **Post** /inventory.Inventory/UpdateStockItem | 



## InventoryAdjustQty

> InventoryStockItem InventoryAdjustQty(ctx).Body(body).Execute()



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
    body := *openapiclient.NewInventoryAdjustQtyRequest() // InventoryAdjustQtyRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InventoryApi.InventoryAdjustQty(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.InventoryAdjustQty``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InventoryAdjustQty`: InventoryStockItem
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.InventoryAdjustQty`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInventoryAdjustQtyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**InventoryAdjustQtyRequest**](InventoryAdjustQtyRequest.md) |  | 

### Return type

[**InventoryStockItem**](InventoryStockItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryAdjustQtyCommitted

> InventoryStockItem InventoryAdjustQtyCommitted(ctx).Body(body).Execute()

adjust operations are intended for increment/decrement quantities fields by a certain amount

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
    body := *openapiclient.NewInventoryAdjustQtyCommittedRequest() // InventoryAdjustQtyCommittedRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InventoryApi.InventoryAdjustQtyCommitted(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.InventoryAdjustQtyCommitted``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InventoryAdjustQtyCommitted`: InventoryStockItem
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.InventoryAdjustQtyCommitted`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInventoryAdjustQtyCommittedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**InventoryAdjustQtyCommittedRequest**](InventoryAdjustQtyCommittedRequest.md) |  | 

### Return type

[**InventoryStockItem**](InventoryStockItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryCreateStockItem

> InventoryStockItem InventoryCreateStockItem(ctx).Body(body).Execute()



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
    body := *openapiclient.NewInventoryCreateStockItemRequest() // InventoryCreateStockItemRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InventoryApi.InventoryCreateStockItem(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.InventoryCreateStockItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InventoryCreateStockItem`: InventoryStockItem
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.InventoryCreateStockItem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInventoryCreateStockItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**InventoryCreateStockItemRequest**](InventoryCreateStockItemRequest.md) |  | 

### Return type

[**InventoryStockItem**](InventoryStockItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryGetQtySalable

> InventoryGetQtySalableResponse InventoryGetQtySalable(ctx).Body(body).Execute()



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
    body := *openapiclient.NewInventoryGetQtySalableRequest() // InventoryGetQtySalableRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InventoryApi.InventoryGetQtySalable(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.InventoryGetQtySalable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InventoryGetQtySalable`: InventoryGetQtySalableResponse
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.InventoryGetQtySalable`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInventoryGetQtySalableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**InventoryGetQtySalableRequest**](InventoryGetQtySalableRequest.md) |  | 

### Return type

[**InventoryGetQtySalableResponse**](InventoryGetQtySalableResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryGetStockItem

> InventoryStockItem InventoryGetStockItem(ctx).Body(body).Execute()



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
    body := *openapiclient.NewInventoryGetStockItemRequest() // InventoryGetStockItemRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InventoryApi.InventoryGetStockItem(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.InventoryGetStockItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InventoryGetStockItem`: InventoryStockItem
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.InventoryGetStockItem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInventoryGetStockItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**InventoryGetStockItemRequest**](InventoryGetStockItemRequest.md) |  | 

### Return type

[**InventoryStockItem**](InventoryStockItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryListStockItems

> InventoryListStockItemsResponse InventoryListStockItems(ctx).Body(body).Execute()



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
    body := *openapiclient.NewInventoryListStockItemsRequest() // InventoryListStockItemsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InventoryApi.InventoryListStockItems(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.InventoryListStockItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InventoryListStockItems`: InventoryListStockItemsResponse
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.InventoryListStockItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInventoryListStockItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**InventoryListStockItemsRequest**](InventoryListStockItemsRequest.md) |  | 

### Return type

[**InventoryListStockItemsResponse**](InventoryListStockItemsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryListStockItemsBySkus

> InventoryListStockItemsBySkusResponse InventoryListStockItemsBySkus(ctx).Body(body).Execute()



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
    body := *openapiclient.NewInventoryListStockItemsBySkusRequest() // InventoryListStockItemsBySkusRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InventoryApi.InventoryListStockItemsBySkus(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.InventoryListStockItemsBySkus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InventoryListStockItemsBySkus`: InventoryListStockItemsBySkusResponse
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.InventoryListStockItemsBySkus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInventoryListStockItemsBySkusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**InventoryListStockItemsBySkusRequest**](InventoryListStockItemsBySkusRequest.md) |  | 

### Return type

[**InventoryListStockItemsBySkusResponse**](InventoryListStockItemsBySkusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryRebalanceStockQtys

> InventoryStockItem InventoryRebalanceStockQtys(ctx).Body(body).Execute()

rebalance operation is intended for apply or discard a certain amount of qty_committed to qty and old qty_committed

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
    body := *openapiclient.NewInventoryRebalanceStockQtysRequest() // InventoryRebalanceStockQtysRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InventoryApi.InventoryRebalanceStockQtys(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.InventoryRebalanceStockQtys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InventoryRebalanceStockQtys`: InventoryStockItem
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.InventoryRebalanceStockQtys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInventoryRebalanceStockQtysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**InventoryRebalanceStockQtysRequest**](InventoryRebalanceStockQtysRequest.md) |  | 

### Return type

[**InventoryStockItem**](InventoryStockItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InventoryUpdateStockItem

> InventoryStockItem InventoryUpdateStockItem(ctx).Body(body).Execute()



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
    body := *openapiclient.NewInventoryUpdateStockItemRequest() // InventoryUpdateStockItemRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InventoryApi.InventoryUpdateStockItem(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.InventoryUpdateStockItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InventoryUpdateStockItem`: InventoryStockItem
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.InventoryUpdateStockItem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInventoryUpdateStockItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**InventoryUpdateStockItemRequest**](InventoryUpdateStockItemRequest.md) |  | 

### Return type

[**InventoryStockItem**](InventoryStockItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

