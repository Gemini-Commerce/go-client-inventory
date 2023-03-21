/*
inventory.proto

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: version not set
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package inventory

import (
	"encoding/json"
)

// InventoryListStockItemsBySkusResponse struct for InventoryListStockItemsBySkusResponse
type InventoryListStockItemsBySkusResponse struct {
	StockItems []InventoryStockItem `json:"stockItems,omitempty"`
}

// NewInventoryListStockItemsBySkusResponse instantiates a new InventoryListStockItemsBySkusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryListStockItemsBySkusResponse() *InventoryListStockItemsBySkusResponse {
	this := InventoryListStockItemsBySkusResponse{}
	return &this
}

// NewInventoryListStockItemsBySkusResponseWithDefaults instantiates a new InventoryListStockItemsBySkusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryListStockItemsBySkusResponseWithDefaults() *InventoryListStockItemsBySkusResponse {
	this := InventoryListStockItemsBySkusResponse{}
	return &this
}

// GetStockItems returns the StockItems field value if set, zero value otherwise.
func (o *InventoryListStockItemsBySkusResponse) GetStockItems() []InventoryStockItem {
	if o == nil || isNil(o.StockItems) {
		var ret []InventoryStockItem
		return ret
	}
	return o.StockItems
}

// GetStockItemsOk returns a tuple with the StockItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryListStockItemsBySkusResponse) GetStockItemsOk() ([]InventoryStockItem, bool) {
	if o == nil || isNil(o.StockItems) {
    return nil, false
	}
	return o.StockItems, true
}

// HasStockItems returns a boolean if a field has been set.
func (o *InventoryListStockItemsBySkusResponse) HasStockItems() bool {
	if o != nil && !isNil(o.StockItems) {
		return true
	}

	return false
}

// SetStockItems gets a reference to the given []InventoryStockItem and assigns it to the StockItems field.
func (o *InventoryListStockItemsBySkusResponse) SetStockItems(v []InventoryStockItem) {
	o.StockItems = v
}

func (o InventoryListStockItemsBySkusResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.StockItems) {
		toSerialize["stockItems"] = o.StockItems
	}
	return json.Marshal(toSerialize)
}

type NullableInventoryListStockItemsBySkusResponse struct {
	value *InventoryListStockItemsBySkusResponse
	isSet bool
}

func (v NullableInventoryListStockItemsBySkusResponse) Get() *InventoryListStockItemsBySkusResponse {
	return v.value
}

func (v *NullableInventoryListStockItemsBySkusResponse) Set(val *InventoryListStockItemsBySkusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryListStockItemsBySkusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryListStockItemsBySkusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryListStockItemsBySkusResponse(val *InventoryListStockItemsBySkusResponse) *NullableInventoryListStockItemsBySkusResponse {
	return &NullableInventoryListStockItemsBySkusResponse{value: val, isSet: true}
}

func (v NullableInventoryListStockItemsBySkusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryListStockItemsBySkusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

