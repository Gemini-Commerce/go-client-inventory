# Go API client for inventory

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: version not set
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
import inventory "github.com/GIT_USER_ID/GIT_REPO_ID"
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

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*InventoryApi* | [**InventoryAdjustQty**](docs/InventoryApi.md#inventoryadjustqty) | **Post** /inventory.Inventory/AdjustQty | 
*InventoryApi* | [**InventoryAdjustQtyCommitted**](docs/InventoryApi.md#inventoryadjustqtycommitted) | **Post** /inventory.Inventory/AdjustQtyCommitted | adjust operations are intended for increment/decrement quantities fields by a certain amount
*InventoryApi* | [**InventoryCreateStockItem**](docs/InventoryApi.md#inventorycreatestockitem) | **Post** /inventory.Inventory/CreateStockItem | 
*InventoryApi* | [**InventoryGetQtySalable**](docs/InventoryApi.md#inventorygetqtysalable) | **Post** /inventory.Inventory/GetQtySalable | 
*InventoryApi* | [**InventoryGetStockItem**](docs/InventoryApi.md#inventorygetstockitem) | **Post** /inventory.Inventory/GetStockItem | 
*InventoryApi* | [**InventoryListStockItems**](docs/InventoryApi.md#inventoryliststockitems) | **Post** /inventory.Inventory/ListStockItems | 
*InventoryApi* | [**InventoryListStockItemsBySkus**](docs/InventoryApi.md#inventoryliststockitemsbyskus) | **Post** /inventory.Inventory/ListStockItemsBySkus | 
*InventoryApi* | [**InventoryRebalanceStockQtys**](docs/InventoryApi.md#inventoryrebalancestockqtys) | **Post** /inventory.Inventory/RebalanceStockQtys | rebalance operation is intended for apply or discard a certain amount of qty_committed to qty and old qty_committed
*InventoryApi* | [**InventoryUpdateStockItem**](docs/InventoryApi.md#inventoryupdatestockitem) | **Post** /inventory.Inventory/UpdateStockItem | 


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



