/*
inventory Service

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
Contact: info@gemini-commerce.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package inventory

import (
	"encoding/json"
)

// InventoryAdjustQtyCommittedRequest struct for InventoryAdjustQtyCommittedRequest
type InventoryAdjustQtyCommittedRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	Sku *string `json:"sku,omitempty"`
	QtyAdjust *int32 `json:"qtyAdjust,omitempty"`
}

// NewInventoryAdjustQtyCommittedRequest instantiates a new InventoryAdjustQtyCommittedRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryAdjustQtyCommittedRequest() *InventoryAdjustQtyCommittedRequest {
	this := InventoryAdjustQtyCommittedRequest{}
	return &this
}

// NewInventoryAdjustQtyCommittedRequestWithDefaults instantiates a new InventoryAdjustQtyCommittedRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryAdjustQtyCommittedRequestWithDefaults() *InventoryAdjustQtyCommittedRequest {
	this := InventoryAdjustQtyCommittedRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *InventoryAdjustQtyCommittedRequest) GetTenantId() string {
	if o == nil || isNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryAdjustQtyCommittedRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || isNil(o.TenantId) {
    return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *InventoryAdjustQtyCommittedRequest) HasTenantId() bool {
	if o != nil && !isNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *InventoryAdjustQtyCommittedRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *InventoryAdjustQtyCommittedRequest) GetSku() string {
	if o == nil || isNil(o.Sku) {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryAdjustQtyCommittedRequest) GetSkuOk() (*string, bool) {
	if o == nil || isNil(o.Sku) {
    return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *InventoryAdjustQtyCommittedRequest) HasSku() bool {
	if o != nil && !isNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *InventoryAdjustQtyCommittedRequest) SetSku(v string) {
	o.Sku = &v
}

// GetQtyAdjust returns the QtyAdjust field value if set, zero value otherwise.
func (o *InventoryAdjustQtyCommittedRequest) GetQtyAdjust() int32 {
	if o == nil || isNil(o.QtyAdjust) {
		var ret int32
		return ret
	}
	return *o.QtyAdjust
}

// GetQtyAdjustOk returns a tuple with the QtyAdjust field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryAdjustQtyCommittedRequest) GetQtyAdjustOk() (*int32, bool) {
	if o == nil || isNil(o.QtyAdjust) {
    return nil, false
	}
	return o.QtyAdjust, true
}

// HasQtyAdjust returns a boolean if a field has been set.
func (o *InventoryAdjustQtyCommittedRequest) HasQtyAdjust() bool {
	if o != nil && !isNil(o.QtyAdjust) {
		return true
	}

	return false
}

// SetQtyAdjust gets a reference to the given int32 and assigns it to the QtyAdjust field.
func (o *InventoryAdjustQtyCommittedRequest) SetQtyAdjust(v int32) {
	o.QtyAdjust = &v
}

func (o InventoryAdjustQtyCommittedRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !isNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	if !isNil(o.QtyAdjust) {
		toSerialize["qtyAdjust"] = o.QtyAdjust
	}
	return json.Marshal(toSerialize)
}

type NullableInventoryAdjustQtyCommittedRequest struct {
	value *InventoryAdjustQtyCommittedRequest
	isSet bool
}

func (v NullableInventoryAdjustQtyCommittedRequest) Get() *InventoryAdjustQtyCommittedRequest {
	return v.value
}

func (v *NullableInventoryAdjustQtyCommittedRequest) Set(val *InventoryAdjustQtyCommittedRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryAdjustQtyCommittedRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryAdjustQtyCommittedRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryAdjustQtyCommittedRequest(val *InventoryAdjustQtyCommittedRequest) *NullableInventoryAdjustQtyCommittedRequest {
	return &NullableInventoryAdjustQtyCommittedRequest{value: val, isSet: true}
}

func (v NullableInventoryAdjustQtyCommittedRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryAdjustQtyCommittedRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


