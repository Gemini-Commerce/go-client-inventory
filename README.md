# Go API client for inventory

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v1
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import inventory "github.com/gemini-commerce/go-client-inventory"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), inventory.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), inventory.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), inventory.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), inventory.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://inventory.api.gogemini.io*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*InventoryApi* | [**AdjustQty**](docs/InventoryApi.md#adjustqty) | **Post** /inventory.Inventory/AdjustQty | Adjust Quantity
*InventoryApi* | [**AdjustQtyCommitted**](docs/InventoryApi.md#adjustqtycommitted) | **Post** /inventory.Inventory/AdjustQtyCommitted | Adjust Quantity Committed
*InventoryApi* | [**CreateStockItem**](docs/InventoryApi.md#createstockitem) | **Post** /inventory.Inventory/CreateStockItem | Create Stock Item
*InventoryApi* | [**GetQtySalable**](docs/InventoryApi.md#getqtysalable) | **Post** /inventory.Inventory/GetQtySalable | Get Salable Quantity
*InventoryApi* | [**GetStockItem**](docs/InventoryApi.md#getstockitem) | **Post** /inventory.Inventory/GetStockItem | Get Stock Item
*InventoryApi* | [**ListStockItems**](docs/InventoryApi.md#liststockitems) | **Post** /inventory.Inventory/ListStockItems | List Stock Items
*InventoryApi* | [**ListStockItemsBySkus**](docs/InventoryApi.md#liststockitemsbyskus) | **Post** /inventory.Inventory/ListStockItemsBySkus | List Stock Items by SKUs
*InventoryApi* | [**RebalanceStockQtys**](docs/InventoryApi.md#rebalancestockqtys) | **Post** /inventory.Inventory/RebalanceStockQtys | Rebalance Stock Quantities
*InventoryApi* | [**UpdateStockItem**](docs/InventoryApi.md#updatestockitem) | **Post** /inventory.Inventory/UpdateStockItem | Update Stock Item


## Documentation For Models

 - [InventoryAdjustQtyCommittedRequest](docs/InventoryAdjustQtyCommittedRequest.md)
 - [InventoryAdjustQtyRequest](docs/InventoryAdjustQtyRequest.md)
 - [InventoryCreateStockItemRequest](docs/InventoryCreateStockItemRequest.md)
 - [InventoryGetQtySalableRequest](docs/InventoryGetQtySalableRequest.md)
 - [InventoryGetQtySalableResponse](docs/InventoryGetQtySalableResponse.md)
 - [InventoryGetStockItemRequest](docs/InventoryGetStockItemRequest.md)
 - [InventoryListStockItemsBySkusRequest](docs/InventoryListStockItemsBySkusRequest.md)
 - [InventoryListStockItemsBySkusResponse](docs/InventoryListStockItemsBySkusResponse.md)
 - [InventoryListStockItemsRequest](docs/InventoryListStockItemsRequest.md)
 - [InventoryListStockItemsResponse](docs/InventoryListStockItemsResponse.md)
 - [InventoryRebalanceStockQtysRequest](docs/InventoryRebalanceStockQtysRequest.md)
 - [InventoryStockItem](docs/InventoryStockItem.md)
 - [InventoryUpdateStockItemRequest](docs/InventoryUpdateStockItemRequest.md)
 - [ProtobufAny](docs/ProtobufAny.md)
 - [RpcStatus](docs/RpcStatus.md)
 - [UpdateStockItemRequestPayload](docs/UpdateStockItemRequestPayload.md)


## Documentation For Authorization

 Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

info@gemini-commerce.com

