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

// InventoryRebalanceStockQtysRequest struct for InventoryRebalanceStockQtysRequest
type InventoryRebalanceStockQtysRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	Sku *string `json:"sku,omitempty"`
	QtyCommitted *int32 `json:"qtyCommitted,omitempty"`
}

// NewInventoryRebalanceStockQtysRequest instantiates a new InventoryRebalanceStockQtysRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryRebalanceStockQtysRequest() *InventoryRebalanceStockQtysRequest {
	this := InventoryRebalanceStockQtysRequest{}
	return &this
}

// NewInventoryRebalanceStockQtysRequestWithDefaults instantiates a new InventoryRebalanceStockQtysRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryRebalanceStockQtysRequestWithDefaults() *InventoryRebalanceStockQtysRequest {
	this := InventoryRebalanceStockQtysRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *InventoryRebalanceStockQtysRequest) GetTenantId() string {
	if o == nil || isNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryRebalanceStockQtysRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || isNil(o.TenantId) {
    return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *InventoryRebalanceStockQtysRequest) HasTenantId() bool {
	if o != nil && !isNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *InventoryRebalanceStockQtysRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *InventoryRebalanceStockQtysRequest) GetSku() string {
	if o == nil || isNil(o.Sku) {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryRebalanceStockQtysRequest) GetSkuOk() (*string, bool) {
	if o == nil || isNil(o.Sku) {
    return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *InventoryRebalanceStockQtysRequest) HasSku() bool {
	if o != nil && !isNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *InventoryRebalanceStockQtysRequest) SetSku(v string) {
	o.Sku = &v
}

// GetQtyCommitted returns the QtyCommitted field value if set, zero value otherwise.
func (o *InventoryRebalanceStockQtysRequest) GetQtyCommitted() int32 {
	if o == nil || isNil(o.QtyCommitted) {
		var ret int32
		return ret
	}
	return *o.QtyCommitted
}

// GetQtyCommittedOk returns a tuple with the QtyCommitted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryRebalanceStockQtysRequest) GetQtyCommittedOk() (*int32, bool) {
	if o == nil || isNil(o.QtyCommitted) {
    return nil, false
	}
	return o.QtyCommitted, true
}

// HasQtyCommitted returns a boolean if a field has been set.
func (o *InventoryRebalanceStockQtysRequest) HasQtyCommitted() bool {
	if o != nil && !isNil(o.QtyCommitted) {
		return true
	}

	return false
}

// SetQtyCommitted gets a reference to the given int32 and assigns it to the QtyCommitted field.
func (o *InventoryRebalanceStockQtysRequest) SetQtyCommitted(v int32) {
	o.QtyCommitted = &v
}

func (o InventoryRebalanceStockQtysRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !isNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	if !isNil(o.QtyCommitted) {
		toSerialize["qtyCommitted"] = o.QtyCommitted
	}
	return json.Marshal(toSerialize)
}

type NullableInventoryRebalanceStockQtysRequest struct {
	value *InventoryRebalanceStockQtysRequest
	isSet bool
}

func (v NullableInventoryRebalanceStockQtysRequest) Get() *InventoryRebalanceStockQtysRequest {
	return v.value
}

func (v *NullableInventoryRebalanceStockQtysRequest) Set(val *InventoryRebalanceStockQtysRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryRebalanceStockQtysRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryRebalanceStockQtysRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryRebalanceStockQtysRequest(val *InventoryRebalanceStockQtysRequest) *NullableInventoryRebalanceStockQtysRequest {
	return &NullableInventoryRebalanceStockQtysRequest{value: val, isSet: true}
}

func (v NullableInventoryRebalanceStockQtysRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryRebalanceStockQtysRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

