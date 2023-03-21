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

// InventoryListStockItemsResponse struct for InventoryListStockItemsResponse
type InventoryListStockItemsResponse struct {
	StockItems []InventoryStockItem `json:"stockItems,omitempty"`
	// Token to retrieve the next page of results, or empty if there are no more results in the list.
	NextPageToken *string `json:"nextPageToken,omitempty"`
}

// NewInventoryListStockItemsResponse instantiates a new InventoryListStockItemsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryListStockItemsResponse() *InventoryListStockItemsResponse {
	this := InventoryListStockItemsResponse{}
	return &this
}

// NewInventoryListStockItemsResponseWithDefaults instantiates a new InventoryListStockItemsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryListStockItemsResponseWithDefaults() *InventoryListStockItemsResponse {
	this := InventoryListStockItemsResponse{}
	return &this
}

// GetStockItems returns the StockItems field value if set, zero value otherwise.
func (o *InventoryListStockItemsResponse) GetStockItems() []InventoryStockItem {
	if o == nil || isNil(o.StockItems) {
		var ret []InventoryStockItem
		return ret
	}
	return o.StockItems
}

// GetStockItemsOk returns a tuple with the StockItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryListStockItemsResponse) GetStockItemsOk() ([]InventoryStockItem, bool) {
	if o == nil || isNil(o.StockItems) {
    return nil, false
	}
	return o.StockItems, true
}

// HasStockItems returns a boolean if a field has been set.
func (o *InventoryListStockItemsResponse) HasStockItems() bool {
	if o != nil && !isNil(o.StockItems) {
		return true
	}

	return false
}

// SetStockItems gets a reference to the given []InventoryStockItem and assigns it to the StockItems field.
func (o *InventoryListStockItemsResponse) SetStockItems(v []InventoryStockItem) {
	o.StockItems = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *InventoryListStockItemsResponse) GetNextPageToken() string {
	if o == nil || isNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryListStockItemsResponse) GetNextPageTokenOk() (*string, bool) {
	if o == nil || isNil(o.NextPageToken) {
    return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *InventoryListStockItemsResponse) HasNextPageToken() bool {
	if o != nil && !isNil(o.NextPageToken) {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *InventoryListStockItemsResponse) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o InventoryListStockItemsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.StockItems) {
		toSerialize["stockItems"] = o.StockItems
	}
	if !isNil(o.NextPageToken) {
		toSerialize["nextPageToken"] = o.NextPageToken
	}
	return json.Marshal(toSerialize)
}

type NullableInventoryListStockItemsResponse struct {
	value *InventoryListStockItemsResponse
	isSet bool
}

func (v NullableInventoryListStockItemsResponse) Get() *InventoryListStockItemsResponse {
	return v.value
}

func (v *NullableInventoryListStockItemsResponse) Set(val *InventoryListStockItemsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryListStockItemsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryListStockItemsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryListStockItemsResponse(val *InventoryListStockItemsResponse) *NullableInventoryListStockItemsResponse {
	return &NullableInventoryListStockItemsResponse{value: val, isSet: true}
}

func (v NullableInventoryListStockItemsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryListStockItemsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


