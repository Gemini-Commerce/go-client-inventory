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

// InventoryListStockItemsRequest struct for InventoryListStockItemsRequest
type InventoryListStockItemsRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	// The maximum number of items to return.
	PageSize *int64 `json:"pageSize,omitempty"`
	// The next_page_token value returned from a previous List request, if any.
	PageToken *string `json:"pageToken,omitempty"`
}

// NewInventoryListStockItemsRequest instantiates a new InventoryListStockItemsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryListStockItemsRequest() *InventoryListStockItemsRequest {
	this := InventoryListStockItemsRequest{}
	return &this
}

// NewInventoryListStockItemsRequestWithDefaults instantiates a new InventoryListStockItemsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryListStockItemsRequestWithDefaults() *InventoryListStockItemsRequest {
	this := InventoryListStockItemsRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *InventoryListStockItemsRequest) GetTenantId() string {
	if o == nil || isNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryListStockItemsRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || isNil(o.TenantId) {
    return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *InventoryListStockItemsRequest) HasTenantId() bool {
	if o != nil && !isNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *InventoryListStockItemsRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *InventoryListStockItemsRequest) GetPageSize() int64 {
	if o == nil || isNil(o.PageSize) {
		var ret int64
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryListStockItemsRequest) GetPageSizeOk() (*int64, bool) {
	if o == nil || isNil(o.PageSize) {
    return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *InventoryListStockItemsRequest) HasPageSize() bool {
	if o != nil && !isNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int64 and assigns it to the PageSize field.
func (o *InventoryListStockItemsRequest) SetPageSize(v int64) {
	o.PageSize = &v
}

// GetPageToken returns the PageToken field value if set, zero value otherwise.
func (o *InventoryListStockItemsRequest) GetPageToken() string {
	if o == nil || isNil(o.PageToken) {
		var ret string
		return ret
	}
	return *o.PageToken
}

// GetPageTokenOk returns a tuple with the PageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryListStockItemsRequest) GetPageTokenOk() (*string, bool) {
	if o == nil || isNil(o.PageToken) {
    return nil, false
	}
	return o.PageToken, true
}

// HasPageToken returns a boolean if a field has been set.
func (o *InventoryListStockItemsRequest) HasPageToken() bool {
	if o != nil && !isNil(o.PageToken) {
		return true
	}

	return false
}

// SetPageToken gets a reference to the given string and assigns it to the PageToken field.
func (o *InventoryListStockItemsRequest) SetPageToken(v string) {
	o.PageToken = &v
}

func (o InventoryListStockItemsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !isNil(o.PageSize) {
		toSerialize["pageSize"] = o.PageSize
	}
	if !isNil(o.PageToken) {
		toSerialize["pageToken"] = o.PageToken
	}
	return json.Marshal(toSerialize)
}

type NullableInventoryListStockItemsRequest struct {
	value *InventoryListStockItemsRequest
	isSet bool
}

func (v NullableInventoryListStockItemsRequest) Get() *InventoryListStockItemsRequest {
	return v.value
}

func (v *NullableInventoryListStockItemsRequest) Set(val *InventoryListStockItemsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryListStockItemsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryListStockItemsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryListStockItemsRequest(val *InventoryListStockItemsRequest) *NullableInventoryListStockItemsRequest {
	return &NullableInventoryListStockItemsRequest{value: val, isSet: true}
}

func (v NullableInventoryListStockItemsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryListStockItemsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

