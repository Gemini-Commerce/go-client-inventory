# \InventoryAPI

All URIs are relative to *https://inventory.api.gogemini.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdjustQty**](InventoryAPI.md#AdjustQty) | **Post** /inventory.Inventory/AdjustQty | Adjust Quantity
[**AdjustQtyCommitted**](InventoryAPI.md#AdjustQtyCommitted) | **Post** /inventory.Inventory/AdjustQtyCommitted | Adjust Quantity Committed
[**CreateStockItem**](InventoryAPI.md#CreateStockItem) | **Post** /inventory.Inventory/CreateStockItem | Create Stock Item
[**GetQtySalable**](InventoryAPI.md#GetQtySalable) | **Post** /inventory.Inventory/GetQtySalable | Get Salable Quantity
[**GetStockItem**](InventoryAPI.md#GetStockItem) | **Post** /inventory.Inventory/GetStockItem | Get Stock Item
[**ListStockItems**](InventoryAPI.md#ListStockItems) | **Post** /inventory.Inventory/ListStockItems | List Stock Items
[**ListStockItemsBySkus**](InventoryAPI.md#ListStockItemsBySkus) | **Post** /inventory.Inventory/ListStockItemsBySkus | List Stock Items by SKUs
[**RebalanceStockQtys**](InventoryAPI.md#RebalanceStockQtys) | **Post** /inventory.Inventory/RebalanceStockQtys | Rebalance Stock Quantities
[**UpdateStockItem**](InventoryAPI.md#UpdateStockItem) | **Post** /inventory.Inventory/UpdateStockItem | Update Stock Item



## AdjustQty

> InventoryStockItem AdjustQty(ctx).Body(body).Execute()

Adjust Quantity



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/gemini-commerce/go-client-inventory"
)

func main() {
    body := *openapiclient.NewInventoryAdjustQtyRequest() // InventoryAdjustQtyRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InventoryAPI.AdjustQty(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryAPI.AdjustQty``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdjustQty`: InventoryStockItem
    fmt.Fprintf(os.Stdout, "Response from `InventoryAPI.AdjustQty`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdjustQtyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**InventoryAdjustQtyRequest**](InventoryAdjustQtyRequest.md) |  | 

### Return type

[**InventoryStockItem**](InventoryStockItem.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdjustQtyCommitted

> InventoryStockItem AdjustQtyCommitted(ctx).Body(body).Execute()

Adjust Quantity Committed



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/gemini-commerce/go-client-inventory"
)

func main() {
    body := *openapiclient.NewInventoryAdjustQtyCommittedRequest() // InventoryAdjustQtyCommittedRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InventoryAPI.AdjustQtyCommitted(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryAPI.AdjustQtyCommitted``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdjustQtyCommitted`: InventoryStockItem
    fmt.Fprintf(os.Stdout, "Response from `InventoryAPI.AdjustQtyCommitted`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdjustQtyCommittedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**InventoryAdjustQtyCommittedRequest**](InventoryAdjustQtyCommittedRequest.md) |  | 

### Return type

[**InventoryStockItem**](InventoryStockItem.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateStockItem

> InventoryStockItem CreateStockItem(ctx).Body(body).Execute()

Create Stock Item

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/gemini-commerce/go-client-inventory"
)

func main() {
    body := *openapiclient.NewInventoryCreateStockItemRequest() // InventoryCreateStockItemRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InventoryAPI.CreateStockItem(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryAPI.CreateStockItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateStockItem`: InventoryStockItem
    fmt.Fprintf(os.Stdout, "Response from `InventoryAPI.CreateStockItem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateStockItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**InventoryCreateStockItemRequest**](InventoryCreateStockItemRequest.md) |  | 

### Return type

[**InventoryStockItem**](InventoryStockItem.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQtySalable

> InventoryGetQtySalableResponse GetQtySalable(ctx).Body(body).Execute()

Get Salable Quantity

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/gemini-commerce/go-client-inventory"
)

func main() {
    body := *openapiclient.NewInventoryGetQtySalableRequest() // InventoryGetQtySalableRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InventoryAPI.GetQtySalable(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryAPI.GetQtySalable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetQtySalable`: InventoryGetQtySalableResponse
    fmt.Fprintf(os.Stdout, "Response from `InventoryAPI.GetQtySalable`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetQtySalableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**InventoryGetQtySalableRequest**](InventoryGetQtySalableRequest.md) |  | 

### Return type

[**InventoryGetQtySalableResponse**](InventoryGetQtySalableResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStockItem

> InventoryStockItem GetStockItem(ctx).Body(body).Execute()

Get Stock Item

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/gemini-commerce/go-client-inventory"
)

func main() {
    body := *openapiclient.NewInventoryGetStockItemRequest() // InventoryGetStockItemRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InventoryAPI.GetStockItem(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryAPI.GetStockItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStockItem`: InventoryStockItem
    fmt.Fprintf(os.Stdout, "Response from `InventoryAPI.GetStockItem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStockItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**InventoryGetStockItemRequest**](InventoryGetStockItemRequest.md) |  | 

### Return type

[**InventoryStockItem**](InventoryStockItem.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStockItems

> InventoryListStockItemsResponse ListStockItems(ctx).Body(body).Execute()

List Stock Items

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/gemini-commerce/go-client-inventory"
)

func main() {
    body := *openapiclient.NewInventoryListStockItemsRequest() // InventoryListStockItemsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InventoryAPI.ListStockItems(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryAPI.ListStockItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListStockItems`: InventoryListStockItemsResponse
    fmt.Fprintf(os.Stdout, "Response from `InventoryAPI.ListStockItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListStockItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**InventoryListStockItemsRequest**](InventoryListStockItemsRequest.md) |  | 

### Return type

[**InventoryListStockItemsResponse**](InventoryListStockItemsResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStockItemsBySkus

> InventoryListStockItemsBySkusResponse ListStockItemsBySkus(ctx).Body(body).Execute()

List Stock Items by SKUs

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/gemini-commerce/go-client-inventory"
)

func main() {
    body := *openapiclient.NewInventoryListStockItemsBySkusRequest() // InventoryListStockItemsBySkusRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InventoryAPI.ListStockItemsBySkus(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryAPI.ListStockItemsBySkus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListStockItemsBySkus`: InventoryListStockItemsBySkusResponse
    fmt.Fprintf(os.Stdout, "Response from `InventoryAPI.ListStockItemsBySkus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListStockItemsBySkusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**InventoryListStockItemsBySkusRequest**](InventoryListStockItemsBySkusRequest.md) |  | 

### Return type

[**InventoryListStockItemsBySkusResponse**](InventoryListStockItemsBySkusResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RebalanceStockQtys

> InventoryStockItem RebalanceStockQtys(ctx).Body(body).Execute()

Rebalance Stock Quantities

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/gemini-commerce/go-client-inventory"
)

func main() {
    body := *openapiclient.NewInventoryRebalanceStockQtysRequest() // InventoryRebalanceStockQtysRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InventoryAPI.RebalanceStockQtys(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryAPI.RebalanceStockQtys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RebalanceStockQtys`: InventoryStockItem
    fmt.Fprintf(os.Stdout, "Response from `InventoryAPI.RebalanceStockQtys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRebalanceStockQtysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**InventoryRebalanceStockQtysRequest**](InventoryRebalanceStockQtysRequest.md) |  | 

### Return type

[**InventoryStockItem**](InventoryStockItem.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStockItem

> InventoryStockItem UpdateStockItem(ctx).Body(body).Execute()

Update Stock Item

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/gemini-commerce/go-client-inventory"
)

func main() {
    body := *openapiclient.NewInventoryUpdateStockItemRequest() // InventoryUpdateStockItemRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InventoryAPI.UpdateStockItem(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryAPI.UpdateStockItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateStockItem`: InventoryStockItem
    fmt.Fprintf(os.Stdout, "Response from `InventoryAPI.UpdateStockItem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStockItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**InventoryUpdateStockItemRequest**](InventoryUpdateStockItemRequest.md) |  | 

### Return type

[**InventoryStockItem**](InventoryStockItem.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

