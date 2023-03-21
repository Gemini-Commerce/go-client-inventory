/*
product/product.proto

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: version not set
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package inventory

import (
	"encoding/json"
)

// EntitymanagerListAttributeOptionsRequest struct for EntitymanagerListAttributeOptionsRequest
type EntitymanagerListAttributeOptionsRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	ListCode *string `json:"listCode,omitempty"`
}

// NewEntitymanagerListAttributeOptionsRequest instantiates a new EntitymanagerListAttributeOptionsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitymanagerListAttributeOptionsRequest() *EntitymanagerListAttributeOptionsRequest {
	this := EntitymanagerListAttributeOptionsRequest{}
	return &this
}

// NewEntitymanagerListAttributeOptionsRequestWithDefaults instantiates a new EntitymanagerListAttributeOptionsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitymanagerListAttributeOptionsRequestWithDefaults() *EntitymanagerListAttributeOptionsRequest {
	this := EntitymanagerListAttributeOptionsRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *EntitymanagerListAttributeOptionsRequest) GetTenantId() string {
	if o == nil || isNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerListAttributeOptionsRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || isNil(o.TenantId) {
    return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *EntitymanagerListAttributeOptionsRequest) HasTenantId() bool {
	if o != nil && !isNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *EntitymanagerListAttributeOptionsRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetListCode returns the ListCode field value if set, zero value otherwise.
func (o *EntitymanagerListAttributeOptionsRequest) GetListCode() string {
	if o == nil || isNil(o.ListCode) {
		var ret string
		return ret
	}
	return *o.ListCode
}

// GetListCodeOk returns a tuple with the ListCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitymanagerListAttributeOptionsRequest) GetListCodeOk() (*string, bool) {
	if o == nil || isNil(o.ListCode) {
    return nil, false
	}
	return o.ListCode, true
}

// HasListCode returns a boolean if a field has been set.
func (o *EntitymanagerListAttributeOptionsRequest) HasListCode() bool {
	if o != nil && !isNil(o.ListCode) {
		return true
	}

	return false
}

// SetListCode gets a reference to the given string and assigns it to the ListCode field.
func (o *EntitymanagerListAttributeOptionsRequest) SetListCode(v string) {
	o.ListCode = &v
}

func (o EntitymanagerListAttributeOptionsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !isNil(o.ListCode) {
		toSerialize["listCode"] = o.ListCode
	}
	return json.Marshal(toSerialize)
}

type NullableEntitymanagerListAttributeOptionsRequest struct {
	value *EntitymanagerListAttributeOptionsRequest
	isSet bool
}

func (v NullableEntitymanagerListAttributeOptionsRequest) Get() *EntitymanagerListAttributeOptionsRequest {
	return v.value
}

func (v *NullableEntitymanagerListAttributeOptionsRequest) Set(val *EntitymanagerListAttributeOptionsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitymanagerListAttributeOptionsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitymanagerListAttributeOptionsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitymanagerListAttributeOptionsRequest(val *EntitymanagerListAttributeOptionsRequest) *NullableEntitymanagerListAttributeOptionsRequest {
	return &NullableEntitymanagerListAttributeOptionsRequest{value: val, isSet: true}
}

func (v NullableEntitymanagerListAttributeOptionsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitymanagerListAttributeOptionsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

