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

// InventoryGetQtySalableRequest struct for InventoryGetQtySalableRequest
type InventoryGetQtySalableRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	Sku *string `json:"sku,omitempty"`
}

// NewInventoryGetQtySalableRequest instantiates a new InventoryGetQtySalableRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryGetQtySalableRequest() *InventoryGetQtySalableRequest {
	this := InventoryGetQtySalableRequest{}
	return &this
}

// NewInventoryGetQtySalableRequestWithDefaults instantiates a new InventoryGetQtySalableRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryGetQtySalableRequestWithDefaults() *InventoryGetQtySalableRequest {
	this := InventoryGetQtySalableRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *InventoryGetQtySalableRequest) GetTenantId() string {
	if o == nil || isNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryGetQtySalableRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || isNil(o.TenantId) {
    return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *InventoryGetQtySalableRequest) HasTenantId() bool {
	if o != nil && !isNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *InventoryGetQtySalableRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *InventoryGetQtySalableRequest) GetSku() string {
	if o == nil || isNil(o.Sku) {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryGetQtySalableRequest) GetSkuOk() (*string, bool) {
	if o == nil || isNil(o.Sku) {
    return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *InventoryGetQtySalableRequest) HasSku() bool {
	if o != nil && !isNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *InventoryGetQtySalableRequest) SetSku(v string) {
	o.Sku = &v
}

func (o InventoryGetQtySalableRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !isNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	return json.Marshal(toSerialize)
}

type NullableInventoryGetQtySalableRequest struct {
	value *InventoryGetQtySalableRequest
	isSet bool
}

func (v NullableInventoryGetQtySalableRequest) Get() *InventoryGetQtySalableRequest {
	return v.value
}

func (v *NullableInventoryGetQtySalableRequest) Set(val *InventoryGetQtySalableRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryGetQtySalableRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryGetQtySalableRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryGetQtySalableRequest(val *InventoryGetQtySalableRequest) *NullableInventoryGetQtySalableRequest {
	return &NullableInventoryGetQtySalableRequest{value: val, isSet: true}
}

func (v NullableInventoryGetQtySalableRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryGetQtySalableRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


