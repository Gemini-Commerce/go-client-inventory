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

// checks if the InventoryListStockItemsBySkusRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InventoryListStockItemsBySkusRequest{}

// InventoryListStockItemsBySkusRequest struct for InventoryListStockItemsBySkusRequest
type InventoryListStockItemsBySkusRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	Skus []string `json:"skus,omitempty"`
}

// NewInventoryListStockItemsBySkusRequest instantiates a new InventoryListStockItemsBySkusRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryListStockItemsBySkusRequest() *InventoryListStockItemsBySkusRequest {
	this := InventoryListStockItemsBySkusRequest{}
	return &this
}

// NewInventoryListStockItemsBySkusRequestWithDefaults instantiates a new InventoryListStockItemsBySkusRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryListStockItemsBySkusRequestWithDefaults() *InventoryListStockItemsBySkusRequest {
	this := InventoryListStockItemsBySkusRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *InventoryListStockItemsBySkusRequest) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryListStockItemsBySkusRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *InventoryListStockItemsBySkusRequest) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *InventoryListStockItemsBySkusRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetSkus returns the Skus field value if set, zero value otherwise.
func (o *InventoryListStockItemsBySkusRequest) GetSkus() []string {
	if o == nil || IsNil(o.Skus) {
		var ret []string
		return ret
	}
	return o.Skus
}

// GetSkusOk returns a tuple with the Skus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryListStockItemsBySkusRequest) GetSkusOk() ([]string, bool) {
	if o == nil || IsNil(o.Skus) {
		return nil, false
	}
	return o.Skus, true
}

// HasSkus returns a boolean if a field has been set.
func (o *InventoryListStockItemsBySkusRequest) HasSkus() bool {
	if o != nil && !IsNil(o.Skus) {
		return true
	}

	return false
}

// SetSkus gets a reference to the given []string and assigns it to the Skus field.
func (o *InventoryListStockItemsBySkusRequest) SetSkus(v []string) {
	o.Skus = v
}

func (o InventoryListStockItemsBySkusRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InventoryListStockItemsBySkusRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.Skus) {
		toSerialize["skus"] = o.Skus
	}
	return toSerialize, nil
}

type NullableInventoryListStockItemsBySkusRequest struct {
	value *InventoryListStockItemsBySkusRequest
	isSet bool
}

func (v NullableInventoryListStockItemsBySkusRequest) Get() *InventoryListStockItemsBySkusRequest {
	return v.value
}

func (v *NullableInventoryListStockItemsBySkusRequest) Set(val *InventoryListStockItemsBySkusRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryListStockItemsBySkusRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryListStockItemsBySkusRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryListStockItemsBySkusRequest(val *InventoryListStockItemsBySkusRequest) *NullableInventoryListStockItemsBySkusRequest {
	return &NullableInventoryListStockItemsBySkusRequest{value: val, isSet: true}
}

func (v NullableInventoryListStockItemsBySkusRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryListStockItemsBySkusRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


