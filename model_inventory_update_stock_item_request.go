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

// checks if the InventoryUpdateStockItemRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InventoryUpdateStockItemRequest{}

// InventoryUpdateStockItemRequest struct for InventoryUpdateStockItemRequest
type InventoryUpdateStockItemRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	Sku *string `json:"sku,omitempty"`
	Payload *UpdateStockItemRequestPayload `json:"payload,omitempty"`
	PayloadMask *string `json:"payloadMask,omitempty"`
}

// NewInventoryUpdateStockItemRequest instantiates a new InventoryUpdateStockItemRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryUpdateStockItemRequest() *InventoryUpdateStockItemRequest {
	this := InventoryUpdateStockItemRequest{}
	return &this
}

// NewInventoryUpdateStockItemRequestWithDefaults instantiates a new InventoryUpdateStockItemRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryUpdateStockItemRequestWithDefaults() *InventoryUpdateStockItemRequest {
	this := InventoryUpdateStockItemRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *InventoryUpdateStockItemRequest) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryUpdateStockItemRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *InventoryUpdateStockItemRequest) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *InventoryUpdateStockItemRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *InventoryUpdateStockItemRequest) GetSku() string {
	if o == nil || IsNil(o.Sku) {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryUpdateStockItemRequest) GetSkuOk() (*string, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *InventoryUpdateStockItemRequest) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *InventoryUpdateStockItemRequest) SetSku(v string) {
	o.Sku = &v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *InventoryUpdateStockItemRequest) GetPayload() UpdateStockItemRequestPayload {
	if o == nil || IsNil(o.Payload) {
		var ret UpdateStockItemRequestPayload
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryUpdateStockItemRequest) GetPayloadOk() (*UpdateStockItemRequestPayload, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *InventoryUpdateStockItemRequest) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given UpdateStockItemRequestPayload and assigns it to the Payload field.
func (o *InventoryUpdateStockItemRequest) SetPayload(v UpdateStockItemRequestPayload) {
	o.Payload = &v
}

// GetPayloadMask returns the PayloadMask field value if set, zero value otherwise.
func (o *InventoryUpdateStockItemRequest) GetPayloadMask() string {
	if o == nil || IsNil(o.PayloadMask) {
		var ret string
		return ret
	}
	return *o.PayloadMask
}

// GetPayloadMaskOk returns a tuple with the PayloadMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryUpdateStockItemRequest) GetPayloadMaskOk() (*string, bool) {
	if o == nil || IsNil(o.PayloadMask) {
		return nil, false
	}
	return o.PayloadMask, true
}

// HasPayloadMask returns a boolean if a field has been set.
func (o *InventoryUpdateStockItemRequest) HasPayloadMask() bool {
	if o != nil && !IsNil(o.PayloadMask) {
		return true
	}

	return false
}

// SetPayloadMask gets a reference to the given string and assigns it to the PayloadMask field.
func (o *InventoryUpdateStockItemRequest) SetPayloadMask(v string) {
	o.PayloadMask = &v
}

func (o InventoryUpdateStockItemRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InventoryUpdateStockItemRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.PayloadMask) {
		toSerialize["payloadMask"] = o.PayloadMask
	}
	return toSerialize, nil
}

type NullableInventoryUpdateStockItemRequest struct {
	value *InventoryUpdateStockItemRequest
	isSet bool
}

func (v NullableInventoryUpdateStockItemRequest) Get() *InventoryUpdateStockItemRequest {
	return v.value
}

func (v *NullableInventoryUpdateStockItemRequest) Set(val *InventoryUpdateStockItemRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryUpdateStockItemRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryUpdateStockItemRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryUpdateStockItemRequest(val *InventoryUpdateStockItemRequest) *NullableInventoryUpdateStockItemRequest {
	return &NullableInventoryUpdateStockItemRequest{value: val, isSet: true}
}

func (v NullableInventoryUpdateStockItemRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryUpdateStockItemRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


