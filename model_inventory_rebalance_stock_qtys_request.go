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

// checks if the InventoryRebalanceStockQtysRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InventoryRebalanceStockQtysRequest{}

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
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryRebalanceStockQtysRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *InventoryRebalanceStockQtysRequest) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
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
	if o == nil || IsNil(o.Sku) {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryRebalanceStockQtysRequest) GetSkuOk() (*string, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *InventoryRebalanceStockQtysRequest) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
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
	if o == nil || IsNil(o.QtyCommitted) {
		var ret int32
		return ret
	}
	return *o.QtyCommitted
}

// GetQtyCommittedOk returns a tuple with the QtyCommitted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryRebalanceStockQtysRequest) GetQtyCommittedOk() (*int32, bool) {
	if o == nil || IsNil(o.QtyCommitted) {
		return nil, false
	}
	return o.QtyCommitted, true
}

// HasQtyCommitted returns a boolean if a field has been set.
func (o *InventoryRebalanceStockQtysRequest) HasQtyCommitted() bool {
	if o != nil && !IsNil(o.QtyCommitted) {
		return true
	}

	return false
}

// SetQtyCommitted gets a reference to the given int32 and assigns it to the QtyCommitted field.
func (o *InventoryRebalanceStockQtysRequest) SetQtyCommitted(v int32) {
	o.QtyCommitted = &v
}

func (o InventoryRebalanceStockQtysRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InventoryRebalanceStockQtysRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	if !IsNil(o.QtyCommitted) {
		toSerialize["qtyCommitted"] = o.QtyCommitted
	}
	return toSerialize, nil
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


